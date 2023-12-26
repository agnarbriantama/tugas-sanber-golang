package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func displayPhone(index int, phone string, wg *sync.WaitGroup){
	defer wg.Done()

	time.Sleep(time.Duration(index) * time.Second)

	fmt.Printf("%v. %v\n", index+1, phone)
}



func main() {
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

	sort.Strings(phones)

	var wg sync.WaitGroup

	for index, phone := range phones {

		wg.Add(1)
		go displayPhone(index, phone, &wg)
	}

	wg.Wait()

}