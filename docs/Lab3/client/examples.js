// This file contains examples of scenarios implementation using
// the SDK for channels management.

const menu = require('./menu/client');

const client = menu.Client('http://localhost:9080');

// Scenario 1: Display available channels.
client.listMenu()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Available dishes:');
        console.log(list)
    })
    .catch((e) => {
        console.log(`Problem listing available dishes: ${e.message}`);
    });

