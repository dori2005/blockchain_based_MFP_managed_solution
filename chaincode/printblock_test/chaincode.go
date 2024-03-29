package main

import ( //몰라 겁나 인포트
	"bytes"
	"encoding/json"
	"fmt" //입출력

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type PrinterInfo struct {
	Id        string `json:"id"`  //프린터 ID
	Ip        string `json:"url"` //프린터 IP 데이터
	Black     int    `json:"black,string"`
	Cyan      int    `json:"cyan,string"`
	Magenta   int    `json:"magenta,string"`
	Yellow    int    `json:"yellow,string"`
	Drum      int    `json:"drum,string"`
	ErrorCode int    `json:"Errorcode,string`
	Paper     int    `json:"Paer,string`
}
type SimpleAsset struct {
}

func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// 프린터 ID와 IP를 입력받아 저장 나머지 값들은 0으로 초기화
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("인자의 개수가 맞지 않습니다.")
	}

	printerVal := PrinterInfo{}
	printerVal.Id = args[0]
	printerVal.Ip = args[1]
	printerVal.Black = 0
	printerVal.Cyan = 0
	printerVal.Magenta = 0
	printerVal.Yellow = 0
	printerVal.Drum = 0
	printerVal.ErrorCode = 0
	printerVal.Paper = 0

	printerValByte, _ := json.Marshal(printerVal)
	// We store the key and the value on the ledger
	err := stub.PutState(args[0], printerValByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("장부에 데이터를 생성하지 못했습니다: %s", args[0]))
	}

	return shim.Success(nil)
}

// 인보크 함수
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fn, args := stub.GetFunctionAndParameters()

	if fn == "enrollPrint" { //프린터 등록
		return t.enrollPrint(stub, args)
	} else if fn == "updatePrint" { // 업데이트 프린텅
		return t.updatePrint(stub, args)
	} else if fn == "query" { // 업데이트 프린텅
		return t.query(stub, args)
	} else if fn == "queryAll" {
		return t.queryAll(stub)
	}
 
	return shim.Error("함수이름이 올바르지 않습니다. 예) \"enrollPrint\" \"updatePrint\"")
}

// 프린터 등록 초기화는 0으로
func (t *SimpleAsset) enrollPrint(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("인자의 개수가 맞지 않습니다.")
	}

	printerVal := PrinterInfo{}
	printerVal.Id = args[0]
	printerVal.Ip = args[1]

	printerVal.Black = 10
	printerVal.Cyan = 20
	printerVal.Magenta = 30
	printerVal.Yellow = 50
	printerVal.Drum = 30
	printerVal.ErrorCode = 0
	printerVal.Paper = 123

	printerValByte, _ := json.Marshal(printerVal)
	// We store the key and the value on the ledger
	err := stub.PutState(args[0], printerValByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("장부에 데이터를 생성하지 못했습니다: %s", args[0]))
	}
	return shim.Success(nil)
}

func (t *SimpleAsset) updatePrint(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("매개변수 개수가 다름니다.")
	}

	ValBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("해당 ID에서 정보를 가져오지 못했습니다.:" + err.Error())
	} else if ValBytes == nil {
		return shim.Error("데이터가 존재하지 않습니다.")
	}

	printerVal := PrinterInfo{}
	err = json.Unmarshal(ValBytes, &printerVal)
	if err != nil {
		return shim.Error(err.Error())
	}
	updateVal := PrinterInfo{}

	updateVal.Black = 10
	updateVal.Cyan = printerVal.Cyan - 1
	updateVal.Magenta = printerVal.Magenta - 1
	updateVal.Yellow = printerVal.Yellow - 1
	updateVal.Drum = printerVal.Drum - 1
	updateVal.ErrorCode = 0
	updateVal.Paper = printerVal.Paper + 3

	printerValByte, _ := json.Marshal(updateVal)
	err = stub.PutState(args[0], printerValByte)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(printerValByte)
}
func (t *SimpleAsset) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	printerVal := PrinterInfo{}
	if len(args) != 1 {
		return shim.Error("매개변수 개수가 다름니다.")
	}

	ValBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("해당 ID에서 정보를 가져오지 못했습니다.:" + err.Error())
	} else if ValBytes == nil {
		return shim.Error("데이터가 존재하지 않습니다.")
	}

	err = json.Unmarshal(ValBytes, &printerVal)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(ValBytes)
}


func (s *SimpleAsset) queryAll(stub shim.ChaincodeStubInterface) pb.Response {
	startKey := ""
	endKey := "zzzzzzzzzzzz"

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
