////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Package widget implements things that the producer can manufacturer.
package widget

import (
	"errors"
	"strings"
	"time"

	. "github.com/abitofhelp/go-helpers/string"
)

// Type Widget is a thing that is branded with a name and its timestamp of creation.s
type Widget struct {
	name         string
	serialNumber uint64
	createdUtc   time.Time
	age          uint
}

// Function NewWidget is a factory that creates an initialized Widget.
// Parameter name is the name to brand to an instance of a Widget.
// Returns an initialized widget or error.
func NewWidget(name string, serialNumber uint64) (*Widget, error) {
	widget := &Widget{}
	widget.age = 0
	if widget == nil {
		return nil, errors.New("failed to create an instance of Widget")
	}

	err := widget.SetName(name)
	if err != nil {
		return nil, err
	}

	err = widget.SetSerialNumber(serialNumber)
	if err != nil {
		return nil, err
	}

	err = widget.SetCreatedUtc(time.Now().UTC())
	if err != nil {
		return nil, err
	}

	return widget, nil
}

// Method Name gets the name from the instance of a Widget.
func (w Widget) Name() string {
	return w.name
}

// Method SetName sets the value of the name in the instance.
// If there is an error, an error is returned, otherwise nil.
func (w *Widget) SetName(name string) error {

	if name == "" {
		return errors.New("the name cannot be empty")
	}
	name = CleanStringForPlatform(name)
	w.name = name

	return nil
}

// Method SerialNumber gets the serial number from the instance of a Widget.
func (w Widget) SerialNumber() uint64 {
	return w.serialNumber
}

// Method SetSerialNumber sets the value of the serial number in the instance.
// If there is an error, an error is returned, otherwise nil.
func (w *Widget) SetSerialNumber(serialNumber uint64) error {
	w.serialNumber = serialNumber
	return nil
}

// Method Created gets the UTC date/time when the instance was created.
func (w Widget) CreatedUtc() time.Time {
	return w.createdUtc
}

// Method SetCreatedUtc sets the value of the createdUtc in the instance.
// If there is an error, an error is returned, otherwise nil.
func (w *Widget) SetCreatedUtc(createdUtc time.Time) error {
	if createdUtc.IsZero() {
		return errors.New("the createdUtc cannot be zero")
	}

	utc := CleanStringForPlatform(createdUtc.Location().String())
	if strings.Compare(utc, "UTC") != 0 {
		return errors.New("the createdUtc value must be in UTC")
	}

	w.createdUtc = createdUtc

	return nil
}
