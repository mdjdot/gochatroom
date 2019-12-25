package main

import (
	"fmt"
	"net"

	"github.com/mdjdot/gochatroom/client/processes"
)

func login() (net.Conn, error) {
	fmt.Print("请输入用户ID：")
	_, err := fmt.Scanln(&userID)
	if err != nil {
		return nil, err
	}
	fmt.Print("请输入用户密码：")
	fmt.Scanln(&userPWD)

	fmt.Printf("你输入的userid=%d pwd=%s\n", userID, userPWD)
	conn, err := processes.ProcessLogin(userID, userPWD)
	if err != nil {
		return nil, err
	}
	return conn, err
}

func register() error {
	fmt.Print("请输入用户ID：")
	n, err := fmt.Scanln(&userID)
	if err != nil {
		fmt.Println(n, err)
		fmt.Println("用户ID错误，请重新选择")
		return err
	}
	fmt.Print("请输入用户密码：")
	fmt.Scanln(&userPWD)

	fmt.Printf("你输入的userid=%d pwd=%s\n", userID, userPWD)
	err = processes.ProcessRegister(userID, userPWD)
	return err
}

func communite(net.Conn) error {
	return nil
}
