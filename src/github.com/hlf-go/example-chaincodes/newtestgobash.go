package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("/root/testPQC.sh").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("output is %s\n", out)

}
