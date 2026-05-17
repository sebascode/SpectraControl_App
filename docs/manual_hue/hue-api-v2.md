---
title: "Hue CLIP API"
source: "developers.meethue.com"
url: "https://developers.meethue.com/develop/hue-api-v2/api-reference/"
date_saved: "2026-05-09T20:41:16.124Z"
word_count: "42412"
reading_time: "213 min"
---

### /resource

API to retrieve all API resources

### /resource/light

API to manage light services. These are offered by devices with lighting capabilities.

get

List all resources of this type

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of LightGet)*

    **Items**: LightGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **owner**: *required(object)*

        Owner of the service, in case the owner service is deleted, the service also gets deleted

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **type**: *required(light)*
    -   **metadata**: *required(object)*

        additional metadata including a user given name

        -   **name**: *required(string – minLength: 1 – maxLength: 32)*

            Human readable name of a resource

        -   **archetype**: *required(one of unknown\_archetype, classic\_bulb, sultan\_bulb, flood\_bulb, spot\_bulb, candle\_bulb, luster\_bulb, pendant\_round, pendant\_long, ceiling\_round, ceiling\_square, floor\_shade, floor\_lantern, table\_shade, recessed\_ceiling, recessed\_floor, single\_spot, double\_spot, table\_wash, wall\_lantern, wall\_shade, flexible\_lamp, ground\_spot, wall\_spot, plug, hue\_go, hue\_lightstrip, hue\_iris, hue\_bloom, bollard, wall\_washer, hue\_play, vintage\_bulb, vintage\_candle\_bulb, ellipse\_bulb, triangle\_bulb, small\_globe\_bulb, large\_globe\_bulb, edison\_bulb, christmas\_tree, string\_light, hue\_centris, hue\_lightstrip\_tv, hue\_lightstrip\_pc, hue\_tube, hue\_signe, pendant\_spot, ceiling\_horizontal, ceiling\_tube, up\_and\_down, up\_and\_down\_up, up\_and\_down\_down, hue\_floodlight\_camera, twilight, twilight\_front, twilight\_back, hue\_play\_wallwasher, hue\_omniglow, hue\_neon, string\_globe, string\_permanent)*

            Deprecated: use metadata on device level

        -   **fixed\_mired**: *(integer – minimum: 50 – maximum: 1000)*

            A fixed mired value of the white lamp

            **Example**:

            ```
            233
            ```

        -   **function**: *required(one of functional, decorative, mixed, unknown)*

            Function of the lightservice

    -   **product\_data**: *(object)*

        Factory defaults of the product data

        -   **name**: *(string – minLength: 1 – maxLength: 32)*

            Name of the lightservice, only available for devices with multiple lightservices

        -   **archetype**: *(one of unknown\_archetype, classic\_bulb, sultan\_bulb, flood\_bulb, spot\_bulb, candle\_bulb, luster\_bulb, pendant\_round, pendant\_long, ceiling\_round, ceiling\_square, floor\_shade, floor\_lantern, table\_shade, recessed\_ceiling, recessed\_floor, single\_spot, double\_spot, table\_wash, wall\_lantern, wall\_shade, flexible\_lamp, ground\_spot, wall\_spot, plug, hue\_go, hue\_lightstrip, hue\_iris, hue\_bloom, bollard, wall\_washer, hue\_play, vintage\_bulb, vintage\_candle\_bulb, ellipse\_bulb, triangle\_bulb, small\_globe\_bulb, large\_globe\_bulb, edison\_bulb, christmas\_tree, string\_light, hue\_centris, hue\_lightstrip\_tv, hue\_lightstrip\_pc, hue\_tube, hue\_signe, pendant\_spot, ceiling\_horizontal, ceiling\_tube, up\_and\_down, up\_and\_down\_up, up\_and\_down\_down, hue\_floodlight\_camera, twilight, twilight\_front, twilight\_back, hue\_play\_wallwasher, hue\_omniglow, hue\_neon, string\_globe, string\_permanent)*

            Archetype of the lightservice, only available for devices with multiple lightservices

        -   **function**: *required(one of functional, decorative, mixed, unknown)*

            Function of the lightservice

    -   **identify**: *required(object)*
    -   **service\_id**: *required(integer – minimum: 0)*

        Service identification number. 0 indicates service of a single instance

    -   **on**: *required(object)*
        -   **on**: *required(boolean)*

            On/Off state of the light on=true, off=false

    -   **dimming**: *(object)*

        -   **brightness**: *required(number – maximum: 100)*

            Brightness percentage. Value 0 is the lowest possible brightness.

            **Example**:

            ```
            80
            ```

        -   **min\_dim\_level**: *(number – minimum: 0 – maximum: 100)*

            Percentage of the maximum lumen the device outputs on minimum brightness


        **Example**:

        ```json
        {
          "brightness": 80
        }
        ```

    -   **dimming\_delta**: *(object)*

        **Example**:

        ```json
        {
          "action": "up",
          "brightness_delta": 10.5
        }
        ```

    -   **color\_temperature**: *(object)*

        -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

            color temperature in mirek or null when the light color is not in the ct spectrum

            **Example**:

            ```
            233
            ```

        -   **mirek\_valid**: *required(boolean)*

            Indication whether the value presented in mirek is valid

        -   **mirek\_schema**: *required(object)*
            -   **mirek\_minimum**: *required(integer – minimum: 50 – maximum: 1000)*

                minimum color temperature this light supports

                **Example**:

                ```
                233
                ```

            -   **mirek\_maximum**: *required(integer – minimum: 50 – maximum: 1000)*

                maximum color temperature this light supports

                **Example**:

                ```
                233
                ```


        **Example**:

        ```json
        {
          "mirek_schema": {
            "mirek_minimum": 153,
            "mirek_maximum": 500
          },
          "mirek": 202,
          "mirek_valid": true
        }
        ```

    -   **color\_temperature\_delta**: *(object)*

        **Example**:

        ```json
        {
          "action": "up",
          "mirek_delta": 200
        }
        ```

    -   **color**: *(object)*

        -   **xy**: *required(object)*

            CIE XY gamut position

            -   **x**: *required(number – minimum: 0 – maximum: 1)*

                X position in color gamut

            -   **y**: *required(number – minimum: 0 – maximum: 1)*

                Y position in color gamut


            **Example**:

            ```json
            {
              "x": 0.369,
              "y": 0.445
            }
            ```

        -   **gamut**: *(object)*

            Color gamut of color bulb. Some bulbs do not properly return the Gamut information. In this case this is not present.

            -   **red**: *required(object)*

                CIE XY gamut position

                -   **x**: *required(number – minimum: 0 – maximum: 1)*

                    X position in color gamut

                -   **y**: *required(number – minimum: 0 – maximum: 1)*

                    Y position in color gamut


                **Example**:

                ```json
                {
                  "x": 0.369,
                  "y": 0.445
                }
                ```

            -   **green**: *required(object)*

                CIE XY gamut position

                -   **x**: *required(number – minimum: 0 – maximum: 1)*

                    X position in color gamut

                -   **y**: *required(number – minimum: 0 – maximum: 1)*

                    Y position in color gamut


                **Example**:

                ```json
                {
                  "x": 0.369,
                  "y": 0.445
                }
                ```

            -   **blue**: *required(object)*

                CIE XY gamut position

                -   **x**: *required(number – minimum: 0 – maximum: 1)*

                    X position in color gamut

                -   **y**: *required(number – minimum: 0 – maximum: 1)*

                    Y position in color gamut


                **Example**:

                ```json
                {
                  "x": 0.369,
                  "y": 0.445
                }
                ```

        -   **gamut\_type**: *required(one of A, B, C, other)*

            The gammut types supported by hue – A Gamut of early Philips color-only products – B Limited gamut of first Hue color products – C Richer color gamut of Hue white and color ambiance products – other Color gamut of non-hue products with non-hue gamuts resp w/o gamut


        **Example**:

        ```json
        {
          "xy": {
            "x": 0.6915,
            "y": 0.3083
          },
          "gamut_type": "C",
          "gamut": {
            "red": {
              "x": 0.6915,
              "y": 0.3083
            },
            "green": {
              "x": 0.17,
              "y": 0.7
            },
            "blue": {
              "x": 0.1532,
              "y": 0.0475
            }
          }
        }
        ```

    -   **dynamics**: *(object)*
        -   **status**: *required(one of dynamic\_palette, none)*

            Current status of the lamp with dynamics.

        -   **status\_values**: *required(array of SupportedDynamicStatus)*

            Statuses in which a lamp could be when playing dynamics.

        -   **speed**: *required(number – minimum: 0 – maximum: 1)*

            speed of dynamic palette. The speed is only valid if the status is dynamic\_palette.

        -   **speed\_valid**: *required(boolean)*

            Indicates whether the value presented in speed is valid

    -   **alert**: *(object)*
        -   **action\_values**: *required(array of AlertEffectType)*

            Alert effects that the light supports.

    -   **signaling**: *(object)*

        Feature containing signaling properties.

        -   **signal\_values**: *required(array of SupportedSignals)*

            Signals that the light supports.

        -   **status**: *(object)*

            Indicates status of active signal. Not available when inactive.

            -   **signal**: *required(one of no\_signal, on\_off, on\_off\_color, alternating)*

                Indicates which signal is currently active.

            -   **estimated\_end**: *required(datetime)*

                Timestamp indicating when the active signal is expected to end. Value is not set if there is no\_signal

            -   **colors**: *required(array of ColorFeatureBasicGet)*

                Colors that were provided for the active effect.

                **Items**: ColorFeatureBasicGet

                -   **xy**: *required(object)*

                    CIE XY gamut position

                    -   **x**: *required(number – minimum: 0 – maximum: 1)*

                        X position in color gamut

                    -   **y**: *required(number – minimum: 0 – maximum: 1)*

                        Y position in color gamut


                    **Example**:

                    ```json
                    {
                      "x": 0.369,
                      "y": 0.445
                    }
                    ```


                **Example**:

                ```json
                {
                  "xy": {
                    "x": 0.6915,
                    "y": 0.3083
                  }
                }
                ```

    -   **mode**: *required(one of normal, streaming)*

        Mode the light is currently in

    -   **gradient**: *(object)*

        Basic feature containing gradient properties.

        -   **points**: *required(array of GradientPointGet – maxItems: 5)*

            Collection of gradients points. For control of the gradient points through a PUT a minimum of 2 points need to be provided.

            **Items**: GradientPointGet

            -   **color**: *required(object)*

                -   **xy**: *required(object)*

                    CIE XY gamut position

                    -   **x**: *required(number – minimum: 0 – maximum: 1)*

                        X position in color gamut

                    -   **y**: *required(number – minimum: 0 – maximum: 1)*

                        Y position in color gamut


                    **Example**:

                    ```json
                    {
                      "x": 0.369,
                      "y": 0.445
                    }
                    ```


                **Example**:

                ```json
                {
                  "xy": {
                    "x": 0.6915,
                    "y": 0.3083
                  }
                }
                ```


        -   **mode**: *required(one of interpolated\_palette, interpolated\_palette\_mirrored, random\_pixelated, segmented\_palette)*

            Mode in which the points are currently being deployed. If not provided during PUT/POST it will be defaulted to interpolated\_palette

        -   **points\_capable**: *required(integer)*

            Number of color points that gradient lamp is capable of showing with gradience.

        -   **mode\_values**: *required(array of SupportedGradientModes)*

            Modes a gradient device can deploy the gradient palette of colors

        -   **pixel\_count**: *(integer)*

            Number of pixels in the device

    -   **effects**: *(object)*

        Deprecated: use effects\_v2

        -   **status\_values**: *required(array of SupportedEffects)*

            Possible status values in which a light could be when playing an effect.

        -   **status**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*

            Current status values the light is in regarding effects

        -   **effect\_values**: *required(array of SupportedEffects)*

            Possible effect values you can set in a light.

    -   **effects\_v2**: *(object)*
        -   **action**: *required(object)*
            -   **effect\_values**: *required(array of SupportedEffects)*

                Possible effect values you can set in a light.

        -   **status**: *required(object)*
            -   **effect**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*
            -   **effect\_values**: *required(array of SupportedEffects)*
            -   **parameters**: *(object)*
                -   **color**: *(object)*

                    -   **xy**: *required(object)*

                        CIE XY gamut position

                        -   **x**: *required(number – minimum: 0 – maximum: 1)*

                            X position in color gamut

                        -   **y**: *required(number – minimum: 0 – maximum: 1)*

                            Y position in color gamut


                        **Example**:

                        ```json
                        {
                          "x": 0.369,
                          "y": 0.445
                        }
                        ```


                    **Example**:

                    ```json
                    {
                      "xy": {
                        "x": 0.6915,
                        "y": 0.3083
                      }
                    }
                    ```

                -   **color\_temperature**: *(object)*

                    -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

                        color temperature in mirek or null when the light color is not in the ct spectrum

                        **Example**:

                        ```json
                        233
                        ```

                    -   **mirek\_valid**: *required(boolean)*

                        Indication whether the value presented in mirek is valid


                    **Example**:

                    ```json
                    {
                      "mirek": 202,
                      "mirek_valid": true
                    }
                    ```

                -   **speed**: *required(number – minimum: 0 – maximum: 1)*
    -   **timed\_effects**: *(object)*
        -   **status\_values**: *required(array of SupportedTimedEffects)*

            Possible status values in which a light could be when playing a timed effect.

        -   **status**: *required(one of sunrise, sunset, no\_effect)*

            Current status values the light is in regarding timed effects

        -   **effect\_values**: *required(array of SupportedTimedEffects)*

            Possible timed effect values you can set in a light.

    -   **powerup**: *(object)*

        Feature containing properties to configure powerup behaviour of a lightsource.

        -   **preset**: *required(one of safety, powerfail, last\_on\_state, custom)*

            When setting the custom preset the additional properties can be set. For all other presets, no other properties can be included.

        -   **configured**: *required(boolean)*

            Indicates if the shown values have been configured in the lightsource.

        -   **on**: *required(object)*
            -   **mode**: *required(one of on, toggle, previous)*

                State to activate after powerup. On will use the value specified in the “on” property. When setting mode “on”, the on property must be included. Toggle will alternate between on and off on each subsequent power toggle. Previous will return to the state it was in before powering off.

            -   **on**: *(object)*
                -   **on**: *required(boolean)*

                    On/Off state of the light on=true, off=false

        -   **dimming**: *(object)*
            -   **mode**: *required(one of dimming, previous)*

                Dimming will set the brightness to the specified value after power up. When setting mode “dimming”, the dimming property must be included. Previous will set brightness to the state it was in before powering off.

            -   **dimming**: *(object)*

                -   **brightness**: *required(number – maximum: 100)*

                    Brightness percentage. Value 0 is the lowest possible brightness.

                    **Example**:

                    ```
                    80
                    ```


                **Example**:

                ```json
                {
                  "brightness": 80
                }
                ```

        -   **color**: *(object)*
            -   **mode**: *required(one of color\_temperature, color, previous)*

                State to activate after powerup. Availability of “color\_temperature” and “color” modes depend on the capabilities of the lamp. Colortemperature will set the colortemperature to the specified value after power up. When setting color\_temperature, the color\_temperature property must be included Color will set the color tot he specified value after power up. When setting color mode, the color property must be included Previous will set color to the state it was in before powering off.

            -   **color\_temperature**: *(object)*

                -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

                    color temperature in mirek or null when the light color is not in the ct spectrum

                    **Example**:

                    ```
                    233
                    ```


                **Example**:

                ```json
                {
                  "mirek": 202
                }
                ```

            -   **color**: *(object)*

                -   **xy**: *required(object)*

                    CIE XY gamut position

                    -   **x**: *required(number – minimum: 0 – maximum: 1)*

                        X position in color gamut

                    -   **y**: *required(number – minimum: 0 – maximum: 1)*

                        Y position in color gamut


                    **Example**:

                    ```json
                    {
                      "x": 0.369,
                      "y": 0.445
                    }
                    ```


                **Example**:

                ```json
                {
                  "xy": {
                    "x": 0.6915,
                    "y": 0.3083
                  }
                }
                ```

    -   **content\_configuration**: *(object)*

        Configuration parameters that affect how content is deployed on a product

        -   **orientation**: *(object)*

            Orientation for content deployed on pixelated products when applicable (e.g. sunrise, candle)

            -   **status**: *required(one of set, changing)*
            -   **configurable**: *required(boolean)*

                Indicates if the product allows modifying the configuration attribute

            -   **orientation**: *required(one of horizontal, vertical)*
        -   **order**: *(object)*

            Order in which the content is represented on the pixels

            -   **status**: *required(one of set, changing)*
            -   **configurable**: *required(boolean)*

                Indicates if the product allows modifying the configuration attribute

            -   **order**: *required(one of forward, reversed)*
    -   **geometry**: *(object)*

        Feature describing the geometry a light service.

        -   **pixel\_positions**: *(array of PixelPosition – maxItems: 5)*

            **Items**: PixelPosition

            -   **index**: *required(integer – minimum: 0)*

                Index of the pixel to which the position belongs, use 0 for non-gradient lights

            -   **position**: *required(object)*

                Position of the pixel in 3D space. Range of coordinates is -100 to 100 meters.

                -   **x**: *required(number)*
                -   **y**: *required(number)*
                -   **z**: *required(number)*



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 406

Query parameter has invalid value.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


API calls for resource which cannot be deleted (PUT and GET on resource instance)

put

get

### URI Parameters

-   **id**: *required(string)*

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **type**: *(light)*
-   **metadata**: *(object)*

    additional metadata including a user given name

    -   **name**: *(string – minLength: 1 – maxLength: 32)*

        Human readable name of a resource

    -   **archetype**: *(one of unknown\_archetype, classic\_bulb, sultan\_bulb, flood\_bulb, spot\_bulb, candle\_bulb, luster\_bulb, pendant\_round, pendant\_long, ceiling\_round, ceiling\_square, floor\_shade, floor\_lantern, table\_shade, recessed\_ceiling, recessed\_floor, single\_spot, double\_spot, table\_wash, wall\_lantern, wall\_shade, flexible\_lamp, ground\_spot, wall\_spot, plug, hue\_go, hue\_lightstrip, hue\_iris, hue\_bloom, bollard, wall\_washer, hue\_play, vintage\_bulb, vintage\_candle\_bulb, ellipse\_bulb, triangle\_bulb, small\_globe\_bulb, large\_globe\_bulb, edison\_bulb, christmas\_tree, string\_light, hue\_centris, hue\_lightstrip\_tv, hue\_lightstrip\_pc, hue\_tube, hue\_signe, pendant\_spot, ceiling\_horizontal, ceiling\_tube, up\_and\_down, up\_and\_down\_up, up\_and\_down\_down, hue\_floodlight\_camera, twilight, twilight\_front, twilight\_back, hue\_play\_wallwasher, hue\_omniglow, hue\_neon, string\_globe, string\_permanent)*

        Deprecated: use metadata on device level

    -   **function**: *(one of functional, decorative, mixed, unknown)*

        Function of the lightservice

