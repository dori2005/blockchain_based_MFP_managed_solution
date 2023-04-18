'use strict';

var fs = require('fs');
const { FileSystemWallet, Gateway } = require('fabric-network');
const path = require('path');

const ccpPath = path.resolve(__dirname, '..', '..', 'basic-network', 'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

exports.queryB = async function (func, name, ip) {
    try {

        // 지갑에서 신원 선택
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // 등록된 사용자인지 확인
        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // 게이트웨이에 연결
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: false } });

        // 네트워크에 접속
        const network = await gateway.getNetwork('mychannel');

        // 스마트 컨트랙트 요청
        const contract = network.getContract('testcc');

        // 트랜잭션 Submit
        if (ip == null)
          var result = await contract.evaluateTransaction(func, name);
        else
          var result = await contract.evaluateTransaction(func, name, ip);

        //const result = await contract.evaluateTransaction('query','dori');

        // 프로세스 응답
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        return result;

    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        process.exit(1);
    }
};