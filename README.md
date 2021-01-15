# ark-notify
A discord bot to notify events what's going on your *Ark: Survival Evolved* server. Use webhook to send notifications. Requires to watch entire server logs.

## Features
This program can track your logs on ARK game server. When new lines are comming, the program send a notification with added lines using webhook.

Here is an example of logs:
```
[2021.01.15-15.50.51:148][ 41]2021.01.15_15.50.51: k5342 left this ARK!
[2021.01.15-15.51.13:035][693]2021.01.15_15.51.13: k5342 joined this ARK!
[2021.01.15-15.53.06:953][ 99]2021.01.15_15.53.06: k5342 left this ARK!
```

## Requirements
- Go 1.13 or later

See go.mod file for additional dependencies.

## Installation
1. Clone this repository
1. Set environment variables based on sample.env
1. Build binary with `go build`

- `ARK_LOG_DIR=`  
Set *directory* path to your ARK server logs are stored on. (example: `/home/steam/ARK/ShooterGame/Saved/Logs/`)
- `DISCORD_WEBHOOK_URL`  
Set Webhook URL where you want to notify. You need to issue this URL manually from Discord App.

## LICENSE
The MIT License (MIT) Copyright (c) 2021 k5342

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
