/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"ipna/cmd"
	"ipna/pkg/api"
)

func main() {
	api.GenerateIndex()
	cmd.Execute()
}
