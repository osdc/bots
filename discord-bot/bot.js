require('dotenv').config()
const axios = require('axios');

const Discord = require('discord.js')
const bot = new Discord.Client()
const token = process.env.TOKEN_OSDC

bot.on('ready', () => {
    console.log('The bot is online!!!!')
})


// A set of valid commands 

var isValid = new Set(["!help" , "!xkcd" , "!telegram" , "!facebook" , "!twitter" , "!irc" , "!blog" , "!website" , "!github" , "!blog" , "!instagram"])

//Greeting message for a New user!!!

bot.on('guildMemberAdd',member => {
    const channel = member.guild.channels.cache.find(ch => ch.name === 'member-log');
    if(!channel) return ;
    message.channel.send(`Hello ${message.author} Welcome to OSDC Discord Server..Please introduce yourself!!`)
})

//help command

bot.on('message',message => {
        if(message.content === '!help')
        {
            const commandEmbedded = new Discord.MessageEmbed()
            .setTitle('Here are the list of commands ')
            .addFields(
                {name : '!help', value: "To view the list of commands " ,inline:false},
                {name : '!website', value: "Visit our Website" ,inline:false}, 
                {name : '!facebook', value: "Follow us on Facebook " ,inline:false},
                {name : '!twitter', value: "Check us out on Twitter " ,inline:false},
                {name : '!github', value: "Visit our Github Repository " ,inline:false},
                {name : '!telegram', value: "Join our telegram channel " ,inline:false},
                {name : '!xkcd', value: "To get an xkcd comic " ,inline:false},
                {name : '!irc', value: "Find us on IRC :) " ,inline:false},
                {name : '!blog', value: "Get the link of OSDC blog" ,inline:false},
                {name : '!instagram', value: "Follow us on instagram :) " ,inline:false},
            )
            message.channel.send(commandEmbedded)
        }
})

//xkcd comic command
//It first randomly generate a number between 100 to 2000 and then sends the http request with the generated number to fetch comic 


bot.on('message',message => {
if(message.content === '!xkcd')
   { 
    let comicNo = Math.floor(Math.random() * (2000 - 100 + 1) + 100)
    
    axios.get(`http://xkcd.com/${comicNo}/info.0.json`).then(resp => {
    
        message.channel.send(resp.data.img)
    });
   }
    
})

//Social media commands

bot.on('message',message => { 

   if(message.content === '!website')
   {    
    const website = new Discord.MessageEmbed().setTitle('Visit our Website').setURL('https://osdc.netlify.app/')
    message.channel.send(website)
    return false;
   }
   else if(message.content === '!twitter')
   {
        const twitter =  new Discord.MessageEmbed().setTitle('Check us out on Twitter').setURL('https://twitter.com/osdcjiit')
        message.channel.send(twitter)
        return false
   }
   else  if(message.content === '!facebook')
   {
        const facebook = new Discord.MessageEmbed().setTitle('Follow us on Facebook').setURL('https://www.facebook.com/JIIT-OSDC-169171359799320/')
        message.channel.send(facebook)
   }
   else  if(message.content === '!github')
    {
        const github = new Discord.MessageEmbed().setTitle(' Take a look at our cool projects ').setURL('https://github.com/osdc')
        message.channel.send(github)
    }
    else  if(message.content === '!telegram')
    {
        const telegram = new Discord.MessageEmbed().setTitle(' Join our Telegram Channel').setURL('https://t.me/jiitosdc')
        message.channel.send(telegram)
    } 
    else  if(message.content === '!irc')
    {
        const irc = new Discord.MessageEmbed().setTitle(' Join us on IRC server of Freenode at #jiit-lug').setURL('https://github.com/osdc/community-committee/wiki/IRC')
        message.channel.send(irc)
    } 
    else  if(message.content === '!blog')
    {
        const blog = new Discord.MessageEmbed().setTitle('Blogs written by the folks at the Open Source Developers Community').setURL('https://osdcblog.netlify.com/')
        message.channel.send(blog)
    } 
    
    else  if(message.content === '!instagram')
    {
        message.channel.send('https://tenor.com/view/dont-do-that-avengers-black-panther-we-dont-do-that-here-gif-12042935')
    }

    else if(message.content.length > 0 && isValid.has(message.content) === false && message.content[0] === '!')
    {
        const sorry = new Discord.MessageEmbed().setTitle('Command Not Found , try !help for reference')    
        message.channel.send(sorry)
    }
        
})


 bot.login(token)
