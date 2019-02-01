package main

import (
	"fmt"

	"./pb"
	"github.com/golang/protobuf/proto"
)

func test_protobuf() {

	data := pb.User{
		Id:       12345,
		Username: "softkk",
		Password: "qq123",
	}

	dataBuffer, _ := proto.Marshal(&data)

	var user pb.User
	proto.Unmarshal(dataBuffer, &user)
	fmt.Println(user.Id, " ", user.Username, " ", user.Password)

}
