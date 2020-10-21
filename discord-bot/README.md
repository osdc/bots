# Discord Bot 
A chat bot for Discord, built on top of discord.js.

## Installation

This bot runs on [node.js](https://nodejs.org/en/). You will need at least node 12.xx

## GET THE DISCORD-API TOKEN

1. Go to [Discord Developer Portal](https://discord.com/developers/applications) and login with your Discord Account.
2. Create a New Application.
3. Click on Add Bot in the Bot section.
4. You’ll get your Bot API token under the token title
5. Copy it and save it in a file named as /.env/ in your project folder.

## ADD BOT TO YOUR TEST SERVER

1. Go to OAuth2 section in your application
2. Select bot in the scopes menu and Administrator in bot permission menu.
3. A Link will be generated in the scope menu, copy it and paste it in your browser URL tab.
4. Select your test server in the drop down box...
 

## Build & Execute the bot

1. Export the Discord API Token by running export TOKEN_OSDC=<your-token> in the OSDC-Bot directory using terminal.
2. Run following commands from terminal
```bash
npm install
node bot.js
```
