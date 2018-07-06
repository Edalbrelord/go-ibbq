package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-ble/ble"
	"github.com/sworisbreathing/go-iBBQ/ibbq"
)

func main() {
	var err error
	fmt.Println("initializing context")
	ctx1, cancel := context.WithCancel(context.Background())
	defer cancel()
	registerInterruptHandler(cancel)
	ctx := ble.WithSigHandler(ctx1, cancel)
	fmt.Println("context initialized")
	var bbq ibbq.Ibbq
	fmt.Println("instantiating ibbq struct")
	if bbq, err = ibbq.NewIbbq(ctx); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("instantiated ibbq struct")
	fmt.Println("connecting to device")
	done := make(chan struct{})
	if err = bbq.Connect(done); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("Connected to device")
	<-ctx.Done()
	if err = bbq.Disconnect(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("waiting for device to send disconnect signal")
	<-done
}