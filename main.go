package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println()
	date, _ := exec.Command("date").Output()
	fmt.Println(string(date)) 