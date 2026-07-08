package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func order(wait int, proces string, orderSucces *[]string, wg *sync.WaitGroup) {
	timeWait := strconv.Itoa(wait)
	time.Sleep(2 * time.Second)
	*orderSucces = append(*orderSucces, proces+" dalam "+timeWait+" Detik")
	defer wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	orderSucces := []string{}

	type request struct {
		urut    sync.Mutex
		request string
		time    int
	}
	askOrder := []request{
		{
			urut:    1,
			request: "Nasi Goreng",
			time:    5,
		},
		{
			urut:    2,
			request: "Mie Goreng",
			time:    2,
		},
		{
			urut:    3,
			request: "Kopi Luak",
			time:    1,
		},
		{
			urut:    4,
			request: "Jus Apel",
			time:    1,
		},
		{
			urut:    5,
			request: "Nasi Mandhi",
			time:    3,
		},
	}

	for x := range askOrder {
		wg.Add(1)
		request := askOrder[x].request
		time := askOrder[x].time
		fmt.Printf("Proses (%d) [%s]\n", x+1, request)
		go order(time, request, &orderSucces, &wg)
	}

	wg.Wait()

	for i := range orderSucces {
		fmt.Printf("%d. %s\n", i+1, orderSucces[i])
	}
}