-   **identify**: *(object)*
    -   **action**: *required(identify)*
        -   identify: Triggers a visual identification sequence, current implemented as (which can change in the future): Bridge performs Zigbee LED identification cycles for 5 seconds Lights perform one breathe cycle Sensors perform LED identification cycles for 15 seconds
    -   **duration**: *(integer)*

        Duration in milliseconds to perform the identify cycle.

        **Example**:

        ```
        800
        ```

-   **on**: *(object)*
    -   **on**: *(boolean)*

        On/Off state of the light on=true, off=false

-   **dimming**: *(object)*

    -   **brightness**: *(number – maximum: 100)*

        Brightness percentage. Value 0 is the lowest possible brightness.

        **Example**:

        ```
        80
        ```


    **Example**:

    ```json
    {
      "brightness": 80
    }
    ```

-   **dimming\_delta**: *(object)*

    -   **action**: *required(one of up, down, stop)*

        The delta action to apply

    -   **brightness\_delta**: *(number – maximum: 100)*

        Brightness percentage of full-scale increase delta to current dimlevel. Clip at Max-level or Min-level.

        **Example**:

        ```
        20
        ```


    **Example**:

    ```json
    {
      "action": "up",
      "brightness_delta": 10.5
    }
    ```

-   **color\_temperature**: *(object)*

    -   **mirek**: *(integer – minimum: 50 – maximum: 1000)*

        color temperature in mirek or null when the light color is not in the ct spectrum

        **Example**:

        ```
        233
        ```


    **Example**:

    ```json
    {
      "mirek_schema": {
        "mirek_minimum": 153,
        "mirek_maximum": 500
      },
      "mirek": 202,
      "mirek_valid": true
    }
    ```

-   **color\_temperature\_delta**: *(object)*

    -   **action**: *required(one of up, down, stop)*

        The delta action to apply

    -   **mirek\_delta**: *(integer – maximum: 950)*

        Mirek delta to current mirek. Clip at mirek\_minimum and mirek\_maximum of mirek\_schema.

        **Example**:

        ```
        10
        ```


    **Example**:

    ```json
    {
      "action": "up",
      "mirek_delta": 200
    }
    ```

-   **color**: *(object)*

    -   **xy**: *(object)*

        CIE XY gamut position

        -   **x**: *required(number – minimum: 0 – maximum: 1)*

            X position in color gamut

        -   **y**: *required(number – minimum: 0 – maximum: 1)*

            Y position in color gamut


        **Example**:

        ```json
        {
          "x": 0.369,
          "y": 0.445
        }
        ```


    **Example**:

    ```json
    {
      "xy": {
        "x": 0.6915,
        "y": 0.3083
      },
      "gamut_type": "C",
      "gamut": {
        "red": {
          "x": 0.6915,
          "y": 0.3083
        },
        "green": {
          "x": 0.17,
          "y": 0.7
        },
        "blue": {
          "x": 0.1532,
          "y": 0.0475
        }
      }
    }
    ```

-   **dynamics**: *(object)*
    -   **duration**: *(integer)*

        Duration of a light transition in ms.

        **Example**:

        ```
        800
        ```

    -   **speed**: *(number – minimum: 0 – maximum: 1)*

        speed of dynamic palette. The speed is only valid if the status is dynamic\_palette.

-   **alert**: *(object)*
    -   **action**: *required(breathe)*

        Alert to set the light to

-   **signaling**: *(object)*

    Feature containing signaling properties.

    -   **signal**: *required(one of no\_signal, on\_off, on\_off\_color, alternating)*

        Signal to set the light to

    -   **duration**: *required(integer)*

        Duration in milliseconds. Maximum value is 65534000 ms and a stepsize of 1 second. Values inbetween steps will be rounded. Duration is ignored for no\_signal.

        **Example**:

        ```
        800
        ```

    -   **colors**: *(array of ColorFeatureBasicPut – minItems: 1 – maxItems: 2)*

        List of colors to apply to the signal (not supported by all signals)

        **Items**: ColorFeatureBasicPut

        -   **xy**: *(object)*

            CIE XY gamut position

            -   **x**: *required(number – minimum: 0 – maximum: 1)*

                X position in color gamut

            -   **y**: *required(number – minimum: 0 – maximum: 1)*

                Y position in color gamut


            **Example**:

            ```json
            {
              "x": 0.369,
              "y": 0.445
            }
            ```


        **Example**:

        ```json
        {
          "xy": {
            "x": 0.6915,
            "y": 0.3083
          }
        }
        ```

-   **gradient**: *(object)*

    Basic feature containing gradient properties.

    -   **points**: *required(array of GradientPointPut – maxItems: 5)*

        Collection of gradients points. For control of the gradient points through a PUT a minimum of 2 points need to be provided.

        **Items**: GradientPointPut

        -   **color**: *required(object)*

            -   **xy**: *(object)*

                CIE XY gamut position

                -   **x**: *required(number – minimum: 0 – maximum: 1)*

                    X position in color gamut

                -   **y**: *required(number – minimum: 0 – maximum: 1)*

                    Y position in color gamut


                **Example**:

                ```json
                {
                  "x": 0.369,
                  "y": 0.445
                }
                ```


            **Example**:

            ```json
            {
              "xy": {
                "x": 0.6915,
                "y": 0.3083
              }
            }
            ```


    -   **mode**: *(one of interpolated\_palette, interpolated\_palette\_mirrored, random\_pixelated, segmented\_palette)*

        Mode in which the points are currently being deployed. If not provided during PUT/POST it will be defaulted to interpolated\_palette

-   **effects**: *(object)*

    Deprecated: use effects\_v2

    -   **effect**: *(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*

        Effect to set the light to.

-   **effects\_v2**: *(object)*
    -   **action**: *(object)*
        -   **effect**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*

            Effect to set the light to.

        -   **parameters**: *(object)*
            -   **color**: *(object)*

                -   **xy**: *(object)*

                    CIE XY gamut position

                    -   **x**: *required(number – minimum: 0 – maximum: 1)*

                        X position in color gamut

                    -   **y**: *required(number – minimum: 0 – maximum: 1)*

                        Y position in color gamut


                    **Example**:

                    ```json
                    {
                      "x": 0.369,
                      "y": 0.445
                    }
                    ```


                **Example**:

                ```json
                {
                  "xy": {
                    "x": 0.6915,
                    "y": 0.3083
                  }
                }
                ```

            -   **color\_temperature**: *(object)*

                -   **mirek**: *(integer – minimum: 50 – maximum: 1000)*

                    color temperature in mirek or null when the light color is not in the ct spectrum

                    **Example**:

                    ```json
                    233
                    ```


                **Example**:

                ```json
                {
                  "mirek": 202,
                  "mirek_valid": true
                }
                ```

            -   **speed**: *(number – minimum: 0 – maximum: 1)*
-   **timed\_effects**: *(object)*
    -   **effect**: *(one of sunrise, sunset, no\_effect)*

        Timed effect to set to the light.

    -   **duration**: *(integer)*

        Duration in milliseconds. Duration is mandatory when timed effect is set except for no\_effect. Resolution decreases for a larger duration. e.g Effects with duration smaller than a minute will be rounded to a resolution of 1s, while effects with duration larger than an hour will be arounded up to a resolution of 300s. Duration has a max of 21600000 ms.

        **Example**:

        ```
        800
        ```

-   **powerup**: *(object)*

    Feature containing properties to configure powerup behaviour of a lightsource.

    -   **preset**: *required(one of safety, powerfail, last\_on\_state, custom)*

        When setting the custom preset the additional properties can be set. For all other presets, no other properties can be included.

    -   **on**: *(object)*
        -   **mode**: *required(one of on, toggle, previous)*

            State to activate after powerup. On will use the value specified in the “on” property. When setting mode “on”, the on property must be included. Toggle will alternate between on and off on each subsequent power toggle. Previous will return to the state it was in before powering off.

        -   **on**: *(object)*
            -   **on**: *(boolean)*

                On/Off state of the light on=true, off=false

    -   **dimming**: *(object)*
        -   **mode**: *required(one of dimming, previous)*

            Dimming will set the brightness to the specified value after power up. When setting mode “dimming”, the dimming property must be included. Previous will set brightness to the state it was in before powering off.

        -   **dimming**: *(object)*

            -   **brightness**: *(number – maximum: 100)*

                Brightness percentage. Value 0 is the lowest possible brightness.

                **Example**:

                ```
                80
                ```


            **Example**:

            ```json
            {
              "brightness": 80
            }
            ```

    -   **color**: *(object)*
        -   **mode**: *required(one of color\_temperature, color, previous)*

            State to activate after powerup. Availability of “color\_temperature” and “color” modes depend on the capabilities of the lamp. Colortemperature will set the colortemperature to the specified value after power up. When setting color\_temperature, the color\_temperature property must be included Color will set the color tot he specified value after power up. When setting color mode, the color property must be included Previous will set color to the state it was in before powering off.

        -   **color\_temperature**: *(object)*

            -   **mirek**: *(integer – minimum: 50 – maximum: 1000)*

                color temperature in mirek or null when the light color is not in the ct spectrum

                **Example**:

                ```
                233
                ```


            **Example**:

            ```json
            {
              "mirek": 202
            }
            ```

        -   **color**: *(object)*

            -   **xy**: *(object)*

                CIE XY gamut position

                -   **x**: *required(number – minimum: 0 – maximum: 1)*

                    X position in color gamut

                -   **y**: *required(number – minimum: 0 – maximum: 1)*

                    Y position in color gamut


                **Example**:

                ```json
                {
                  "x": 0.369,
                  "y": 0.445
                }
                ```


            **Example**:

            ```json
            {
              "xy": {
                "x": 0.6915,
                "y": 0.3083
              }
            }
            ```

-   **content\_configuration**: *(object)*

    Configuration parameters that affect how content is deployed on a product

    -   **orientation**: *(object)*

        Orientation for content deployed on pixelated products when applicable (e.g. sunrise, candle)

        -   **orientation**: *(one of horizontal, vertical)*
    -   **order**: *(object)*

        Order in which the content is represented on the pixels

        -   **order**: *(one of forward, reversed)*
-   **geometry**: *(object)*

    Feature describing the geometry a light service.

    -   **pixel\_positions**: *(array of PixelPosition – maxItems: 5)*

        **Items**: PixelPosition

        -   **index**: *required(integer – minimum: 0)*

            Index of the pixel to which the position belongs, use 0 for non-gradient lights

        -   **position**: *required(object)*

            Position of the pixel in 3D space. Range of coordinates is -100 to 100 meters.

            -   **x**: *required(number)*
            -   **y**: *required(number)*
            -   **z**: *required(number)*


## HTTP status code 200

Request was on resource path was successful.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 207

Request was partially successful, resource has been updated. Failed requests and errors are listed in error object.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 400

Request was on resource path was unsuccessful, body of request contains errors.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of LightGet)*

    **Items**: LightGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **owner**: *required(object)*

        Owner of the service, in case the owner service is deleted, the service also gets deleted

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **type**: *required(light)*
    -   **metadata**: *required(object)*

        additional metadata including a user given name

        -   **name**: *required(string – minLength: 1 – maxLength: 32)*

            Human readable name of a resource

        -   **archetype**: *required(one of unknown\_archetype, classic\_bulb, sultan\_bulb, flood\_bulb, spot\_bulb, candle\_bulb, luster\_bulb, pendant\_round, pendant\_long, ceiling\_round, ceiling\_square, floor\_shade, floor\_lantern, table\_shade, recessed\_ceiling, recessed\_floor, single\_spot, double\_spot, table\_wash, wall\_lantern, wall\_shade, flexible\_lamp, ground\_spot, wall\_spot, plug, hue\_go, hue\_lightstrip, hue\_iris, hue\_bloom, bollard, wall\_washer, hue\_play, vintage\_bulb, vintage\_candle\_bulb, ellipse\_bulb, triangle\_bulb, small\_globe\_bulb, large\_globe\_bulb, edison\_bulb, christmas\_tree, string\_light, hue\_centris, hue\_lightstrip\_tv, hue\_lightstrip\_pc, hue\_tube, hue\_signe, pendant\_spot, ceiling\_horizontal, ceiling\_tube, up\_and\_down, up\_and\_down\_up, up\_and\_down\_down, hue\_floodlight\_camera, twilight, twilight\_front, twilight\_back, hue\_play\_wallwasher, hue\_omniglow, hue\_neon, string\_globe, string\_permanent)*

            Deprecated: use metadata on device level

        -   **fixed\_mired**: *(integer – minimum: 50 – maximum: 1000)*

            A fixed mired value of the white lamp

            **Example**:

            ```
            233
            ```

        -   **function**: *required(one of functional, decorative, mixed, unknown)*

            Function of the lightservice

    -   **product\_data**: *(object)*

        Factory defaults of the product data

        -   **name**: *(string – minLength: 1 – maxLength: 32)*

            Name of the lightservice, only available for devices with multiple lightservices

        -   **archetype**: *(one of unknown\_archetype, classic\_bulb, sultan\_bulb, flood\_bulb, spot\_bulb, candle\_bulb, luster\_bulb, pendant\_round, pendant\_long, ceiling\_round, ceiling\_square, floor\_shade, floor\_lantern, table\_shade, recessed\_ceiling, recessed\_floor, single\_spot, double\_spot, table\_wash, wall\_lantern, wall\_shade, flexible\_lamp, ground\_spot, wall\_spot, plug, hue\_go, hue\_lightstrip, hue\_iris, hue\_bloom, bollard, wall\_washer, hue\_play, vintage\_bulb, vintage\_candle\_bulb, ellipse\_bulb, triangle\_bulb, small\_globe\_bulb, large\_globe\_bulb, edison\_bulb, christmas\_tree, string\_light, hue\_centris, hue\_lightstrip\_tv, hue\_lightstrip\_pc, hue\_tube, hue\_signe, pendant\_spot, ceiling\_horizontal, ceiling\_tube, up\_and\_down, up\_and\_down\_up, up\_and\_down\_down, hue\_floodlight\_camera, twilight, twilight\_front, twilight\_back, hue\_play\_wallwasher, hue\_omniglow, hue\_neon, string\_globe, string\_permanent)*

            Archetype of the lightservice, only available for devices with multiple lightservices

        -   **function**: *required(one of functional, decorative, mixed, unknown)*

            Function of the lightservice

    -   **identify**: *required(object)*
    -   **service\_id**: *required(integer – minimum: 0)*

        Service identification number. 0 indicates service of a single instance

    -   **on**: *required(object)*
        -   **on**: *required(boolean)*

            On/Off state of the light on=true, off=false

    -   **dimming**: *(object)*

        -   **brightness**: *required(number – maximum: 100)*

            Brightness percentage. Value 0 is the lowest possible brightness.

            **Example**:

            ```
            80
            ```

        -   **min\_dim\_level**: *(number – minimum: 0 – maximum: 100)*

            Percentage of the maximum lumen the device outputs on minimum brightness


        **Example**:

        ```json
        {
          "brightness": 80
        }
        ```

    -   **dimming\_delta**: *(object)*

        **Example**:

        ```json
        {
          "action": "up",
          "brightness_delta": 10.5
        }
        ```

    -   **color\_temperature**: *(object)*

        -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

            color temperature in mirek or null when the light color is not in the ct spectrum

            **Example**:

            ```
            233
            ```

        -   **mirek\_valid**: *required(boolean)*

            Indication whether the value presented in mirek is valid

        -   **mirek\_schema**: *required(object)*
            -   **mirek\_minimum**: *required(integer – minimum: 50 – maximum: 1000)*

                minimum color temperature this light supports

                **Example**:

                ```
                233
                ```

            -   **mirek\_maximum**: *required(integer – minimum: 50 – maximum: 1000)*

                maximum color temperature this light supports

                **Example**:

                ```
                233
                ```


        **Example**:

        ```json
        {
          "mirek_schema": {
            "mirek_minimum": 153,
            "mirek_maximum": 500
          },
          "mirek": 202,
          "mirek_valid": true
        }
        ```

    -   **color\_temperature\_delta**: *(object)*

        **Example**:

        ```json
        {
          "action": "up",
          "mirek_delta": 200
        }
        ```

    -   **color**: *(object)*

        -   **xy**: *required(object)*

            CIE XY gamut position

            -   **x**: *required(number – minimum: 0 – maximum: 1)*

                X position in color gamut

            -   **y**: *required(number – minimum: 0 – maximum: 1)*

                Y position in color gamut


            **Example**:

            ```json
            {
              "x": 0.369,
              "y": 0.445
            }
            ```

        -   **gamut**: *(object)*

            Color gamut of color bulb. Some bulbs do not properly return the Gamut information. In this case this is not present.

            -   **red**: *required(object)*

                CIE XY gamut position

                -   **x**: *required(number – minimum: 0 – maximum: 1)*

                    X position in color gamut

                -   **y**: *required(number – minimum: 0 – maximum: 1)*

                    Y position in color gamut


                **Example**:

                ```json
                {
                  "x": 0.369,
                  "y": 0.445
                }
                ```

            -   **green**: *required(object)*

                CIE XY gamut position

                -   **x**: *required(number – minimum: 0 – maximum: 1)*

                    X position in color gamut

                -   **y**: *required(number – minimum: 0 – maximum: 1)*

                    Y position in color gamut


                **Example**:

                ```json
                {
                  "x": 0.369,
                  "y": 0.445
                }
                ```

            -   **blue**: *required(object)*

                CIE XY gamut position

                -   **x**: *required(number – minimum: 0 – maximum: 1)*

                    X position in color gamut

                -   **y**: *required(number – minimum: 0 – maximum: 1)*

                    Y position in color gamut


                **Example**:

                ```json
                {
                  "x": 0.369,
                  "y": 0.445
                }
                ```

        -   **gamut\_type**: *required(one of A, B, C, other)*

            The gammut types supported by hue – A Gamut of early Philips color-only products – B Limited gamut of first Hue color products – C Richer color gamut of Hue white and color ambiance products – other Color gamut of non-hue products with non-hue gamuts resp w/o gamut


        **Example**:

        ```json
        {
          "xy": {
            "x": 0.6915,
            "y": 0.3083
          },
          "gamut_type": "C",
          "gamut": {
            "red": {
              "x": 0.6915,
              "y": 0.3083
            },
            "green": {
              "x": 0.17,
              "y": 0.7
            },
            "blue": {
              "x": 0.1532,
              "y": 0.0475
            }
          }
        }
        ```

    -   **dynamics**: *(object)*
        -   **status**: *required(one of dynamic\_palette, none)*

            Current status of the lamp with dynamics.

        -   **status\_values**: *required(array of SupportedDynamicStatus)*

            Statuses in which a lamp could be when playing dynamics.

        -   **speed**: *required(number – minimum: 0 – maximum: 1)*

            speed of dynamic palette. The speed is only valid if the status is dynamic\_palette.

        -   **speed\_valid**: *required(boolean)*

            Indicates whether the value presented in speed is valid

    -   **alert**: *(object)*
        -   **action\_values**: *required(array of AlertEffectType)*

            Alert effects that the light supports.

    -   **signaling**: *(object)*

        Feature containing signaling properties.

        -   **signal\_values**: *required(array of SupportedSignals)*

            Signals that the light supports.

        -   **status**: *(object)*

            Indicates status of active signal. Not available when inactive.

            -   **signal**: *required(one of no\_signal, on\_off, on\_off\_color, alternating)*

                Indicates which signal is currently active.

            -   **estimated\_end**: *required(datetime)*

                Timestamp indicating when the active signal is expected to end. Value is not set if there is no\_signal

            -   **colors**: *required(array of ColorFeatureBasicGet)*

                Colors that were provided for the active effect.

                **Items**: ColorFeatureBasicGet

                -   **xy**: *required(object)*

                    CIE XY gamut position

                    -   **x**: *required(number – minimum: 0 – maximum: 1)*

                        X position in color gamut

                    -   **y**: *required(number – minimum: 0 – maximum: 1)*

                        Y position in color gamut


                    **Example**:

                    ```json
                    {
                      "x": 0.369,
                      "y": 0.445
                    }
                    ```


                **Example**:

                ```json
                {
                  "xy": {
                    "x": 0.6915,
                    "y": 0.3083
                  }
                }
                ```

    -   **mode**: *required(one of normal, streaming)*

        Mode the light is currently in

    -   **gradient**: *(object)*

        Basic feature containing gradient properties.

        -   **points**: *required(array of GradientPointGet – maxItems: 5)*

            Collection of gradients points. For control of the gradient points through a PUT a minimum of 2 points need to be provided.

            **Items**: GradientPointGet

            -   **color**: *required(object)*

                -   **xy**: *required(object)*

                    CIE XY gamut position

                    -   **x**: *required(number – minimum: 0 – maximum: 1)*

                        X position in color gamut

                    -   **y**: *required(number – minimum: 0 – maximum: 1)*

                        Y position in color gamut


                    **Example**:

                    ```json
                    {
                      "x": 0.369,
                      "y": 0.445
                    }
                    ```


                **Example**:

                ```json
                {
                  "xy": {
                    "x": 0.6915,
                    "y": 0.3083
                  }
                }
                ```


        -   **mode**: *required(one of interpolated\_palette, interpolated\_palette\_mirrored, random\_pixelated, segmented\_palette)*

            Mode in which the points are currently being deployed. If not provided during PUT/POST it will be defaulted to interpolated\_palette

        -   **points\_capable**: *required(integer)*

            Number of color points that gradient lamp is capable of showing with gradience.

        -   **mode\_values**: *required(array of SupportedGradientModes)*

            Modes a gradient device can deploy the gradient palette of colors

        -   **pixel\_count**: *(integer)*

            Number of pixels in the device

    -   **effects**: *(object)*

        Deprecated: use effects\_v2

        -   **status\_values**: *required(array of SupportedEffects)*

            Possible status values in which a light could be when playing an effect.

        -   **status**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*

            Current status values the light is in regarding effects

        -   **effect\_values**: *required(array of SupportedEffects)*

            Possible effect values you can set in a light.

    -   **effects\_v2**: *(object)*
        -   **action**: *required(object)*
            -   **effect\_values**: *required(array of SupportedEffects)*

                Possible effect values you can set in a light.

        -   **status**: *required(object)*
            -   **effect**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*
            -   **effect\_values**: *required(array of SupportedEffects)*
            -   **parameters**: *(object)*
                -   **color**: *(object)*

                    -   **xy**: *required(object)*

                        CIE XY gamut position

                        -   **x**: *required(number – minimum: 0 – maximum: 1)*

                            X position in color gamut

                        -   **y**: *required(number – minimum: 0 – maximum: 1)*

                            Y position in color gamut


                        **Example**:

                        ```json
                        {
                          "x": 0.369,
                          "y": 0.445
                        }
                        ```


                    **Example**:

                    ```json
                    {
                      "xy": {
                        "x": 0.6915,
                        "y": 0.3083
                      }
                    }
                    ```

                -   **color\_temperature**: *(object)*

                    -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

                        color temperature in mirek or null when the light color is not in the ct spectrum

                        **Example**:

                        ```json
                        233
                        ```

                    -   **mirek\_valid**: *required(boolean)*

                        Indication whether the value presented in mirek is valid


                    **Example**:

                    ```json
                    {
                      "mirek": 202,
                      "mirek_valid": true
                    }
                    ```

                -   **speed**: *required(number – minimum: 0 – maximum: 1)*
    -   **timed\_effects**: *(object)*
        -   **status\_values**: *required(array of SupportedTimedEffects)*

            Possible status values in which a light could be when playing a timed effect.

        -   **status**: *required(one of sunrise, sunset, no\_effect)*

            Current status values the light is in regarding timed effects

        -   **effect\_values**: *required(array of SupportedTimedEffects)*

            Possible timed effect values you can set in a light.

    -   **powerup**: *(object)*

        Feature containing properties to configure powerup behaviour of a lightsource.

        -   **preset**: *required(one of safety, powerfail, last\_on\_state, custom)*

            When setting the custom preset the additional properties can be set. For all other presets, no other properties can be included.

        -   **configured**: *required(boolean)*

            Indicates if the shown values have been configured in the lightsource.

        -   **on**: *required(object)*
            -   **mode**: *required(one of on, toggle, previous)*

                State to activate after powerup. On will use the value specified in the “on” property. When setting mode “on”, the on property must be included. Toggle will alternate between on and off on each subsequent power toggle. Previous will return to the state it was in before powering off.

            -   **on**: *(object)*
                -   **on**: *required(boolean)*

                    On/Off state of the light on=true, off=false

        -   **dimming**: *(object)*
            -   **mode**: *required(one of dimming, previous)*

                Dimming will set the brightness to the specified value after power up. When setting mode “dimming”, the dimming property must be included. Previous will set brightness to the state it was in before powering off.

            -   **dimming**: *(object)*

                -   **brightness**: *required(number – maximum: 100)*

                    Brightness percentage. Value 0 is the lowest possible brightness.

                    **Example**:

                    ```
                    80
                    ```


                **Example**:

                ```json
                {
                  "brightness": 80
                }
                ```

        -   **color**: *(object)*
            -   **mode**: *required(one of color\_temperature, color, previous)*

                State to activate after powerup. Availability of “color\_temperature” and “color” modes depend on the capabilities of the lamp. Colortemperature will set the colortemperature to the specified value after power up. When setting color\_temperature, the color\_temperature property must be included Color will set the color tot he specified value after power up. When setting color mode, the color property must be included Previous will set color to the state it was in before powering off.

            -   **color\_temperature**: *(object)*

                -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

                    color temperature in mirek or null when the light color is not in the ct spectrum

                    **Example**:

                    ```
                    233
                    ```


                **Example**:

                ```json
                {
                  "mirek": 202
                }
                ```

            -   **color**: *(object)*

                -   **xy**: *required(object)*

                    CIE XY gamut position

                    -   **x**: *required(number – minimum: 0 – maximum: 1)*

                        X position in color gamut

                    -   **y**: *required(number – minimum: 0 – maximum: 1)*

                        Y position in color gamut


                    **Example**:

                    ```json
                    {
                      "x": 0.369,
                      "y": 0.445
                    }
                    ```


                **Example**:

                ```json
                {
                  "xy": {
                    "x": 0.6915,
                    "y": 0.3083
                  }
                }
                ```

    -   **content\_configuration**: *(object)*

        Configuration parameters that affect how content is deployed on a product

        -   **orientation**: *(object)*

            Orientation for content deployed on pixelated products when applicable (e.g. sunrise, candle)

            -   **status**: *required(one of set, changing)*
            -   **configurable**: *required(boolean)*

                Indicates if the product allows modifying the configuration attribute

            -   **orientation**: *required(one of horizontal, vertical)*
        -   **order**: *(object)*

            Order in which the content is represented on the pixels

            -   **status**: *required(one of set, changing)*
            -   **configurable**: *required(boolean)*

                Indicates if the product allows modifying the configuration attribute

            -   **order**: *required(one of forward, reversed)*
    -   **geometry**: *(object)*

        Feature describing the geometry a light service.

        -   **pixel\_positions**: *(array of PixelPosition – maxItems: 5)*

            **Items**: PixelPosition

            -   **index**: *required(integer – minimum: 0)*

                Index of the pixel to which the position belongs, use 0 for non-gradient lights

            -   **position**: *required(object)*

                Position of the pixel in 3D space. Range of coordinates is -100 to 100 meters.

                -   **x**: *required(number)*
                -   **y**: *required(number)*
                -   **z**: *required(number)*



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### /resource/scene

