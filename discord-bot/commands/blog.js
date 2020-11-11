const Discord = require('discord.js')

module.exports = {
    name: '!blog',
    description: 'Blog writen by folks at OSDC',
    execute(message, args) {
        const blog = new Discord.MessageEmbed().setTitle('Blogs written by the folks at the Open Source Developers Community').setURL('https://osdcblog.netlify.com/')
        message.channel.send(blog)
    },
};

