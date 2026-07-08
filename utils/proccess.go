package utils

import (
	"fmt"
	"strconv"
	"time"
)

func Worker(urut int, wait int, proces string) {
	fmt.Printf("Membuat Gorutine proses %s\n", proces)
	for {
		Mut.Lock()
		if Counter == urut {
			fmt.Printf("Memproses pesanan %s\n", proces)
			timeWait := strconv.Itoa(wait)
			fmt.Printf("---[ %s Selesai dalam %s Detik ]---\n", proces, timeWait)
			x := 0
			for x < wait {
				time.Sleep(1 * time.Second)
				fmt.Println(x + 1)
				x++
			}
			// time.Sleep(time.Duration(wait) * time.Second)
			OrderSuccesSlice = append(OrderSuccesSlice, proces+" dalam "+timeWait+" Detik")
			OrderSucces <- proces + " Disajikan\n\n"
			Wait.Done()
			time.Sleep(1 * time.Millisecond) // Jedaa proses Kak klo tidak jeda lebih cepat lanjut proses selanjutnya daripada receiver menerima data
			Counter++
			Mut.Unlock()
			break
		}
		Mut.Unlock()
	}
}
