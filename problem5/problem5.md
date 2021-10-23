# Problem 5 - Switching Vendor

Type: open problem

## Scenario

Currently, Kitabisa has its own notification service. We use this service to send sms notification to user when user has created donation. Unfortunately, our 3rd party vendor which sends SMS to user is unrealiable but it's the cheapest. So we need to add another 3rd party vendors which are more reliable, even though it's less cheaper.

We need our system to be interchangeable among any 3rd parties. So that we can switch our service to different vendor quickly.

[Use this link to see use case](https://gudang-dev.imgix.net/images/51c7bd15-0c63-11eb-a792-8e8de1254850_DAA5B08071A1347.png)

Please create a simple service to implement this scenario. Whenever it need, use the right design pattern.

## Expected Feature

- Send notification.

## Terms & Conditions

- Keep your code clean and maintainable, even when we have >20 3rd party vendors used in our system.
- Vendor can be configurable via config file (toml/yaml/env).
- JSON payload & response is up to you.
- To simulate 3rd party vendor, it can be either just a library or another service.
- Add unit test.
- Add necessary documentation (readme/how to, api doc, diagram, etc).

## Out of scope

- Authentication for our notification service.

## Bonus point

- Implementation in Go.
- Implement feature flag. So that vendor can be switched without restarting the system. See here for more information about feature flagging [#1](https://martinfowler.com/articles/feature-toggles.html) [#2](https://en.wikipedia.org/wiki/Feature_toggle)
- Run your app in docker (or even kubernetes).