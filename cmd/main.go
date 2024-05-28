package main

import (
	"fmt"
	"github.com/projectdiscovery/goflags"
	"github.com/tongchengbin/cdnceck/pkg/cdncheck"
	"net"
	"os"
)

type Options struct {
	IP string
}

func ParseOptions() *Options {
	options := &Options{}
	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription(`AppFinger is a web application fingerprint scanner.`)
	flagSet.CreateGroup("AppFinger", "AppFinger",
		flagSet.StringVarP(&options.IP, "ip", "i", "", "IP"),
	)
	if err := flagSet.Parse(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return options
}

func main() {
	options := ParseOptions()
	client, err := cdncheck.NewWithOpts(3, nil)
	if err != nil {
		return
	}
	match, val, itemType, err := client.Check(net.ParseIP(options.IP))
	if err != nil {
		println(err.Error())
		return
	}
	if match {
		println("match", val, itemType)
	} else {
		println("no match")
	}
}
