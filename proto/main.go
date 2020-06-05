package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"

	"GoStuff/proto/example"
	"GoStuff/proto/tutorial"
)

func main11() {
	open, err := os.Open("test3")
	fmt.Println(err)
	all, err := ioutil.ReadAll(open)
	fmt.Println(err)

	d := &example.Data3{}
	err = proto.Unmarshal(all, d)
	fmt.Println(err)
	fmt.Println(marshaler.MarshalToString(d))
}

var marshaler = jsonpb.Marshaler{
	EmitDefaults: false,
}

func main() {

	fv := float32(0)
	pNil := &example.Data2{
		P: &example.Person2{Value: nil},
	}
	fmt.Println(marshaler.MarshalToString(pNil))
	fmt.Println(write("test1", pNil))

	pZero := &example.Data2{
		P: &example.Person2{Value: &fv},
	}
	fmt.Println(marshaler.MarshalToString(pZero))
	fmt.Println(write("test2", pZero))

	p1Nil := &example.Data3{
		P: &example.Person3{
			Value:  1,
			Value2: nil,
		},
	}
	fmt.Println(marshaler.MarshalToString(p1Nil))
	fmt.Println(write("test3", p1Nil))

	p1Zero := &example.Data3{
		P: &example.Person3{
			Value: 0,
			Value2: &wrappers.FloatValue{
				Value: 0,
			},
		},
	}
	fmt.Println(marshaler.MarshalToString(p1Zero))
	fmt.Println(write("test4", p1Zero))

	p1NoZero := &example.Data3{
		P: &example.Person3{
			Value: 1,
			Value2: &wrappers.FloatValue{
				Value: 1,
			},
		},
	}
	fmt.Println(marshaler.MarshalToString(p1NoZero))
	fmt.Println(write("test5", p1NoZero))

	p1NoStrZero := &example.Data3{
		P: &example.Person3{
			StrValue: "",
			StrValue2: &wrappers.StringValue{
				Value: "",
			},
		},
	}
	fmt.Println(marshaler.MarshalToString(p1NoStrZero))
	fmt.Println(write("test6", p1NoStrZero))

	p1NoStrZero1 := &example.Data3{
		P: &example.Person3{
			StrValue: "a",
			StrValue2: &wrappers.StringValue{
				Value: "a",
			},
		},
	}
	fmt.Println(marshaler.MarshalToString(p1NoStrZero1))
	fmt.Println(write("test8", p1NoStrZero1))

	str := ""
	pStrZero := &example.Data2{
		P: &example.Person2{StrValue: &str},
	}
	fmt.Println(marshaler.MarshalToString(pStrZero))
	fmt.Println(write("test7", pStrZero))
}

func write(s string, data proto.Message) error {
	//proto.ma
	bytes, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	fmt.Println(ioutil.WriteFile(s, bytes, os.ModeAppend))
	return nil
}

func main1() {

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
