const Discord = require('discord.js')

module.exports = {
    name: '!irc',
    description: 'irc Link',
    execute(message, args) {
        const irc = new Discord.MessageEmbed().setTitle(' Join us on IRC server of Freenode at #jiit-lug').setURL('https://github.com/osdc/community-committee/wiki/IRC')
        message.channel.send(irc)
    },
};