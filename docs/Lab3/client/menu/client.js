const http = require('../common/http.js');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        listMenu: () => client.get('/menu'),
    }

};

module.exports = { Client };
