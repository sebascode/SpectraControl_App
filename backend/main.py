"""
SpectraControl Backend - FastAPI
Controla Philips Hue Bridge via API local
"""

import asyncio
import threading
import subprocess
import re
import json
import os
from typing import Optional

import httpx
from fastapi import FastAPI, HTTPException, WebSocket, WebSocketDisconnect
from fastapi.middleware.cors import CORSMiddleware
from fastapi.staticfiles import StaticFiles
from pydantic import BaseModel

# ──────────────────────────────────────────────
# CONFIG
# ──────────────────────────────────────────────
_HERE = os.path.dirname(os.path.abspath(__file__))

BRIDGE_IP = os.getenv("HUE_BRIDGE_IP", "")
API_KEY   = os.getenv("HUE_API_KEY", "")

app = FastAPI(title="SpectraControl")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_methods=["*"],
    allow_headers=["*"],
)

# ──────────────────────────────────────────────
# SYNC STATE
# ──────────────────────────────────────────────
sync_state = {
    "running": False,
    "interval": 1.0,
    "group_ids": [],
    "thread": None,
}

# ──────────────────────────────────────────────
# MODELS
# ──────────────────────────────────────────────
class BridgeConfig(BaseModel):
    ip: str
    api_key: str

class LightState(BaseModel):
    on: Optional[bool] = None
    bri: Optional[int] = None   # 1-254
    hue: Optional[int] = None   # 0-65535
    sat: Optional[int] = None   # 0-254
    xy: Optional[list[float]] = None
    ct: Optional[int] = None    # color temp mired 153-500
    transitiontime: Optional[int] = 4  # x100ms

class GroupAction(BaseModel):
    on: Optional[bool] = None
    bri: Optional[int] = None
    ct: Optional[int] = None
    xy: Optional[list[float]] = None

class SyncConfig(BaseModel):
    enabled: bool
    group_ids: list[str] = []
    interval: float = 1.0

class ColorPayload(BaseModel):
    hex_color: str      # e.g. "#FF6A00"
    bri: Optional[int] = 200
    transitiontime: Optional[int] = 4

# ──────────────────────────────────────────────
# HELPERS
# ──────────────────────────────────────────────
def rgb_to_xy(r: int, g: int, b: int) -> list[float]:
    """RGB (0-255) → CIE XY usando gamma correction y matriz Philips Hue."""
    r, g, b = r / 255.0, g / 255.0, b / 255.0
    r = pow((r + 0.055) / 1.055, 2.4) if r > 0.04045 else r / 12.92
    g = pow((g + 0.055) / 1.055, 2.4) if g > 0.04045 else g / 12.92
    b = pow((b + 0.055) / 1.055, 2.4) if b > 0.04045 else b / 12.92
    X = r * 0.664511 + g * 0.154324 + b * 0.162028
    Y = r * 0.283881 + g * 0.668433 + b * 0.047685
    Z = r * 0.000088 + g * 0.072310 + b * 0.986039
    total = X + Y + Z
    if total == 0:
        return [0.0, 0.0]
    return [round(X / total, 4), round(Y / total, 4)]


def hex_to_rgb(hex_color: str) -> tuple[int, int, int]:
    hex_color = hex_color.lstrip("#")
    return tuple(int(hex_color[i:i+2], 16) for i in (0, 2, 4))


async def hue_request(method: str, path: str, data: dict = None):
    if not BRIDGE_IP or not API_KEY:
        raise HTTPException(status_code=503, detail="Bridge no configurado. Ve a /api/config primero.")
    url = f"http://{BRIDGE_IP}/api/{API_KEY}/{path}"
    async with httpx.AsyncClient(timeout=5.0) as client:
        if method == "GET":
            resp = await client.get(url)
        elif method == "PUT":
            resp = await client.put(url, json=data)
        elif method == "POST":
            resp = await client.post(url, json=data)
        else:
            raise ValueError(f"Método no soportado: {method}")
    if resp.status_code != 200:
        raise HTTPException(status_code=resp.status_code, detail=resp.text)
    return resp.json()


