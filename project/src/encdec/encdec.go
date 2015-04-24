package encdec

import(
	"fmt"
	"def"
	"encoding/json"
)

func encodeMsg(msg def.Msg) []byte {
	encMsg, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error encoding msg: ", err)
	}
	return encMsg
}

func decodeMsg(msg []byte) def.Msg {
	var msg_rec Message
	err := json.Unmarshal(msg[:len(msg)], &msg_rec)
	if err != nil {
		fmt.Println("Error decoding msg: ", err)
	}
	return msg_rec
}