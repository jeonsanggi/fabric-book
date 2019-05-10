// ExpressJS Setup
const express = require('express');
const app = express();
var bodyParser = require('body-parser');
// Constants
const PORT = 8080;
const HOST = '10.0.2.15';


 // Hyperledger Bridge
const { FileSystemWallet, Gateway } = require('fabric-network');
const fs = require('fs');
const path = require('path');
const ccpPath = path.resolve(__dirname, '..', '..', 'book-network', 'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

app.use(bodyParser.urlencoded({ extended: false }));

// Index page
app.get('/', function (req, res) {
  fs.readFile('./application/index.html', function (error, data) {
              res.send(data.toString());

  });
});

// Query book handle
// localhost:8080/api/querycar?carno=CAR5
// Change book owner page
app.get('/api/querybook', function (req, res) {
  fs.readFile('./application/querybook.html', function (error, data) {
              res.send(data.toString());
  });
});
app.post('/api/querybook/', async function (req, res) {
                // create the key value store as defined in the fabric-client/config/default.json 'key-value-store' setting
    try {
	var bookname = req.body.bookname;
  var location = req.body.location;

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);

        // Check to see if we've already enrolled the user.
        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }
        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: false } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('library');

        // Evaluate the specified transaction.
        const result = await contract.evaluateTransaction('queryBook', bookname, location);

        //console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        res.status(200).json({response: JSON.parse(result.toString())});
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(400).json(error);
    }
});

// Create book page
app.get('/api/createbook', function (req, res) {
  fs.readFile('./application/createbook.html', function (error, data) {
              res.send(data.toString());
  });
});
// Create book handle
app.post('/api/createbook/', async function (req, res) {
    try {
	var bookname = req.body.bookname;
	var author = req.body.author;
	var publisher = req.body.publisher;
	var location = req.body.location;
  var library = req.body.library;

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);

        // Check to see if we've already enrolled the user.
        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }
        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: false } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('library');

        // Submit the specified transaction.
        await contract.submitTransaction('createBook', bookname, author, publisher, location, library);
        console.log('Transaction has been submitted');

        // Disconnect from the gateway.
        await gateway.disconnect();

        res.status(200).json({response: 'Transaction has been submitted'});

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        res.status(400).json(error);
    }

});

// server start
app.listen(PORT, HOST);
console.log(`Running on http://10.0.2.15:${PORT}`);
