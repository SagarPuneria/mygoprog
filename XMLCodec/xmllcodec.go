package xmlCodec

import (
	"encoding/xml"
	"fmt"

	"golang.org/x/net/websocket"
)

func xmlMarshal(v interface{}) (msg []byte, payloadType byte, err error) {
	//buff := &bytes.Buffer{}
	fmt.Println("xmlMarshal func v:", v)
	fmt.Println("xmlMarshal func msg:", string(msg))
	msg, err = xml.Marshal(v)
	//msgRet := buff.Bytes()
	fmt.Println("xmlMarshal func msg:", string(msg))
	fmt.Println("xmlMarshal func websocket.TextFrame:", websocket.TextFrame)
	return msg, websocket.TextFrame, err
}

func xmlUnmarshal(msg []byte, payloadType byte, v interface{}) (err error) {
	// r := bytes.NewBuffer(msg)
	fmt.Println("xmlUnmarshal func msg:", string(msg))
	fmt.Println("xmlUnmarshal func v:", v)
	err = xml.Unmarshal(msg, v)
	fmt.Println("xmlUnmarshal func v:", v)
	return err
}

var XMLCodec = websocket.Codec{xmlMarshal, xmlUnmarshal}
