const Discord = require('discord.js')

module.exports = {
    name: '!instagram',
    description: 'instagram Link',
    execute(message, args) {
        message.channel.send('https://tenor.com/view/dont-do-that-avengers-black-panther-we-dont-do-that-here-gif-12042935')
    },
};