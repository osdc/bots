const Discord = require('discord.js')

module.exports = {
    name: '!github',
    description: 'Github Link',
    execute(message, args) {
        const github = new Discord.MessageEmbed().setTitle(' Take a look at our cool projects ').setURL('https://github.com/osdc')
        message.channel.send(github)
    },
};