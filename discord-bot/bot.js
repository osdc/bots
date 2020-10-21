require('dotenv').config()
let XMLHttpRequest = require('xmlhttprequest').XMLHttpRequest
const Discord = require('discord.js')
const bot = new Discord.Client()
const token = process.env.TOKEN_OSDC
const Http = new XMLHttpRequest();

const regex = 'https://imgs.xkcd.com/comics'
const png = 'png'

bot.on('ready', () => {
    console.log('The bot is online!!!!')
})


//Greeting message for a New user!!!

bot.on('guildMemberAdd',member => {
    const channel = member.guild.channels.cache.find(ch => ch.name === 'member-log');
    if(!channel) return ;
    message.channel.send(`Hello ${message.author} Welcome to OSDC discord Server..Please introduce yourself!!`)
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
                {name : '!xkcd', value: "To get a Comic " ,inline:false},
                {name : '!irc', value: "Find us on IRC :) " ,inline:false},
                {name : '!blog', value: "Get the link of OSDC blog" ,inline:false},
                {name : '!instagram', value: "Follow us on instagram :) " ,inline:false},

            )
            message.channel.send(commandEmbedded)
        }
})

//xkcd comic command

bot.on('message',message => {
if(message.content === '!xkcd')
{ 
    let comicNo = Math.floor(Math.random() * (2000 - 100 + 1) + 100)
    const url='https://xkcd.com/' + comicNo;
    Http.open("GET", url);
    Http.send();

    Http.onload = () => {
    let endpoint = 0
    let pos = Http.responseText.indexOf(regex,2000);
        
    let arr2 = Http.responseText.indexOf('png',pos + 1)
    if(arr2 - pos < 100) {
         endpoint = arr2 + 3
        }
    else {
        let arr3 = Http.responseText.indexOf('jpg',pos+ 1);
        endpoint = arr3 + 3 ;
     }
    message.channel.send(Http.responseText.substring(pos,endpoint))
    }
    
}
})

//Social media commands

bot.on('message',message => { 
   if(message.content === '!website')
   {    
    const website = new Discord.MessageEmbed().setTitle('Visit our Website').setURL('https://osdc.github.io/')
    message.channel.send(website)
   }
   if(message.content === '!twitter')
   {
        const twitter =  new Discord.MessageEmbed().setTitle('Check us out on Twitter').setURL('https://twitter.com/osdcjiit')
        message.channel.send(twitter)
   }
   if(message.content === '!facebook')
   {
        const facebook = new Discord.MessageEmbed().setTitle('Follow us on Facebook').setURL('https://www.facebook.com/JIIT-OSDC-169171359799320/')
        message.channel.send(facebook)
   }
    if(message.content === '!github')
    {
        const github = new Discord.MessageEmbed().setTitle('Visit you Github Repository ').setURL('https://github.com/osdc')
        message.channel.send(github)
    }
    if(message.content === '!telegram')
    {
        const telegram = new Discord.MessageEmbed().setTitle(' Join our Telegram Channel').setURL('https://t.me/jiitosdc')
        message.channel.send(telegram)
    } 
    if(message.content === '!irc')
    {
        const irc = new Discord.MessageEmbed().setTitle(' Join us on IRC server of Freenode at #jiit-lug').setURL('https://github.com/osdc/community-committee/wiki/IRC')
        message.channel.send(irc)
    } 
    if(message.content === '!blog')
    {
        const blog = new Discord.MessageEmbed().setTitle('Blogs written by the folks at the Open Source Developers Club').setURL('https://osdcblog.netlify.com/')
        message.channel.send(blog)
    } 
    
    if(message.content === '!instagram')
    {
        message.channel.send('https://tenor.com/view/dont-do-that-avengers-black-panther-we-dont-do-that-here-gif-12042935')
    }
        
})


 bot.login(token)