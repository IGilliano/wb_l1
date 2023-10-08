package main

import (
	"fmt"
	"strconv"
)

//Реализовать паттерн «адаптер» на любом примере.

type dataStorage interface {
	showData()
}

type DVD struct {
	data string
}

func (d *DVD) showData() {
	fmt.Println(d.data)
}

type memoryCart struct {
	data int
}

func (mc *memoryCart) transferData() int {
	return mc.data
}

// Adapter
type cardReader struct {
	mc *memoryCart
}

func (c *cardReader) showData() {
	fmt.Println(strconv.Itoa(c.mc.transferData()))
}

type PC struct {
}

func (p *PC) insertDataStorage(s dataStorage) {
	s.showData()
}

func main() {
	pc := PC{}
	dvd := DVD{data: "DVD data"}
	mc := memoryCart{data: 11010001101001}
	cr := cardReader{mc: &mc}

	pc.insertDataStorage(&dvd)
	pc.insertDataStorage(&cr)
}
