package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("helo")
	out, err := exec.Command("go", "version").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
