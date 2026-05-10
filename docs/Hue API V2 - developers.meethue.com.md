---
title: "Hue API V2"
author: "Philips Hue Developer Program"
source: "Philips Hue Developer Program"
url: "https://developers.meethue.com/develop/hue-api-v2/"
date_saved: "2026-05-09T20:44:03.141Z"
date_published: "2021-09-14T02:31:33+00:00"
word_count: "519"
reading_time: "3 min"
description: "Where to start The best place to start depends on whether you are an existing Hue developer with knowledge on the V1 API, or a completely new Hue developer. Existing developers: Carefully read the migration guide as it explains the differences between the V1 and V2 API, before checking out the complete API reference. New […]"
---

### Where to start

The best place to start depends on whether you are an existing Hue developer with knowledge on the V1 API, or a completely new Hue developer.

-   Existing developers: Carefully read the [migration guide](https://developers.meethue.com/develop/hue-api-v2/migration-guide-to-the-new-hue-api/) as it explains the differences between the V1 and V2 API, before checking out the complete [API reference](https://developers.meethue.com/develop/hue-api-v2/api-reference/).
-   New developers: Have a look at the [getting started](https://developers.meethue.com/develop/hue-api-v2/getting-started/) and [core concepts](https://developers.meethue.com/develop/hue-api-v2/core-concepts/), before exploring the complete [API reference](https://developers.meethue.com/develop/hue-api-v2/api-reference/). Also don’t forget to make use of the tips and tricks in the [application design guidance](https://developers.meethue.com/develop/application-design-guidance/).

### Maturity state

The V2 API has been released in production and is the recommended API to use for any application. Do take these notes into account:

-   Missing features: the API still has two missing features: app registration, and programmable rule engine. For app registration, you should for now keeping using the V1 API as described in our getting started guide. We do not recommend new apps to use the V1 rule engine as it will not be available long term.
-   Deprecated items: a few items that were used during our own internal development remain visible on the API even though they have since been replaced. Be careful to not use any properties that are marked as deprecated in the api reference.

### Free to publish

Signify (the company behind Philips Hue) has a quite progressive policy with Hue which will remain unchanged for the V2 API. As you are free to create with our product, we think it should also be you who profits from your work. What you produce you own and are free to give away or sell. And, it’s up to you whether and on what terms you choose to commercialize your product. There’s a little catch – this also means that everything connected with use of your product is your responsibility. Signify will not accept liability if your product causes harm, for example. So use your powers for good!

Furthermore, while we say “what is yours is yours”, on the flip side we also say, “what is ours is ours”. Here, we mean the software, trademarks documentation, and any other materials we provide to help you develop Hue apps.

For example, you may refer to “Hue” and “Philips” in plain text but you aren’t allowed to use “Hue” or “Philips” branding in any logo or graphics. Also important to note is that the interface specifications “API” belong to Signify. Imagine you are working on an app and you come up with a brilliant idea for an improvement in the API or our materials. If you suggest any improvements to us and we adopt them, they become part of the platform used by everyone, and will belong to us.

And a final note: as the interface between your apps and the Hue platform will evolve over time, we will do our level best to maintain backwards compatibility and will inform you with enough time, before we roll out updates. That’s why, to keep everyone up to date, we ask you to register your app/s in the whitelist and make it clear they are your creation.