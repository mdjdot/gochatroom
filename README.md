# 海量用户即时通信系统项目
## 需求分析-->设计阶段-->编码实现-->测试阶段-->实施

	1. 需求分析
		1. 用户注册
		2. 用户登录
		3. 显示在线用户列表
		4. 群聊（广播）
		5. 点对点聊天
		6. 离线留言

	2. 界面设计
	----------------欢迎登陆多人聊天系统：----------------
				1 登录聊天系统
				2 注册用户
				3 退出系统
				
				请选择（1-3）：
	---------------------------------------------------------
	1
	登录…
	请输入用户id：
	100
	请输入用户密码：
	200
	
	你输入的 userid=10 pwd=200

	3. 项目开发前技术准备
	server（进程，协程，端口） -->  数据库（mysql，oracle，redis，mongodb，memcache）
		|
		|
		V
	clientA  clientB  clientN    
	
	4. 实现显示客户端登录菜单
	----------------欢迎登陆多人聊天系统：----------------
				1 登录聊天系统
				2 注册用户
				3 退出系统
				
				请选择（1-3）：
	---------------------------------------------------------
	1
	登录…
	请输入用户id：
	100
	请输入用户密码：
	200
	
	你输入的 userid=10 pwd=200

	5. 实现用户登录
	完成用户验证，用户id=100，pwd=123456可以登录，其他用户不能登录
		1. 完成客户端可以发送消息长度，服务器端可以正常收到该长度值
		2. 完成客户端可以发送消息本身，服务器端可以正常接受到消息，并根据客户端发送的消息（LoginMes）判断用户的合法性，并返回相应的LoginResMes
		3. 发送流程
			i. 先创建一个Message 的结构体
			ii. Mes.Type=登录消息类型
			iii. Mes.Data=登录消息的内容（序列化）
			iv. 对mes进行序列化
			v. 在网络传输中，防止丢包
				1) 先给服务器发送mes的长度有多少个字节n
				2) 再发送消息本身
		4. 接收流程
			i. 接收到客户端发送的长度len
			ii. 根据接收到的长度len，再接收消息本身
			iii. 接收时要判断实际接收到的消息内容是否等于len
			iv. 如不相等，需要纠错
			v. 取到反序列化的mes
			vi. 取出mes.Data 中的LoginMes
			vii. 根据比较结果返回Mess
			viii. 发送给客户端
		5. 重构
		Client/
			main/
			model/
			process/
			utils/
				
	6. 实现注册用户
	7. 实现登录时能返回当前的在线用户
	8. 实现登录后可以群聊
	9. 扩展功能
		1. 实现私聊（点对点聊天）
		2. 如果一个登录用户离线，就把这个人从在先列表去掉
		3. 实现离线留言，在群聊时，如果某个用户没有在线，当登录后，可以接收离线的消息
        4. 发送一个文件