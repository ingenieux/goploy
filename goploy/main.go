package main

import (
	"fmt"
	"github.com/ingenieux/goploy"
	"os"
)

func main() {
	if 2 != len(os.Args) {
		fmt.Printf("Usage: goploy [commitId]\n")
		os.Exit(1)
	}

	commitId := os.Args[1]

	deploy, err := goploy.NewDeployment()

	if nil != err {
		fmt.Printf("Oops: %q\n", err)
		os.Exit(1)
	}

	err = deploy.CommitId(commitId)

	if nil != err {
		fmt.Printf("Oops: %q\n", err)
		os.Exit(1)
	}

	fmt.Println(deploy.GetPushURL())
}
