package main

import(
	"fmt"
	"encoding/json"
)

type Message struct {
	N		int
	Str1	string
	Str2	[]string
}

func encodeMsg(msgs Message) []byte {
	encMsg, err := json.Marshal(msgs)
	if err != nil {
		fmt.Println("Error encoding msg: ", err)
	}
	return encMsg
}

func decodeMsg(msg []byte) Message {
	var msg_rec Message
	err := json.Unmarshal(msg[:len(msg)], &msg_rec)
	if err != nil {
		fmt.Println("Error decoding msg: ", err)
	}
	return msg_rec
}

func main() {
	msgs := Message {
		N:		1,
		Str1:	"s1",
		Str2:	[]string{"s2", "s3", "s4"},
	}
	fmt.Println(msgs)
	fmt.Println(encodeMsg(msgs))
	fmt.Println(decodeMsg(encodeMsg(msgs)))
}