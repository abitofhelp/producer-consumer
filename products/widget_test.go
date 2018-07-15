////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a BSD-style  license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Package widget implements things that the producer can manufacturer.
package widget_test

import (
	"reflect"
	"testing"
	"time"

	. "github.com/abitofhelp/producer-consumer/products"
)

func TestNewWidget(t *testing.T) {
	type args struct {
		name         string
		serialNumber uint64
	}
	tests := []struct {
		name    string
		args    args
		want    *Widget
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWidget(tt.args.name, tt.args.serialNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWidget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWidget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWidget_Name(t *testing.T) {
	type fields struct {
		name         string
		serialNumber uint64
		createdUtc   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, _ := NewWidget(tt.fields.name, tt.fields.serialNumber)

			if got := w.Name(); got != tt.want {
				t.Errorf("Widget.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWidget_SetName(t *testing.T) {
	type fields struct {
		name         string
		serialNumber uint64
		createdUtc   time.Time
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, _ := NewWidget(tt.fields.name, tt.fields.serialNumber)
			if err := w.SetName(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Widget.SetName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWidget_SerialNumber(t *testing.T) {
	type fields struct {
		name         string
		serialNumber uint64
		createdUtc   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, _ := NewWidget(tt.fields.name, tt.fields.serialNumber)
			if got := w.SerialNumber(); got != tt.want {
				t.Errorf("Widget.SerialNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWidget_SetSerialNumber(t *testing.T) {
	type fields struct {
		name         string
		serialNumber uint64
		createdUtc   time.Time
	}
	type args struct {
		serialNumber uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, _ := NewWidget(tt.fields.name, tt.fields.serialNumber)
			if err := w.SetSerialNumber(tt.args.serialNumber); (err != nil) != tt.wantErr {
				t.Errorf("Widget.SetSerialNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWidget_CreatedUtc(t *testing.T) {
	type fields struct {
		name         string
		serialNumber uint64
		createdUtc   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, _ := NewWidget(tt.fields.name, tt.fields.serialNumber)
			if got := w.CreatedUtc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Widget.CreatedUtc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWidget_SetCreatedUtc(t *testing.T) {
	type fields struct {
		name         string
		serialNumber uint64
		createdUtc   time.Time
	}
	type args struct {
		createdUtc time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, _ := NewWidget(tt.fields.name, tt.fields.serialNumber)
			if err := w.SetCreatedUtc(tt.args.createdUtc); (err != nil) != tt.wantErr {
				t.Errorf("Widget.SetCreatedUtc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
