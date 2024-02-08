package main

import ( //몰라 겁나 인포트
	"encoding/json"
	"fmt" //입출력

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Info struct {
	Num     	int    `json:"num,string"`	// 소모품 보충시 보충 양
}
type SimpleAsset struct {
}

func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// 프린터 ID와 IP를 입력받아 저장 나머지 값들은 0으로 초기화
	_, args := stub.GetFunctionAndParameters()
	if len(args) != 0 {
		return shim.Error("인자의 개수가 맞지 않습니다.")
	}

	Val := Info{}
	Val.Num = 0

	ValByte, _ := json.Marshal(Val)
	// We store the key and the value on the ledger
	err := stub.PutState("num", ValByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("장부에 데이터를 생성하지 못했습니다: %s", args[0]))
	}

	return shim.Success(ValByte)
}

// 인보크 함수
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fn, args := stub.GetFunctionAndParameters()

	if fn == "check" { //프린터 등록
		return t.check(stub, args)
	} else if fn == "checkPlus" { // 업데이트 프린텅
		return t.checkPlus(stub, args)
	} else if fn == "checkZero" { // 업데이트 프린텅
		return t.checkZero(stub, args)
	}


	return shim.Error("함수이름이 올바르지 않습니다. 예) \"enrollPrint\" \"updatePrint\"")
}

// 프린터 등록 초기화는 0으로
func (t *SimpleAsset) check(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 0 {
		return shim.Error("인자의 개수가 맞지 않습니다.")
	}


	Val := Info{}
	ValBytes, err := stub.GetState("num")
	if err != nil {
		return shim.Error("해당 ID에서 정보를 가져오지 못했습니다.:" + err.Error())
	} else if ValBytes == nil {
		return shim.Error("데이터가 존재하지 않습니다.")
	}

	err = json.Unmarshal(ValBytes, &Val)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(ValBytes)
}

func (t *SimpleAsset) checkPlus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	Val := Info{}
	if len(args) != 0 {
		return shim.Error("매개변수 개수가 다름니다.")
	}

	ValBytes, err := stub.GetState("num")
	if err != nil {
		return shim.Error("해당 ID에서 정보를 가져오지 못했습니다.:" + err.Error())
	} else if ValBytes == nil {
		return shim.Error("데이터가 존재하지 않습니다.")
	}

	err = json.Unmarshal(ValBytes, &Val)
	if err != nil {
		return shim.Error(err.Error())
	}

	Val2 := Info{}
	Val2.Num = Val.Num + 1
	Val2Byte, _ := json.Marshal(Val2)
	// We store the key and the value on the ledger
	err = stub.PutState("num", Val2Byte)
	if err != nil {
		return shim.Error(fmt.Sprintf("장부에 데이터를 생성하지 못했습니다: %s", args[0]))
	}

	return shim.Success(Val2Byte)
}

func (t *SimpleAsset) checkZero(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	Val := Info{}
	if len(args) != 0 {
		return shim.Error("매개변수 개수가 다름니다.")
	}

	ValBytes, err := stub.GetState("num")
	if err != nil {
		return shim.Error("해당 ID에서 정보를 가져오지 못했습니다.:" + err.Error())
	} else if ValBytes == nil {
		return shim.Error("데이터가 존재하지 않습니다.")
	}

	err = json.Unmarshal(ValBytes, &Val)
	if err != nil {
		return shim.Error(err.Error())
	}

	Val2 := Info{}
	Val2.Num = 0
	Val2Byte, _ := json.Marshal(Val2)
	// We store the key and the value on the ledger
	err = stub.PutState("num", Val2Byte)
	if err != nil {
		return shim.Error(fmt.Sprintf("장부에 데이터를 생성하지 못했습니다: %s", args[0]))
	}

	return shim.Success(Val2Byte)
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
