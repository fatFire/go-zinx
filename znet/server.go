package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {
	// server name
	Name string
	// server ip version
	IPVersion string
	IP        string
	Port      int
}

func (s *Server) Start() {
	// 启动服务监听端口
	fmt.Println("gtcp is starting... ")

	go func() {
		// 不让start函数阻塞，因为启动后，我们想回到serve函数中，处理其他任务
		// 开启go程，不断循环阻塞监听新的连接请求
		listen, err := net.Listen(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))

		if err != nil {
			fmt.Println("net listen err : ", err)
			return
		}
		fmt.Println("start server success, address: ", s.IP, s.Port)

		// 循环不断的监听客户端请求
		for {
			// 建立简介
			conn, err := listen.Accept()
			if err != nil {
				fmt.Println("listen accept err:", err)
				continue
			}
			// go程处理业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err", err)
						continue
					}
					//回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()

		}
	}()

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()

	select {}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
