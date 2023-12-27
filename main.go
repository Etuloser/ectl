package main

import (
	"ectl/internal/zsh"
	"flag"
	"fmt"
)

func main() {
	var action string
	flag.StringVar(&action, "install", "", "")
	flag.Parse()
	switch action {
	case "zsh":
		err := zsh.InstallOhMyZsh()
		if err != nil {
			fmt.Println(err)
		}
	}
}
