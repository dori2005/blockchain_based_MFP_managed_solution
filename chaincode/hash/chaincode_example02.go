/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"fmt"
	"crypto/x509"
	"crypto/x509/pkix"
	"crypto/rand"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	var id string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	id = args[0]
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)

	subject := pkix.Name{
		Organization: [] string {"test Organization"},
		OrganizationalUnit: [] string {"test"},
		CommonName: "Go Web Programming",
	}

	template := x509.Certificate{
		SerialNumber:serialNumber,
		Subject:subject,
		NotBefore:time.Now(),
		NotAfter:time.Now().Add(365*24*time.Hour),
		KeyUsage:x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:[] x509.ExtKeyUsage {x509.ExtKeyUsageServerAuth},
		IPAddresses:[] net.IP {net.ParseIP("127.0.0.1")},
	}

	curve := elliptic.P256()
	pk, _:= ecdsa.GenerateKey(curve, rand.Reader)

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	certOut, _ := os.Create("cert.pen")
	pem.Encode(certOut, &pem.Block{Type:"CERTIFICATE", Bytes:derBytes})
	certOut.Close()

	keyOut, _ := os.Create("key.pem")
	keyBytes, _ := x509.MarshalECPrivateKey(pk)
	pem.Encode(keyOut, &pem.Block{Type:"ECDSA PRIVATE KEY", Bytes: keyBytes})
	keyOut.Close()

	err = stub.PutState(id, derBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(derBytes)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "up" {
		return t.up(stub, args)
	}else if function == "down" {
		return t.down(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) up(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	id = args[0]
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)

	subject := pkix.Name{
		Organization: [] string {"test Organization"},
		OrganizationalUnit: [] string {"test"},
		CommonName: "Go Web Programming",
	}

	template := x509.Certificate{
		SerialNumber:serialNumber,
		Subject:subject,
		NotBefore:time.Now(),
		NotAfter:time.Now().Add(365*24*time.Hour),
		KeyUsage:x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:[] x509.ExtKeyUsage {x509.ExtKeyUsageServerAuth},
		IPAddresses:[] net.IP {net.ParseIP("127.0.0.1")},
	}

	curve := elliptic.P256()
	pk, _:= ecdsa.GenerateKey(curve, rand.Reader)

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	certOut, _ := os.Create("cert.pen")
	pem.Encode(certOut, &pem.Block{Type:"CERTIFICATE", Bytes:derBytes})
	certOut.Close()

	keyOut, _ := os.Create("key.pem")
	keyBytes, _ := x509.MarshalECPrivateKey(pk)
	pem.Encode(keyOut, &pem.Block{Type:"ECDSA PRIVATE KEY", Bytes: keyBytes})
	keyOut.Close()

	err = stub.PutState(id, derBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(derBytes)
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) down(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	id = args[0]

	certBytes, err := stub.GetState(id)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(certBytes)
}
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
