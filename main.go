package main

import (
	"fmt"

	_ "github.com/onsi/ginkgo"
	"golang.org/x/mod/semver"
)

func main() {
	fmt.Println(semver.Build("foo"))
}
