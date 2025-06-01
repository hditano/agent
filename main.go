package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/hditano/agent/presenter"
)

type response struct {
	Page int   `json:"page"`
	Data []any `json:"vms"`
}

var jsonData response
var mutex sync.Mutex

func main() {

	var wg sync.WaitGroup

	fmt.Println("Starting GoRoutine")

	wg.Add(2)
	go drawTable(&wg)
	go jsonParsing(&wg)

	wg.Wait()

	fmt.Printf("%v", jsonData)
}

func drawTable(wg *sync.WaitGroup) {

	defer wg.Done()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:

			ticker := time.NewTicker(1 * time.Second)
			defer ticker.Stop()

			presenter.MainTable()
			presenter.HardwareTable()
			presenter.CpuTable()

		}
	}
}

func jsonParsing(wg *sync.WaitGroup) {

	defer wg.Done()
	mutex.Lock()
	res := &response{
		Page: 1,
		Data: []any{"test1", 123, "test3", 12.32},
	}
	mutex.Unlock()
	rest1B, _ := json.Marshal(res)
	fmt.Print(string(rest1B))
}
