## New feature
<!-- Details about the new feature -->

## Directions/Approach
<!-- Any approach/direction how to implement it -->

## First-time Contributors

**If it is the first time that you contribute to the bot, follow these steps:**

1. You need to have **git & golang** available on your machine. 
    - To install Git - https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
    - To setup golang, follow the doc here - https://golang.org/doc/install
2. Write a comment in this issue thread to let other possible contributors know that you are working on this bug. For eg : `Hey all, I would like to work on this issue.`
3. Ping [Botfather](https://t.me/BotFather) on Telegram and make your instance of OSDC-Bot bot by selecting `/newbot` from the options it provides.
4. Copy the `TELEGRAM_TOKEN` provided by Botfather. 
5. Fork this repo.
6. Run `git clone https://github.com/<YOUR_USERNAME>/bots.git && cd bots/telegram-bot`
7. If you have installed golang, run `go build .` 
8. Wait ⏳
9. Run `export TELEGRAM_TOKEN=<botfather_token>`
10. Now, run `./telegram-bot`. The bot would be running at the user handle provided by you.
11. Make changes according to the issue. Test the working of the changes.
12. Commit the changes and then run `git push origin master` and then open a PR! 🎉
