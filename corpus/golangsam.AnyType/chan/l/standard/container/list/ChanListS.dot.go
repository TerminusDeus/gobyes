// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	list "container/list"
)

// ListSChan represents a
// bidirectional
// channel
type ListSChan interface {
	ListSROnlyChan // aka "<-chan" - receive only
	ListSSOnlyChan // aka "chan<-" - send only
}

// ListSROnlyChan represents a
// receive-only
// channel
type ListSROnlyChan interface {
	RequestListS() (dat []*list.List)        // the receive function - aka "MyListS := <-MyListSROnlyChan"
	TryListS() (dat []*list.List, open bool) // the multi-valued comma-ok receive function - aka "MyListS, ok := <-MyListSROnlyChan"
}

// ListSSOnlyChan represents a
// send-only
// channel
type ListSSOnlyChan interface {
	ProvideListS(dat []*list.List) // the send function - aka "MyKind <- some ListS"
}

// DChListS is a demand channel
type DChListS struct {
	dat chan []*list.List
	req chan struct{}
}

// MakeDemandListSChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandListSChan() *DChListS {
	d := new(DChListS)
	d.dat = make(chan []*list.List)
	d.req = make(chan struct{})
	return d
}

// MakeDemandListSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandListSBuff(cap int) *DChListS {
	d := new(DChListS)
	d.dat = make(chan []*list.List, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideListS is the send function - aka "MyKind <- some ListS"
func (c *DChListS) ProvideListS(dat []*list.List) {
	<-c.req
	c.dat <- dat
}

// RequestListS is the receive function - aka "some ListS <- MyKind"
func (c *DChListS) RequestListS() (dat []*list.List) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryListS is the comma-ok multi-valued form of RequestListS and
// reports whether a received value was sent before the ListS channel was closed.
func (c *DChListS) TryListS() (dat []*list.List, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
