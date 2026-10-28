---
title: "Cloud2Cloud Getting Started"
author: "Philips Hue Developer Program"
source: "Philips Hue Developer Program"
url: "https://developers.meethue.com/develop/hue-api-v2/cloud2cloud-getting-started/"
date_saved: "2026-05-09T20:42:32.262Z"
date_published: "2022-05-05T07:49:43+00:00"
word_count: "431"
reading_time: "3 min"
description: "The Hue system uses OAuth2.0 to authorize Cloud to Cloud integrations. To get credentials, you need to log into your developer account, select your username on the top right, click on “Remote Hue API appids”, and select “Add new Remote Hue API app”. This will grant you a client id and client secret that you’ll […]"
---

# Cloud2Cloud Getting Started

The Hue system uses OAuth2.0 to authorize Cloud to Cloud integrations. To get credentials, you need to log into your developer account, select your username on the top right, click on “Remote Hue API appids”, and select “Add new Remote Hue API app”. This will grant you a client id and client secret that you’ll need in the next steps.

We start with an authorization request in a web browser.

```
GET https://api.meethue.com/v2/oauth2/authorize?client_id=<clientid>&response_type=code&state=<state>&redirect_uri=<uri>
```

The user will log in to their Philips Hue account, is requested to approve the access request, and then gets redirected to the redirect\_uri you specified when registering your application. If the user approved the access request, then the response contains an authorization code (which in the next step can be exchanged for an access token) and the state parameter as query parameters. If the user did not approve the request, the response contains an error message.

In the next step you exchange the code for a token.

POST https://api.meethue.com/v2/oauth2/token

Authorization: Basic <base64(clientid:clientsecret)>

Content-Type: application/x-www-form-urlencoded

grant\_type=authorization\_code&code=<code>&redirect\_uri=<uri>

POST https://api.meethue.com/v2/oauth2/token Authorization: Basic <base64(clientid:clientsecret)> Content-Type: application/x-www-form-urlencoded grant\_type=authorization\_code&code=<code>&redirect\_uri=<uri>

```
POST https://api.meethue.com/v2/oauth2/token
Authorization: Basic <base64(clientid:clientsecret)>
Content-Type: application/x-www-form-urlencoded
grant_type=authorization_code&code=<code>&redirect_uri=<uri>
```

Example response:

200 OK

Content-Type: application/json

{

"access\_token":"<access token>",

"expires\_in":604799,

"refresh\_token":"<refresh token>",

"token\_type":"bearer"

}

200 OK Content-Type: application/json { "access\_token":"<access token>", "expires\_in":604799, "refresh\_token":"<refresh token>", "token\_type":"bearer" }

```
200 OK
Content-Type: application/json
{
    "access_token":"<access token>",
    "expires_in":604799,
    "refresh_token":"<refresh token>",
    "token_type":"bearer"
}
```

The complete list of options for OAUTH2 including how to refresh tokens is described [here](https://developers.meethue.com/develop/hue-api/remote-authentication-oauth/).

To finalize the authorization we need to execute these two additional requests using the access\_token as a bearer token in the Authorization header:

PUT https://api.meethue.com/route/api/0/config

Authorization: Bearer <access\_token>

Content-Type: application/json

{"linkbutton":true}

PUT https://api.meethue.com/route/api/0/config Authorization: Bearer <access\_token> Content-Type: application/json {"linkbutton":true}

```
PUT https://api.meethue.com/route/api/0/config
Authorization: Bearer <access_token>
Content-Type: application/json
{"linkbutton":true}
```

POST https://api.meethue.com/route/api

Authorization: Bearer <access\_token>

Content-Type: application/json

{"devicetype":"<your-application-name>"}

POST https://api.meethue.com/route/api Authorization: Bearer <access\_token> Content-Type: application/json {"devicetype":"<your-application-name>"}

```
POST https://api.meethue.com/route/api
Authorization: Bearer <access_token>
Content-Type: application/json
{"devicetype":"<your-application-name>"}
```

This last call will return the username, which you can save and use as the application key in the API requests:

\[

{

"success": {

"username": “\*\*\*\*\*"

}

}

\]

\[ { "success": { "username": “\*\*\*\*\*" } } \]

```
[
  {
    "success": {
      "username": “*****"
    }
  }
]
```

Now you can use the full Hue API V2 described in the core concepts and API reference sections. An example API request to list all devices is the following:

<table><tbody><tr><td>Address</td><td><code>https://api.meethue.com/route/clip/v2/resource/device</code></td></tr><tr><td>Method</td><td><code>GET</code></td></tr><tr><td>Headers</td><td><code>Authorization: Bearer &lt;access_token&gt;</code><br><code>hue-application-key: &lt;username&gt;</code></td></tr></tbody></table>

From here on you can follow the getting started guide for local control from the [controlling a light](https://developers.meethue.com/develop/hue-api-v2/getting-started/#controlling-a-light) section. Just in every example replace the bridge ip address by the Cloud API base path https://api.meethue.com/route.