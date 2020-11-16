const Discord = require('discord.js')

module.exports = {
    name: '!website',
    description: 'Website of Discord bot',
    execute(message, args) {
        const website = new Discord.MessageEmbed().setTitle('Visit our Website').setURL('https://osdc.netlify.app/')
        message.channel.send(website)
    },
};