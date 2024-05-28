package main

import (
	"github.com/tongchengbin/cdnceck/pkg/cdncheck"
	"net"
)

func main() {
	client, err := cdncheck.NewWithOpts(3, nil)
	if err != nil {
		return
	}
	match, val, itemType, err := client.Check(net.ParseIP("8.8.8.8"))
	if err != nil {
		return
	}
	println(match, val, itemType)
}
