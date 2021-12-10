package test

import (
	"fmt"
	"io"
	"net"
	"testing"
	"zinx/pkg/znet"
)

func TestDPServer(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen error:", err)
		return
	}

	go DPClientTest()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("server accept error:", err)
			continue
		}

		go func(conn net.Conn) {
			dp := znet.NewDataPack()
			for {
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

					fmt.Println("==> Recv Msg: ID=", msg.Id, ", Len=", msg.DataLen, ", data=", string(msg.Data))
				}
			}
		}(conn)
	}
}

func DPClientTest() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}

	dp := znet.NewDataPack()

	msg1 := &znet.Message{
		Id:      0,
		DataLen: 5,
		Data:    []byte{'H', 'E', 'L', 'L', 'O'},
	}

	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client packa msg1 err:", err)
		return
	}

	msg2 := &znet.Message{
		Id:      1,
		DataLen: 7,
		Data:    []byte("World!!"),
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg2 err:", err)
		return
	}

	sendData1 = append(sendData1, sendData2...)
	conn.Write(sendData1)
	select {}
}
