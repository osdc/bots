const Discord = require('discord.js')

module.exports = {
    name: '!facebook',
    description: 'Facebook Page',
    execute(message, args) {
        const facebook = new Discord.MessageEmbed().setTitle('Follow us on Facebook').setURL('https://www.facebook.com/JIIT-OSDC-169171359799320/')
        message.channel.send(facebook)
    },
};