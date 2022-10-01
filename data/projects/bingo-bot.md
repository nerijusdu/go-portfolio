# Bingo Bot

This is a slack bot that we use at my workplace to track failed releases.

<div class="img-md">
![Bingo bot in action](/data/images/bingo-bot.png)
</div><!---->

## Tech stack

This bot is built entirely in Go using [Slacker](https://github.com/shomali11/slacker#preparing-your-slack-app) package to communicate with slack API and [Go Graphics](https://github.com/fogleman/gg) package to generate an image of a bingo board.

## Challenges

This app is connecting to slack using web sockets. I though that if I wanted to have this bot in multiple slack workspaces I would just spin up multiple go routines each listening to a socket from different workspace and respond using access token from SQLite database. However, if you have multiple sockets connected slack sends events to random sockets, and usually the go routine listening for workspace A gets a message from workspace B, and tries to respond using access token for workspace A, then errors happen.


TL;DR; you can only connect single slack workspace per instance.
