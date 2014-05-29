package main

import (
	"fmt"
	"github.com/benschw/consul-clb-go/clb"
)

func Example_RoundRobinLookup() {
	srvName := "foo.service.fliglio.com"
	c := clb.NewClb("8.8.8.8", "53", clb.RoundRobin)
	address, err := c.GetAddress(srvName)
	if err != nil {
		fmt.Print(err)
	}

	if address.Port == 8001 {
		fmt.Printf("%s", address)
	} else {
		address2, err := c.GetAddress(srvName)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("%s", address2)
	}
	// Output: 0.1.2.3:8001

}
