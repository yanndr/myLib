package main

import "myLib/cmd"

var Version = "v0.1-dev"

func main() {
	cmd.Execute(Version)
}
