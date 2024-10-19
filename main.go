package main

import (
	"fmt"

	"github.com/hantsaniala/uploader/pkg/uploader"
)

var Version string

const (
	banner = `
	░█░█░█▀█░█░░░█▀█░█▀█░█▀▄░█▀▀░█▀▄
	░█░█░█▀▀░█░░░█░█░█▀█░█░█░█▀▀░█▀▄
	░▀▀▀░▀░░░▀▀▀░▀▀▀░▀░▀░▀▀░░▀▀▀░▀░▀
			%s
`
)

func main() {
	fmt.Printf(banner, Version)
	uploader.Run()

	fmt.Println("Press Entrer key quit!")
	fmt.Scanln() // wait for Enter Key
}
