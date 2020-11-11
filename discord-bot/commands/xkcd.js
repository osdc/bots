const axios = require('axios');

module.exports = {
    name: '!xkcd',
    description: 'A xkcd Comic ',
    execute(message, args) {
        let comicNo = Math.floor(Math.random() * (2000 - 100 + 1) + 100)
        axios.get(`http://xkcd.com/${comicNo}/info.0.json`).then(resp => {
            message.channel.send(resp.data.img)
        });
    },
};