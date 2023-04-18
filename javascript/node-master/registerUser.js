/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { FileSystemWallet, Gateway, X509WalletMixin } = require('fabric-network');
const fs = require('fs');
const path = require('path');

const ccpPath = path.resolve(__dirname, '..', '..', 'basic-network', 'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

exports.registerUser = async function (id, pw) {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path ww: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const userExists = await wallet.exists(id);
        if (userExists) {
            console.log(`An identity for the user "${id}" already exists in the wallet`);
            return;
        }

        // Check to see if we've already enrolled the admin user.
        const adminExists = await wallet.exists('admin');
        if (!adminExists) {
            console.log('An identity for the admin user "admin" does not exist in the wallet');
            console.log('Run the enrollAdmin.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'admin', discovery: { enabled: false } });

        console.log(`3`);
        // Get the CA client object from the gateway for interacting with the CA.
        const ca = gateway.getClient().getCertificateAuthority();
        const adminIdentity = gateway.getCurrentIdentity();

        console.log(`4`);

        // Register the user, enroll the user, and import the new identity into the wallet.
        const secret = await ca.register({ affiliation: 'org1.department1', enrollmentID: id, role: 'client' }, adminIdentity);
        console.log(secret);
        const enrollment = await ca.enroll({ enrollmentID: id, enrollmentSecret: secret });
        const userIdentity = X509WalletMixin.createIdentity('Org1MSP', enrollment.certificate, enrollment.key.toBytes());
        wallet.import(id, userIdentity);
        // // Register the user, enroll the user, and import the new identity into the wallet.
        // const secret = await ca.register({ affiliation: 'org1.department1', enrollmentID: 'user2', role: 'client'}, adminIdentity);
        // console.log(secret);
        // // console.log(`PW(secret) : ${id}(${secret})`);
        // const enrollment = await ca.enroll({ enrollmentID: 'user2', enrollmentSecret: secret });
        
        // const userIdentity = X509WalletMixin.createIdentity('Org1MSP', enrollment.certificate, enrollment.key.toBytes());
        // wallet.import('user2', userIdentity);
        console.log('Successfully registered and enrolled admin user "user1" and imported it into the wallet');

    } catch (error) {
        console.error(`Failed to register user "${id}": ${error}`);
        process.exit(1);
    }
    return;
};
