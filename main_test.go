////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a BSD-style  license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Package main contains the entry point for the application and configuration of its environment.
package main

import (
	"sync"
	"testing"

	. "github.com/abitofhelp/producer-consumer/products"
	// github.com/stretchr/testify/assert
	// github.com/stretchr/testify/mock
	// github.com/stretchr/testify/http
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_producer(t *testing.T) {
	type args struct {
		widgets   chan<- Widget
		waitGroup *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			producer(tt.args.widgets, tt.args.waitGroup)
		})
	}
}

func Test_consumer(t *testing.T) {
	type args struct {
		widgets   <-chan Widget
		waitGroup *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			consumer(tt.args.widgets, tt.args.waitGroup)
		})
	}
}
