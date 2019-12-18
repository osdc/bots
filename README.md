# osdc-bots

Compilation of all the bots made for simplifying the repetetive tasks on various community platforms of OSDC.

# Telegram Bot Setup Guide
## Setup Golang
* Open [Getting Started - The Go Programming Language](https://golang.org/doc/install#install)
* Follow the setup instructions as per your Operating System.
* Do make sure you have the Go version 1.13.
* Test your installation - https://golang.org/doc/install#testing

## Clone the bots repository
```
$ git clone https://github.com/osdc/bots
$ cd bots
$ git checkout -b <new-branch-name>
```

## Get the TELEGRAM_API Token
1. Ping `@BotFather` on Telegram.
2. Send the message `/start`
3. Read the instructions and make a new bot for your personal testing by `/newbot` command.
4. Give a suitable name and username to the bot.
5. You’ll get your bot API token as `89xxxxxxx:xxxxxxxxxxxxAAHmf32ZghS-cqxBLfnkUx9VwoXeOIRlnUQ`.
6. Copy it and save it in a file named as `/.env/` in `telegram-bot` directory inside your project.

## Build & Execute the bot
1. In the `telegram-bots` directory, run `go build .`
2. Export the Telegram API Token by running `export TELEGRAM_TOKEN=<your-token>`  in the `telegram-bots` directory using terminal.
3. The build creates an executable with the name `telegram-bot`. Run it using `./telegram-bot`.
4. Ping your bot at @<username_set>. You’ll find the bot up and running.










