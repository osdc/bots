require('dotenv').config()

const fs = require('fs');

const Discord = require('discord.js')

const bot = new Discord.Client()
bot.commands = new Discord.Collection();

// To filter out the files which are other than js
const commandFiles = fs.readdirSync('./commands').filter(file => file.endsWith('.js'));

for (const file of commandFiles) {
    const command = require(`./commands/${file}`);
    bot.commands.set(command.name, command);
}

const token = process.env.TOKEN_OSDC

bot.on('ready', () => {
    console.log('The bot is online!!!!')
})

//Greeting message for a New user!!!

bot.on('guildMemberAdd', member => {
    const channel = member.guild.channels.cache.find(ch => ch.name === 'member-log');
    if (!channel) return;
    message.channel.send(`Hello ${message.author} Welcome to OSDC Discord Server..Please introduce yourself!!`)
})

// dynamic command handling

bot.on('message', message => {

    const args = message.content.trim().split(/ +/);
    const command = args.shift().toLowerCase();

    // If a command is not present , log the default message
    if (!bot.commands.has(command)) {
        if (command[0] === "!")
            bot.commands.get("!invalid").execute(message, args);
        return;
    }

    // otherwise execute that command
    try {
        bot.commands.get(command).execute(message, args);

    } catch (error) {
        console.error(error);
        message.reply('there was an error trying to execute that command!');
    }   
});


bot.login(token)