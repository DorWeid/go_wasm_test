package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	leButton := js.Global().Get("document").Call("createElement", "button").Call("setAttribute", "id", "leButton")
	fmt.Println("%v", leButton)
	doc := js.Global().Get("document")
	fmt.Println("%v", doc)
	js.Global().Get("document").Get("body").Call("appendChild", leButton)
}
