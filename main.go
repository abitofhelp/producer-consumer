////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a BSD-style  license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Package main contains the entry point for the application and configuration of its environment.
package main

import (
	"fmt"
	. "gitlab.com/abitofhelpinc/producer-consumer/products"
	"strconv"
	"sync"
)

const (
	kOrderSize = 1000

	kChannelSize = 10
)

// Function main is the entry point for the application and will configure its environment.

// The following function demonstrates a producer-consumer pattern using a buffered channel
// and using a WaitGroup.
func main() {

	var (
		// The variable wb is a WaitGroup and is used to know when all goroutines have completed their work.
		wg sync.WaitGroup

		// The variable widgets is the channel from producer to consumer.
		widgets = make(chan Widget, kChannelSize)
	)

	// Function main() is implicitly a goroutine, so we will bump the goroutine counter in the WaitGroup.
	wg.Add(1)

	// The producer and consumer functions are launched as goroutines, so they will
	// run concurrently.
	go producer(widgets, &wg)
	go consumer(widgets, &wg)

	// Waiting for each of the goroutines to complete.  Internally, Wait() is waiting for
	// the count of goroutines to be zero.
	wg.Wait()

	fmt.Println("All done!")
}

// Function producer generates widgets and sends them to the consumer.
func producer(widgets chan<- Widget, waitGroup *sync.WaitGroup) {
	// Generate widgets. Since the channel only has room for kChannelSize widgets
	// at a time, the producer will block when the channel is full.
	for i := uint64(0); i < kOrderSize; i++ {

		// Create a widget...
		w, err := NewWidget(strconv.FormatUint(i, 10), i)
		if err != nil {
			fmt.Println(err)
		}

		// We bump the goroutine counter in the WaitGroup for each widget that is placed into the channel.
		waitGroup.Add(1)

		// Sending the widget...
		widgets <- *w
	}

	// Signal that we are done producing...
	defer waitGroup.Done()
}

// Function consumer receives widgets and consumes them.
func consumer(widgets <-chan Widget, waitGroup *sync.WaitGroup) {
	// Loop until all of the widgets have been consumed...
	for widget := range widgets {

		fmt.Println(widget.Name())

		// Decrement the counter after consuming a widget.
		waitGroup.Done()
	}
}
