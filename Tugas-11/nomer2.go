package main

import (
	"fmt"
	"sync"
)

var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

func getMovies(moviesChannel chan string, movies ...string) {
	defer close(moviesChannel)

	for _, movie := range movies {
		moviesChannel <- movie
	}
}

func main() {
	moviesChannel := make(chan string)

	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		getMovies(moviesChannel, movies...)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range moviesChannel {
			fmt.Println(value)
		}
	}()

	wg.Wait()
}