# ──────────────────────────────────────────────
# SYNC LOOP
# NOTE: This legacy loop uses mss/xrandr which don't work on KDE Wayland.
# The proper sync path is: frontend getDisplayMedia() → canvas sampling → /ws/color WebSocket.
# ──────────────────────────────────────────────
def sync_loop():
    loop = asyncio.new_event_loop()
    asyncio.set_event_loop(loop)

    async def _sync():
        while sync_state["running"]:
            group_ids = sync_state["group_ids"]
            if not group_ids:
                await asyncio.sleep(sync_state["interval"])
                continue
            try:
                groups_data = await hue_request("GET", "groups")
                all_lights = []
                for gid in group_ids:
                    if gid in groups_data:
                        all_lights.extend(groups_data[gid].get("lights", []))
                if not all_lights:
                    await asyncio.sleep(sync_state["interval"])
                    continue

                result = subprocess.run(["xrandr", "--current"], capture_output=True, text=True, timeout=2)
                screen_w, screen_h = 1920, 1080
                for line in result.stdout.splitlines():
                    match = re.search(r"(\d+)x(\d+)\+0\+0", line)
                    if match:
                        screen_w, screen_h = int(match.group(1)), int(match.group(2))
                        break

                n = len(all_lights)
                tasks = []
                for i, light_id in enumerate(all_lights):
                    px = int(screen_w * (i + 1) / (n + 1))
                    py = screen_h // 2
                    tasks.append(
                        hue_request("PUT", f"lights/{light_id}/state", {
                            "on": True, "xy": [0.3, 0.3], "bri": 200, "transitiontime": 3
                        })
                    )
                await asyncio.gather(*tasks, return_exceptions=True)
            except Exception as e:
                print(f"[sync] error: {e}")
            await asyncio.sleep(sync_state["interval"])

    loop.run_until_complete(_sync())


# ──────────────────────────────────────────────
# ROUTES — Config
# ──────────────────────────────────────────────
@app.get("/api/config")
async def get_config():
    return {
        "bridge_ip": BRIDGE_IP,
        "has_api_key": bool(API_KEY),
        "configured": bool(BRIDGE_IP and API_KEY),
    }


@app.post("/api/config")
async def set_config(cfg: BridgeConfig):
    global BRIDGE_IP, API_KEY
    BRIDGE_IP = cfg.ip
    API_KEY   = cfg.api_key
    with open(os.path.join(_HERE, ".hue_config"), "w") as f:
        json.dump({"ip": BRIDGE_IP, "api_key": API_KEY}, f)
    return {"ok": True}


@app.post("/api/discover")
async def discover_bridge():
    async with httpx.AsyncClient(timeout=5.0) as client:
        try:
            resp = await client.get("https://discovery.meethue.com/")
            return {"bridges": resp.json()}
        except Exception as e:
            raise HTTPException(status_code=503, detail=str(e))


@app.post("/api/pair")
async def pair_bridge(body: dict):
    ip = body.get("ip")
    if not ip:
        raise HTTPException(status_code=400, detail="Falta 'ip'")
    async with httpx.AsyncClient(timeout=10.0) as client:
        resp = await client.post(
            f"http://{ip}/api",
            json={"devicetype": "spectra_control#linux"}
        )
    result = resp.json()
    if isinstance(result, list) and "success" in result[0]:
        return {"api_key": result[0]["success"]["username"]}
    elif isinstance(result, list) and "error" in result[0]:
        raise HTTPException(status_code=403, detail=result[0]["error"]["description"])
    raise HTTPException(status_code=500, detail=str(result))


# ──────────────────────────────────────────────
# ROUTES — Lights
# ──────────────────────────────────────────────
@app.get("/api/lights")
async def get_lights():
    data = await hue_request("GET", "lights")
    return {"lights": [
        {
            "id": lid,
            "name": info["name"],
            "type": info.get("type", ""),
            "on": info["state"]["on"],
            "bri": info["state"].get("bri", 254),
            "reachable": info["state"].get("reachable", False),
            "color_mode": info["state"].get("colormode"),
            "xy": info["state"].get("xy", [0.3, 0.3]),
            "ct": info["state"].get("ct", 300),
            "hue": info["state"].get("hue", 0),
            "sat": info["state"].get("sat", 0),
        }
        for lid, info in data.items()
    ]}


@app.put("/api/lights/{light_id}/state")
async def set_light_state(light_id: str, state: LightState):
    payload = {k: v for k, v in state.dict().items() if v is not None}
    return {"result": await hue_request("PUT", f"lights/{light_id}/state", payload)}


@app.put("/api/lights/{light_id}/color")
async def set_light_color(light_id: str, payload: ColorPayload):
    r, g, b = hex_to_rgb(payload.hex_color)
    return {"result": await hue_request("PUT", f"lights/{light_id}/state", {
        "on": True, "xy": rgb_to_xy(r, g, b),
        "bri": payload.bri, "transitiontime": payload.transitiontime,
    })}


# ──────────────────────────────────────────────
# ROUTES — Groups (Rooms)
# ──────────────────────────────────────────────
@app.get("/api/groups")
async def get_groups():
    data = await hue_request("GET", "groups")
    return {"groups": [
        {
            "id": gid,
            "name": info["name"],
            "type": info.get("type", ""),
            "lights": info.get("lights", []),
            "on": info["action"].get("on", False),
            "bri": info["action"].get("bri", 254),
            "ct": info["action"].get("ct", 300),
            "xy": info["action"].get("xy", [0.3, 0.3]),
        }
        for gid, info in data.items()
    ]}


