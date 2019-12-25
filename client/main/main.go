package main

import (
	"fmt"

	"github.com/mdjdot/gochatroom/client/processes"
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
			fmt.Print("请输入用户ID：")
			n, err := fmt.Scanln(&userID)
			if err != nil {
				fmt.Println(n, err)
				fmt.Println("用户ID错误，请重新选择")
				continue
			}
			fmt.Print("请输入用户密码：")
			fmt.Scanln(&userPWD)

			fmt.Printf("你输入的userid=%d pwd=%s\n", userID, userPWD)
			processes.ProcessConn(userID, userPWD)

			return
		case 2:
			fmt.Println("注册")
			return
		case 3:
			fmt.Println("退出")
			return
		default:
			fmt.Println("选择错误，请重新选择")
		}
	}
}