API to manage scenes. Scenes are used to store and recall settings for a group of lights.

post

create a new resource of this type

get

List all resources of this type

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **actions**: *required(array of ActionPost)*

    List of actions to be executed synchronously on recall

    **Items**: ActionPost

    -   **target**: *required(object)*

        The identifier of the light to execute the action on

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **action**: *required(object)*

        the action to be executed on recall

        -   **on**: *(object)*
            -   **on**: *required(boolean)*

                On/Off state of the light on=true, off=false

        -   **dimming**: *(object)*

            -   **brightness**: *required(number – maximum: 100)*

                Brightness percentage. Value 0 is the lowest possible brightness.

                **Example**:

                ```
                80
                ```


            **Example**:

            ```json
            {
              "brightness": 80
            }
            ```

        -   **color**: *(object)*

            -   **xy**: *required(object)*

                CIE XY gamut position

                -   **x**: *required(number – minimum: 0 – maximum: 1)*

                    X position in color gamut

                -   **y**: *required(number – minimum: 0 – maximum: 1)*

                    Y position in color gamut


                **Example**:

                ```json
                {
                  "x": 0.369,
                  "y": 0.445
                }
                ```


            **Example**:

            ```json
            {
              "xy": {
                "x": 0.6915,
                "y": 0.3083
              }
            }
            ```

        -   **color\_temperature**: *(object)*

            -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

                color temperature in mirek or null when the light color is not in the ct spectrum

                **Example**:

                ```
                233
                ```


            **Example**:

            ```json
            {
              "mirek": 202
            }
            ```

        -   **gradient**: *(object)*

            Basic feature containing gradient properties.

            -   **points**: *required(array of GradientPointPost – maxItems: 5)*

                Collection of gradients points. For control of the gradient points through a PUT a minimum of 2 points need to be provided.

                **Items**: GradientPointPost

                -   **color**: *required(object)*

                    -   **xy**: *required(object)*

                        CIE XY gamut position

                        -   **x**: *required(number – minimum: 0 – maximum: 1)*

                            X position in color gamut

                        -   **y**: *required(number – minimum: 0 – maximum: 1)*

                            Y position in color gamut


                        **Example**:

                        ```json
                        {
                          "x": 0.369,
                          "y": 0.445
                        }
                        ```


                    **Example**:

                    ```json
                    {
                      "xy": {
                        "x": 0.6915,
                        "y": 0.3083
                      }
                    }
                    ```


            -   **mode**: *(one of interpolated\_palette, interpolated\_palette\_mirrored, random\_pixelated, segmented\_palette)*

                Mode in which the points are currently being deployed. If not provided during PUT/POST it will be defaulted to interpolated\_palette

        -   **effects**: *(object)*

            Deprecated: use effects\_v2

            -   **effect**: *(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*
        -   **effects\_v2**: *(object)*

            Basic feature containing effects v2 properties

            -   **action**: *required(object)*
                -   **effect**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*
                -   **parameters**: *(object)*
                    -   **color**: *(object)*

                        -   **xy**: *required(object)*

                            CIE XY gamut position

                            -   **x**: *required(number – minimum: 0 – maximum: 1)*

                                X position in color gamut

                            -   **y**: *required(number – minimum: 0 – maximum: 1)*

                                Y position in color gamut


                            **Example**:

                            ```json
                            {
                              "x": 0.369,
                              "y": 0.445
                            }
                            ```


                        **Example**:

                        ```json
                        {
                          "xy": {
                            "x": 0.6915,
                            "y": 0.3083
                          }
                        }
                        ```

                    -   **color\_temperature**: *(object)*

                        -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

                            color temperature in mirek or null when the light color is not in the ct spectrum

                            **Example**:

                            ```json
                            233
                            ```


                        **Example**:

                        ```json
                        {
                          "mirek": 202
                        }
                        ```

                    -   **speed**: *(number – minimum: 0 – maximum: 1)*
        -   **dynamics**: *(object)*
            -   **duration**: *(integer)*

                Duration of a light transition in ms.

                **Example**:

                ```
                800
                ```


    **Example**:

    ```json
    {
      "target": {
        "rid": "08221fb4-ce0c-4914-845d-50aa5033f126",
        "rtype": "light"
      },
      "action": {
        "on": {
          "on": true
        },
        "dimming": {
          "brightness": 30
        },
        "color": {
          "xy": {
            "x": 0.6915,
            "y": 0.3083
          }
        },
        "color_temperature": {
          "mirek": 200
        },
        "effects": {
          "effect": "candle"
        },
        "effects_v2": {
          "action": {
            "effect": "fire",
            "parameters": {
              "color": {
                "xy": {
                  "x": 0.6915,
                  "y": 0.3083
                }
              },
              "color_temperature": {
                "mirek": 250
              },
              "speed": 0.3
            }
          }
        },
        "gradient": {
          "points": [
            {
              "color": {
                "xy": {
                  "x": 0.6915,
                  "y": 0.3083
                }
              }
            },
            {
              "color": {
                "xy": {
                  "x": 0.1532,
                  "y": 0.4431
                }
              }
            }
          ],
          "mode": "interpolated_palette"
        },
        "dynamics": {
          "duration": 400
        }
      }
    }
    ```

-   **palette**: *(object)*

    Group of colors that describe the palette of colors to be used when playing dynamics

    -   **color**: *required(array of ColorPalettePost – minItems: 0 – maxItems: 9)*

        **Items**: ColorPalettePost

        -   **color**: *required(object)*

            -   **xy**: *required(object)*

                CIE XY gamut position

                -   **x**: *required(number – minimum: 0 – maximum: 1)*

                    X position in color gamut

                -   **y**: *required(number – minimum: 0 – maximum: 1)*

                    Y position in color gamut


                **Example**:

                ```json
                {
                  "x": 0.369,
                  "y": 0.445
                }
                ```


            **Example**:

            ```json
            {
              "xy": {
                "x": 0.6915,
                "y": 0.3083
              }
            }
            ```

        -   **dimming**: *required(object)*

            -   **brightness**: *required(number – maximum: 100)*

                Brightness percentage. Value 0 is the lowest possible brightness.

                **Example**:

                ```json
                80
                ```


            **Example**:

            ```json
            {
              "brightness": 80
            }
            ```


    -   **dimming**: *required(array of DimmingFeatureBasicPost – minItems: 0 – maxItems: 1)*

        **Items**: DimmingFeatureBasicPost

        -   **brightness**: *required(number – maximum: 100)*

            Brightness percentage. Value 0 is the lowest possible brightness.

            **Example**:

            ```json
            80
            ```


        **Example**:

        ```json
        {
          "brightness": 80
        }
        ```

    -   **color\_temperature**: *required(array of ColorTemperaturePalettePost – minItems: 0 – maxItems: 1)*

        **Items**: ColorTemperaturePalettePost

        -   **color\_temperature**: *required(object)*

            -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

                color temperature in mirek or null when the light color is not in the ct spectrum

                **Example**:

                ```json
                233
                ```


            **Example**:

            ```json
            {
              "mirek": 202
            }
            ```

        -   **dimming**: *required(object)*

            -   **brightness**: *required(number – maximum: 100)*

                Brightness percentage. Value 0 is the lowest possible brightness.

                **Example**:

                ```json
                80
                ```


            **Example**:

            ```json
            {
              "brightness": 80
            }
            ```


    -   **effects**: *(array of EffectFeatureBasicPost – minItems: 0 – maxItems: 3)*

        Deprecated. use effects\_v2

        **Items**: EffectFeatureBasicPost

        -   **effect**: *(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*

    -   **effects\_v2**: *(array of EffectV2FeatureBasicPost – minItems: 0 – maxItems: 3)*

        **Items**: EffectV2FeatureBasicPost

        -   **action**: *required(object)*
            -   **effect**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*
            -   **parameters**: *(object)*
                -   **color**: *(object)*

                    -   **xy**: *required(object)*

                        CIE XY gamut position

                        -   **x**: *required(number – minimum: 0 – maximum: 1)*

                            X position in color gamut

                        -   **y**: *required(number – minimum: 0 – maximum: 1)*

                            Y position in color gamut


                        **Example**:

                        ```json
                        {
                          "x": 0.369,
                          "y": 0.445
                        }
                        ```


                    **Example**:

                    ```json
                    {
                      "xy": {
                        "x": 0.6915,
                        "y": 0.3083
                      }
                    }
                    ```

                -   **color\_temperature**: *(object)*

                    -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

                        color temperature in mirek or null when the light color is not in the ct spectrum

                        **Example**:

                        ```json
                        233
                        ```


                    **Example**:

                    ```json
                    {
                      "mirek": 202
                    }
                    ```

                -   **speed**: *(number – minimum: 0 – maximum: 1)*

-   **type**: *(scene)*

    Type of the supported resources

-   **metadata**: *required(object)*
    -   **name**: *required(string – minLength: 1 – maxLength: 32)*

        Human readable name of a resource

    -   **image**: *(object)*

        Reference with unique identifier for the image representing the scene only accepting “rtype”: “public\_image” on creation

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **appdata**: *(string – minLength: 1 – maxLength: 16)*

        Application specific data. Free format string.

-   **group**: *required(object)*

    Group associated with this Scene. All services in the group are part of this scene. If the group is changed the scene is update (e.g. light added/removed)

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource

-   **speed**: *(number – minimum: 0 – maximum: 1)*

    Speed of dynamic palette for this scene

-   **auto\_dynamic**: *(boolean)*

    Indicates whether to automatically start the scene dynamically on active recall


## HTTP status code 201

Resource has been created

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 400

Request was on resource path was unsuccessful, body of request contains errors.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of SceneGet)*

    **Items**: SceneGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **actions**: *required(array of ActionGet)*

        List of actions to be executed synchronously on recall

        **Items**: ActionGet

        -   **target**: *required(object)*

            The identifier of the light to execute the action on

            -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

                The unique id of the referenced resource

            -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

                The type of the referenced resource

        -   **action**: *required(object)*

            the action to be executed on recall

            -   **on**: *(object)*
                -   **on**: *required(boolean)*

                    On/Off state of the light on=true, off=false

            -   **dimming**: *(object)*

                -   **brightness**: *required(number – maximum: 100)*

                    Brightness percentage. Value 0 is the lowest possible brightness.

                    **Example**:

                    ```
                    80
                    ```


                **Example**:

                ```json
                {
                  "brightness": 80
                }
                ```

            -   **color**: *(object)*

                -   **xy**: *required(object)*

                    CIE XY gamut position

                    -   **x**: *required(number – minimum: 0 – maximum: 1)*

                        X position in color gamut

                    -   **y**: *required(number – minimum: 0 – maximum: 1)*

                        Y position in color gamut


                    **Example**:

                    ```json
                    {
                      "x": 0.369,
                      "y": 0.445
                    }
                    ```


                **Example**:

                ```json
                {
                  "xy": {
                    "x": 0.6915,
                    "y": 0.3083
                  }
                }
                ```

            -   **color\_temperature**: *(object)*

                -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

                    color temperature in mirek or null when the light color is not in the ct spectrum

                    **Example**:

                    ```
                    233
                    ```


                **Example**:

                ```json
                {
                  "mirek": 202
                }
                ```

            -   **gradient**: *(object)*

                Basic feature containing gradient properties.

                -   **points**: *required(array of GradientPointGet – maxItems: 5)*

                    Collection of gradients points. For control of the gradient points through a PUT a minimum of 2 points need to be provided.

                    **Items**: GradientPointGet

                    -   **color**: *required(object)*

                        -   **xy**: *required(object)*

                            CIE XY gamut position

                            -   **x**: *required(number – minimum: 0 – maximum: 1)*

                                X position in color gamut

                            -   **y**: *required(number – minimum: 0 – maximum: 1)*

                                Y position in color gamut


                            **Example**:

                            ```json
                            {
                              "x": 0.369,
                              "y": 0.445
                            }
                            ```


                        **Example**:

                        ```json
                        {
                          "xy": {
                            "x": 0.6915,
                            "y": 0.3083
                          }
                        }
                        ```


                -   **mode**: *required(one of interpolated\_palette, interpolated\_palette\_mirrored, random\_pixelated, segmented\_palette)*

                    Mode in which the points are currently being deployed. If not provided during PUT/POST it will be defaulted to interpolated\_palette

            -   **effects**: *(object)*

                Deprecated: use effects\_v2

                -   **effect**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*
            -   **effects\_v2**: *(object)*

                Basic feature containing effects v2 properties

                -   **action**: *required(object)*
                    -   **effect**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*
                    -   **parameters**: *(object)*
                        -   **color**: *(object)*

                            -   **xy**: *required(object)*

                                CIE XY gamut position

                                -   **x**: *required(number – minimum: 0 – maximum: 1)*

                                    X position in color gamut

                                -   **y**: *required(number – minimum: 0 – maximum: 1)*

                                    Y position in color gamut


                                **Example**:

                                ```json
                                {
                                  "x": 0.369,
                                  "y": 0.445
                                }
                                ```


                            **Example**:

                            ```json
                            {
                              "xy": {
                                "x": 0.6915,
                                "y": 0.3083
                              }
                            }
                            ```

                        -   **color\_temperature**: *(object)*

                            -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

                                color temperature in mirek or null when the light color is not in the ct spectrum

                                **Example**:

                                ```json
                                233
                                ```


                            **Example**:

                            ```json
                            {
                              "mirek": 202
                            }
                            ```

                        -   **speed**: *(number – minimum: 0 – maximum: 1)*
            -   **dynamics**: *(object)*
                -   **duration**: *(integer)*

                    Duration of a light transition in ms.

                    **Example**:

                    ```
                    800
                    ```


        **Example**:

        ```json
        {
          "target": {
            "rid": "08221fb4-ce0c-4914-845d-50aa5033f126",
            "rtype": "light"
          },
          "action": {
            "on": {
              "on": true
            },
            "dimming": {
              "brightness": 30
            },
            "color": {
              "xy": {
                "x": 0.6915,
                "y": 0.3083
              }
            },
            "color_temperature": {
              "mirek": 200
            },
            "effects": {
              "effect": "candle"
            },
            "effects_v2": {
              "action": {
                "effect": "fire",
                "parameters": {
                  "color": {
                    "xy": {
                      "x": 0.6915,
                      "y": 0.3083
                    }
                  },
                  "color_temperature": {
                    "mirek": 250
                  },
                  "speed": 0.3
                }
              }
            },
            "gradient": {
              "points": [
                {
                  "color": {
                    "xy": {
                      "x": 0.6915,
                      "y": 0.3083
                    }
                  }
                },
                {
                  "color": {
                    "xy": {
                      "x": 0.1532,
                      "y": 0.4431
                    }
                  }
                }
              ],
              "mode": "interpolated_palette"
            },
            "dynamics": {
              "duration": 400
            }
          }
        }
        ```

    -   **palette**: *(object)*

        Group of colors that describe the palette of colors to be used when playing dynamics

        -   **color**: *required(array of ColorPaletteGet – minItems: 0 – maxItems: 9)*

            **Items**: ColorPaletteGet

            -   **color**: *required(object)*

                -   **xy**: *required(object)*

                    CIE XY gamut position

                    -   **x**: *required(number – minimum: 0 – maximum: 1)*

                        X position in color gamut

                    -   **y**: *required(number – minimum: 0 – maximum: 1)*

                        Y position in color gamut


                    **Example**:

                    ```json
                    {
                      "x": 0.369,
                      "y": 0.445
                    }
                    ```


                **Example**:

                ```json
                {
                  "xy": {
                    "x": 0.6915,
                    "y": 0.3083
                  }
                }
                ```

            -   **dimming**: *required(object)*

                -   **brightness**: *required(number – maximum: 100)*

                    Brightness percentage. Value 0 is the lowest possible brightness.

                    **Example**:

                    ```json
                    80
                    ```


                **Example**:

                ```json
                {
                  "brightness": 80
                }
                ```


        -   **dimming**: *required(array of DimmingFeatureBasicGet – minItems: 0 – maxItems: 1)*

            **Items**: DimmingFeatureBasicGet

            -   **brightness**: *required(number – maximum: 100)*

                Brightness percentage. Value 0 is the lowest possible brightness.

                **Example**:

                ```json
                80
                ```


            **Example**:

            ```json
            {
              "brightness": 80
            }
            ```

        -   **color\_temperature**: *required(array of ColorTemperaturePaletteGet – minItems: 0 – maxItems: 1)*

            **Items**: ColorTemperaturePaletteGet

            -   **color\_temperature**: *required(object)*

                -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

                    color temperature in mirek or null when the light color is not in the ct spectrum

                    **Example**:

                    ```json
                    233
                    ```


                **Example**:

                ```json
                {
                  "mirek": 202
                }
                ```

            -   **dimming**: *required(object)*

                -   **brightness**: *required(number – maximum: 100)*

                    Brightness percentage. Value 0 is the lowest possible brightness.

                    **Example**:

                    ```json
                    80
                    ```


                **Example**:

                ```json
                {
                  "brightness": 80
                }
                ```


        -   **effects**: *required(array of EffectFeatureBasicGet – minItems: 0 – maxItems: 3)*

            Deprecated. use effects\_v2

            **Items**: EffectFeatureBasicGet

            -   **effect**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*

        -   **effects\_v2**: *required(array of EffectV2FeatureBasicGet – minItems: 0 – maxItems: 3)*

            **Items**: EffectV2FeatureBasicGet

            -   **action**: *required(object)*
                -   **effect**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*
                -   **parameters**: *(object)*
                    -   **color**: *(object)*

                        -   **xy**: *required(object)*

                            CIE XY gamut position

                            -   **x**: *required(number – minimum: 0 – maximum: 1)*

                                X position in color gamut

                            -   **y**: *required(number – minimum: 0 – maximum: 1)*

                                Y position in color gamut


                            **Example**:

                            ```json
                            {
                              "x": 0.369,
                              "y": 0.445
                            }
                            ```


                        **Example**:

                        ```json
                        {
                          "xy": {
                            "x": 0.6915,
                            "y": 0.3083
                          }
                        }
                        ```

                    -   **color\_temperature**: *(object)*

                        -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

                            color temperature in mirek or null when the light color is not in the ct spectrum

                            **Example**:

                            ```json
                            233
                            ```


                        **Example**:

                        ```json
                        {
                          "mirek": 202
                        }
                        ```

                    -   **speed**: *(number – minimum: 0 – maximum: 1)*

    -   **recall**: *required(object)*
    -   **type**: *required(scene)*

        Type of the supported resources

    -   **metadata**: *required(object)*
        -   **name**: *required(string – minLength: 1 – maxLength: 32)*

            Human readable name of a resource

        -   **image**: *(object)*

            Reference with unique identifier for the image representing the scene only accepting “rtype”: “public\_image” on creation

            -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

                The unique id of the referenced resource

            -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

                The type of the referenced resource

        -   **appdata**: *(string – minLength: 1 – maxLength: 16)*

            Application specific data. Free format string.

    -   **group**: *required(object)*

        Group associated with this Scene. All services in the group are part of this scene. If the group is changed the scene is update (e.g. light added/removed)

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **speed**: *required(number – minimum: 0 – maximum: 1)*

        Speed of dynamic palette for this scene

    -   **auto\_dynamic**: *required(boolean)*

        Indicates whether to automatically start the scene dynamically on active recall

    -   **status**: *required(object)*

        Consists the information about the current status and last time it is recalled

        -   **active**: *(one of inactive, static, dynamic\_palette)*
        -   **last\_recall**: *(datetime)*


## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 406

Query parameter has invalid value.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


API calls to read, update and delete a resource instance

delete

put

get

### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Resource has been deleted

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### URI Parameters

-   **id**: *required(string)*

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **actions**: *(array of ActionPut)*

    List of actions to be executed synchronously on recall

    **Items**: ActionPut

    -   **target**: *required(object)*

        The identifier of the light to execute the action on

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **action**: *required(object)*

        the action to be executed on recall

        -   **on**: *(object)*
            -   **on**: *(boolean)*

                On/Off state of the light on=true, off=false

        -   **dimming**: *(object)*

            -   **brightness**: *(number – maximum: 100)*

                Brightness percentage. Value 0 is the lowest possible brightness.

                **Example**:

                ```
                80
                ```


            **Example**:

            ```json
            {
              "brightness": 80
            }
            ```

        -   **color**: *(object)*

            -   **xy**: *(object)*

                CIE XY gamut position

                -   **x**: *required(number – minimum: 0 – maximum: 1)*

                    X position in color gamut

                -   **y**: *required(number – minimum: 0 – maximum: 1)*

                    Y position in color gamut


                **Example**:

                ```json
                {
                  "x": 0.369,
                  "y": 0.445
                }
                ```


            **Example**:

            ```json
            {
              "xy": {
                "x": 0.6915,
                "y": 0.3083
              }
            }
            ```

        -   **color\_temperature**: *(object)*

            -   **mirek**: *(integer – minimum: 50 – maximum: 1000)*

                color temperature in mirek or null when the light color is not in the ct spectrum

                **Example**:

                ```
                233
                ```


            **Example**:

            ```json
            {
              "mirek": 202
            }
            ```

        -   **gradient**: *(object)*

            Basic feature containing gradient properties.

            -   **points**: *required(array of GradientPointPut – maxItems: 5)*

                Collection of gradients points. For control of the gradient points through a PUT a minimum of 2 points need to be provided.

                **Items**: GradientPointPut

                -   **color**: *required(object)*

                    -   **xy**: *(object)*

                        CIE XY gamut position

                        -   **x**: *required(number – minimum: 0 – maximum: 1)*

                            X position in color gamut

                        -   **y**: *required(number – minimum: 0 – maximum: 1)*

                            Y position in color gamut


                        **Example**:

                        ```json
                        {
                          "x": 0.369,
                          "y": 0.445
                        }
                        ```


                    **Example**:

                    ```json
                    {
                      "xy": {
                        "x": 0.6915,
                        "y": 0.3083
                      }
                    }
                    ```


            -   **mode**: *(one of interpolated\_palette, interpolated\_palette\_mirrored, random\_pixelated, segmented\_palette)*

                Mode in which the points are currently being deployed. If not provided during PUT/POST it will be defaulted to interpolated\_palette

        -   **effects**: *(object)*

            Deprecated: use effects\_v2

            -   **effect**: *(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*
        -   **effects\_v2**: *(object)*

            Basic feature containing effects v2 properties

            -   **action**: *(object)*
                -   **effect**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*
                -   **parameters**: *(object)*
                    -   **color**: *(object)*

                        -   **xy**: *(object)*

                            CIE XY gamut position

                            -   **x**: *required(number – minimum: 0 – maximum: 1)*

                                X position in color gamut

                            -   **y**: *required(number – minimum: 0 – maximum: 1)*

                                Y position in color gamut


                            **Example**:

                            ```json
                            {
                              "x": 0.369,
                              "y": 0.445
                            }
                            ```


                        **Example**:

                        ```json
                        {
                          "xy": {
                            "x": 0.6915,
                            "y": 0.3083
                          }
                        }
                        ```

                    -   **color\_temperature**: *(object)*

                        -   **mirek**: *(integer – minimum: 50 – maximum: 1000)*

                            color temperature in mirek or null when the light color is not in the ct spectrum

                            **Example**:

                            ```json
                            233
                            ```


                        **Example**:

                        ```json
                        {
                          "mirek": 202
                        }
                        ```

                    -   **speed**: *(number – minimum: 0 – maximum: 1)*
        -   **dynamics**: *(object)*
            -   **duration**: *(integer)*

                Duration of a light transition in ms.

                **Example**:

                ```
                800
                ```


    **Example**:

    ```json
    {
      "target": {
        "rid": "08221fb4-ce0c-4914-845d-50aa5033f126",
        "rtype": "light"
      },
      "action": {
        "on": {
          "on": true
        },
        "dimming": {
          "brightness": 30
        },
        "color": {
          "xy": {
            "x": 0.6915,
            "y": 0.3083
          }
        },
        "color_temperature": {
          "mirek": 200
        },
        "effects": {
          "effect": "candle"
        },
        "effects_v2": {
          "action": {
            "effect": "fire",
            "parameters": {
              "color": {
                "xy": {
                  "x": 0.6915,
                  "y": 0.3083
                }
              },
              "color_temperature": {
                "mirek": 250
              },
              "speed": 0.3
            }
          }
        },
        "gradient": {
          "points": [
            {
              "color": {
                "xy": {
                  "x": 0.6915,
                  "y": 0.3083
                }
              }
            },
            {
              "color": {
                "xy": {
                  "x": 0.1532,
                  "y": 0.4431
                }
              }
            }
          ],
          "mode": "interpolated_palette"
        },
        "dynamics": {
          "duration": 400
        }
      }
    }
    ```

-   **palette**: *(object)*

    Group of colors that describe the palette of colors to be used when playing dynamics

    -   **color**: *required(array of ColorPalettePut – minItems: 0 – maxItems: 9)*

        **Items**: ColorPalettePut

        -   **color**: *required(object)*

            -   **xy**: *(object)*

                CIE XY gamut position

                -   **x**: *required(number – minimum: 0 – maximum: 1)*

                    X position in color gamut

                -   **y**: *required(number – minimum: 0 – maximum: 1)*

                    Y position in color gamut


                **Example**:

                ```json
                {
                  "x": 0.369,
                  "y": 0.445
                }
                ```


            **Example**:

            ```json
            {
              "xy": {
                "x": 0.6915,
                "y": 0.3083
              }
            }
            ```

        -   **dimming**: *required(object)*

            -   **brightness**: *(number – maximum: 100)*

                Brightness percentage. Value 0 is the lowest possible brightness.

                **Example**:

                ```json
                80
                ```


            **Example**:

            ```json
            {
              "brightness": 80
            }
            ```


    -   **dimming**: *required(array of DimmingFeatureBasicPut – minItems: 0 – maxItems: 1)*

        **Items**: DimmingFeatureBasicPut

        -   **brightness**: *(number – maximum: 100)*

            Brightness percentage. Value 0 is the lowest possible brightness.

            **Example**:

            ```json
            80
            ```


        **Example**:

        ```json
        {
          "brightness": 80
        }
        ```

    -   **color\_temperature**: *required(array of ColorTemperaturePalettePut – minItems: 0 – maxItems: 1)*

        **Items**: ColorTemperaturePalettePut

        -   **color\_temperature**: *required(object)*

            -   **mirek**: *(integer – minimum: 50 – maximum: 1000)*

                color temperature in mirek or null when the light color is not in the ct spectrum

                **Example**:

                ```json
                233
                ```


            **Example**:

            ```json
            {
              "mirek": 202
            }
            ```

        -   **dimming**: *required(object)*

            -   **brightness**: *(number – maximum: 100)*

                Brightness percentage. Value 0 is the lowest possible brightness.

                **Example**:

                ```json
                80
                ```


            **Example**:

            ```json
            {
              "brightness": 80
            }
            ```


    -   **effects**: *(array of EffectFeatureBasicPut – minItems: 0 – maxItems: 3)*

        Deprecated. use effects\_v2

        **Items**: EffectFeatureBasicPut

        -   **effect**: *(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*

    -   **effects\_v2**: *(array of EffectV2FeatureBasicPut – minItems: 0 – maxItems: 3)*

        **Items**: EffectV2FeatureBasicPut

        -   **action**: *(object)*
            -   **effect**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*
            -   **parameters**: *(object)*
                -   **color**: *(object)*

                    -   **xy**: *(object)*

                        CIE XY gamut position

                        -   **x**: *required(number – minimum: 0 – maximum: 1)*

                            X position in color gamut

                        -   **y**: *required(number – minimum: 0 – maximum: 1)*

                            Y position in color gamut


                        **Example**:

                        ```json
                        {
                          "x": 0.369,
                          "y": 0.445
                        }
                        ```


                    **Example**:

                    ```json
                    {
                      "xy": {
                        "x": 0.6915,
                        "y": 0.3083
                      }
                    }
                    ```

                -   **color\_temperature**: *(object)*

                    -   **mirek**: *(integer – minimum: 50 – maximum: 1000)*

                        color temperature in mirek or null when the light color is not in the ct spectrum

                        **Example**:

                        ```json
                        233
                        ```


                    **Example**:

                    ```json
                    {
                      "mirek": 202
                    }
                    ```

                -   **speed**: *(number – minimum: 0 – maximum: 1)*

-   **recall**: *(object)*
    -   **action**: *(one of active, dynamic\_palette, static)*

        When writing active, the actions in the scene are executed on the target. dynamic\_palette starts dynamic scene with colors in the Palette object.

    -   **duration**: *(integer)*

        transition to the scene within the timeframe given by duration

        **Example**:

        ```
        800
        ```

    -   **dimming**: *(object)*

        override the scene dimming/brightness

        -   **brightness**: *(number – maximum: 100)*

            Brightness percentage. Value 0 is the lowest possible brightness.

            **Example**:

            ```
            80
            ```


        **Example**:

        ```json
        {
          "brightness": 80
        }
        ```

-   **type**: *(scene)*

    Type of the supported resources

-   **metadata**: *(object)*
    -   **name**: *(string – minLength: 1 – maxLength: 32)*

        Human readable name of a resource

    -   **appdata**: *(string – minLength: 1 – maxLength: 16)*

        Application specific data. Free format string.

-   **speed**: *(number – minimum: 0 – maximum: 1)*

    Speed of dynamic palette for this scene

-   **auto\_dynamic**: *(boolean)*

    Indicates whether to automatically start the scene dynamically on active recall


## HTTP status code 200

Request was on resource path was successful.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 207

Request was partially successful, resource has been updated. Failed requests and errors are listed in error object.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 400

Request was on resource path was unsuccessful, body of request contains errors.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of SceneGet)*

    **Items**: SceneGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **actions**: *required(array of ActionGet)*

        List of actions to be executed synchronously on recall

        **Items**: ActionGet

        -   **target**: *required(object)*

            The identifier of the light to execute the action on

            -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

                The unique id of the referenced resource

            -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

                The type of the referenced resource

        -   **action**: *required(object)*

            the action to be executed on recall

            -   **on**: *(object)*
                -   **on**: *required(boolean)*

                    On/Off state of the light on=true, off=false

            -   **dimming**: *(object)*

                -   **brightness**: *required(number – maximum: 100)*

                    Brightness percentage. Value 0 is the lowest possible brightness.

                    **Example**:

                    ```
                    80
                    ```


                **Example**:

                ```json
                {
                  "brightness": 80
                }
                ```

            -   **color**: *(object)*

                -   **xy**: *required(object)*

                    CIE XY gamut position

                    -   **x**: *required(number – minimum: 0 – maximum: 1)*

                        X position in color gamut

                    -   **y**: *required(number – minimum: 0 – maximum: 1)*

                        Y position in color gamut


                    **Example**:

                    ```json
                    {
                      "x": 0.369,
                      "y": 0.445
                    }
                    ```


                **Example**:

                ```json
                {
                  "xy": {
                    "x": 0.6915,
                    "y": 0.3083
                  }
                }
                ```

            -   **color\_temperature**: *(object)*

                -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

                    color temperature in mirek or null when the light color is not in the ct spectrum

                    **Example**:

                    ```
                    233
                    ```


                **Example**:

                ```json
                {
                  "mirek": 202
                }
                ```

            -   **gradient**: *(object)*

                Basic feature containing gradient properties.

                -   **points**: *required(array of GradientPointGet – maxItems: 5)*

                    Collection of gradients points. For control of the gradient points through a PUT a minimum of 2 points need to be provided.

                    **Items**: GradientPointGet

                    -   **color**: *required(object)*

                        -   **xy**: *required(object)*

                            CIE XY gamut position

                            -   **x**: *required(number – minimum: 0 – maximum: 1)*

                                X position in color gamut

                            -   **y**: *required(number – minimum: 0 – maximum: 1)*

                                Y position in color gamut


                            **Example**:

                            ```json
                            {
                              "x": 0.369,
                              "y": 0.445
                            }
                            ```


                        **Example**:

                        ```json
                        {
                          "xy": {
                            "x": 0.6915,
                            "y": 0.3083
                          }
                        }
                        ```


                -   **mode**: *required(one of interpolated\_palette, interpolated\_palette\_mirrored, random\_pixelated, segmented\_palette)*

                    Mode in which the points are currently being deployed. If not provided during PUT/POST it will be defaulted to interpolated\_palette

            -   **effects**: *(object)*

                Deprecated: use effects\_v2

                -   **effect**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*
            -   **effects\_v2**: *(object)*

                Basic feature containing effects v2 properties

                -   **action**: *required(object)*
                    -   **effect**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*
                    -   **parameters**: *(object)*
                        -   **color**: *(object)*

                            -   **xy**: *required(object)*

                                CIE XY gamut position

                                -   **x**: *required(number – minimum: 0 – maximum: 1)*

                                    X position in color gamut

                                -   **y**: *required(number – minimum: 0 – maximum: 1)*

                                    Y position in color gamut


                                **Example**:

                                ```json
                                {
                                  "x": 0.369,
                                  "y": 0.445
                                }
                                ```


                            **Example**:

                            ```json
                            {
                              "xy": {
                                "x": 0.6915,
                                "y": 0.3083
                              }
                            }
                            ```

                        -   **color\_temperature**: *(object)*

                            -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

                                color temperature in mirek or null when the light color is not in the ct spectrum

                                **Example**:

                                ```json
                                233
                                ```


                            **Example**:

                            ```json
                            {
                              "mirek": 202
                            }
                            ```

                        -   **speed**: *(number – minimum: 0 – maximum: 1)*
            -   **dynamics**: *(object)*
                -   **duration**: *(integer)*

                    Duration of a light transition in ms.

                    **Example**:

                    ```
                    800
                    ```


        **Example**:

        ```json
        {
          "target": {
            "rid": "08221fb4-ce0c-4914-845d-50aa5033f126",
            "rtype": "light"
          },
          "action": {
            "on": {
              "on": true
            },
            "dimming": {
              "brightness": 30
            },
            "color": {
              "xy": {
                "x": 0.6915,
                "y": 0.3083
              }
            },
            "color_temperature": {
              "mirek": 200
            },
            "effects": {
              "effect": "candle"
            },
            "effects_v2": {
              "action": {
                "effect": "fire",
                "parameters": {
                  "color": {
                    "xy": {
                      "x": 0.6915,
                      "y": 0.3083
                    }
                  },
                  "color_temperature": {
                    "mirek": 250
                  },
                  "speed": 0.3
                }
              }
            },
            "gradient": {
              "points": [
                {
                  "color": {
                    "xy": {
                      "x": 0.6915,
                      "y": 0.3083
                    }
                  }
                },
                {
                  "color": {
                    "xy": {
                      "x": 0.1532,
                      "y": 0.4431
                    }
                  }
                }
              ],
              "mode": "interpolated_palette"
            },
            "dynamics": {
              "duration": 400
            }
          }
        }
        ```

    -   **palette**: *(object)*

        Group of colors that describe the palette of colors to be used when playing dynamics

        -   **color**: *required(array of ColorPaletteGet – minItems: 0 – maxItems: 9)*

            **Items**: ColorPaletteGet

            -   **color**: *required(object)*

                -   **xy**: *required(object)*

                    CIE XY gamut position

                    -   **x**: *required(number – minimum: 0 – maximum: 1)*

                        X position in color gamut

                    -   **y**: *required(number – minimum: 0 – maximum: 1)*

                        Y position in color gamut


                    **Example**:

                    ```json
                    {
                      "x": 0.369,
                      "y": 0.445
                    }
                    ```


                **Example**:

                ```json
                {
                  "xy": {
                    "x": 0.6915,
                    "y": 0.3083
                  }
                }
                ```

            -   **dimming**: *required(object)*

                -   **brightness**: *required(number – maximum: 100)*

                    Brightness percentage. Value 0 is the lowest possible brightness.

                    **Example**:

                    ```json
                    80
                    ```


                **Example**:

                ```json
                {
                  "brightness": 80
                }
                ```


        -   **dimming**: *required(array of DimmingFeatureBasicGet – minItems: 0 – maxItems: 1)*

            **Items**: DimmingFeatureBasicGet

            -   **brightness**: *required(number – maximum: 100)*

                Brightness percentage. Value 0 is the lowest possible brightness.

                **Example**:

                ```json
                80
                ```


            **Example**:

            ```json
            {
              "brightness": 80
            }
            ```

        -   **color\_temperature**: *required(array of ColorTemperaturePaletteGet – minItems: 0 – maxItems: 1)*

            **Items**: ColorTemperaturePaletteGet

            -   **color\_temperature**: *required(object)*

                -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

                    color temperature in mirek or null when the light color is not in the ct spectrum

                    **Example**:

                    ```json
                    233
                    ```


                **Example**:

                ```json
                {
                  "mirek": 202
                }
                ```

            -   **dimming**: *required(object)*

                -   **brightness**: *required(number – maximum: 100)*

                    Brightness percentage. Value 0 is the lowest possible brightness.

                    **Example**:

                    ```json
                    80
                    ```


                **Example**:

                ```json
                {
                  "brightness": 80
                }
                ```


        -   **effects**: *required(array of EffectFeatureBasicGet – minItems: 0 – maxItems: 3)*

            Deprecated. use effects\_v2

            **Items**: EffectFeatureBasicGet

            -   **effect**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*

        -   **effects\_v2**: *required(array of EffectV2FeatureBasicGet – minItems: 0 – maxItems: 3)*

            **Items**: EffectV2FeatureBasicGet

            -   **action**: *required(object)*
                -   **effect**: *required(one of prism, opal, glisten, sparkle, fire, candle, underwater, cosmos, sunbeam, enchant, no\_effect)*
                -   **parameters**: *(object)*
                    -   **color**: *(object)*

                        -   **xy**: *required(object)*

                            CIE XY gamut position

                            -   **x**: *required(number – minimum: 0 – maximum: 1)*

                                X position in color gamut

                            -   **y**: *required(number – minimum: 0 – maximum: 1)*

                                Y position in color gamut


                            **Example**:

                            ```json
                            {
                              "x": 0.369,
                              "y": 0.445
                            }
                            ```


                        **Example**:

                        ```json
                        {
                          "xy": {
                            "x": 0.6915,
                            "y": 0.3083
                          }
                        }
                        ```

                    -   **color\_temperature**: *(object)*

                        -   **mirek**: *required(integer – minimum: 50 – maximum: 1000)*

                            color temperature in mirek or null when the light color is not in the ct spectrum

                            **Example**:

                            ```json
                            233
                            ```


                        **Example**:

                        ```json
                        {
                          "mirek": 202
                        }
                        ```

                    -   **speed**: *(number – minimum: 0 – maximum: 1)*

    -   **recall**: *required(object)*
    -   **type**: *required(scene)*

        Type of the supported resources

    -   **metadata**: *required(object)*
        -   **name**: *required(string – minLength: 1 – maxLength: 32)*

            Human readable name of a resource

        -   **image**: *(object)*

            Reference with unique identifier for the image representing the scene only accepting “rtype”: “public\_image” on creation

            -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

                The unique id of the referenced resource

            -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

                The type of the referenced resource

        -   **appdata**: *(string – minLength: 1 – maxLength: 16)*

            Application specific data. Free format string.

    -   **group**: *required(object)*

        Group associated with this Scene. All services in the group are part of this scene. If the group is changed the scene is update (e.g. light added/removed)

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **speed**: *required(number – minimum: 0 – maximum: 1)*

        Speed of dynamic palette for this scene

    -   **auto\_dynamic**: *required(boolean)*

        Indicates whether to automatically start the scene dynamically on active recall

    -   **status**: *required(object)*

        Consists the information about the current status and last time it is recalled

        -   **active**: *(one of inactive, static, dynamic\_palette)*
        -   **last\_recall**: *(datetime)*


## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### /resource/room

API to manage rooms. Rooms group devices and each device can only be part of one room.

post

create a new resource of this type

get

List all resources of this type

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **children**: *required(array of ResourceIdentifier)*

    Child devices/services to group by the derived group

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource


-   **type**: *(room)*

    Type of the supported resources

-   **metadata**: *required(object)*

    configuration object for a room

    -   **name**: *required(string – minLength: 1 – maxLength: 32)*

        Human readable name of a resource

    -   **archetype**: *required(one of living\_room, kitchen, dining, bedroom, kids\_bedroom, bathroom, nursery, recreation, office, gym, hallway, toilet, front\_door, garage, terrace, garden, driveway, carport, home, downstairs, upstairs, top\_floor, attic, guest\_room, staircase, lounge, man\_cave, computer, studio, music, tv, reading, closet, storage, laundry\_room, balcony, porch, barbecue, pool, other)*

        possible archetypes of a room


**Example**:

```json
{
  "type": "room",
  "id": "6adb232c-e80c-49be-a4c8-fd32b2488bf5",
  "metadata": {
    "archetype": "living_room",
    "name": "Living room"
  },
  "services": [
    {
      "rtype": "grouped_light",
      "rid": "c8e1f2e6-f54f-4d20-bcdc-a568ebf61dcc"
    }
  ],
  "children": [
    {
      "rtype": "device",
      "rid": "d7c305a8-73b6-4de4-9eb8-4dc20007c2c6"
    },
    {
      "rtype": "device",
      "rid": "5f742a9f-5c07-4c56-a996-888d620ccf8d"
    },
    {
      "rtype": "device",
      "rid": "5a251edd-45ab-4b97-b4c6-cde3a5abea33"
    }
  ]
}
```

## HTTP status code 201

Resource has been created

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 400

Request was on resource path was unsuccessful, body of request contains errors.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of RoomGet)*

    **Items**: RoomGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **children**: *required(array of ResourceIdentifier)*

        Child devices/services to group by the derived group

        **Items**: ResourceIdentifier

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource


    -   **services**: *required(array of ResourceIdentifier)*

        References all services aggregating control and state of children in the group. This includes all services grouped in the group hierarchy given by child relation This includes all services of a device grouped in the group hierarchy given by child relation Aggregation is per service type, ie every service type which can be grouped has a corresponding definition of grouped type Supported types: – grouped\_light – grouped\_motion – grouped\_light\_level

        **Items**: ResourceIdentifier

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource


    -   **type**: *required(room)*

        Type of the supported resources

    -   **metadata**: *required(object)*

        configuration object for a room

        -   **name**: *required(string – minLength: 1 – maxLength: 32)*

            Human readable name of a resource

        -   **archetype**: *required(one of living\_room, kitchen, dining, bedroom, kids\_bedroom, bathroom, nursery, recreation, office, gym, hallway, toilet, front\_door, garage, terrace, garden, driveway, carport, home, downstairs, upstairs, top\_floor, attic, guest\_room, staircase, lounge, man\_cave, computer, studio, music, tv, reading, closet, storage, laundry\_room, balcony, porch, barbecue, pool, other)*

            possible archetypes of a room


    **Example**:

    ```json
    {
      "type": "room",
      "id": "6adb232c-e80c-49be-a4c8-fd32b2488bf5",
      "metadata": {
        "archetype": "living_room",
        "name": "Living room"
      },
      "services": [
        {
          "rtype": "grouped_light",
          "rid": "c8e1f2e6-f54f-4d20-bcdc-a568ebf61dcc"
        }
      ],
      "children": [
        {
          "rtype": "device",
          "rid": "d7c305a8-73b6-4de4-9eb8-4dc20007c2c6"
        },
        {
          "rtype": "device",
          "rid": "5f742a9f-5c07-4c56-a996-888d620ccf8d"
        },
        {
          "rtype": "device",
          "rid": "5a251edd-45ab-4b97-b4c6-cde3a5abea33"
        }
      ]
    }
    ```


## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 406

Query parameter has invalid value.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


API calls to read, update and delete a resource instance

delete

put

get

### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Resource has been deleted

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### URI Parameters

-   **id**: *required(string)*

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **children**: *(array of ResourceIdentifier)*

    Child devices/services to group by the derived group

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource


-   **type**: *(room)*

    Type of the supported resources

-   **metadata**: *(object)*

    configuration object for a room

    -   **name**: *(string – minLength: 1 – maxLength: 32)*

        Human readable name of a resource

    -   **archetype**: *(one of living\_room, kitchen, dining, bedroom, kids\_bedroom, bathroom, nursery, recreation, office, gym, hallway, toilet, front\_door, garage, terrace, garden, driveway, carport, home, downstairs, upstairs, top\_floor, attic, guest\_room, staircase, lounge, man\_cave, computer, studio, music, tv, reading, closet, storage, laundry\_room, balcony, porch, barbecue, pool, other)*

        possible archetypes of a room


**Example**:

```json
{
  "type": "room",
  "id": "6adb232c-e80c-49be-a4c8-fd32b2488bf5",
  "metadata": {
    "archetype": "living_room",
    "name": "Living room"
  },
  "services": [
    {
      "rtype": "grouped_light",
      "rid": "c8e1f2e6-f54f-4d20-bcdc-a568ebf61dcc"
    }
  ],
  "children": [
    {
      "rtype": "device",
      "rid": "d7c305a8-73b6-4de4-9eb8-4dc20007c2c6"
    },
    {
      "rtype": "device",
      "rid": "5f742a9f-5c07-4c56-a996-888d620ccf8d"
    },
    {
      "rtype": "device",
      "rid": "5a251edd-45ab-4b97-b4c6-cde3a5abea33"
    }
  ]
}
```

## HTTP status code 200

Request was on resource path was successful.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 207

Request was partially successful, resource has been updated. Failed requests and errors are listed in error object.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 400

Request was on resource path was unsuccessful, body of request contains errors.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of RoomGet)*

    **Items**: RoomGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **children**: *required(array of ResourceIdentifier)*

        Child devices/services to group by the derived group

        **Items**: ResourceIdentifier

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource


    -   **services**: *required(array of ResourceIdentifier)*

        References all services aggregating control and state of children in the group. This includes all services grouped in the group hierarchy given by child relation This includes all services of a device grouped in the group hierarchy given by child relation Aggregation is per service type, ie every service type which can be grouped has a corresponding definition of grouped type Supported types: – grouped\_light – grouped\_motion – grouped\_light\_level

        **Items**: ResourceIdentifier

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource


    -   **type**: *required(room)*

        Type of the supported resources

    -   **metadata**: *required(object)*

        configuration object for a room

        -   **name**: *required(string – minLength: 1 – maxLength: 32)*

            Human readable name of a resource

        -   **archetype**: *required(one of living\_room, kitchen, dining, bedroom, kids\_bedroom, bathroom, nursery, recreation, office, gym, hallway, toilet, front\_door, garage, terrace, garden, driveway, carport, home, downstairs, upstairs, top\_floor, attic, guest\_room, staircase, lounge, man\_cave, computer, studio, music, tv, reading, closet, storage, laundry\_room, balcony, porch, barbecue, pool, other)*

            possible archetypes of a room


    **Example**:

    ```json
    {
      "type": "room",
      "id": "6adb232c-e80c-49be-a4c8-fd32b2488bf5",
      "metadata": {
        "archetype": "living_room",
        "name": "Living room"
      },
      "services": [
        {
          "rtype": "grouped_light",
          "rid": "c8e1f2e6-f54f-4d20-bcdc-a568ebf61dcc"
        }
      ],
      "children": [
        {
          "rtype": "device",
          "rid": "d7c305a8-73b6-4de4-9eb8-4dc20007c2c6"
        },
        {
          "rtype": "device",
          "rid": "5f742a9f-5c07-4c56-a996-888d620ccf8d"
        },
        {
          "rtype": "device",
          "rid": "5a251edd-45ab-4b97-b4c6-cde3a5abea33"
        }
      ]
    }
    ```


## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### /resource/zone

API to manage zones. Zones group services and each service can be part of multiple zones.

post

create a new resource of this type

get

List all resources of this type

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **children**: *required(array of ResourceIdentifier)*

    Child devices/services to group by the derived group

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource


-   **type**: *(zone)*

    Type of the supported resources

-   **metadata**: *required(object)*

    configuration object for a room

    -   **name**: *required(string – minLength: 1 – maxLength: 32)*

        Human readable name of a resource

    -   **archetype**: *required(one of living\_room, kitchen, dining, bedroom, kids\_bedroom, bathroom, nursery, recreation, office, gym, hallway, toilet, front\_door, garage, terrace, garden, driveway, carport, home, downstairs, upstairs, top\_floor, attic, guest\_room, staircase, lounge, man\_cave, computer, studio, music, tv, reading, closet, storage, laundry\_room, balcony, porch, barbecue, pool, other)*

        possible archetypes of a room


**Example**:

```json
{
  "type": "zone",
  "id": "99e2f2bb-a05f-42be-accb-85b241a95ae8",
  "metadata": {
    "archetype": "downstairs",
    "name": "Downstairs"
  },
  "services": [
    {
      "rtype": "grouped_light",
      "rid": "33a7b1df-46b5-493e-b04a-1aa214e31e41"
    }
  ],
  "children": [
    {
      "rtype": "light",
      "rid": "d7c305a8-73b6-4de4-9eb8-4dc20007c2c6"
    },
    {
      "rtype": "light",
      "rid": "5f742a9f-5c07-4c56-a996-888d620ccf8d"
    },
    {
      "rtype": "light",
      "rid": "5a251edd-45ab-4b97-b4c6-cde3a5abea33"
    }
  ]
}
```

