package main

import (
	"fmt"
	"github.com/cloudfoundry/gosigar"
)

func main() {
	uptime := sigar.Uptime{}
	uptime.Get()

	var sig sigar.ConcreteSigar

	avg, err := sig.GetLoadAverage()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Average load is %f\n", avg.One)
}
