package utils

import "sync"

type Request struct {
	Urut    int
	Request string
	Time    int
}

var BookOrders []Request = []Request{
	{
		Urut:    4,
		Request: "Jus Apel",
		Time:    1,
	},
	{
		Urut:    2,
		Request: "Mie Goreng",
		Time:    2,
	},
	{
		Urut:    1,
		Request: "Nasi Goreng",
		Time:    5,
	},
	{
		Urut:    3,
		Request: "Kopi Luak",
		Time:    1,
	},
	{
		Urut:    5,
		Request: "Nasi Mandhi",
		Time:    3,
	},
}

var Wait sync.WaitGroup = sync.WaitGroup{}

var Mut sync.Mutex = sync.Mutex{}

var OrderSuccesSlice []string = []string{}
var OrderSucces chan string = make(chan string)

var Counter int = -1
