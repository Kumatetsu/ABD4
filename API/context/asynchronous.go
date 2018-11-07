/*
 * File: asynchronous.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 4th November 2018 11:03:55 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 5th November 2018 3:49:02 am
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package context

import (
	"ABD4/API/iserial"
	"ABD4/API/utils"
	"math"
	"sync"
)

// define RequiredPools limit maximum simultaneous request to elasticsearch at 200
func (ctx *AppContext) defineRequiredPools(qty int) int {
	pools := int(math.Trunc(float64(qty) / float64(ctx.Opts.GetBatch())))
	if qty%ctx.Opts.GetBatch() > 0 {
		pools = pools + 1
	}
	if pools > 200 {
		pools = 200
	}
	return pools
}

// AsynchronousBatching will split serializable slice into segment of GetBatch() size
// each batch will be indexed in his own process
func (ctx *AppContext) AsynchronousBatching(serializables []iserial.Serializable, index, t string) error {
	var err error
	qty := len(serializables)
	// max 200
	pools := ctx.defineRequiredPools(qty)
	ctx.Log.Info.Printf("%s %d pools", utils.Use().GetStack(ctx.IndexArrayData), pools)
	signals := make([]chan bool, pools)
	// permet de synchroniser le tableau de channels
	// merged va recevoir les input et être fermé lorsque tous seront passés
	merged := ctx.MergeSignals(signals)
	errChan := make(chan error)
	// voir context/elastic.go pour les fonctions d'indexations asynchrone
	ctx.poolIndexation(serializables, index, t, pools, errChan, signals)
	count := 1
	ctx.Log.Info.Printf("%s before select ", utils.Use().GetStack(ctx.IndexArrayData))
	for count <= pools {
		select {
		case check, more := <-merged:
			ctx.Log.Info.Printf("%s received %v", utils.Use().GetStack(ctx.IndexArrayData), check)
			count++
			if count == pools {
				ctx.Log.Info.Printf("%s all data indexed, total: %d ", utils.Use().GetStack(ctx.IndexArrayData), qty)
			}
			if !more {
				ctx.Log.Info.Printf("%s no more pools to report, channel is closed ", utils.Use().GetStack(ctx.IndexArrayData))
			}
		case err, more := <-errChan:
			if err == nil && !more {
				ctx.Log.Info.Printf("%s err channel is closed, everything goes well", utils.Use().GetStack(ctx.IndexArrayData))
			}
		}
	}
	return err
}

// MergeSignals take a slice of channel boolean
// return a simple channel but two goroutines will listen
// one on each signal in slice, the other on a waitGroup waiting for
// len(signals) call on waitGroup.Done(). Everytime a signal receive a bool
// merged receive one also and waitGroupe call Done method
// when len(signals) calls on Done are made, merged channel is closed
// On the other side, a simple listen on merged allow to synchronise a process
func (ctx *AppContext) MergeSignals(signals []chan bool) chan bool {
	var waitGroup sync.WaitGroup
	merged := make(chan bool)
	for i := range signals {
		signals[i] = make(chan bool)
	}
	waitGroup.Add(len(signals))
	for i, sig := range signals {
		go func(sig chan bool, i int) {
			select {
			case merged <- <-sig:
				{
					waitGroup.Done()
				}
			}
		}(sig, i)
	}
	go func() {
		waitGroup.Wait()
		close(merged)
	}()
	return merged
}

// Tools for asynchronous tricks
func (ctx *AppContext) Lock() {
	ctx.mutex.Lock()
}

func (ctx *AppContext) Unlock() {
	ctx.mutex.Unlock()
}

func (ctx *AppContext) ResetGorout() {
	ctx.gorout = 0
	ctx.Log.Info.Printf("%s we have now %d simultaneous goroutines", utils.Use().GetStack(ctx.ResetGorout), ctx.gorout)
}

func (ctx *AppContext) AddGorout() {
	ctx.gorout++
	ctx.Log.Info.Printf("%s we have now %d simultaneous goroutines", utils.Use().GetStack(ctx.AddGorout), ctx.gorout)
}

func (ctx *AppContext) RmGorout() {
	ctx.gorout--
	ctx.Log.Info.Printf("%s we have now %d simultaneous goroutines", utils.Use().GetStack(ctx.RmGorout), ctx.gorout)
}

func (ctx *AppContext) GoroutOverflow() bool {
	ctx.Log.Info.Printf("%s we have now %d simultaneous goroutines on %d tolerate", utils.Use().GetStack(ctx.GoroutOverflow), ctx.gorout, ctx.Opts.GetGorout())
	return ctx.Opts.GetGorout() < ctx.gorout
}
