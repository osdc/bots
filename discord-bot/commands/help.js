const Discord = require('discord.js')

module.exports = {
    name: '!help',
    description: 'Set of available commands',
    execute(message, args) {

        const commandEmbedded = new Discord.MessageEmbed()
            .setTitle('Here are the list of commands ')
            .addFields({
                name: '!help',
                value: "To view the list of commands ",
                inline: false
            }, {
                name: '!website',
                value: "Visit our Website",
                inline: false
            }, {
                name: '!facebook',
                value: "Follow us on Facebook ",
                inline: false
            }, {
                name: '!twitter',
                value: "Check us out on Twitter ",
                inline: false
            }, {
                name: '!github',
                value: "Visit our Github Repository ",
                inline: false
            }, {
                name: '!telegram',
                value: "Join our telegram channel ",
                inline: false
            }, {
                name: '!xkcd',
                value: "To get an xkcd comic ",
                inline: false
            }, {
                name: '!irc',
                value: "Find us on IRC :) ",
                inline: false
            }, {
                name: '!blog',
                value: "Get the link of OSDC blog",
                inline: false
            }, {
                name: '!instagram',
                value: "Follow us on instagram :) ",
                inline: false
            }, )
        message.channel.send(commandEmbedded)
    },
};