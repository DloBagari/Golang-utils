package main

import (
	"fmt"
	person "github.com/Golang-utils/GRCP/protocol_buffer"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	p := person.Person{
		Age:       22,
		Name:      "dlo",
		Developer: false,
		MyAddress: []*person.Address{{
			Street:      "",
			HouseNumber: 0,
		}},
		EyeColour: 0,
	}
	fmt.Println(p.GetAge())

	// the type of p is struct and also it is a proto.Message
	//we can marshall p using proto.Marshall method
	WriteToFile("marshed", &p)
	// read the file back
	ReadFromFile("marshed", &p)
	fmt.Println(p)
	data := ToJson(&p)
	fmt.Println(data)
	FromJson(data, &p)
	fmt.Println(p)
}

func ReadFromFile(fimeNAme string, pb proto.Message) {
	data, errr := ioutil.ReadFile(fimeNAme)
	if errr != nil {
		log.Fatal(errr)
	}
	if err := proto.Unmarshal(data, pb); err != nil {
		log.Fatal(errr)
	}
}

func WriteToFile(fileName string, pb proto.Message) {
	output, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(fileName, output, 0644); err != nil {
		log.Fatal(err)
	}
}

//convert pb message to json
func ToJson(pb proto.Message) string {
	marsher := jsonpb.Marshaler{}
	out, err := marsher.MarshalToString(pb)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func FromJson(data string, pb proto.Message) {
	_ = jsonpb.UnmarshalString(data, pb)
}
