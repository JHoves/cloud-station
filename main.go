package main

import (
	"fmt"
	"github.com/JHoves/cloud-station/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
