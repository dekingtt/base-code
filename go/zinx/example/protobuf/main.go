package main

import (
	"fmt"
	"zinx/example/protobuf/pb"

	"google.golang.org/protobuf/proto"
)

func main() {
	person := &pb.Person{
		Name:   "Alex",
		Age:    16,
		Emails: []string{"abc@test.com"},
		Phones: []*pb.PhoneNumber{
			&pb.PhoneNumber{
				Number: "5123336666",
				Type:   pb.PhoneType_MOBILE,
			},
			&pb.PhoneNumber{
				Number: "2368888888",
				Type:   pb.PhoneType_HOME,
			},
		},
	}
	data, err := proto.Marshal(person)
	if err != nil {
		fmt.Println("marshal err:", err)
	}
	newdata := &pb.Person{}
	err = proto.Unmarshal(data, newdata)
	if err != nil {
		fmt.Println("unmarshal err:", err)
	}
	fmt.Println(newdata)
}
