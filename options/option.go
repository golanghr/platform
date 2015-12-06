/*
Copyright (c) 2015 Golang Croatia
All rights reserved.

The MIT License (MIT)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Package options ...
package options

// Option - Option container and type helper
type Option struct {
	Key   string      `option:"key"`
	Value interface{} `option:"value"`
}

// Bool - Will return option value as bool
func (o *Option) Bool() bool {
	return o.Value.(bool)
}

// UInt - Will return option value as unassigned int
func (o *Option) UInt() uint {
	return o.Value.(uint)
}

// Int - Will return option value as int
func (o *Option) Int() int {
	return o.Value.(int)
}

// UInt8 - Will return option value as unassigned int 8
func (o *Option) UInt8() uint8 {
	return o.Value.(uint8)
}

// Int8 - Will return option value as int 8
func (o *Option) Int8() int8 {
	return o.Value.(int8)
}

// UInt16 - Will return option value as unassigned int 16
func (o *Option) UInt16() uint16 {
	return o.Value.(uint16)
}

// Int16 - Will return option value as int 16
func (o *Option) Int16() int16 {
	return o.Value.(int16)
}

// UInt32 - Will return option value as unassigned int 32
func (o *Option) UInt32() uint32 {
	return o.Value.(uint32)
}

// Int32 - Will return option value as int 32
func (o *Option) Int32() int32 {
	return o.Value.(int32)
}

// UInt64 - Will return option value as unassigned int 64
func (o *Option) UInt64() uint64 {
	return o.Value.(uint64)
}

// Int64 - Will return option value as int 64
func (o *Option) Int64() int64 {
	return o.Value.(int64)
}

// Float - Will return option value as float (we use float64 as "default float")
func (o *Option) Float() float64 {
	return o.Float64()
}

// Float32 - Will return option value as float 32
func (o *Option) Float32() float32 {
	return o.Value.(float32)
}

// Float64 - Will return option value as float 64
func (o *Option) Float64() float64 {
	return o.Value.(float64)
}

// Complex64 - Will return option value as complex 64
func (o *Option) Complex64() complex64 {
	return o.Value.(complex64)
}

// Complex128 - Will return option value as complex 128
func (o *Option) Complex128() complex128 {
	return o.Value.(complex128)
}

// String - Will return option value as string
func (o *Option) String() string {
	return o.Value.(string)
}

// Interface - Will return option value as interface
func (o *Option) Interface() interface{} {
	return o.Value
}
