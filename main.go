package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/docker/docker/pkg/namesgenerator" // https://github.com/moby/moby/blob/master/pkg/namesgenerator/names-generator.go
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

func main() {
	name := get()
	fmt.Printf("%s\n", name)
	hyphen := strings.Replace(name, "_", "-", -1)
	fmt.Printf("%s\n", hyphen)
	fmt.Printf("%s\n", strings.ToUpper(name))
	fmt.Printf("%s\n", strings.ToUpper(hyphen))
}

func get() string {
	name := namesgenerator.GetRandomName(0)
	return name
}
