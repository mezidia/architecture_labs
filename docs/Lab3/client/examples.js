// This file contains examples of scenarios implementation using
// the SDK for channels management.

const channels = require('./menu/client');

const client = channels.Client('http://localhost:8080');

// Scenario 1: Display available channels.
client.listMenu()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Available channels:');
        list.forEach((c) => console.log(c.name));
    })
    .catch((e) => {
        console.log(`Problem listing available channels: ${e.message}`);
    });

