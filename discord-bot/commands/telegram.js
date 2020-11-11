const Discord = require('discord.js')

module.exports = {
    name: '!telegram',
    description: 'Telegram Link',
    execute(message, args) {
        const telegram = new Discord.MessageEmbed().setTitle(' Join our Telegram Channel').setURL('https://t.me/jiitosdc')
        message.channel.send(telegram)
    },
};