import { Client as _Client } from '../common/http.js';

const Client = (baseUrl) => {

    const client = _Client(baseUrl);

    return {
        listMenu: () => client.get('/menu'),
        createOrder: (id, table_id, dishes) => client.post('/menu', { id, table_id, dishes })
    }

};

export default { Client };