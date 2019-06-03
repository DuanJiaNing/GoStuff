package main

import (
	"GoStuff/proto/tutorial"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

func main() {

	// Write
	n := "tom"
	id := int32(12)
	p := tutorial.Person{
		Name: &n,
		Id:   &id,
	}
	ab := &tutorial.AddressBook{People: []*tutorial.Person{&p}}
	bs, _ := proto.Marshal(ab)
	fn := "protofile"
	ioutil.WriteFile(fn, bs, 0644)

	// Read
	bs, _ = ioutil.ReadFile(fn)
	po := &tutorial.AddressBook{}
	proto.Unmarshal(bs, po)
	fmt.Println(po)

}