## HTTP status code 201

Resource has been created

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 400

Request was on resource path was unsuccessful, body of request contains errors.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ZoneGet)*

    **Items**: ZoneGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **children**: *required(array of ResourceIdentifier)*

        Child devices/services to group by the derived group

        **Items**: ResourceIdentifier

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource


    -   **services**: *required(array of ResourceIdentifier)*

        References all services aggregating control and state of children in the group. This includes all services grouped in the group hierarchy given by child relation This includes all services of a device grouped in the group hierarchy given by child relation Aggregation is per service type, ie every service type which can be grouped has a corresponding definition of grouped type Supported types: – grouped\_light – grouped\_motion – grouped\_light\_level

        **Items**: ResourceIdentifier

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource


    -   **type**: *required(zone)*

        Type of the supported resources

    -   **metadata**: *required(object)*

        configuration object for a room

        -   **name**: *required(string – minLength: 1 – maxLength: 32)*

            Human readable name of a resource

        -   **archetype**: *required(one of living\_room, kitchen, dining, bedroom, kids\_bedroom, bathroom, nursery, recreation, office, gym, hallway, toilet, front\_door, garage, terrace, garden, driveway, carport, home, downstairs, upstairs, top\_floor, attic, guest\_room, staircase, lounge, man\_cave, computer, studio, music, tv, reading, closet, storage, laundry\_room, balcony, porch, barbecue, pool, other)*

            possible archetypes of a room


    **Example**:

    ```json
    {
      "type": "zone",
      "id": "99e2f2bb-a05f-42be-accb-85b241a95ae8",
      "metadata": {
        "archetype": "downstairs",
        "name": "Downstairs"
      },
      "services": [
        {
          "rtype": "grouped_light",
          "rid": "33a7b1df-46b5-493e-b04a-1aa214e31e41"
        }
      ],
      "children": [
        {
          "rtype": "light",
          "rid": "d7c305a8-73b6-4de4-9eb8-4dc20007c2c6"
        },
        {
          "rtype": "light",
          "rid": "5f742a9f-5c07-4c56-a996-888d620ccf8d"
        },
        {
          "rtype": "light",
          "rid": "5a251edd-45ab-4b97-b4c6-cde3a5abea33"
        }
      ]
    }
    ```


## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 406

Query parameter has invalid value.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


API calls to read, update and delete a resource instance

delete

put

get

### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Resource has been deleted

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### URI Parameters

-   **id**: *required(string)*

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **children**: *(array of ResourceIdentifier)*

    Child devices/services to group by the derived group

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource


-   **type**: *(zone)*

    Type of the supported resources

-   **metadata**: *(object)*

    configuration object for a room

    -   **name**: *(string – minLength: 1 – maxLength: 32)*

        Human readable name of a resource

    -   **archetype**: *(one of living\_room, kitchen, dining, bedroom, kids\_bedroom, bathroom, nursery, recreation, office, gym, hallway, toilet, front\_door, garage, terrace, garden, driveway, carport, home, downstairs, upstairs, top\_floor, attic, guest\_room, staircase, lounge, man\_cave, computer, studio, music, tv, reading, closet, storage, laundry\_room, balcony, porch, barbecue, pool, other)*

        possible archetypes of a room


**Example**:

```json
{
  "type": "zone",
  "id": "99e2f2bb-a05f-42be-accb-85b241a95ae8",
  "metadata": {
    "archetype": "downstairs",
    "name": "Downstairs"
  },
  "services": [
    {
      "rtype": "grouped_light",
      "rid": "33a7b1df-46b5-493e-b04a-1aa214e31e41"
    }
  ],
  "children": [
    {
      "rtype": "light",
      "rid": "d7c305a8-73b6-4de4-9eb8-4dc20007c2c6"
    },
    {
      "rtype": "light",
      "rid": "5f742a9f-5c07-4c56-a996-888d620ccf8d"
    },
    {
      "rtype": "light",
      "rid": "5a251edd-45ab-4b97-b4c6-cde3a5abea33"
    }
  ]
}
```

## HTTP status code 200

Request was on resource path was successful.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 207

Request was partially successful, resource has been updated. Failed requests and errors are listed in error object.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 400

Request was on resource path was unsuccessful, body of request contains errors.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ZoneGet)*

    **Items**: ZoneGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **children**: *required(array of ResourceIdentifier)*

        Child devices/services to group by the derived group

        **Items**: ResourceIdentifier

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource


    -   **services**: *required(array of ResourceIdentifier)*

        References all services aggregating control and state of children in the group. This includes all services grouped in the group hierarchy given by child relation This includes all services of a device grouped in the group hierarchy given by child relation Aggregation is per service type, ie every service type which can be grouped has a corresponding definition of grouped type Supported types: – grouped\_light – grouped\_motion – grouped\_light\_level

        **Items**: ResourceIdentifier

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource


    -   **type**: *required(zone)*

        Type of the supported resources

    -   **metadata**: *required(object)*

        configuration object for a room

        -   **name**: *required(string – minLength: 1 – maxLength: 32)*

            Human readable name of a resource

        -   **archetype**: *required(one of living\_room, kitchen, dining, bedroom, kids\_bedroom, bathroom, nursery, recreation, office, gym, hallway, toilet, front\_door, garage, terrace, garden, driveway, carport, home, downstairs, upstairs, top\_floor, attic, guest\_room, staircase, lounge, man\_cave, computer, studio, music, tv, reading, closet, storage, laundry\_room, balcony, porch, barbecue, pool, other)*

            possible archetypes of a room


    **Example**:

    ```json
    {
      "type": "zone",
      "id": "99e2f2bb-a05f-42be-accb-85b241a95ae8",
      "metadata": {
        "archetype": "downstairs",
        "name": "Downstairs"
      },
      "services": [
        {
          "rtype": "grouped_light",
          "rid": "33a7b1df-46b5-493e-b04a-1aa214e31e41"
        }
      ],
      "children": [
        {
          "rtype": "light",
          "rid": "d7c305a8-73b6-4de4-9eb8-4dc20007c2c6"
        },
        {
          "rtype": "light",
          "rid": "5f742a9f-5c07-4c56-a996-888d620ccf8d"
        },
        {
          "rtype": "light",
          "rid": "5a251edd-45ab-4b97-b4c6-cde3a5abea33"
        }
      ]
    }
    ```


## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### /resource/bridge_home

API to manage bridge homes. Homes group rooms as well as devices not assigned to a room.

get

List all resources of this type

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of BridgeHomeGet)*

    **Items**: BridgeHomeGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **children**: *required(array of ResourceIdentifier)*

        Child devices/services to group by the derived group

        **Items**: ResourceIdentifier

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource


    -   **services**: *required(array of ResourceIdentifier)*

        References all services aggregating control and state of children in the group. This includes all services grouped in the group hierarchy given by child relation This includes all services of a device grouped in the group hierarchy given by child relation Aggregation is per service type, ie every service type which can be grouped has a corresponding definition of grouped type Supported types: – grouped\_light – grouped\_motion – grouped\_light\_level

        **Items**: ResourceIdentifier

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource


    -   **type**: *required(bridge\_home)*

        Type of the supported resources


    **Example**:

    ```json
    {
      "type": "bridge_home",
      "id": "91a6626c-8b8f-4c4a-a9e8-18a6be332373",
      "services": [
        {
          "rtype": "grouped_light",
          "rid": "709834c7-29e7-4dc9-961e-0f7a3f4c3842"
        }
      ],
      "children": [
        {
          "rtype": "device",
          "rid": "91a6626c-8b8f-4c4a-a9e8-18a6be332373"
        },
        {
          "rtype": "room",
          "rid": "a032577a-dc4e-4eee-97e1-fc2c146a6b23"
        }
      ]
    }
    ```


## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 406

Query parameter has invalid value.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


API calls for resource which can be just retrieved (GET on resource instance)

get

### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of BridgeHomeGet)*

    **Items**: BridgeHomeGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **children**: *required(array of ResourceIdentifier)*

        Child devices/services to group by the derived group

        **Items**: ResourceIdentifier

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource


    -   **services**: *required(array of ResourceIdentifier)*

        References all services aggregating control and state of children in the group. This includes all services grouped in the group hierarchy given by child relation This includes all services of a device grouped in the group hierarchy given by child relation Aggregation is per service type, ie every service type which can be grouped has a corresponding definition of grouped type Supported types: – grouped\_light – grouped\_motion – grouped\_light\_level

        **Items**: ResourceIdentifier

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource


    -   **type**: *required(bridge\_home)*

        Type of the supported resources


    **Example**:

    ```json
    {
      "type": "bridge_home",
      "id": "91a6626c-8b8f-4c4a-a9e8-18a6be332373",
      "services": [
        {
          "rtype": "grouped_light",
          "rid": "709834c7-29e7-4dc9-961e-0f7a3f4c3842"
        }
      ],
      "children": [
        {
          "rtype": "device",
          "rid": "91a6626c-8b8f-4c4a-a9e8-18a6be332373"
        },
        {
          "rtype": "room",
          "rid": "a032577a-dc4e-4eee-97e1-fc2c146a6b23"
        }
      ]
    }
    ```


## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### /resource/grouped_light

API to manage grouped light services. These are offered by rooms, zones, and homes.

get

List all resources of this type

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of GroupedLightGet)*

    **Items**: GroupedLightGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **owner**: *required(object)*

        Owner of the service, in case the owner service is deleted, the service also gets deleted

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **type**: *required(grouped\_light)*
    -   **on**: *(object)*

        Joined on control & aggregated on state.| – “on” is true if any light in the group is on

        -   **on**: *required(boolean)*

            On/Off state of the light on=true, off=false

    -   **dimming**: *(object)*

        Joined dimming control – “dimming.brightness” contains average brightness of group containing turned-on lights only.

        -   **brightness**: *required(number – maximum: 100)*

            Brightness percentage. Value 0 is the lowest possible brightness.

            **Example**:

            ```
            80
            ```


        **Example**:

        ```json
        {
          "brightness": 80
        }
        ```

    -   **dimming\_delta**: *(object)*

        **Example**:

        ```json
        {
          "action": "up",
          "brightness_delta": 10.5
        }
        ```

    -   **color\_temperature**: *(object)*

        Joined color temperature control

        **Example**:

        ```json
        {
          "mirek": 202
        }
        ```

    -   **color\_temperature\_delta**: *(object)*

        **Example**:

        ```json
        {
          "action": "up",
          "mirek_delta": 200
        }
        ```

    -   **color**: *(object)*

        Joined color control

        **Example**:

        ```json
        {
          "xy": {
            "x": 0.6915,
            "y": 0.3083
          }
        }
        ```

    -   **alert**: *(object)*

        Joined alert control

        -   **action\_values**: *required(array of AlertEffectType)*

            Alert effects that the light supports.

    -   **signaling**: *(object)*

        Feature containing basic signaling properties.

        -   **signal\_values**: *required(array of SupportedSignals)*

            Signals that the light supports.

    -   **dynamics**: *(object)*


## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 406

Query parameter has invalid value.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


API calls for resource which cannot be deleted (PUT and GET on resource instance)

put

get

### URI Parameters

-   **id**: *required(string)*

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **type**: *(grouped\_light)*
-   **on**: *(object)*

    Joined on control & aggregated on state.| – “on” is true if any light in the group is on

    -   **on**: *(boolean)*

        On/Off state of the light on=true, off=false

-   **dimming**: *(object)*

    Joined dimming control – “dimming.brightness” contains average brightness of group containing turned-on lights only.

    -   **brightness**: *(number – maximum: 100)*

        Brightness percentage. Value 0 is the lowest possible brightness.

        **Example**:

        ```
        80
        ```


    **Example**:

    ```json
    {
      "brightness": 80
    }
    ```

-   **dimming\_delta**: *(object)*

    -   **action**: *required(one of up, down, stop)*

        The delta action to apply

    -   **brightness\_delta**: *(number – maximum: 100)*

        Brightness percentage of full-scale increase delta to current dimlevel. Clip at Max-level or Min-level.

        **Example**:

        ```
        20
        ```


    **Example**:

    ```json
    {
      "action": "up",
      "brightness_delta": 10.5
    }
    ```

-   **color\_temperature**: *(object)*

    Joined color temperature control

    -   **mirek**: *(integer – minimum: 50 – maximum: 1000)*

        color temperature in mirek or null when the light color is not in the ct spectrum

        **Example**:

        ```
        233
        ```


    **Example**:

    ```json
    {
      "mirek": 202
    }
    ```

-   **color\_temperature\_delta**: *(object)*

    -   **action**: *required(one of up, down, stop)*

        The delta action to apply

    -   **mirek\_delta**: *(integer – maximum: 950)*

        Mirek delta to current mirek. Clip at mirek\_minimum and mirek\_maximum of mirek\_schema.

        **Example**:

        ```
        10
        ```


    **Example**:

    ```json
    {
      "action": "up",
      "mirek_delta": 200
    }
    ```

-   **color**: *(object)*

    Joined color control

    -   **xy**: *(object)*

        CIE XY gamut position

        -   **x**: *required(number – minimum: 0 – maximum: 1)*

            X position in color gamut

        -   **y**: *required(number – minimum: 0 – maximum: 1)*

            Y position in color gamut


        **Example**:

        ```json
        {
          "x": 0.369,
          "y": 0.445
        }
        ```


    **Example**:

    ```json
    {
      "xy": {
        "x": 0.6915,
        "y": 0.3083
      }
    }
    ```

-   **alert**: *(object)*

    Joined alert control

    -   **action**: *required(breathe)*

        Alert to set the light to

-   **signaling**: *(object)*

    Feature containing basic signaling properties.

    -   **signal**: *required(one of no\_signal, on\_off, on\_off\_color, alternating)*

        Signal to set the light to

    -   **duration**: *required(integer)*

        Duration in milliseconds. Maximum value is 65534000 ms and a stepsize of 1 second. Values inbetween steps will be rounded. Duration is ignored for no\_signal.

        **Example**:

        ```
        800
        ```

    -   **colors**: *(array of ColorFeatureBasicPut – minItems: 1 – maxItems: 2)*

        List of colors to apply to the signal (not supported by all signals)

        **Items**: ColorFeatureBasicPut

        -   **xy**: *(object)*

            CIE XY gamut position

            -   **x**: *required(number – minimum: 0 – maximum: 1)*

                X position in color gamut

            -   **y**: *required(number – minimum: 0 – maximum: 1)*

                Y position in color gamut


            **Example**:

            ```json
            {
              "x": 0.369,
              "y": 0.445
            }
            ```


        **Example**:

        ```json
        {
          "xy": {
            "x": 0.6915,
            "y": 0.3083
          }
        }
        ```

-   **dynamics**: *(object)*
    -   **duration**: *(integer)*

        Duration of a light transition in ms.

        **Example**:

        ```
        800
        ```


## HTTP status code 200

Request was on resource path was successful.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 207

Request was partially successful, resource has been updated. Failed requests and errors are listed in error object.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 400

