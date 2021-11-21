const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        listMenu: () => client.get('/menu'),
    }

};

module.exports = { Client };
