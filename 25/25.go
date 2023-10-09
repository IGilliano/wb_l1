package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func main() {
	fmt.Println(time.Now())
	sleep1(3 * time.Second)
	fmt.Println(time.Now())
	sleep2(5 * time.Second)
	fmt.Println(time.Now())

}

func sleep1(d time.Duration) {
	<-time.After(d)
}

func sleep2(d time.Duration) {
	n := time.Now().Add(d)
	for time.Now().Before(n) {
	}
}
