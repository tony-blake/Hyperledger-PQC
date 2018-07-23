package main

import (
	"fmt"
	"log"
	"os/exec"
)

func f() {
	out, err := exec.Command("/Users/tonyblake/go/src/github.com/hlf-go/example-chaincodes/minimalcc2/testPQC.sh").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("output is %s\n", out)

}
