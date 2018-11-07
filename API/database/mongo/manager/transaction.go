/*
 * File: transaction.go
 * Project: ABD4/VMD Escape Game
 * File Created: Monday, 15th October 2018 10:05:19 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Saturday, 3rd November 2018 2:19:01 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package manager

import (
	"ABD4/API/model"
	"ABD4/API/utils"
	"fmt"
	"strconv"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
 MUST IMPLEMENT

type IDataManager interface {
	GetDB() interface{}
	SetDB(dbObject interface{})
	GetEntity() string
	SetEntity(entity string)
	GeDBName() string
	SetDBName(dbName string)
}

type IUserManager interface {
	IDataManager
	FindAll() ([]*model.User, error)
	FindBy(key, value string) ([]*model.User, error)
	FindOneBy(key, value string) (*model.User, error)
	RemoveBy(key, value string) (int, error)
	RemoveAll() (int, error)
	Create(user *model.User) error
}

*/
type TransactionManager struct {
	session *mgo.Session
	dbName  string
	entity  string
}

func (tm *TransactionManager) parseObjectIds(tx []*model.Transaction) {
	for _, transaction := range tx {
		transaction.ID = bson.ObjectId.Hex(transaction.ObjectID)
	}
}

// IDataManager implementation
func (tm *TransactionManager) Init(params map[string]string) error {
	mandatories := [2]string{"dbName", "entity"}
	for _, key := range mandatories {
		if params[key] == "" {
			return fmt.Errorf("%s missing mandatory: %s", utils.Use().GetStack(tm.Init), key)
		}
	}
	tm.SetDBName(params["dbName"])
	tm.SetEntity(params["entity"])
	return nil
}

func (tm *TransactionManager) SetDB(dbObject interface{}) error {
	var ok bool

	tm.session, ok = dbObject.(*mgo.Session)
	if !ok {
		return fmt.Errorf("%s database object can't be casted in *mgo.Session", utils.Use().GetStack(tm.SetDB))
	}
	return nil
}

func (tm TransactionManager) GetDB() interface{} {
	return tm.session
}

func (tm *TransactionManager) SetEntity(entity string) {
	tm.entity = entity
}

func (tm TransactionManager) GetEntity() string {
	return tm.entity
}

func (tm *TransactionManager) SetDBName(dbName string) {
	tm.dbName = dbName
}

func (tm TransactionManager) GetDBName() string {
	return tm.dbName
}

// TransactionManager implementation

func (tm TransactionManager) FindAll() ([]*model.Transaction, error) {
	c := tm.session.DB(tm.dbName).C(tm.entity)
	results := []*model.Transaction{}
	err := c.Find(bson.M{}).All(&results)
	if err != nil {
		return nil, fmt.Errorf("%s find: %s", utils.Use().GetStack(tm.FindAll), err.Error())
	}
	return results, nil
}

func (tm TransactionManager) FindOneBy(param map[string]string) (*model.Transaction, error) {
	result := &model.Transaction{}
	c := tm.session.DB(tm.dbName).C(tm.entity)
	err := c.Find(utils.Use().MapToBSON(param)).One(result)
	if err != nil {
		return nil, fmt.Errorf("%s find: %s", utils.Use().GetStack(tm.FindOneBy), err.Error())
	}
	return result, nil
}

func (tm TransactionManager) FindBy(param map[string]string) ([]*model.Transaction, error) {
	results := []*model.Transaction{}
	c := tm.session.DB(tm.dbName).C(tm.entity)
	err := c.Find(utils.Use().MapToBSON(param)).All(&results)
	if err != nil {
		return nil, fmt.Errorf("%s find: %s", utils.Use().GetStack(tm.FindBy), err.Error())
	}
	return results, nil
}

func (tm TransactionManager) Create(transaction *model.Transaction) (*model.Transaction, error) {
	transaction.ObjectID = bson.NewObjectId()
	transaction.ID = bson.ObjectId.Hex(transaction.ObjectID)
	c := tm.session.DB(tm.dbName).C(tm.entity)
	err := c.Insert(transaction)
	if err != nil {
		return nil, fmt.Errorf("%s Insert: %s", utils.Use().GetStack(tm.Create), err.Error())
	}
	return transaction, nil
}

func (tm TransactionManager) RemoveAll() (int, error) {
	info := &mgo.ChangeInfo{}
	c := tm.session.
		DB(tm.dbName).
		C(tm.entity)
	info, err := c.RemoveAll(bson.M{})
	return info.Removed, err
}

func (tm TransactionManager) RemoveBy(param map[string]string) (int, error) {
	info := &mgo.ChangeInfo{}
	c := tm.session.DB(tm.dbName).C(tm.entity)
	info, err := c.RemoveAll(utils.Use().MapToBSON(param))
	return info.Removed, err
}

// faker
func getDummyTransactions(n int) []model.Transaction {
	var it int
	tx := []model.Transaction{}
	for it = 0; it < n; it++ {
		reservation := []model.Reservation{
			{
				Spectateur: model.Spectateur{
					Civilite: "Monsieur",
					Nom:      "Carmine",
					Prenom:   "Art",
					Age:      64 + it,
				},
				Tarif: "Senior",
			}, {
				Spectateur: model.Spectateur{
					Civilite: "Madame",
					Nom:      "Nya",
					Prenom:   "Kayla",
					Age:      22 + it,
				},
				Tarif: "Plein tarif",
			},
		}

		game := model.Game{
			Nom:     "Interminable attente chez le medecin",
			Jour:    "2018-09-0" + strconv.Itoa(it),
			Horaire: "05:30",
			VR:      "Non",
		}

		acheteur := model.Acheteur{
			Civilite: "Monsieur",
			Nom:      "Carmine",
			Prenom:   "Art",
			Age:      64 + it,
			Email:    "dummyMail@test.com",
		}

		transaction := model.Transaction{
			Acheteur:    acheteur,
			Game:        game,
			Reservation: reservation,
		}
		tx = append(tx, transaction)
	}
	return tx
}

func (tm TransactionManager) FakePersist(n int) error {
	var iterator int
	var err error

	c := tm.session.DB(tm.dbName).C(tm.entity)
	users := getDummyTransactions(n)
	for iterator = 0; iterator < n; iterator++ {
		err = c.Insert(users[iterator])
		if err != nil {
			return fmt.Errorf("%s Insert: %s", utils.Use().GetStack(tm.FakePersist), err.Error())
		}
	}
	return nil
}
