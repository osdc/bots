
# Contributing to the bot 🔥✨

Contributions are always welcome, no matter how large or small!🙂

### Create an account on Github

In order to contribute, you need to have an account on Github. Go to 👉 https://github.com to create an account.

### Install Git

* In order to contribute, you need to have Git (a version control software) installed in your machine.
* Refer this 👉 https://docs.github.com/en/github/getting-started-with-github/set-up-git#setting-up-git to install and setup Git 🚀.

### Install Go

* The Telegram Bot🤖 is written in Go. Thus, in order to install Go, follow the doc here 👉 https://golang.org/doc/install

### Install MongoDB

* The bot uses MongoDB as its database. In order to install MongoDB, refer the guide here 👉 https://docs.mongodb.com/manual/installation/#mongodb-community-edition-installation-tutorials

## Fork and clone this repository

* Fork this repository using the button in the top-right corner of the page. Refer https://docs.github.com/en/github/getting-started-with-github/fork-a-repo for more details.

* Having forked the repository, clone the repository to your local machine using the below command in the terminal :
```
$ git clone https://github.com/YOUR-GITHUB-USERNAME/bots
```

* Having cloned the copy to your local machine, enter into the `telegram-bot` directory using the `cd` command.
```
$ cd bots/telegram-bot
```

* Great, you are now present in the source code of the project. You can take a look at the contents of the project using the `ls` command.
```
$ ls
```

## Setting up the bot

1. Ping [Botfather](https://telegram.me/botfather) on Telegram and make your instance of OSDC-Bot 🤖 bot by selecting `/newbot` from the options provided.
2. Copy the `TELEGRAM_TOKEN` provided by Botfather.
3. Make sure you have followed all the above steps and are in the `telegram-bot` directory.
4. If you have installed golang, run `go build .`
5. Wait ⏳
6. Run `export TELEGRAM_TOKEN=<TELEGRAM_TOKEN>`
7. Now, run `./telegram-bot`. The bot would be running at the username provided by you on telegram. 🚀
8. If you would like to make some changes and contribute to the bot, follow the steps below.

## Making Pull-Requests (Contributions)

Having setup the bot and tested its working, if you want to contribute to it, follow the steps below :

1. Make a new branch of the project using the `git checkout` command :
```
$ git checkout -b "Name-of-the-branch"
```
2. Make changes according to the issue. Test the working of the changes.
3. Add the changes to staging area using the `git add` command :
```
$ git add .
```
4. Commit the changes made using the `git commit` commad :
```
$ git commit -m "Commit-message"
```
5. Push the changes to your branch on Github using the `git push` command :
```
$ git push -u origin "Name-of-the-branch-from-step-1"
```
6. Then, go to your forked repository and make a Pull Request 🎉. Refer [this](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request) for more details.

## Resources 📚

### Git and Github

* Egghead Course on [How to Contribute to an Open Source Project on GitHub](https://egghead.io/courses/how-to-contribute-to-an-open-source-project-on-github) by Kent C. Dodds.
* [Learn Git](https://www.codecademy.com/learn/learn-git) by Codecademy
* [Github Learning Lab](https://lab.github.com/)

### Golang

* [A Tour of Go](https://tour.golang.org/)
* [Learn Go](https://www.codecademy.com/learn/learn-go) by Codecademy
* [Gophercises](https://gophercises.com/)

## Chat 🔊

* Feel free to check out the `#jiit-lug` channel on IRC or on our Telegram channel. We are always happy to help out!
