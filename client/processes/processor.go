package processes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/mdjdot/gochatroom/common"
)

// ProcessLogin 处理是否能够登录
func ProcessLogin(userID int, userPWD string) (net.Conn, error) {
	var buf [1024]byte
	var loginResult common.LoginRespMessage
	loginData, err := json.Marshal(common.User{
		UserID:  userID,
		UserPWD: userPWD,
	})
	if err != nil {
		return nil, err
	}

	loginMes, err := json.Marshal(common.Message{
		Type: common.LoginMessage,
		Data: string(loginData),
	})
	if err != nil {
		return nil, err
	}
	conn, err := processConn()

	_, err = conn.Write(loginMes)
	if err != nil {
		return nil, err
	}
	fmt.Println("登录中...")
	time.Sleep(5 * time.Second) // 等待5秒，等服务端返回消息

	n, err := conn.Read(buf[:])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf[:n], &loginResult)
	if err != nil {
		return nil, err
	}
	if loginResult.Result == false {
		return nil, errors.New("登录失败")
	}
	return conn, nil
}

func processConn() (net.Conn, error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8321")
	if err != nil {
		return nil, err
	}
	return conn, err
}

// ProcessCommunite 完成通信工作
func ProcessCommunite(conn net.Conn) error {
	defer conn.Close()

	var line string
	var buf [1024]byte
	var req common.Message
	var err error

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
			return err
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
			return err
		}

		fmt.Printf("服务器：%s\n", string(buf[:n]))
	}
	fmt.Println("通话已结束")
	return err
}

// ProcessRegister 完成注册工作
func ProcessRegister(userID int, userPWD string) error {
	var buf [1024]byte
	var registerMes []byte
	var registerResult common.LoginRespMessage

	registerData, err := json.Marshal(common.User{
		UserID:  userID,
		UserPWD: userPWD,
	})
	if err != nil {
		return err
	}

	registerMes, err = json.Marshal(common.Message{
		Type: common.RegisterMessage,
		Data: string(registerData),
	})
	if err != nil {
		return err
	}
	conn, err := processConn()

	_, err = conn.Write(registerMes)
	if err != nil {
		return err
	}
	fmt.Println("注册中...")
	time.Sleep(5 * time.Second) // 等待5秒，等服务端返回消息

	n, err := conn.Read(buf[:])
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf[:n], &registerResult)
	if err != nil {
		return err
	}
	if registerResult.Result == false {
		return errors.New("注册失败")
	}
	return nil
}
