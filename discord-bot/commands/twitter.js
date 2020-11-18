const Discord = require('discord.js')

module.exports = {
    name: '!twitter',
    description: 'Twitter Link',
    execute(message, args) {
        const twitter = new Discord.MessageEmbed().setTitle('Check us out on Twitter').setURL('https://twitter.com/osdcjiit')
        message.channel.send(twitter)
    },
};