---
title: "Migration Guide to the new Hue API"
author: "Philips Hue Developer Program"
source: "Philips Hue Developer Program"
url: "https://developers.meethue.com/develop/hue-api-v2/migration-guide-to-the-new-hue-api/"
date_saved: "2026-05-09T20:42:45.240Z"
date_published: "2021-01-14T21:50:44+00:00"
word_count: "3617"
reading_time: "19 min"
description: "HTTPS The V2 API is only available on HTTPS (i.e. not on HTTP). Details on how to properly implement HTTPS are given in our using HTTPS section. The CURL examples in this guide use the –insecure flag for readability, but this should not be used in production. Local Bridge Discovery Bridge discovery on the local […]"
---

## HTTPS

The V2 API is only available on HTTPS (i.e. not on HTTP). Details on how to properly implement HTTPS are given in our [using HTTPS](https://developers.meethue.com/develop/application-design-guidance/using-https/) section. The CURL examples in this guide use the –insecure flag for readability, but this should not be used in production.

## Local Bridge Discovery

Bridge discovery on the local network has remained mostly unchanged i.e. we still recommend the use mDNS. However, with the release of API V2 we have now officially deprecated support for UPnP as an alternative method. Detailed guidance on local bridge discovery can be found at [Hue Bridge Discovery](https://developers.meethue.com/develop/application-design-guidance/hue-bridge-discovery/).

## API Discovery

To discover whether the CLIPv2 API is available on a bridge, you can currently check the bridge software version on CLIPv1 /api/config endpoint to be at least **1948086000**.

## Application key

The ‘username’ for bridge access has been renamed to ‘application key’ to emphasize it is a key that must be kept secret, and it is moved from a path parameter to a new ‘hue-application-key’ header. The same username retrieved on the V1 API will remain valid to be used as application key on the V2 API. In the future there will be another way to authorize on the V2 API. For referencing an application within a resource (e.g. to indicate resource ownership), the V2 api uses application ids (which are not secret). Retrieving your own application id can be done by performing a GET on `/auth/v1` endpoint with the hue-application-key header. When authenticated, the bridge will respond with 200 OK and a ‘hue-application-id’ response **header**.

## Endpoints

The V1 API was hosted under the `/api` endpoint, whereas the V2 api is hosted under the new `/clip/v2` endpoint. An overview of all resources can be retrieved from the `/clip/v2/resource` endpoint, and the list of resources from a specific type can be queried through `/clip/v2/resource/<resourcetype>`. Commonly used resource types are ‘device’, ‘bridge’, ‘light’, ‘scene’, ‘room’, ‘motion’, ‘button’. Individual resources can be accessed through `/clip/v2/resource/<resourcetype>/<resourceid>` endpoints.

## Example request to retrieve light list

On V1 the following GET request would retrieve the list of lights:

`curl --insecure -X GET 'https://<ipaddress>/api/<username>/lights'`

Combining the new endpoint with the new header results in the following equivalent request on V2:

`curl --insecure -H 'hue-application-key: <appkey (== username)>' -X GET 'https://<ipaddress>/clip/v2/resource/light'`

However, there is a conceptual difference: on the V2 API the devices and it’s services are separated, so the /light endpoint requests the list of light *services*, whereas the /device endpoint request the list of all devices, of which some may expose one or more light services. More on that later in the references section.

An example response with only a single light service might look like this:

{

"errors": \[\],

"data": \[

{

"type": "light",

"id": "c6b028c8-076e-4817-92b1-bcb0cbb78783",

"id\_v1": "/lights/21",

"metadata": {

"name": "Hue downlight right"

},

"on": {

"on": true

},

"dimming": {

"brightness": 100.0

},

"color\_temperature": {

"mirek": 366

},

"color": {

"gamut": {

"blue": {

"x": 0.1532,

"y": 0.0475

},

"green": {

"x": 0.17,

"y": 0.7

},

"red": {

"x": 0.6915,

"y": 0.3083

}

},

"gamut\_type": "C",

"xy": {

"x": 0.4575,

"y": 0.4099

}

}

}

\]

}

{ "errors": \[\], "data": \[ { "type": "light", "id": "c6b028c8-076e-4817-92b1-bcb0cbb78783", "id\_v1": "/lights/21", "metadata": { "name": "Hue downlight right" }, "on": { "on": true }, "dimming": { "brightness": 100.0 }, "color\_temperature": { "mirek": 366 }, "color": { "gamut": { "blue": { "x": 0.1532, "y": 0.0475 }, "green": { "x": 0.17, "y": 0.7 }, "red": { "x": 0.6915, "y": 0.3083 } }, "gamut\_type": "C", "xy": { "x": 0.4575, "y": 0.4099 } } } \] }

```
{
    "errors": [],
    "data": [
        {
            "type": "light",
            "id": "c6b028c8-076e-4817-92b1-bcb0cbb78783",
            "id_v1": "/lights/21",
            "metadata": {
                "name": "Hue downlight right"
            },
            "on": {
                "on": true
            },
            "dimming": {
                "brightness": 100.0
            },
            "color_temperature": {
                "mirek": 366
            },
            "color": {
                "gamut": {
                    "blue": {
                        "x": 0.1532,
                        "y": 0.0475
                    },
                    "green": {
                        "x": 0.17,
                        "y": 0.7
                    },
                    "red": {
                        "x": 0.6915,
                        "y": 0.3083
                    }
                },
                "gamut_type": "C",
                "xy": {
                    "x": 0.4575,
                    "y": 0.4099
                }
            }
        }
    ]
}
```

## Identifiers

To prevent reuse and duplication of identifiers, the V2 API is using Universally Unique Identifiers (UUIDs) for all resources. This does mean that the same resources have a different id on the V2 API than on the V1 API.

Temporarily the resources on V2 will have an id\_v1 field that can be used to find the respective identifier on the V1 API. This would typically be used for a one-time migration to overwrite ids that you may have stored in the context of your application, and it could temporarily be used to support an intermediate application version that uses parts of both APIs.

There are however two caveats. First, some resources on the V2 API do not have a V1 equivalent, and second, the id\_v1 field (as well as the V1 API) will eventually be removed. Therefore your application has to gracefully handle the case that a resource does not have an id\_v1 and/or that the id\_v1 cannot actually be found on the V1 API.

## Example request to change light state

On V1 the following PUT request would set a lights brightness to 100%:

`curl --insecure -X PUT 'https://<ipaddress>/api/<appkey>/lights/<v1-id>/state' -H 'Content-Type: application/json' --data-raw '{"bri": 254}'`

Whereas on the V2 API this looks like:

`curl --insecure -X PUT 'https://<ipaddress>/clip/v2/resource/light/<v2-id>' -H 'hue-application-key: <appkey>' -H 'Content-Type: application/json' --data-raw '{"dimming": {"brightness": 100}}'`

As you can see it is conceptually very similar, however as mentioned the id is different on V2 than on V1, and object structures, key names, and value ranges are different as well. Where possible, we give attributes more meaningful units than in CLIPv1. Note for example that the brightness range is now in percent rather than up to 254. Similar situations apply to temperature in degree Celsius, time in seconds or milliseconds, and so forth.

## Return Codes

The V1 API returned 200 OK as the HTTP status code even when the request failed, so the error would have to be detected by interpreting the response body. The V2 API now returns proper HTTP response codes in case of errors (such as `401 Unauthorized` if your application key is invalid or `404 Not Found` if a resource does not exist). The full list of possible error codes can be found in the [API reference](https://developers.meethue.com/develop/hue-api-v2/api-reference/).

## References

The V2 API is set up in such a way that all resources of the same type are grouped together under the `/resource/<resourcetype>` endpoint, but all those resources commonly reference each other, which is done in a standardized way by indicating the resource type (`rtype`) and resource id (`rid`).

A typical usage is in a single physical device that hosts multiple services. An existing example is the Philips Hue Motion sensor which has a `motion`, `light_level`, and `temperature` service, but theoretically any combination can be supported such as an integrated device with two independently controllable light points and a motion sensor.

This means that the information of the device itself can be found under the `/device` resource endpoint, but it then contains a services array which references for example the light and motion resources, for which the details can be found under the `/light` and `/motion` resource endpoints respectively. Other services the device might have, such as a ZigBee radio (`zigbee_connectivy`) or battery (`device_power`) are modeled in the same way.

Below is an illustrative example of how that could look.

{

"type": "device",

"id": "7b839dff-c2d2-4f90-9509-fea4b461b30d",

"id\_v1": "/lights/21",

"metadata": {

"archetype": "fake\_example",

"name": "User given name"

},

"product\_data": {

"manufacturer\_name": "Signify Netherlands B.V.",

"model\_id": "XXX001",

"product\_name": "Fixed product name",

"software\_version": "x.y.z"

},

"services": \[

{

"rid": "c6b028c8-076e-4817-92b1-bcb0cbb78783",

"rtype": "light"

},

{

"rid": "6cb990a3-7a5c-4c56-9e19-421952405a29",

"rtype": "light"

},

{

"rid": "f7a7f522-22b2-4978-9cf9-f37ed8b53058",

"rtype": "motion"

},

{

"rid": "6ec5432f-fb66-4b07-88de-bb0087a0e33d",

"rtype": "zigbee\_connectivity"

}

\]

}

{ "type": "device", "id": "7b839dff-c2d2-4f90-9509-fea4b461b30d", "id\_v1": "/lights/21", "metadata": { "archetype": "fake\_example", "name": "User given name" }, "product\_data": { "manufacturer\_name": "Signify Netherlands B.V.", "model\_id": "XXX001", "product\_name": "Fixed product name", "software\_version": "x.y.z" }, "services": \[ { "rid": "c6b028c8-076e-4817-92b1-bcb0cbb78783", "rtype": "light" }, { "rid": "6cb990a3-7a5c-4c56-9e19-421952405a29", "rtype": "light" }, { "rid": "f7a7f522-22b2-4978-9cf9-f37ed8b53058", "rtype": "motion" }, { "rid": "6ec5432f-fb66-4b07-88de-bb0087a0e33d", "rtype": "zigbee\_connectivity" } \] }

```
{
    "type": "device",
    "id": "7b839dff-c2d2-4f90-9509-fea4b461b30d",
    "id_v1": "/lights/21",
    "metadata": {
        "archetype": "fake_example",
        "name": "User given name"
    },
    "product_data": {
        "manufacturer_name": "Signify Netherlands B.V.",
        "model_id": "XXX001",
        "product_name": "Fixed product name",
        "software_version": "x.y.z"
    },
    "services": [
        {
            "rid": "c6b028c8-076e-4817-92b1-bcb0cbb78783",
            "rtype": "light"
        },
        {
            "rid": "6cb990a3-7a5c-4c56-9e19-421952405a29",
            "rtype": "light"
        },
        {
            "rid": "f7a7f522-22b2-4978-9cf9-f37ed8b53058",
            "rtype": "motion"
        },
        {
            "rid": "6ec5432f-fb66-4b07-88de-bb0087a0e33d",
            "rtype": "zigbee_connectivity"
        }
    ]
}
```

## Grouping

Grouping is used to give structure to the system, and offer joint control of group members. We distinguish two types of groups: rooms and zones. Rooms group devices based on their physical location, which means each device can only be part of one room, and if a device is in a room then logically all services of that device must be in that same room. Zones group services based on anything that makes sense for the use-case, meaning that services can be part of multiple zones, and any subset of services can be part of a zone.

The resources that are grouped by a room or zone to create the structure are referenced in the `children` array. Much like on a device level, the services used to control group members as a whole, are referenced in the `services` array. For example, the `grouped_light` service can be used to turn on all lights in the group with a multicast command.

Example of a room:

{

"type": "room",

"id": "708d8a89-5d05-408f-b43c-830fbff8316e",

"id\_v1": "/groups/1",

"metadata": {

"archetype": "living\_room",

"name": "Living room"

},

"children": \[

{

"rid": "a91cde76-1d98-400c-873d-12f241f26145",

"rtype": "device"

},

{

"rid": "25f1f7e4-e409-4b64-a1d7-8186916de2d6",

"rtype": "device"

}

\],

"services": \[

{

"rid": "27a6cc29-57e3-4e3e-b83d-f9cc33cc9629",

"rtype": "grouped\_light"

}

\]

}

{ "type": "room", "id": "708d8a89-5d05-408f-b43c-830fbff8316e", "id\_v1": "/groups/1", "metadata": { "archetype": "living\_room", "name": "Living room" }, "children": \[ { "rid": "a91cde76-1d98-400c-873d-12f241f26145", "rtype": "device" }, { "rid": "25f1f7e4-e409-4b64-a1d7-8186916de2d6", "rtype": "device" } \], "services": \[ { "rid": "27a6cc29-57e3-4e3e-b83d-f9cc33cc9629", "rtype": "grouped\_light" } \] }

```
{
    "type": "room",
    "id": "708d8a89-5d05-408f-b43c-830fbff8316e",
    "id_v1": "/groups/1",
    "metadata": {
        "archetype": "living_room",
        "name": "Living room"
    },
    "children": [
        {
            "rid": "a91cde76-1d98-400c-873d-12f241f26145",
            "rtype": "device"
        },
        {
            "rid": "25f1f7e4-e409-4b64-a1d7-8186916de2d6",
            "rtype": "device"
        }
    ],
    "services": [
        {
            "rid": "27a6cc29-57e3-4e3e-b83d-f9cc33cc9629",
            "rtype": "grouped_light"
        }
    ]
}
```

Example of a zone:

{

"type": "zone",

"id": "218bd1d6-fc0f-4288-bcc4-f01675d630ea",

"id\_v1": "/groups/4",

"metadata": {

"archetype": "computer",

"name": "Gaming zone"

},

"children": \[

{

"rid": "2e8297bb-b54c-455f-937e-726ea83df687",

"rtype": "light"

},

{

"rid": "fecdff3b-a5f2-417b-9918-32f1f328d995",

"rtype": "light"

}

\],

"services": \[

{

"rid": "1aa448ee-3299-411e-a2f2-a09a777a5bf1",

"rtype": "grouped\_light"

}

\]

}

{ "type": "zone", "id": "218bd1d6-fc0f-4288-bcc4-f01675d630ea", "id\_v1": "/groups/4", "metadata": { "archetype": "computer", "name": "Gaming zone" }, "children": \[ { "rid": "2e8297bb-b54c-455f-937e-726ea83df687", "rtype": "light" }, { "rid": "fecdff3b-a5f2-417b-9918-32f1f328d995", "rtype": "light" } \], "services": \[ { "rid": "1aa448ee-3299-411e-a2f2-a09a777a5bf1", "rtype": "grouped\_light" } \] }

```
{
    "type": "zone",
    "id": "218bd1d6-fc0f-4288-bcc4-f01675d630ea",
    "id_v1": "/groups/4",
    "metadata": {
        "archetype": "computer",
        "name": "Gaming zone"
    },
    "children": [
        {
            "rid": "2e8297bb-b54c-455f-937e-726ea83df687",
            "rtype": "light"
        },
        {
            "rid": "fecdff3b-a5f2-417b-9918-32f1f328d995",
            "rtype": "light"
        }
    ],
    "services": [
        {
            "rid": "1aa448ee-3299-411e-a2f2-a09a777a5bf1",
            "rtype": "grouped_light"
        }
    ]
}
```

## Scenes

Scenes have remained relatively similar between the V1 and V2 API. The two major differences are in recalling a scene. In the V1 API that used to be done via a group:

`curl --insecure -X PUT 'https://<ipaddress>/api/<appkey>/groups/0/action -H 'Content-Type: application/json' --data-raw '{"scene": "<v1-id>"}'`

Whereas on the V2 API this is done on the scene directly:

`curl --insecure -X PUT 'https://<ipaddress>/clip/v2/resource/scene/<v2-id>' -H 'hue-application-key: <appkey>' -H 'Content-Type: application/json' --data-raw '{"recall": {"action": "active"}}'`

And additionally you can now recall the scene in dynamic mode:

`curl --insecure -X PUT 'https://<ipaddress>/clip/v2/resource/scene/<v2-id>' -H 'hue-application-key: <appkey>' -H 'Content-Type: application/json' --data-raw '{"recall": {"action": "dynamic_palette"}}'`

## Entertainment

To support gradient lights, the V2 API entertainment configuration revolves around channels rather than lights. A light can have multiple channels (like a gradient light strip), and a channel could contain multiple light points. This difference is handled within the Hue System, so streaming to a list of lights with their locations on V1 is conceptually the same as streaming to a list of channels with their locations on V2. One difference is that V1 had a maximum of 10 lights whereas V2 allows for a maximum of 20 channels.

Below shows an example entertainment area for V1.

{

"2": {

"name": "Entertainment area 1",

"type": "Entertainment",

"class": "TV",

"locations": {

"1": \[-0.1, 0.8, -0.8\], //Play Bar (B)

"4": \[-0.1, 0.8, 0.0\] //Gradient Strip (C)

},

"stream": {

"active": false

}

}

}

{ "2": { "name": "Entertainment area 1", "type": "Entertainment", "class": "TV", "locations": { "1": \[-0.1, 0.8, -0.8\], //Play Bar (B) "4": \[-0.1, 0.8, 0.0\] //Gradient Strip (C) }, "stream": { "active": false } } }

```
{
    "2": {
        "name": "Entertainment area 1",
        "type": "Entertainment",
        "class": "TV",
        "locations": {
            "1": [-0.1,  0.8, -0.8], //Play Bar (B)
            "4": [-0.1,  0.8,  0.0]  //Gradient Strip (C)
        },
        "stream": {
            "active": false
        }
    }
}
```

And below is how the same entertainment area would be exposed on V2.

{

"metadata": { "name": "Entertainment area 1" },

"type": "entertainment\_configuration",

"configuration\_type": "screen",

"id": "1a8d99cc-967b-44f2-9202-43f976c0fa6b",

"id\_v1": "/groups/2",

"channels": \[

{

"channel\_id": 0, //Play Bar (B)

"position": { "x": -0.1, "y": 0.8, "z": -0.8 }

},

{

"channel\_id": 1, //Gradient Strip (BL)

"position": { "x": -0.4, "y": 0.8, "z": -0.4 }

},

{

"channel\_id": 2, //Gradient Strip (L)

"position": { "x": -0.4, "y": 0.8, "z": 0.0 }

},

{

"channel\_id": 3, //Gradient Strip (TL)

"position": { "x": -0.4, "y": 0.8, "z": 0.4 }

},

{

"channel\_id": 4, //Gradient Strip (T)

"position": { "x": 0.0, "y": 0.8, "z": 0.4 }

},

{

"channel\_id": 5, //Gradient Strip (TR)

"position": { "x": 0.4, "y": 0.8, "z": 0.4 }

},

{

"channel\_id": 6, //Gradient Strip (R)

"position": { "x": 0.4, "y": 0.8, "z": 0.0 }

},

{

"channel\_id": 7, //Gradient Strip (BR)

"position": { "x": 0.4, "y": 0.8, "z": -0.4 }

}

\],

"status": "inactive"

}

{ "metadata": { "name": "Entertainment area 1" }, "type": "entertainment\_configuration", "configuration\_type": "screen", "id": "1a8d99cc-967b-44f2-9202-43f976c0fa6b", "id\_v1": "/groups/2", "channels": \[ { "channel\_id": 0, //Play Bar (B) "position": { "x": -0.1, "y": 0.8, "z": -0.8 } }, { "channel\_id": 1, //Gradient Strip (BL) "position": { "x": -0.4, "y": 0.8, "z": -0.4 } }, { "channel\_id": 2, //Gradient Strip (L) "position": { "x": -0.4, "y": 0.8, "z": 0.0 } }, { "channel\_id": 3, //Gradient Strip (TL) "position": { "x": -0.4, "y": 0.8, "z": 0.4 } }, { "channel\_id": 4, //Gradient Strip (T) "position": { "x": 0.0, "y": 0.8, "z": 0.4 } }, { "channel\_id": 5, //Gradient Strip (TR) "position": { "x": 0.4, "y": 0.8, "z": 0.4 } }, { "channel\_id": 6, //Gradient Strip (R) "position": { "x": 0.4, "y": 0.8, "z": 0.0 } }, { "channel\_id": 7, //Gradient Strip (BR) "position": { "x": 0.4, "y": 0.8, "z": -0.4 } } \], "status": "inactive" }

```
{
    "metadata": { "name": "Entertainment area 1" },
    "type": "entertainment_configuration",
    "configuration_type": "screen",
    "id": "1a8d99cc-967b-44f2-9202-43f976c0fa6b",
    "id_v1": "/groups/2",
    "channels": [
        {
            "channel_id": 0, //Play Bar (B)
            "position": { "x": -0.1, "y": 0.8, "z": -0.8 }
        },
        {
            "channel_id": 1, //Gradient Strip (BL)
            "position": { "x": -0.4, "y": 0.8, "z": -0.4 }
        },
        {
            "channel_id": 2, //Gradient Strip (L)
            "position": { "x": -0.4, "y": 0.8, "z":  0.0 }
        },
        {
            "channel_id": 3, //Gradient Strip (TL)
            "position": { "x": -0.4, "y": 0.8, "z":  0.4 }
        },
        {
            "channel_id": 4, //Gradient Strip (T)
            "position": { "x":  0.0, "y": 0.8, "z":  0.4 }
        },
        {
            "channel_id": 5, //Gradient Strip (TR)
            "position": { "x":  0.4, "y": 0.8, "z":  0.4 }
        },
        {
            "channel_id": 6, //Gradient Strip (R)
            "position": { "x":  0.4, "y": 0.8, "z":  0.0 }
        },
        {
            "channel_id": 7, //Gradient Strip (BR)
            "position": { "x":  0.4, "y": 0.8, "z": -0.4 }
        }
    ],
    "status": "inactive"
}
```

The examples above include only the fields relevant for an application which uses streaming. The entertainment configuration contains more items relevant for an application that wants to create and update entertainment configurations. These can be found in the API reference, but we instead recommend referring users to the official Philips Hue app for creating and modifying entertainment areas.

The HueStream UDP/DTLS protocol has a respective slight format change from V1 to V2. An example V1 message could look as follows.

{

"HueStream", //protocol

0x01, 0x00, //version 1.0

0x07, //sequence number 7

0x00, 0x00, //reserved

0x00, //color mode RGB

0x00, //reserved

0x00, 0x00, 0x01, //light ID 1

0xff, 0xff, 0x00, 0x00, 0x00, 0x00, //red

0x00, 0x00, 0x04, //light ID 4

0x00, 0x00, 0x00, 0x00, 0xff, 0xff //blue

}

{ "HueStream", //protocol 0x01, 0x00, //version 1.0 0x07, //sequence number 7 0x00, 0x00, //reserved 0x00, //color mode RGB 0x00, //reserved 0x00, 0x00, 0x01, //light ID 1 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, //red 0x00, 0x00, 0x04, //light ID 4 0x00, 0x00, 0x00, 0x00, 0xff, 0xff //blue }

```
{
    "HueStream", //protocol
    0x01, 0x00, //version 1.0
    0x07, //sequence number 7
    0x00, 0x00, //reserved
    0x00, //color mode RGB
    0x00, //reserved
    0x00, 0x00, 0x01, //light ID 1
    0xff, 0xff, 0x00, 0x00, 0x00, 0x00, //red
    0x00, 0x00, 0x04, //light ID 4
    0x00, 0x00, 0x00, 0x00, 0xff, 0xff //blue
}
```

Whereas an example V2 message targets channels instead of lights by a combination of entertainment configuration id and channel id (and has an incremented version number).

{

"HueStream", //protocol

0x02, 0x00, //version 2.0

0x07, //sequence number 7

0x00, 0x00, //reserved

0x00, //color mode RGB

0x00, //reserved

"1a8d99cc-967b-44f2-9202-43f976c0fa6b", //entertainment configuration id

0x00, //channel id 0

0xff, 0xff, 0x00, 0x00, 0x00, 0x00, //red

0x01, //channel id 1

0x00, 0x00, 0xff, 0xff, 0x00, 0x00 //green

0x02, //channel id 2

0x00, 0x00, 0x00, 0x00, 0xff, 0xff //blue

0x03, //channel id 3

0xff, 0xff, 0xff, 0xff, 0xff, 0xff //white

//etc for channel ids 4-7

}

{ "HueStream", //protocol 0x02, 0x00, //version 2.0 0x07, //sequence number 7 0x00, 0x00, //reserved 0x00, //color mode RGB 0x00, //reserved "1a8d99cc-967b-44f2-9202-43f976c0fa6b", //entertainment configuration id 0x00, //channel id 0 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, //red 0x01, //channel id 1 0x00, 0x00, 0xff, 0xff, 0x00, 0x00 //green 0x02, //channel id 2 0x00, 0x00, 0x00, 0x00, 0xff, 0xff //blue 0x03, //channel id 3 0xff, 0xff, 0xff, 0xff, 0xff, 0xff //white //etc for channel ids 4-7 }

```
{
    "HueStream", //protocol
    0x02, 0x00, //version 2.0
    0x07, //sequence number 7
    0x00, 0x00, //reserved
    0x00, //color mode RGB
    0x00, //reserved
    "1a8d99cc-967b-44f2-9202-43f976c0fa6b", //entertainment configuration id
    0x00, //channel id 0
    0xff, 0xff, 0x00, 0x00, 0x00, 0x00, //red
    0x01, //channel id 1
    0x00, 0x00, 0xff, 0xff, 0x00, 0x00 //green
    0x02, //channel id 2
    0x00, 0x00, 0x00, 0x00, 0xff, 0xff //blue
    0x03, //channel id 3
    0xff, 0xff, 0xff, 0xff, 0xff, 0xff //white
    //etc for channel ids 4-7
}
```

The last critical change is that the PSK identity has changed from ‘username’ (i.e. ‘hue-application-key’) to ‘hue-application-id’ (which can be retrieved from the `/auth/v1` endpoint as indicated in the application key section). Everything else like port, cipher suite and client key remain unchanged.

Applications that are using the Hue Entertainment Development Kit (EDK), can just intake the latest version without any changes to the application logic. The EDK supports both the V1 and V2 API including migration internally, without breaking changes on the EDK interface.

We advise prioritizing moving over any entertainment use cases you may have in your application to V2, as V1 cannot target the individual segments of gradient lights, which could result in a confusing user experience.

## Rule engine

The rule engine is currently not available on the V2 API. There is a related concept of ‘behavior scripts’ used for automations in the official Hue app, which on the one hand is capable of executing more advanced behaviors, but on the other hand does currently not support uploading custom behaviors created by 3rd parties. That means you can list instances of existing predefined behaviors on V2, but creating custom behaviors remains on V1 only for now.

## Event Stream

The V2 API supports proactive notifications on changes through Server-Sent Events (SSE) under the `/eventstream` endpoint:

`curl --insecure -N -H 'hue-application-key: <appkey>' -H 'Accept: text/event-stream' https://<ipaddress>/eventstream/clip/v2`

Events have an id, timestamp, type (‘update’, ‘add’, ‘delete’, ‘error’), and data field which contains the changed properties of the resource in the same format as a GET response on the same resource type. The following is an example event stream that would result from turning a light on and off:

`id: 1617322504:0`
`data: [{"creationtime":"2021-04-02T00:15:04Z","data":[{"id":"4413c8fd-6643-48b5-ad02-59453edf8a61","id_v1":"/lights/1","on":{"on":true},"type":"light"}],"id":"19845c30-2e4c-4205-a7b4-8bd496f3407d","type":"update"}]`

`id: 1617322505:0`
`data: [{"creationtime":"2021-04-02T00:15:05Z","data":[{"id":"4413c8fd-6643-48b5-ad02-59453edf8a61","id_v1":"/lights/1","on":{"on":false},"type":"light"}],"id":"bea68344-a36c-4bfd-a658-97830e4e2b1a","type":"update"}]`

On HTTP1.1, you will need a separate connection for the SSE request and regular requests, but we recommend using HTTP2 to multiplex them over a single connection which is more resource efficient.

Currently there is a 1 second rate limit on the amount of event containers the Bridge will send. If the same property has changed twice within that timeframe, you only get the last state. If multiple resources have changed within that timeframe, you will get multiple events grouped in a single container.

## Remote Access

The previous examples were all assuming access on the local network, but if your application (also) uses remote access then there are few extra things to keep in mind. The remote api endpoint `https://api.meethue.com/bridge` points directly to the `/api` endpoint on the bridge, so to access CLIPv2 you need to use the new remote api endpoint `https://api.meethue.com/route` which routes requests to the bridges starting from the root level. This means you can access both CLIPv1 on `/api` and CLIPv2 on `/clip/v2`. Below is an example request to get the CLIPv2 list of light services through the remote API.

`curl -X GET 'https://api.meethue.com/route/clip/v2/resource/light' -H 'hue-application-key: <appkey>' -H 'Authorization: Bearer <access token>'`

To obtain the remote access token, your app must use the latest OAUTH2 endpoints (`/v2/oauth2/authorize` and `/v2/oauth2/token`) as described on [Remote Authentication OAuth2.0](https://developers.meethue.com/develop/hue-api/remote-authentication-oauth/). Please double check that your application is not using the legacy OAUTH2 endpoints (`/oauth2/auth`, `/oauth2/token` and `oauth2/refresh`) anymore, because they have long been deprecated and will be disabled in the near future. Tokens retrieved on the legacy endpoint remain valid to be refreshed on the V2 endpoint.

Event stream is currently not yet released on the remote API.