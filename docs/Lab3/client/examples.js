import { Client } from './menu/client';

const client = Client('http://localhost:9080');

client.listMenu()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Available dishes:');
        console.log(list)
    })
    .catch((e) => {
        console.log(`Problem listing available dishes: ${e.message}`);
    });

client.createOrder(1, 50, [2, 1, 3])
    .then(resp => {
        console.log('=== Scenario 2 ===');
        console.log('Added new order status:', resp)
    })

.catch(err => {
    console.log('=== Scenario 2 ===');
    console.log(`Error trying to set order: ${err}`)
});