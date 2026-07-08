package main

import (
	"fmt"
	"goroutine-control/utils"
	"time"
)

func main() {
	keyLockUrut := 0
	addCounter := 0
	for x := range utils.BookOrders {
		request := utils.BookOrders[x].Request
		time := utils.BookOrders[x].Time
		urut := utils.BookOrders[x].Urut
		utils.Wait.Add(1)
		go utils.Worker(&keyLockUrut, urut, time, request)
		addCounter++
	}

	if addCounter == len(utils.BookOrders) {
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Printf("\n\n")
		keyLockUrut = 1
	}

	defer utils.Wait.Wait()

	fmt.Println(<-utils.OrderSucces)
	fmt.Println(<-utils.OrderSucces)
	fmt.Println(<-utils.OrderSucces)
	fmt.Println(<-utils.OrderSucces)
	fmt.Println(<-utils.OrderSucces)

	fmt.Printf("\nMelihat urutan daftar pesanan selesai...\n\n")
	time.Sleep(2 * time.Millisecond)

	for x := range utils.OrderSuccesSlice {
		fmt.Printf("%d. %s\n", x+1, utils.OrderSuccesSlice[x])
	}
}
