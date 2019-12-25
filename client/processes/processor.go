package processes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/mdjdot/gochatroom/common"
)

// processLogin 处理是否能够登录
func processLogin(conn net.Conn, userID int, userPWD string) error {
	var buf [1024]byte
	var loginResult common.LoginRespMessage
	loginData, err := json.Marshal(common.User{
		UserID:  userID,
		UserPWD: userPWD,
	})
	if err != nil {
		fmt.Println("json.Marshal login data err=", err)
		return err
	}

	loginMes, err := json.Marshal(common.Message{
		Type: common.LoginMessage,
		Data: string(loginData),
	})
	if err != nil {
		fmt.Println("json.Marshal login message err=", err)
		return err
	}

	_, err = conn.Write(loginMes)
	if err != nil {
		fmt.Println("net.Conn.Write err=", err)
		return err
	}
	fmt.Println("登录中...")
	time.Sleep(5 * time.Second) // 等待5秒，等服务端返回消息

	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("net.Conn.Read err=", err)
		return err
	}
	err = json.Unmarshal(buf[:n], &loginResult)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return err
	}
	if loginResult.Result == false {
		fmt.Println("账号或密码不正确，请重新登录")
		return errors.New("账号或密码不正确，请重新登录")
	}
	return nil
}

// ProcessConn 处理客户端的连接
func ProcessConn(userID int, userPWD string) {
	var req common.Message
	conn, err := net.Dial("tcp", "127.0.0.1:8321")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	var buf [1024]byte
	var line string

	err = processLogin(conn, userID, userPWD)
	if err != nil {
		fmt.Println("登录失败，err=", err)
		return
	}

	fmt.Println("已和服务端进行连接，请输入消息：")
	for {
		n, err := fmt.Scanln(&line)
		if err != nil {
			fmt.Println("输入消息出现问题，请重新输入，err=", err)
			continue
		}
		req = common.Message{
			Type: common.RequestMessage,
			Data: line,
		}
		reqByte, err := json.Marshal(req)
		if err != nil {
			fmt.Println("json.Marshal err=", err)
			return
		}
		n, err = conn.Write(reqByte)
		if err != nil {
			fmt.Println("发送消息出现问题，err=", err)
			break
		}
		fmt.Printf("本机：%s\n", line)

		time.Sleep(3 * time.Second)
		n, err = conn.Read(buf[:])
		if err != nil {
			fmt.Println("net.Conn.Read err=", err)
			return
		}

		fmt.Printf("服务器：%s\n", string(buf[:n]))
	}
	fmt.Println("通话已结束")
	return
}
