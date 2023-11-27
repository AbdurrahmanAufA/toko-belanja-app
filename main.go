package main

import "toko-belanja/handler"

type MyString string

func (ms *MyString) ChangeName(name MyString) {

}

func main() {
	handler.SeedAdmin()

	handler.StartApp()

}
