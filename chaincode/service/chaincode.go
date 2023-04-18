package main

import ( //몰라 겁나 인포트
	"encoding/json"
	"fmt" //입출력

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type ServiceInfo struct {
	Sid        	string `json:"sid"`  //프린터 ID
	Pid        	string `json:"pid"` //프린터 IP 데이터
	ServiceCode	int    `json:"serviceCode,string`	// 토너 보충 : 0, 종이 보충 : 1 고장수리 : 2
	Num     	int    `json:"num,string`	// 소모품 보충시 보충 양
}
type SimpleAsset struct {
}

func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// 프린터 ID와 IP를 입력받아 저장 나머지 값들은 0으로 초기화
	args := stub.GetStringArgs()
	if len(args) != 5 {
		return shim.Error("인자의 개수가 맞지 않습니다.")
	}

	ServiceVal := ServiceInfo{}
	ServiceVal.Sid = args[0]
	ServiceVal.Pid = args[1]
	ServiceVal.ServiceCode = args[2]
	ServiceVal.Num = args[3]

	ServiceValByte, _ := json.Marshal(ServiceVal)
	// We store the key and the value on the ledger
	err := stub.PutState(args[4], ServiceValByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("장부에 데이터를 생성하지 못했습니다: %s", args[0]))
	}

	return shim.Success(nil)
}

// 인보크 함수
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fn, args := stub.GetFunctionAndParameters()

	if fn == "putService" { //프린터 등록
		return t.putService(stub, args)
	} else if fn == "query" { // 업데이트 프린텅
		return t.query(stub, args)
	}

	return shim.Error("함수이름이 올바르지 않습니다. 예) \"enrollPrint\" \"updatePrint\"")
}

// 프린터 등록 초기화는 0으로
func (t *SimpleAsset) putService(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	args := stub.GetStringArgs()
	if len(args) != 5 {
		return shim.Error("인자의 개수가 맞지 않습니다.")
	}

	ServiceVal := ServiceInfo{}
	ServiceVal.Sid = args[0]
	ServiceVal.Pid = args[1]
	ServiceVal.ServiceCode = args[2]
	ServiceVal.Num = args[3]

	ServiceValByte, _ := json.Marshal(ServiceVal)
	// We store the key and the value on the ledger
	err := stub.PutState(args[4], ServiceValByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("장부에 데이터를 생성하지 못했습니다: %s", args[0]))
	}

	return shim.Success(nil)
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

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
