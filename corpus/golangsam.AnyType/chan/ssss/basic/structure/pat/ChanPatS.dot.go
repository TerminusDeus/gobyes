// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pat

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakePatSChan returns a new open channel
// (simply a 'chan []struct{}' that is).
//
// Note: No 'PatS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPatSPipelineStartsHere := MakePatSChan()
//	// ... lot's of code to design and build Your favourite "myPatSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPatSPipelineStartsHere <- drop
//	}
//	close(myPatSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipePatSBuffer) the channel is unbuffered.
//
func MakePatSChan() chan []struct{} {
	return make(chan []struct{})
}

// ChanPatS returns a channel to receive all inputs before close.
func ChanPatS(inp ...[]struct{}) chan []struct{} {
	out := make(chan []struct{})
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanPatSSlice returns a channel to receive all inputs before close.
func ChanPatSSlice(inp ...[][]struct{}) chan []struct{} {
	out := make(chan []struct{})
	go func() {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}()
	return out
}

// ChanPatSFuncNok returns a channel to receive all results of act until nok before close.
func ChanPatSFuncNok(act func() ([]struct{}, bool)) <-chan []struct{} {
	out := make(chan []struct{})
	go func() {
		defer close(out)
		for {
			res, ok := act() // Apply action
			if !ok {
				return
			}
			out <- res
		}
	}()
	return out
}

// ChanPatSFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanPatSFuncErr(act func() ([]struct{}, error)) <-chan []struct{} {
	out := make(chan []struct{})
	go func() {
		defer close(out)
		for {
			res, err := act() // Apply action
			if err != nil {
				return
			}
			out <- res
		}
	}()
	return out
}

// JoinPatS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPatS(out chan<- []struct{}, inp ...[]struct{}) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}()
	return done
}

// JoinPatSSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPatSSlice(out chan<- []struct{}, inp ...[][]struct{}) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
		done <- struct{}{}
	}()
	return done
}

// JoinPatSChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPatSChan(out chan<- []struct{}, inp <-chan []struct{}) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}()
	return done
}

// DonePatS returns a channel to receive one signal before close after inp has been drained.
func DonePatS(inp <-chan []struct{}) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}()
	return done
}

// DonePatSSlice returns a channel which will receive a slice
// of all the PatSs received on inp channel before close.
// Unlike DonePatS, a full slice is sent once, not just an event.
func DonePatSSlice(inp <-chan []struct{}) chan [][]struct{} {
	done := make(chan [][]struct{})
	go func() {
		defer close(done)
		PatSS := [][]struct{}{}
		for i := range inp {
			PatSS = append(PatSS, i)
		}
		done <- PatSS
	}()
	return done
}

// DonePatSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DonePatSFunc(inp <-chan []struct{}, act func(a []struct{})) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a []struct{}) { return }
	}
	go func() {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}()
	return done
}

// PipePatSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipePatSBuffer(inp <-chan []struct{}, cap int) chan []struct{} {
	out := make(chan []struct{}, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipePatSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipePatSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipePatSFunc(inp <-chan []struct{}, act func(a []struct{}) []struct{}) chan []struct{} {
	out := make(chan []struct{})
	if act == nil {
		act = func(a []struct{}) []struct{} { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipePatSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipePatSFork(inp <-chan []struct{}) (chan []struct{}, chan []struct{}) {
	out1 := make(chan []struct{})
	out2 := make(chan []struct{})
	go func() {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}()
	return out1, out2
}

// PatSTube is the signature for a pipe function.
type PatSTube func(inp <-chan []struct{}, out <-chan []struct{})

// PatSDaisy returns a channel to receive all inp after having passed thru tube.
func PatSDaisy(inp <-chan []struct{}, tube PatSTube) (out <-chan []struct{}) {
	cha := make(chan []struct{})
	go tube(inp, cha)
	return cha
}

// PatSDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func PatSDaisyChain(inp <-chan []struct{}, tubes ...PatSTube) (out <-chan []struct{}) {
	cha := inp
	for i := range tubes {
		cha = PatSDaisy(cha, tubes[i])
	}
	return cha
}

/*
func sendOneInto(snd chan<- int) {
	defer close(snd)
	snd <- 1 // send a 1
}

func sendTwoInto(snd chan<- int) {
	defer close(snd)
	snd <- 1 // send a 1
	snd <- 2 // send a 2
}

var fun = func(left chan<- int, right <-chan int) { left <- 1 + <-right }

func main() {
	leftmost := make(chan int)
	right := daisyChain(leftmost, fun, 10000) // the chain - right to left!
	go sendTwoInto(right)
	fmt.Println(<-leftmost)
}
*/