@app.put("/api/groups/{group_id}/action")
async def set_group_action(group_id: str, action: GroupAction):
    payload = {k: v for k, v in action.dict().items() if v is not None}
    return {"result": await hue_request("PUT", f"groups/{group_id}/action", payload)}


@app.put("/api/groups/{group_id}/color")
async def set_group_color(group_id: str, payload: ColorPayload):
    r, g, b = hex_to_rgb(payload.hex_color)
    return {"result": await hue_request("PUT", f"groups/{group_id}/action", {
        "on": True, "xy": rgb_to_xy(r, g, b),
        "bri": payload.bri, "transitiontime": payload.transitiontime,
    })}


# ──────────────────────────────────────────────
# ROUTES — Sync
# ──────────────────────────────────────────────
@app.get("/api/sync")
async def get_sync_status():
    return {
        "running": sync_state["running"],
        "interval": sync_state["interval"],
        "group_ids": sync_state["group_ids"],
    }


@app.post("/api/sync")
async def set_sync(cfg: SyncConfig):
    if cfg.enabled and not sync_state["running"]:
        sync_state.update({"running": True, "group_ids": cfg.group_ids, "interval": max(0.5, cfg.interval)})
        t = threading.Thread(target=sync_loop, daemon=True)
        t.start()
        sync_state["thread"] = t
        return {"ok": True, "running": True}
    elif not cfg.enabled and sync_state["running"]:
        sync_state.update({"running": False, "group_ids": []})
        return {"ok": True, "running": False}
    elif cfg.enabled and sync_state["running"]:
        sync_state.update({"group_ids": cfg.group_ids, "interval": max(0.5, cfg.interval)})
        return {"ok": True, "running": True, "updated": True}
    return {"ok": True, "running": sync_state["running"]}


# ──────────────────────────────────────────────
# ROUTES — WebSocket screen sync
# ──────────────────────────────────────────────
@app.websocket("/ws/color")
async def ws_color(websocket: WebSocket):
    await websocket.accept()
    print("[ws/color] cliente conectado")
    group_ids: list[str] = []
    bri: int = 200
    try:
        while True:
            data = await websocket.receive_json()
            if "group_ids" in data:
                group_ids = data["group_ids"]
                print(f"[ws/color] grupos: {group_ids}")
            if "bri" in data:
                bri = int(data["bri"])
            tt = int(data.get("transitiontime", 2))
            if not BRIDGE_IP or not API_KEY:
                continue

            if "lights" in data:
                # Formato por luz: [{"id": "1", "r": 128, "g": 64, "b": 200}, ...]
                tasks = [
                    hue_request("PUT", f"lights/{l['id']}/state", {
                        "on": True,
                        "xy": rgb_to_xy(int(l["r"]), int(l["g"]), int(l["b"])),
                        "bri": int(data.get("bri", bri)),
                        "transitiontime": tt,
                    })
                    for l in data["lights"]
                ]
            else:
                # Formato legacy: un color para el grupo completo
                r, g, b = data.get("r"), data.get("g"), data.get("b")
                if r is None or not group_ids:
                    continue
                xy = rgb_to_xy(int(r), int(g), int(b))
                tasks = [
                    hue_request("PUT", f"groups/{gid}/action",
                        {"on": True, "xy": xy, "bri": bri, "transitiontime": tt})
                    for gid in group_ids
                ]

            results = await asyncio.gather(*tasks, return_exceptions=True)
            for res in results:
                if isinstance(res, Exception):
                    print(f"[ws/color] error: {res}")
    except WebSocketDisconnect:
        print("[ws/color] cliente desconectado")
    except Exception as e:
        print(f"[ws/color] excepción: {e}")


# ──────────────────────────────────────────────
# Startup
# ──────────────────────────────────────────────
def load_persisted_config():
    global BRIDGE_IP, API_KEY
    config_path = os.path.join(_HERE, ".hue_config")
    if os.path.exists(config_path):
        try:
            with open(config_path) as f:
                cfg = json.load(f)
            BRIDGE_IP = cfg.get("ip", BRIDGE_IP)
            API_KEY   = cfg.get("api_key", API_KEY)
        except Exception:
            pass

load_persisted_config()

frontend_path = os.path.join(_HERE, "..", "frontend")
if os.path.exists(frontend_path):
    app.mount("/", StaticFiles(directory=frontend_path, html=True), name="frontend")
