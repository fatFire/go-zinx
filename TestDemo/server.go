package main

import "github.com/fatFire/gtcp/znet"

func main() {
	s := znet.NewServer("[zinx 0,1]")
	s.Serve()
}
