const http = require('../common/http.js');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        listMenu: () => client.get('/menu'),
        createOrder: (id, table_id, dishes) => client.post('/menu', { id, table_id, dishes })
    }

};

module.exports = { Client };
