module.exports = {
    name: '!invalid',
    description: 'Website of Discord bot',
    execute(message, args) {
        message.channel.send("Looks like that's an invalid command , Try !help for reference");
    },
};