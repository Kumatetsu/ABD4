/*
 * File: elastic.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 4th November 2018 10:58:56 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 5th November 2018 4:54:32 am
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package context

import (
	"ABD4/API/elasticsearch"
	"ABD4/API/iserial"
	"ABD4/API/utils"
	goctx "context"
	"fmt"
	"math"
	"runtime"
	"time"
)

// indexData take data actually in database and send it to elastic correspondant index
// there is nothing to prevent the reindexation of data
// its easy to fix a bad indexation using rmindex/index/reindex options
func (ctx *AppContext) indexData() error {
	ctx.Log.Info.Printf("%s managers: %s, %s, %s, %s",
		utils.Use().GetStack(ctx.indexData),
		ctx.UserManager.GetEntity(),
		ctx.TransactionManager.GetEntity(),
		ctx.TarifManager.GetEntity(),
		ctx.ThemeManager.GetEntity())
	if ctx.UserManager.GetDBName() == "" ||
		ctx.TransactionManager.GetDBName() == "" ||
		ctx.TarifManager.GetDBName() == "" ||
		ctx.ThemeManager.GetDBName() == "" {
		return fmt.Errorf("%s At list one data manager is missing", utils.Use().GetStack(ctx.indexData))
	}
	users, err := ctx.UserManager.FindAll()
	if err != nil {
		ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.indexData), err.Error())
	}
	var serialUsers []iserial.Serializable
	for _, u := range users {
		serialUsers = append(serialUsers, u.ToES())
	}
	err = ctx.IndexArrayData(serialUsers, USERS, USER)
	if err != nil {
		ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.indexData), err.Error())
	}
	tx, err := ctx.TransactionManager.FindAll()
	if err != nil {
		ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.indexData), err.Error())
	}
	var serialTx []iserial.Serializable
	for _, transaction := range tx {
		serialTx = append(serialTx, transaction.ToES())
	}
	err = ctx.IndexArrayData(serialTx, TXs, TX)
	if err != nil {
		ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.indexData), err.Error())
	}
	tarifs, err := ctx.TarifManager.FindAll()
	if err != nil {
		ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.indexData), err.Error())
	}
	var serialTarifs []iserial.Serializable
	for _, tarif := range tarifs {
		serialTarifs = append(serialTarifs, tarif.ToES())
	}
	err = ctx.IndexArrayData(serialTarifs, TARIFS, TARIF)
	if err != nil {
		ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.indexData), err.Error())
	}
	themes, err := ctx.ThemeManager.FindAll()
	if err != nil {
		ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.indexData), err.Error())
	}
	var serialThemes []iserial.Serializable
	for _, theme := range themes {
		serialThemes = append(serialThemes, theme.ToES())
	}
	err = ctx.IndexArrayData(serialThemes, THEMES, THEME)
	if err != nil {
		ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.indexData), err.Error())
	}
	return nil
}

// IndexData send a iserial.Serializable entity to elasticsearch
// index and t are elasticsearch target index and entity type:
// 'abd4-transactions' and 'transaction' for exemple
func (ctx *AppContext) IndexData(i iserial.Serializable, index, t string) error {
	ctx.SavedTime = ctx.Time
	op, err := ctx.ElasticClient.Index().
		Index(index).
		Type(t).
		BodyJson(i).
		Refresh("true").
		Do(goctx.Background())
	if err != nil {
		ctx.Log.Error.Printf("%s %s", utils.Use().GetStack(ctx.IndexData), err.Error())
		return err
	}
	ctx.Log.Info.Printf("%s insert unit %s to index %s, entity type: %s", utils.Use().GetStack(ctx.IndexData), op.Id, op.Index, op.Type)
	return nil
}

// embedElasticSearch is the entry point to settle an elasticsearch context linked with this api
// User and Transaction entities are indexed
// This can be deactivated using flag embedES=false
func (ctx *AppContext) embedElasticSearch() {
	var err error

	ctx.ElasticClient, err = elasticsearch.Instanciate(ctx.Opts.GetEs())
	if err != nil {
		ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.Instanciate), err.Error())
	}
	// if flag rmindex is passed to .exe, indexes 'users' and 'transactions' are removed
	if ctx.Opts.GetRmindex() {
		ctx.Log.Info.Printf("%s removing indexes", utils.Use().GetStack(ctx.embedElasticSearch))
		for _, index := range INDEXES {
			err = elasticsearch.RemoveIndex(ctx.ElasticClient, index)
			if err != nil {
				ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.Instanciate), err.Error())
			}
		}
	}
	// if index or reindex flag is present
	// index create indexes if don't exist
	// reindex remove and create indexes
	// if reindex, data are pulled from database and send to elasticsearch
	// there is nothing to prevent duplicate in elastic
	if ctx.Opts.GetIndex() || ctx.Opts.GetReindex() {
		if ctx.Opts.GetReindex() {
			ctx.Log.Info.Printf("%s reindexation", utils.Use().GetStack(ctx.embedElasticSearch))
		} else {
			ctx.Log.Info.Printf("%s indexation", utils.Use().GetStack(ctx.embedElasticSearch))
		}
		err = elasticsearch.CreateIndexation(ctx.ElasticClient, ctx.Opts.GetReindex())
		if err != nil {
			ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.Instanciate), err.Error())
		}
		if ctx.Opts.GetReindex() {
			err = ctx.indexData()
			if err != nil {
				ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.Instanciate), err.Error())
			}
		}
	}
}

// indexArrayData call IndexData on each item in serializables
func (ctx *AppContext) indexArrayData(serializables []iserial.Serializable, index, t string) error {
	for _, i := range serializables {
		err := ctx.IndexData(i, index, t)
		if err != nil {
			return err
		}
	}
	return nil
}

// IndexArrayData, call IndexData on an array of iserial.Serializable
// Accept asynchronous process to speed up indexation
func (ctx *AppContext) IndexArrayData(serializables []iserial.Serializable, index, t string) error {
	var err error
	qty := len(serializables)
	ctx.Log.Info.Printf("%s will index %d data", utils.Use().GetStack(ctx.IndexArrayData), qty)
	// si on autorise l'asynchrone et que ca le mérite
	if qty > ctx.Opts.GetBatch() && ctx.Opts.GetAllowAsync() {
		runtime.GOMAXPROCS(4)
		// On lance des goroutines pour accélérer le traitement
		err = ctx.AsynchronousBatching(serializables, index, t)
		runtime.GOMAXPROCS(1)
	} else {
		err = ctx.indexArrayData(serializables, index, t)
	}
	return err
}

// RemoveIndex destroy completely an index
func (ctx *AppContext) RemoveIndex(index string) error {
	return elasticsearch.RemoveIndex(ctx.ElasticClient, index)
}

// processBatch is design to be used in async process
func indexBatchOfSerializables(ctx *AppContext, serializables []iserial.Serializable, index, t string, errCh chan error, success chan bool) {
	ctx.Log.Info.Printf("%s threatig batch of %d data", utils.Use().GetStack(indexBatchOfSerializables), len(serializables))
	for _, i := range serializables {
		err := ctx.IndexData(i, index, t)
		if err != nil {
			ctx.Log.Info.Printf("%s threatBatch: errCh <- err, if log stop here, its a deadlock issue", utils.Use().GetStack(indexBatchOfSerializables))
			errCh <- err
			break
		}
	}
	ctx.Log.Info.Printf("%s threatBatch: success <- true, if log stop here, its a deadlock issue", utils.Use().GetStack(indexBatchOfSerializables))
	success <- true
}

// poolIndexation is design to be used in async process
// on considère que les channels errChan et signals sont sur écoute
// pools doit correspondre au nombre d'entrées dans serializable
// voir ctx.defineRequiredPools
func (ctx *AppContext) poolIndexation(
	serializables []iserial.Serializable,
	index,
	t string,
	pools int,
	errChan chan error,
	signals []chan bool,
) {
	batcher := ctx.Opts.GetBatch()
	// si le nombre de pools définis ne permet pas de respecter le batch souhaité
	if float64(len(serializables))/float64(pools) > float64(batcher) {
		// on définis le batcher comme étant le nombre d'éléments par le nombre total de pools
		batcher = int(math.Trunc(float64(len(serializables)) / float64(pools)))
	}
	// si il y a moins de channels sur écoute que de pools
	// la fonction ne peut pas allez au bout.
	// On tolère un excédent mais cela implique un risque de deadlock
	if pools > len(signals) {
		if errChan == nil {
			return
		}
		errChan <- fmt.Errorf("%s Not enough signals to synchronise pools", utils.Use().GetStack(ctx.poolIndexation))
	}
	for batch := 0; batch < pools; batch++ {
		var wagon []iserial.Serializable
		// le point d'entrée dans la slice est la taille du batcher multiplié par le numéro de la pool
		seg := batch * batcher
		ctx.Log.Info.Printf("%s batch %d for pools %d", utils.Use().GetStack(ctx.poolIndexation), batch, pools)
		// si c'est la dernière fournée, on découpe un wagon de seg : end
		// sinon, on découpe un batch
		if batch == int(pools)-1 {
			ctx.Log.Info.Printf("%s wagon %d to end", utils.Use().GetStack(ctx.poolIndexation), seg)
			wagon = serializables[seg:]
		} else {
			ctx.Log.Info.Printf("%s wagon %d to %d", utils.Use().GetStack(ctx.poolIndexation), seg, seg+batcher)
			wagon = serializables[seg : seg+batcher]
		}
		ctx.AddGorout()
		for ctx.GoroutOverflow() {
			time.Sleep(time.Millisecond * 50)
		}
		// on passe les channels et le numéro de pool en cours par closure (boucle for golang asynchrone)
		// on index en passant le chan d'erreur et le chan signal correspondant au batch
		go func(innerErrChan chan error, innerSignals []chan bool, innerBatch int) {
			ctx.benchmarkStart()
			defer ctx.benchmarkStop(fmt.Sprintf("indexation batch of %d elements", len(wagon)))
			defer ctx.RmGorout()
			ctx.Log.Info.Printf("%s batch %d goroutine threath batch", utils.Use().GetStack(ctx.poolIndexation), innerBatch)
			indexBatchOfSerializables(ctx, wagon, index, t, errChan, innerSignals[innerBatch])
		}(errChan, signals, batch)
	}
}
