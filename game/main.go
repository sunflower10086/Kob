package main

import "snake/cmd"

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