Request was on resource path was unsuccessful, body of request contains errors.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of GroupedLightGet)*

    **Items**: GroupedLightGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **owner**: *required(object)*

        Owner of the service, in case the owner service is deleted, the service also gets deleted

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **type**: *required(grouped\_light)*
    -   **on**: *(object)*

        Joined on control & aggregated on state.| – “on” is true if any light in the group is on

        -   **on**: *required(boolean)*

            On/Off state of the light on=true, off=false

    -   **dimming**: *(object)*

        Joined dimming control – “dimming.brightness” contains average brightness of group containing turned-on lights only.

        -   **brightness**: *required(number – maximum: 100)*

            Brightness percentage. Value 0 is the lowest possible brightness.

            **Example**:

            ```
            80
            ```


        **Example**:

        ```json
        {
          "brightness": 80
        }
        ```

    -   **dimming\_delta**: *(object)*

        **Example**:

        ```json
        {
          "action": "up",
          "brightness_delta": 10.5
        }
        ```

    -   **color\_temperature**: *(object)*

        Joined color temperature control

        **Example**:

        ```json
        {
          "mirek": 202
        }
        ```

    -   **color\_temperature\_delta**: *(object)*

        **Example**:

        ```json
        {
          "action": "up",
          "mirek_delta": 200
        }
        ```

    -   **color**: *(object)*

        Joined color control

        **Example**:

        ```json
        {
          "xy": {
            "x": 0.6915,
            "y": 0.3083
          }
        }
        ```

    -   **alert**: *(object)*

        Joined alert control

        -   **action\_values**: *required(array of AlertEffectType)*

            Alert effects that the light supports.

    -   **signaling**: *(object)*

        Feature containing basic signaling properties.

        -   **signal\_values**: *required(array of SupportedSignals)*

            Signals that the light supports.

    -   **dynamics**: *(object)*


## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### /resource/device

API to manage devices. Devices have device level properties and offer services such as light. Bridge device cannot be deleted.

get

List all resources of this type

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of DeviceGet)*

    **Items**: DeviceGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **type**: *required(device)*
    -   **product\_data**: *required(object)*

        -   **model\_id**: *required(string)*

            unique identification of device model

        -   **manufacturer\_name**: *required(string)*

            Name of device manufacturer

        -   **product\_name**: *required(string)*

            Name of the product.

        -   **product\_archetype**: *required(one of bridge\_v2, bridge\_v3, unknown\_archetype, classic\_bulb, sultan\_bulb, flood\_bulb, spot\_bulb, candle\_bulb, luster\_bulb, pendant\_round, pendant\_long, ceiling\_round, ceiling\_square, floor\_shade, floor\_lantern, table\_shade, recessed\_ceiling, recessed\_floor, single\_spot, double\_spot, table\_wash, wall\_lantern, wall\_shade, flexible\_lamp, ground\_spot, wall\_spot, plug, hue\_go, hue\_lightstrip, hue\_iris, hue\_bloom, bollard, wall\_washer, hue\_play, hue\_chime, vintage\_bulb, vintage\_candle\_bulb, ellipse\_bulb, triangle\_bulb, small\_globe\_bulb, large\_globe\_bulb, edison\_bulb, christmas\_tree, string\_light, hue\_centris, hue\_lightstrip\_tv, hue\_lightstrip\_pc, hue\_tube, hue\_signe, pendant\_spot, ceiling\_horizontal, ceiling\_tube, up\_and\_down, up\_and\_down\_up, up\_and\_down\_down, hue\_floodlight\_camera, twilight, twilight\_front, twilight\_back, hue\_play\_wallwasher, hue\_omniglow, hue\_neon, string\_globe, string\_permanent)*

            Archetype of the product

        -   **certified**: *required(boolean)*

            This device is Hue certified

        -   **software\_version**: *required(string – pattern: \\d+\\.\\d+\\.\\d+)*

            Software version of the product

        -   **hardware\_platform\_type**: *(string)*

            Hardware type; identified by Manufacturer code and ImageType


        **Example**:

        ```json
        {
          "model_id": "LCB001",
          "manufacturer_name": "Signify Netherlands B.V.",
          "product_name": "Hue color downlight",
          "product_archetype": "flood_bulb",
          "certified": true,
          "software_version": "1.122.8",
          "hardware_platform_type": "100b-112"
        }
        ```

    -   **metadata**: *required(object)*

        additional metadata including a user given name

        -   **name**: *required(string – minLength: 1 – maxLength: 32)*

            Human readable name of a resource

        -   **archetype**: *required(one of bridge\_v2, bridge\_v3, unknown\_archetype, classic\_bulb, sultan\_bulb, flood\_bulb, spot\_bulb, candle\_bulb, luster\_bulb, pendant\_round, pendant\_long, ceiling\_round, ceiling\_square, floor\_shade, floor\_lantern, table\_shade, recessed\_ceiling, recessed\_floor, single\_spot, double\_spot, table\_wash, wall\_lantern, wall\_shade, flexible\_lamp, ground\_spot, wall\_spot, plug, hue\_go, hue\_lightstrip, hue\_iris, hue\_bloom, bollard, wall\_washer, hue\_play, hue\_chime, vintage\_bulb, vintage\_candle\_bulb, ellipse\_bulb, triangle\_bulb, small\_globe\_bulb, large\_globe\_bulb, edison\_bulb, christmas\_tree, string\_light, hue\_centris, hue\_lightstrip\_tv, hue\_lightstrip\_pc, hue\_tube, hue\_signe, pendant\_spot, ceiling\_horizontal, ceiling\_tube, up\_and\_down, up\_and\_down\_up, up\_and\_down\_down, hue\_floodlight\_camera, twilight, twilight\_front, twilight\_back, hue\_play\_wallwasher, hue\_omniglow, hue\_neon, string\_globe, string\_permanent)*

            By default archetype given by manufacturer. Can be changed by user.

    -   **identify**: *required(object)*
    -   **usertest**: *(object)*
        -   **status**: *required(one of set, changing)*
        -   **usertest**: *required(boolean)*

            Activates or extends user usertest mode of device for 120 seconds. false deactivates usertest mode. In usertest mode, devices report changes in state faster and indicate state changes on device LED (if applicable)

    -   **device\_mode**: *(object)*
        -   **status**: *required(one of set, changing)*
        -   **mode**: *required(one of switch\_single\_rocker, switch\_single\_pushbutton, switch\_dual\_rocker, switch\_dual\_pushbutton)*

            current mode (on read) or requested mode (on write) of the device

        -   **mode\_values**: *required(array of DeviceModeType)*

            the modes that the device supports

    -   **services**: *required(array of ResourceIdentifier)*

        References all services providing control and state of the device.

        **Items**: ResourceIdentifier

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource




## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 406

Query parameter has invalid value.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


API calls to read, update and delete a resource instance

delete

put

get

### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Resource has been deleted

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### URI Parameters

-   **id**: *required(string)*

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **type**: *(device)*
-   **metadata**: *(object)*

    additional metadata including a user given name

    -   **name**: *(string – minLength: 1 – maxLength: 32)*

        Human readable name of a resource

    -   **archetype**: *(one of bridge\_v2, bridge\_v3, unknown\_archetype, classic\_bulb, sultan\_bulb, flood\_bulb, spot\_bulb, candle\_bulb, luster\_bulb, pendant\_round, pendant\_long, ceiling\_round, ceiling\_square, floor\_shade, floor\_lantern, table\_shade, recessed\_ceiling, recessed\_floor, single\_spot, double\_spot, table\_wash, wall\_lantern, wall\_shade, flexible\_lamp, ground\_spot, wall\_spot, plug, hue\_go, hue\_lightstrip, hue\_iris, hue\_bloom, bollard, wall\_washer, hue\_play, hue\_chime, vintage\_bulb, vintage\_candle\_bulb, ellipse\_bulb, triangle\_bulb, small\_globe\_bulb, large\_globe\_bulb, edison\_bulb, christmas\_tree, string\_light, hue\_centris, hue\_lightstrip\_tv, hue\_lightstrip\_pc, hue\_tube, hue\_signe, pendant\_spot, ceiling\_horizontal, ceiling\_tube, up\_and\_down, up\_and\_down\_up, up\_and\_down\_down, hue\_floodlight\_camera, twilight, twilight\_front, twilight\_back, hue\_play\_wallwasher, hue\_omniglow, hue\_neon, string\_globe, string\_permanent)*

        By default archetype given by manufacturer. Can be changed by user.

-   **identify**: *(object)*
    -   **action**: *required(identify)*
        -   identify: Triggers a visual identification sequence, current implemented as (which can change in the future): Bridge performs Zigbee LED identification cycles for 5 seconds Lights perform one breathe cycle Sensors perform LED identification cycles for 15 seconds
    -   **duration**: *(integer)*

        Duration in milliseconds to perform the identify cycle.

        **Example**:

        ```json
        800
        ```

-   **usertest**: *(object)*
    -   **usertest**: *(boolean)*

        Activates or extends user usertest mode of device for 120 seconds. false deactivates usertest mode. In usertest mode, devices report changes in state faster and indicate state changes on device LED (if applicable)

-   **device\_mode**: *(object)*
    -   **mode**: *required(one of switch\_single\_rocker, switch\_single\_pushbutton, switch\_dual\_rocker, switch\_dual\_pushbutton)*

        current mode (on read) or requested mode (on write) of the device


## HTTP status code 200

Request was on resource path was successful.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 207

Request was partially successful, resource has been updated. Failed requests and errors are listed in error object.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 400

Request was on resource path was unsuccessful, body of request contains errors.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of DeviceGet)*

    **Items**: DeviceGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **type**: *required(device)*
    -   **product\_data**: *required(object)*

        -   **model\_id**: *required(string)*

            unique identification of device model

        -   **manufacturer\_name**: *required(string)*

            Name of device manufacturer

        -   **product\_name**: *required(string)*

            Name of the product.

        -   **product\_archetype**: *required(one of bridge\_v2, bridge\_v3, unknown\_archetype, classic\_bulb, sultan\_bulb, flood\_bulb, spot\_bulb, candle\_bulb, luster\_bulb, pendant\_round, pendant\_long, ceiling\_round, ceiling\_square, floor\_shade, floor\_lantern, table\_shade, recessed\_ceiling, recessed\_floor, single\_spot, double\_spot, table\_wash, wall\_lantern, wall\_shade, flexible\_lamp, ground\_spot, wall\_spot, plug, hue\_go, hue\_lightstrip, hue\_iris, hue\_bloom, bollard, wall\_washer, hue\_play, hue\_chime, vintage\_bulb, vintage\_candle\_bulb, ellipse\_bulb, triangle\_bulb, small\_globe\_bulb, large\_globe\_bulb, edison\_bulb, christmas\_tree, string\_light, hue\_centris, hue\_lightstrip\_tv, hue\_lightstrip\_pc, hue\_tube, hue\_signe, pendant\_spot, ceiling\_horizontal, ceiling\_tube, up\_and\_down, up\_and\_down\_up, up\_and\_down\_down, hue\_floodlight\_camera, twilight, twilight\_front, twilight\_back, hue\_play\_wallwasher, hue\_omniglow, hue\_neon, string\_globe, string\_permanent)*

            Archetype of the product

        -   **certified**: *required(boolean)*

            This device is Hue certified

        -   **software\_version**: *required(string – pattern: \\d+\\.\\d+\\.\\d+)*

            Software version of the product

        -   **hardware\_platform\_type**: *(string)*

            Hardware type; identified by Manufacturer code and ImageType


        **Example**:

        ```json
        {
          "model_id": "LCB001",
          "manufacturer_name": "Signify Netherlands B.V.",
          "product_name": "Hue color downlight",
          "product_archetype": "flood_bulb",
          "certified": true,
          "software_version": "1.122.8",
          "hardware_platform_type": "100b-112"
        }
        ```

    -   **metadata**: *required(object)*

        additional metadata including a user given name

        -   **name**: *required(string – minLength: 1 – maxLength: 32)*

            Human readable name of a resource

        -   **archetype**: *required(one of bridge\_v2, bridge\_v3, unknown\_archetype, classic\_bulb, sultan\_bulb, flood\_bulb, spot\_bulb, candle\_bulb, luster\_bulb, pendant\_round, pendant\_long, ceiling\_round, ceiling\_square, floor\_shade, floor\_lantern, table\_shade, recessed\_ceiling, recessed\_floor, single\_spot, double\_spot, table\_wash, wall\_lantern, wall\_shade, flexible\_lamp, ground\_spot, wall\_spot, plug, hue\_go, hue\_lightstrip, hue\_iris, hue\_bloom, bollard, wall\_washer, hue\_play, hue\_chime, vintage\_bulb, vintage\_candle\_bulb, ellipse\_bulb, triangle\_bulb, small\_globe\_bulb, large\_globe\_bulb, edison\_bulb, christmas\_tree, string\_light, hue\_centris, hue\_lightstrip\_tv, hue\_lightstrip\_pc, hue\_tube, hue\_signe, pendant\_spot, ceiling\_horizontal, ceiling\_tube, up\_and\_down, up\_and\_down\_up, up\_and\_down\_down, hue\_floodlight\_camera, twilight, twilight\_front, twilight\_back, hue\_play\_wallwasher, hue\_omniglow, hue\_neon, string\_globe, string\_permanent)*

            By default archetype given by manufacturer. Can be changed by user.

    -   **identify**: *required(object)*
    -   **usertest**: *(object)*
        -   **status**: *required(one of set, changing)*
        -   **usertest**: *required(boolean)*

            Activates or extends user usertest mode of device for 120 seconds. false deactivates usertest mode. In usertest mode, devices report changes in state faster and indicate state changes on device LED (if applicable)

    -   **device\_mode**: *(object)*
        -   **status**: *required(one of set, changing)*
        -   **mode**: *required(one of switch\_single\_rocker, switch\_single\_pushbutton, switch\_dual\_rocker, switch\_dual\_pushbutton)*

            current mode (on read) or requested mode (on write) of the device

        -   **mode\_values**: *required(array of DeviceModeType)*

            the modes that the device supports

    -   **services**: *required(array of ResourceIdentifier)*

        References all services providing control and state of the device.

        **Items**: ResourceIdentifier

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource




## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### /resource/bridge

API to manage the bridge

get

List all resources of this type

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of BridgeGet)*

    **Items**: BridgeGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **owner**: *required(object)*

        Owner of the service, in case the owner service is deleted, the service also gets deleted

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **type**: *required(bridge)*

        Type of the supported resources

    -   **bridge\_id**: *required(string)*

        Unique identifier of the bridge as printed on the device. Lower case.

    -   **time\_zone**: *required(object)*
        -   **time\_zone**: *required(string)*

            Time zone where the user's home is located (as Olson ID).

    -   **import**: *(object)*
        -   **origin**: *required(string)*

            Bridge ID (in lower case) where the imported data originates from.

        -   **time**: *required(datetime)*

            UTC date and time when the import took place.


    **Example**:

    ```json
    {
      "type": "bridge",
      "id": "91a6626c-8b8f-4c4a-a9e8-18a6be332373",
      "owner": {
        "rid": "08221fb4-ce0c-4914-845d-50aa5033f126",
        "rtype": "device"
      },
      "bridge_id": "001788fffe21055d",
      "time_zone": {
        "time_zone": "Europe/Brussels"
      },
      "import": {
        "origin": "c42996fffec02aea",
        "time": "2024-11-28T10:19:33.637Z"
      }
    }
    ```


## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 406

Query parameter has invalid value.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


API calls for resource which cannot be deleted (PUT and GET on resource instance)

put

get

### URI Parameters

-   **id**: *required(string)*

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **type**: *(bridge)*

    Type of the supported resources


**Example**:

```json
{
  "type": "bridge",
  "id": "91a6626c-8b8f-4c4a-a9e8-18a6be332373",
  "owner": {
    "rid": "08221fb4-ce0c-4914-845d-50aa5033f126",
    "rtype": "device"
  },
  "bridge_id": "001788fffe21055d",
  "time_zone": {
    "time_zone": "Europe/Brussels"
  },
  "import": {
    "origin": "c42996fffec02aea",
    "time": "2024-11-28T10:19:33.637Z"
  }
}
```

## HTTP status code 200

Request was on resource path was successful.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 207

Request was partially successful, resource has been updated. Failed requests and errors are listed in error object.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 400

Request was on resource path was unsuccessful, body of request contains errors.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of BridgeGet)*

    **Items**: BridgeGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **owner**: *required(object)*

        Owner of the service, in case the owner service is deleted, the service also gets deleted

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **type**: *required(bridge)*

        Type of the supported resources

    -   **bridge\_id**: *required(string)*

        Unique identifier of the bridge as printed on the device. Lower case.

    -   **time\_zone**: *required(object)*
        -   **time\_zone**: *required(string)*

            Time zone where the user's home is located (as Olson ID).

    -   **import**: *(object)*
        -   **origin**: *required(string)*

            Bridge ID (in lower case) where the imported data originates from.

        -   **time**: *required(datetime)*

            UTC date and time when the import took place.


    **Example**:

    ```json
    {
      "type": "bridge",
      "id": "91a6626c-8b8f-4c4a-a9e8-18a6be332373",
      "owner": {
        "rid": "08221fb4-ce0c-4914-845d-50aa5033f126",
        "rtype": "device"
      },
      "bridge_id": "001788fffe21055d",
      "time_zone": {
        "time_zone": "Europe/Brussels"
      },
      "import": {
        "origin": "c42996fffec02aea",
        "time": "2024-11-28T10:19:33.637Z"
      }
    }
    ```


## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### /resource/device_software_update

API to manage device update services. These services are present for devices that support software update.

### /resource/device_power

API to manage device power services. These are offered by battery powered devices.

### /resource/zigbee_connectivity

API to manage zigbee connectivity services. These are offered by zigbee connected devices.

### /resource/zgp_connectivity

API to manage zgp connectivity services. These are offered by zigbee green power devices.

### /resource/zigbee_device_discovery

API to manage zigbee device discovery service. This is offered by the bridge for commissioning.

### /resource/motion

API to manage motion services. These are offered by devices with motion sensing capabilities.

### /resource/service_group

API to manage service group services. These are services that aggregate their distinct children into grouped services. For example, adding one or more motion services as children will result in a grouped\_motion service that aggregates the motions service report.

### /resource/grouped_motion

API to manage grouped motion services. These are offered by service groups with archetype “sensor\_group”.

### /resource/grouped_light_level

API to manage grouped light-level services. These are offered by service groups with archetype “sensor\_group”.

### /resource/camera_motion

API to manage camera\_motion services. These are offered by devices with camera based motion sensing capabilities.

### /resource/temperature

API to manage temperature services. These are offered by devices with temperature sensing capabilities.

### /resource/light_level

API to manage light level services. These are offered by devices with light level sensing capabilities.

### /resource/button

API to manage button services. These are offered by devices with buttons.

get

