package channels

import "fmt"

func Channel() {

	myChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	msg := <-myChannel

	fmt.Println(msg)
}
