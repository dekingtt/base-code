package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/pkg/znet"
)

func main() {
	fmt.Println("Client Test ... start")
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("Failed to connect to server ... ", err)
		return
	}

	for {
		// v0.4
		// _, err = conn.Write([]byte("Hello zinx"))
		// if err != nil {
		// 	fmt.Println("Failed to write, ", err)
		// 	return
		// }
		// buf := make([]byte, 512)
		// cnt, err := conn.Read(buf)
		// if err != nil {
		// 	fmt.Println("Failed to read...", err)
		// 	return
		// }
		// fmt.Printf("server call back : %s, cnt = %d\n", buf, cnt)
		// time.Sleep(1 * time.Second)
		dp := znet.NewDataPack()
		msg, _ := dp.Pack(znet.NewMsgPackage(1, []byte("Zinx v0.6 Client1 Test Message")))
		_, err = conn.Write(msg)
		if err != nil {
			fmt.Println("write error err", err)
			return
		}

		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(conn, headData)
		if err != nil {
			fmt.Println("read head error")
			break
		}
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}

			fmt.Println("==> recv msg: ID = ", msg.Id, ",len=", msg.DataLen, ",data=", string(msg.Data))
		}
		time.Sleep(1 * time.Second)
	}

}