List all resources of this type

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ButtonGet)*

    **Items**: ButtonGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **owner**: *required(object)*

        Owner of the service, in case the owner service is deleted, the service also gets deleted

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **metadata**: *required(object)*

        Metadata describing this resource

        -   **control\_id**: *required(integer – minimum: 0 – maximum: 8)*

            control identifier of the switch which is unique per device. Meaning in combination with type – dots Number of dots – number Number printed on device – other a logical order of controls in switch

    -   **button**: *required(object)*
        -   **last\_event**: *(one of initial\_press, repeat, short\_release, long\_release, double\_short\_release, long\_press)*

            Deprecated. Move to button\_report/event

        -   **button\_report**: *(object)*
            -   **updated**: *required(datetime)*

                last time the value of this property is updated.

            -   **event**: *required(one of initial\_press, repeat, short\_release, long\_release, double\_short\_release, long\_press)*

                event which can be send by a button control

        -   **repeat\_interval**: *(integer)*

            Duration between repeat events when holding the button in milliseconds.

            **Example**:

            ```
            800
            ```

        -   **event\_values**: *(array of ButtonEvent)*

            list of all button events that this device supports

    -   **type**: *required(button)*

        Type of the supported resources



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 406

Query parameter has invalid value.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


API calls for resource which cannot be deleted (PUT and GET on resource instance)

put

get

### URI Parameters

-   **id**: *required(string)*

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **type**: *(button)*

    Type of the supported resources


## HTTP status code 200

Request was on resource path was successful.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 207

Request was partially successful, resource has been updated. Failed requests and errors are listed in error object.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 400

Request was on resource path was unsuccessful, body of request contains errors.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ButtonGet)*

    **Items**: ButtonGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **owner**: *required(object)*

        Owner of the service, in case the owner service is deleted, the service also gets deleted

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **metadata**: *required(object)*

        Metadata describing this resource

        -   **control\_id**: *required(integer – minimum: 0 – maximum: 8)*

            control identifier of the switch which is unique per device. Meaning in combination with type – dots Number of dots – number Number printed on device – other a logical order of controls in switch

    -   **button**: *required(object)*
        -   **last\_event**: *(one of initial\_press, repeat, short\_release, long\_release, double\_short\_release, long\_press)*

            Deprecated. Move to button\_report/event

        -   **button\_report**: *(object)*
            -   **updated**: *required(datetime)*

                last time the value of this property is updated.

            -   **event**: *required(one of initial\_press, repeat, short\_release, long\_release, double\_short\_release, long\_press)*

                event which can be send by a button control

        -   **repeat\_interval**: *(integer)*

            Duration between repeat events when holding the button in milliseconds.

            **Example**:

            ```
            800
            ```

        -   **event\_values**: *(array of ButtonEvent)*

            list of all button events that this device supports

    -   **type**: *required(button)*

        Type of the supported resources



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### /resource/bell_button

API to manage button services. These are offered by devices with buttons.

get

List all resources of this type

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of BellButtonGet)*

    **Items**: BellButtonGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **owner**: *required(object)*

        Owner of the service, in case the owner service is deleted, the service also gets deleted

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **metadata**: *required(object)*

        Metadata describing this resource

        -   **control\_id**: *required(integer – minimum: 0 – maximum: 8)*

            control identifier of the switch which is unique per device. Meaning in combination with type – dots Number of dots – number Number printed on device – other a logical order of controls in switch

    -   **button**: *required(object)*
        -   **last\_event**: *(one of initial\_press, repeat, short\_release, long\_release, double\_short\_release, long\_press)*

            Deprecated. Move to button\_report/event

        -   **button\_report**: *(object)*
            -   **updated**: *required(datetime)*

                last time the value of this property is updated.

            -   **event**: *required(one of initial\_press, repeat, short\_release, long\_release, double\_short\_release, long\_press)*

                event which can be send by a button control

        -   **repeat\_interval**: *(integer)*

            Duration between repeat events when holding the button in milliseconds.

            **Example**:

            ```
            800
            ```

        -   **event\_values**: *(array of ButtonEvent)*

            list of all button events that this device supports

    -   **type**: *required(bell\_button)*

        Type of the supported resources



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 406

Query parameter has invalid value.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


API calls for resource which cannot be deleted (PUT and GET on resource instance)

put

get

### URI Parameters

-   **id**: *required(string)*

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **type**: *(bell\_button)*

    Type of the supported resources


## HTTP status code 200

Request was on resource path was successful.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 207

Request was partially successful, resource has been updated. Failed requests and errors are listed in error object.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 400

Request was on resource path was unsuccessful, body of request contains errors.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of BellButtonGet)*

    **Items**: BellButtonGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **owner**: *required(object)*

        Owner of the service, in case the owner service is deleted, the service also gets deleted

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **metadata**: *required(object)*

        Metadata describing this resource

        -   **control\_id**: *required(integer – minimum: 0 – maximum: 8)*

            control identifier of the switch which is unique per device. Meaning in combination with type – dots Number of dots – number Number printed on device – other a logical order of controls in switch

    -   **button**: *required(object)*
        -   **last\_event**: *(one of initial\_press, repeat, short\_release, long\_release, double\_short\_release, long\_press)*

            Deprecated. Move to button\_report/event

        -   **button\_report**: *(object)*
            -   **updated**: *required(datetime)*

                last time the value of this property is updated.

            -   **event**: *required(one of initial\_press, repeat, short\_release, long\_release, double\_short\_release, long\_press)*

                event which can be send by a button control

        -   **repeat\_interval**: *(integer)*

            Duration between repeat events when holding the button in milliseconds.

            **Example**:

            ```
            800
            ```

        -   **event\_values**: *(array of ButtonEvent)*

            list of all button events that this device supports

    -   **type**: *required(bell\_button)*

        Type of the supported resources



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### /resource/relative_rotary

API to manage relative rotary services. These are offered by devices with rotary capabilities.

### /resource/behavior_script

API to discover available scripts that can be instantiated

### /resource/behavior_instance

API to manage instances of script

### /resource/geofence_client

API for geofencing functionality

### /resource/geolocation

API for setting the geolocation

### /resource/entertainment_configuration

API to manage entertainment configurations, used for Hue Entertainment functionality

### /resource/entertainment

API to manage entertainment services. These are offered by devices with color lighting capabilities.

### /resource/homekit

API to manage homekit service

### /resource/matter

API to manage matter service

get

List all resources of this type

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of MatterGet)*

    **Items**: MatterGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **type**: *required(matter)*

        Type of the supported resources

    -   **max\_fabrics**: *required(integer)*

        Maximum number of fabrics that can exist at a time

    -   **has\_qr\_code**: *required(boolean)*

        Indicates whether a physical QR code is present

    -   **software\_version\_string**: *required(string – pattern: ^\[0-9\]\[A-Za-z0-9\]\*(?:.(?:\[A-Za-z0-9\]+(?:\[-\_\]\[A-Za-z0-9\]+)?)){1,2}$)*

        Indicates the software version of the matter daemon

        **Example**:

        ```
        1.5.0
        ```



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 406

Query parameter has invalid value.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


API calls for resource which cannot be deleted (PUT and GET on resource instance)

put

get

### URI Parameters

-   **id**: *required(string)*

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **type**: *(matter)*

    Type of the supported resources

-   **action**: *required(matter\_reset)*

    matter\_reset: Resets Matter, including removing all fabrics and reset state to factory settings


## HTTP status code 200

Request was on resource path was successful.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 207

Request was partially successful, resource has been updated. Failed requests and errors are listed in error object.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 400

Request was on resource path was unsuccessful, body of request contains errors.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of MatterGet)*

    **Items**: MatterGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **type**: *required(matter)*

        Type of the supported resources

    -   **max\_fabrics**: *required(integer)*

        Maximum number of fabrics that can exist at a time

    -   **has\_qr\_code**: *required(boolean)*

        Indicates whether a physical QR code is present

    -   **software\_version\_string**: *required(string – pattern: ^\[0-9\]\[A-Za-z0-9\]\*(?:.(?:\[A-Za-z0-9\]+(?:\[-\_\]\[A-Za-z0-9\]+)?)){1,2}$)*

        Indicates the software version of the matter daemon

        **Example**:

        ```
        1.5.0
        ```



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### /resource/matter_fabric

API to manage matter fabrics

### /resource/smart_scene

API to manage smart scenes

post

create a new resource of this type

get

List all resources of this type

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **type**: *(smart\_scene)*

    Type of the supported resources

-   **metadata**: *required(object)*
    -   **name**: *required(string – minLength: 1 – maxLength: 32)*

        Human readable name of a resource

    -   **image**: *(object)*

        Reference with unique identifier for the image representing the scene only accepting “rtype”: “public\_image” on creation

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **appdata**: *(string – minLength: 1 – maxLength: 16)*

        Application specific data. Free format string.

-   **group**: *required(object)*

    Group associated with this Scene. All services in the group are part of this scene. If the group is changed the scene is update (e.g. light added/removed)

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource

-   **week\_timeslots**: *required(array of DayTimeslotsPost)*

    information on what is the light state for every timeslot of the day

    **Items**: DayTimeslotsPost

    -   **timeslots**: *required(array of SmartSceneTimeslotPost)*

        **Items**: SmartSceneTimeslotPost

        -   **start\_time**: *required(object)*
            -   **kind**: *required(one of time, sunset)*
            -   **time**: *(object)*

                this property is only used when property “kind” is “time”

                -   **hour**: *required(integer – minimum: 0 – maximum: 23)*
                -   **minute**: *required(integer – minimum: 0 – maximum: 59)*
                -   **second**: *required(integer – minimum: 0 – maximum: 59)*
        -   **target**: *required(object)*

            The identifier of the scene to recall

            -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

                The unique id of the referenced resource

            -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

                The type of the referenced resource


    -   **recurrence**: *required(array of WeekDay)*

-   **transition\_duration**: *(integer)*

    duration of the transition from on one timeslot's scene to the other (defaults to 60000ms)

    **Example**:

    ```
    800
    ```

-   **recall**: *(object)*
    -   **action**: *required(one of activate, deactivate)*

        Activate will start the smart (24h) scene; deactivate will stop it


## HTTP status code 201

Resource has been created

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 400

Request was on resource path was unsuccessful, body of request contains errors.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of SmartSceneGet)*

    **Items**: SmartSceneGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **type**: *required(smart\_scene)*

        Type of the supported resources

    -   **metadata**: *required(object)*
        -   **name**: *required(string – minLength: 1 – maxLength: 32)*

            Human readable name of a resource

        -   **image**: *(object)*

            Reference with unique identifier for the image representing the scene only accepting “rtype”: “public\_image” on creation

            -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

                The unique id of the referenced resource

            -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

                The type of the referenced resource

        -   **appdata**: *(string – minLength: 1 – maxLength: 16)*

            Application specific data. Free format string.

    -   **group**: *required(object)*

        Group associated with this Scene. All services in the group are part of this scene. If the group is changed the scene is update (e.g. light added/removed)

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **week\_timeslots**: *required(array of DayTimeslotsGet)*

        information on what is the light state for every timeslot of the day

        **Items**: DayTimeslotsGet

        -   **timeslots**: *required(array of SmartSceneTimeslotGet)*

            **Items**: SmartSceneTimeslotGet

            -   **start\_time**: *required(object)*
                -   **kind**: *required(one of time, sunset)*
                -   **time**: *required(object)*

                    this property is only used when property “kind” is “time”

                    -   **hour**: *required(integer – minimum: 0 – maximum: 23)*
                    -   **minute**: *required(integer – minimum: 0 – maximum: 59)*
                    -   **second**: *required(integer – minimum: 0 – maximum: 59)*
            -   **target**: *required(object)*

                The identifier of the scene to recall

                -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

                    The unique id of the referenced resource

                -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

                    The type of the referenced resource


        -   **recurrence**: *required(array of WeekDay)*

    -   **transition\_duration**: *required(integer)*

        duration of the transition from on one timeslot's scene to the other (defaults to 60000ms)

        **Example**:

        ```
        800
        ```

    -   **active\_timeslot**: *(object)*

        the active time slot in execution

        -   **timeslot\_id**: *required(integer)*
        -   **weekday**: *required(one of monday, tuesday, wednesday, thursday, friday, saturday, sunday)*
    -   **state**: *required(one of active, inactive)*

        the current state of the smart scene. The default state is `inactive` if no `recall` is provided



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 406

Query parameter has invalid value.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


API calls to read, update and delete a resource instance

delete

put

get

### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Resource has been deleted

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### URI Parameters

-   **id**: *required(string)*

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **type**: *(smart\_scene)*

    Type of the supported resources

-   **metadata**: *(object)*
    -   **name**: *(string – minLength: 1 – maxLength: 32)*

        Human readable name of a resource

    -   **appdata**: *(string – minLength: 1 – maxLength: 16)*

        Application specific data. Free format string.

-   **week\_timeslots**: *(array of DayTimeslotsPut)*

    information on what is the light state for every timeslot of the day

    **Items**: DayTimeslotsPut

    -   **timeslots**: *required(array of SmartSceneTimeslotPut)*

        **Items**: SmartSceneTimeslotPut

        -   **start\_time**: *required(object)*
            -   **kind**: *required(one of time, sunset)*
            -   **time**: *(object)*

                this property is only used when property “kind” is “time”

                -   **hour**: *required(integer – minimum: 0 – maximum: 23)*
                -   **minute**: *required(integer – minimum: 0 – maximum: 59)*
                -   **second**: *(integer – minimum: 0 – maximum: 59)*
        -   **target**: *required(object)*

            The identifier of the scene to recall

            -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

                The unique id of the referenced resource

            -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

                The type of the referenced resource


    -   **recurrence**: *required(array of WeekDay)*

-   **transition\_duration**: *(integer)*

    duration of the transition from on one timeslot's scene to the other (defaults to 60000ms)

    **Example**:

    ```
    800
    ```

-   **recall**: *(object)*
    -   **action**: *(one of activate, deactivate)*

        Activate will start the smart (24h) scene; deactivate will stop it


## HTTP status code 200

Request was on resource path was successful.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 207

Request was partially successful, resource has been updated. Failed requests and errors are listed in error object.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of ResourceIdentifier)*

    **Items**: ResourceIdentifier

    -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        The unique id of the referenced resource

    -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

        The type of the referenced resource



## HTTP status code 400

Request was on resource path was unsuccessful, body of request contains errors.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### URI Parameters

-   **id**: *required(string)*

## HTTP status code 200

Request was on resource path was successful. Result of request is in data.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.


-   **data**: *required(array of SmartSceneGet)*

    **Items**: SmartSceneGet

    -   **id**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

        Unique identifier representing a specific resource instance

    -   **id\_v1**: *(string – pattern: ^(\\/\[a-z\]{4,32}\\/\[0-9a-zA-Z-\]{1,32})?$)*

        Clip v1 resource identifier

    -   **type**: *required(smart\_scene)*

        Type of the supported resources

    -   **metadata**: *required(object)*
        -   **name**: *required(string – minLength: 1 – maxLength: 32)*

            Human readable name of a resource

        -   **image**: *(object)*

            Reference with unique identifier for the image representing the scene only accepting “rtype”: “public\_image” on creation

            -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

                The unique id of the referenced resource

            -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

                The type of the referenced resource

        -   **appdata**: *(string – minLength: 1 – maxLength: 16)*

            Application specific data. Free format string.

    -   **group**: *required(object)*

        Group associated with this Scene. All services in the group are part of this scene. If the group is changed the scene is update (e.g. light added/removed)

        -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

            The unique id of the referenced resource

        -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

            The type of the referenced resource

    -   **week\_timeslots**: *required(array of DayTimeslotsGet)*

        information on what is the light state for every timeslot of the day

        **Items**: DayTimeslotsGet

        -   **timeslots**: *required(array of SmartSceneTimeslotGet)*

            **Items**: SmartSceneTimeslotGet

            -   **start\_time**: *required(object)*
                -   **kind**: *required(one of time, sunset)*
                -   **time**: *required(object)*

                    this property is only used when property “kind” is “time”

                    -   **hour**: *required(integer – minimum: 0 – maximum: 23)*
                    -   **minute**: *required(integer – minimum: 0 – maximum: 59)*
                    -   **second**: *required(integer – minimum: 0 – maximum: 59)*
            -   **target**: *required(object)*

                The identifier of the scene to recall

                -   **rid**: *required(string – pattern: ^\[0-9a-f\]{8}-(\[0-9a-f\]{4}-){3}\[0-9a-f\]{12}$)*

                    The unique id of the referenced resource

                -   **rtype**: *required(one of device, bridge\_home, room, zone, service\_group, light, button, bell\_button, relative\_rotary, temperature, light\_level, motion, camera\_motion, entertainment, contact, tamper, convenience\_area\_motion, security\_area\_motion, speaker, grouped\_light, grouped\_motion, grouped\_light\_level, device\_power, device\_software\_update, zigbee\_connectivity, zgp\_connectivity, bridge, motion\_area\_candidate, wifi\_connectivity, zigbee\_device\_discovery, homekit, matter, matter\_fabric, scene, entertainment\_configuration, public\_image, auth\_v1, behavior\_script, behavior\_instance, geofence\_client, geolocation, smart\_scene, motion\_area\_configuration, clip)*

                    The type of the referenced resource


        -   **recurrence**: *required(array of WeekDay)*

    -   **transition\_duration**: *required(integer)*

        duration of the transition from on one timeslot's scene to the other (defaults to 60000ms)

        **Example**:

        ```
        800
        ```

    -   **active\_timeslot**: *(object)*

        the active time slot in execution

        -   **timeslot\_id**: *required(integer)*
        -   **weekday**: *required(one of monday, tuesday, wednesday, thursday, friday, saturday, sunday)*
    -   **state**: *required(one of active, inactive)*

        the current state of the smart scene. The default state is `inactive` if no `recall` is provided



## HTTP status code 401

Client is not authorized

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 403

Client is not allowed to perform operation on this resource.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 404

Resource or resource path does not exist

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 405

Method is not supported by the resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 409

Request conflict with the current state of the target resource

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 429

Too many requests

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 500

Internal Server Error

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 503

Busy, try again later.

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## HTTP status code 507

Insufficient resources

### Body

**Media type**: application/json

**Type**: object

**Properties**

-   **errors**: *required(array of Error)*

    **Items**: Error

    -   **description**: *required(string)*

        a human-readable explanation specific to this occurrence of the problem.



## Secured by hue-application-key

### Headers

-   **hue-application-key**: *required(string)*

    API token


### /resource/contact

API to manage contact sensor state for general purpose cases. These are offered by devices that have a static “switch” position. Not necessarily a physical switch.

### /resource/tamper

API to manage device tamper state. Offered by devices capable of detecting tampering.

### /resource/motion_area_configuration

API to manage motion area's

### /resource/motion_area_candidate

API to manage motion area services of lights

### /resource/convenience_area_motion

API to motion area based motion sensors for convenience use case

### /resource/security_area_motion

API to motion area based motion sensors for security use case

### /resource/speaker

API to manage speaker services. These are offered by devices with speakers.

### /resource/clip

API to reflect public clip services

### /resource/wifi_connectivity

API to expose wifi\_connectivity status
