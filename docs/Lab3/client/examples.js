const menu = require('./menu/client');

const client = menu.Client('http://localhost:9080');

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

client.createOrder(2, 18, [7, 7, 2, 3])
    .then(resp => {
        console.log('=== Scenario 3 ===');
        console.log('Added new order status:', resp)
    })
    .catch(err => {
    console.log('=== Scenario 3 ===');
    console.log(`Error trying to set order: ${err}`)
});

client.createOrder(3, 3, [5])
    .then(resp => {
        console.log('=== Scenario 4 ===');
        console.log('Added new order status:', resp)
    })
    .catch(err => {
    console.log('=== Scenario 4 ===');
    console.log(`Error trying to set order: ${err}`)
});
