package main

import (
	"github.com/yaronsumel/simpleuser/client/communicator"
	"github.com/yaronsumel/simpleuser/client/parser"
	"github.com/yaronsumel/simpleuser/server/user"
	"log"
	"sync"
)

const maxConnections = 100

func main() {

	// communicator is responsible for sever communication
	cm := communicator.NewCommunicator("http://localhost:8080", maxConnections)

	// chan for getting user entries
	// not buffered channel
	uChan := make(chan *user.Object)
	wg := sync.WaitGroup{}

	// listen on incoming messages
	go func() {
		for {
			select {
			case u := <-uChan:
				wg.Add(1)
				// send events in goroutine
				go func() {
					defer wg.Done()
					cm.SendEvent(u)
				}()
			}
		}
	}()

	// create new parse
	p := parser.NewParser()

	// parse csv into user chan
	if err := p.Parse("data.csv", uChan); err != nil {
		panic(err)
	}

	// wait for all goroutines
	wg.Wait()
	log.Println("Done")

}
