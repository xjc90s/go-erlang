package main

import (
	"github.com/processone/bert"
)

func main() {
	svc := bert.New("http://localhost:5281/rpc/")
	_ = svc.Call("ejabberd_auth", "try_register",
		"john", "localhost", "password")
}