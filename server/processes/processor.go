package processes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/garyburd/redigo/redis"
	"github.com/mdjdot/gochatroom/common"
	"github.com/mdjdot/gochatroom/server/datas"
)

func checkUser(userID int, userPWD string) (bool, error) {
	var user common.User
	conn := datas.Pool.Get()
	users, err := redis.StringMap(conn.Do("hgetall", "users"))
	if err != nil {
		fmt.Println("redis.Conn.Do err", err)
		return false, err
	}
	for _, rduser := range users {
		err = json.Unmarshal([]byte(rduser), &user)
		if err != nil {
			fmt.Println("json.Unmarshal err=", err)
			return false, err
		}
		if user.UserID == userID && user.UserPWD == userPWD {
			return true, nil
		}
	}
	return false, nil
}

func registerUser(data string) error {
	conn := datas.Pool.Get()
	max, err := getTheMaxKey()
	if err != nil {
		return err
	}
	key := strconv.Itoa(max + 1)
	_, err = conn.Do("hset", "users", key, data)
	if err != nil {
		fmt.Println("redis.Conn.Do err", err)
		return err
	}

	return nil

}

func getTheMaxKey() (int, error) {
	var max int
	conn := datas.Pool.Get()
	users, err := redis.StringMap(conn.Do("hgetall", "users"))
	if err != nil {
		fmt.Println("redis.Conn.Do err", err)
		return -1, err
	}
	for key := range users {
		keyi, err := strconv.Atoi(key)
		if err != nil {
			return -1, err
		}
		if max < keyi {
			max = keyi
		}
	}
	return max, nil
}

// processLogin 判断用户是否在数据库中
func processLogin(data string) error {
	var user common.User
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return err
	}
	exist, err := checkUser(user.UserID, user.UserPWD)
	if err != nil || exist == false {
		return errors.New("用户ID或用户密码不正确")
	}
	return nil
}

func processRegister(data string) error {
	err := registerUser(data)
	if err != nil {
		return errors.New("注册失败")
	}
	return nil
}

// ProcessConn 处理客户端的请求
func ProcessConn(conn net.Conn) {
	defer conn.Close()
	var buf [1024]byte
	var mes common.Message
	var loginResp common.LoginRespMessage
	var registerResp common.RegisterRespMessage

	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("net.Conn.Read err=", err)
			return
		}

		err = json.Unmarshal(buf[:n], &mes)
		if err != nil {
			fmt.Println("json.Unmarshal err=", err)
			return
		}

		switch mes.Type {
		case common.LoginMessage:
			err = processLogin(mes.Data)
			if err != nil {
				fmt.Println("登录失败，err=", err)
				loginResp = common.LoginRespMessage{
					Result: false,
				}
			} else {
				loginResp = common.LoginRespMessage{
					Result: true,
				}
			}
			respByte, err := json.Marshal(loginResp)
			if err != nil {
				fmt.Println("json.Marshall err=", err)
				return
			}
			_, err = conn.Write(respByte)
			if err != nil {
				fmt.Println("net.Conn.Write err=", err)
				return
			}
		case common.RegisterMessage:
			err = processRegister(mes.Data)
			if err != nil {
				fmt.Println("注册失败，err=", err)
				registerResp = common.RegisterRespMessage{
					Result: false,
				}
			} else {
				registerResp = common.RegisterRespMessage{
					Result: true,
				}
			}
			respByte, err := json.Marshal(registerResp)
			if err != nil {
				fmt.Println("json.Marshall err=", err)
				return
			}
			_, err = conn.Write(respByte)
			if err != nil {
				fmt.Println("net.Conn.Write err=", err)
				return
			}
		case common.RequestMessage:
			resp := fmt.Sprintf("收到\"%s\"", mes.Data)
			_, err = conn.Write([]byte(resp))
			if err != nil {
				fmt.Println("net.Conn.Write err=", err)
				return
			}
		default:
			fmt.Println("收到不合法的请求")
			conn.Write([]byte("不合法的请求，连接关闭"))
			return
		}
	}
}
