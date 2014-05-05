# slackline-b5

Share channels between Slack accounts!

We are big fans of [Slack][https://slack.com/] and have been using it for a while
now. So I just created a quick Google AppEngine service to link two channels from two Slack organizations.

## How can I set it up?

Dead simple, follow this steps from each organization's account.

 1. Create a channel you want to share with another organization.
 2. Create an Incoming WebHook integration and select the channel you created.
 3. Copy the Incoming WebHook token (you can find it in the left column
    from the integration page).
 4. Create a URL with the following format: ```http://[GAE_APP_ID].appspot.com/bridge/?token=[TOKEN]&domain=[YOUR_SLACK_DOMAIN]``` send it to the person setting up the other organization.
 5. The person setting up the other organization will send you a similar
    URL with their domain and token, create an Outgoing WebHook with
    that URL and the channel you created in step 1.
 6. Update the `app.yaml` and publish your app to GAE

Once you have done this in both organizations, you will have a chat-room
shared by both organizations.

Here you have an example of a Outgoing WebHook URL:

```
http://myapp.appspot.com/bridge/?token=bcaa5867b1d42142b74eDVA4&domain=avengers.slack.com
```

## How does it work?

We are just bridging hooks, we don't store any messages going through
the bridge.

## ToDo

- custom avatars right now.
- better mentions fix

## Credit

Based on [ernesto-jimenez/slackline](https://github.com/ernesto-jimenez/slackline) 

## License

The MIT License (MIT)

Copyright (c) 2014 Christian Sullivan &lt;cs@bodhi5.com&gt;

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
