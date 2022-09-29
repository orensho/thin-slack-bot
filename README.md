# thin Slack Bot

[![Release](https://img.shields.io/github/release/orensho/thin-slack-bot/all.svg)](https://github.com/orensho/thin-slack-bot/latest)
[![PkgGoDev](https://pkg.go.dev/badge/orensho/thin-slack-bot/)](https://github.com/orensho/thin-slack-bot/)

I created this thin Slack bot to demonstrate how easy it is to automate your repetitive actions using a bot

## Description

A thin Slack bot using github.com/shomali11/slacker to fork from for your custom slack bot.

## Required environment variables

| Name                     | Description                                                                                                                                         |
|--------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| SLACK_BOT_TOKEN          | Bot tokens represent a bot associated with the app installed in a workspace                                                                         |
| SLACK_APP_TOKEN          | App-level tokens represent your app across organizations, including installations by all individual users on all workspaces in a given organization | |

# Slack app

## Creation

1. go to https://api.slack.com/apps/
2. create an app from scratch on brcm-bsg-ims-nis-ses workspace
3. generate an app-level token with  connections:write scope and retrieve your app token
4. add collaborators
5. go to 'OAuth & Permissions'
6. add the following Bot token scopes 
   1. app_mentions:read
   2. channels:history 
   3. chat:write 
   4. groups:history 
   5. im:history 
   6. mpim:history
7. add the following Bot token scopes
   1. chat:write
8. go to 'Socket mode'
   1. enable it
   2. go to  'Subscribe to bot events' and add message.im
10. click 'Request to install'
12. retrieve your bot user token
13. got to 'Install app' and install the app

## App
app-name<br />
https://api.slack.com/apps/tbd/ on workspace tbd<br />
App token = xapp-1234<br />
Bot token = xoxb-1234

## App debug
app-name-debug<br />
https://api.slack.com/apps/tbd/ on workspace tbd<br />
App token = xapp-4321<br />
Bot token = xoxb-4321

# Docker 

## Build

makefile build

## Run

docker run -e SLACK_APP_TOKEN="xapp-1234" -e SLACK_BOT_TOKEN="xoxb-1234" myrepo/thin-slack-bot:1.0.0

# Deployment

Your containerized Slack bot should be deployed on a workload to provide availability<br />
It is recommended to create a CI pipeline to automate deployment 


