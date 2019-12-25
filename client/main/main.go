package main

import (
	"fmt"
)

var (
	homeStr = `
----------------欢迎登陆多人聊天系统----------------
                1 登录聊天系统
                2 注册用户
                3 退出登录
                请选择（1-3）：`
	op      int
	userID  int
	userPWD string
)

func main() {
	for {
		fmt.Print(homeStr)
		fmt.Scanln(&op)
		switch op {
		case 1:
			fmt.Println("登录...")
			conn, err := login()
			if err != nil {
				fmt.Println("登录失败，请重新选择。错误：", err)
				continue
			}
			err = communite(conn)
			if err != nil {
				fmt.Println("通信失败，请重新选择。错误：", err)
			}
			continue
		case 2:
			fmt.Println("注册")
			err := register()
			if err != nil {
				fmt.Println("注册失败，请重新选择。错误：", err)
			}
			fmt.Println("注册成功")
			continue
		case 3:
			fmt.Println("退出")
			return
		default:
			fmt.Println("选择错误，请重新选择")
		}
	}
}
