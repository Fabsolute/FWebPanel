package main

import (
	"fwebpanel/api"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go api.Run(wg)

	wg.Wait()
}
