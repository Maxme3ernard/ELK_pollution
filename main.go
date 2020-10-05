package main

import (
	"os"

	"github.com/Maxme3ernard/polutbeat/cmd"

	_ "github.com/Maxme3ernard/polutbeat/include"
	"github.com/Maxme3ernard/polutbeat/beater"
)

var RootCmd = cmd.GenRootCmdWithSettings(beater.New, instance.Settings{Name: "polutbeat"})

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
