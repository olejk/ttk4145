package encdec

import(
	"fmt"
	. "def"
	"encoding/json"
)

func EncodeMsg(msg MSG) []byte {
	encMsg, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error encoding msg: ", err)
	}
	return encMsg
}

func DecodeMsg(msg []byte) MSG {
	var msg_rec MSG
	err := json.Unmarshal(msg[:len(msg)], &msg_rec)
	if err != nil {
		fmt.Println("Error decoding msg: ", err)
	}
	return msg_rec
}