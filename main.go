package main

import (
	"fmt"

	"golang.org/x/mod/semver"
)

func main() {
	fmt.Println(semver.Build("foo"))
}
