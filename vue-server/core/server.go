package core

import (
	"fmt"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	fmt.Printf(`
	欢迎使用 
`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
