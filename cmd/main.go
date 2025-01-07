package main

import (
	_ "xiaoniuds.com/cid/bootstrap"
	"xiaoniuds.com/cid/internal/server"
)

func main() {
	_ = server.NewServer().Run()
}
