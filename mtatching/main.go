package main

import "matching/cmd"

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
