package main

import (
	"os"

	"github.com/Maxme3ernard/polutbeat/cmd"

	_ "github.com/Maxme3ernard/polutbeat/include"
)


func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
