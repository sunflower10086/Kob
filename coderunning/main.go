package main

import "coderunning/cmd"

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
