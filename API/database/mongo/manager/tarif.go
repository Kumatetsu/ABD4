/*
 * File: tarif.go
 * Project: ABD4/VMD Escape Game
 * File Created: Monday, 15th October 2018 10:05:19 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 30th October 2018 9:05:53 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package manager

import (
	"ABD4/API/model"
	"ABD4/API/utils"
	"fmt"
	"time"

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

 type ITarifManager interface {
	IDataManager
	FindAll() ([]*model.Tarif, error)
	FindBy(map[string]string) ([]*model.Tarif, error)
	FindOneBy(map[string]string) (*model.Tarif, error)
	RemoveBy(map[string]string) (int, error)
	RemoveAll() (int, error)
	Create(tx *model.Tarif) (*model.Tarif, error)
	Load(dir string) error
	GetTarifs() map[string]float64
}

*/
type TarifManager struct {
	session *mgo.Session
	dbName  string
	entity  string
	// Description: Prix, représentation en type générique
	// du type privé model.tarif, évite d'exporter le champ
	tarifs map[string]float64
}

func (tm *TarifManager) parseObjectIds(tarifs []*model.Tarif) {
	for _, tarif := range tarifs {
		tarif.ID = bson.ObjectId.Hex(tarif.ObjectID)
	}
}

// IDataManager implementation
func (tm *TarifManager) Init(params map[string]string) error {
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

func (tm *TarifManager) SetDB(dbObject interface{}) error {
	var ok bool

	tm.session, ok = dbObject.(*mgo.Session)
	if !ok {
		return fmt.Errorf("%s database object can't be casted in *mgo.Session", utils.Use().GetStack(tm.SetDB))
	}
	return nil
}

func (tm TarifManager) GetDB() interface{} {
	return tm.session
}

func (tm *TarifManager) SetEntity(entity string) {
	tm.entity = entity
}

func (tm TarifManager) GetEntity() string {
	return tm.entity
}

func (tm *TarifManager) SetDBName(dbName string) {
	tm.dbName = dbName
}

func (tm TarifManager) GetDBName() string {
	return tm.dbName
}

func (tm *TarifManager) GetTarifs() map[string]float64 {
	return tm.tarifs
}

// IUserManager implementation

func (tm TarifManager) FindAll() ([]*model.Tarif, error) {
	c := tm.session.DB(tm.dbName).C(tm.entity)
	results := []*model.Tarif{}
	err := c.Find(bson.M{}).All(&results)
	if err != nil {
		return nil, fmt.Errorf("%s find: %s", utils.Use().GetStack(tm.FindAll), err.Error())
	}
	return results, nil
}

func (tm TarifManager) FindOneBy(param map[string]string) (*model.Tarif, error) {
	result := &model.Tarif{}
	c := tm.session.DB(tm.dbName).C(tm.entity)
	err := c.Find(utils.Use().MapToBSON(param)).One(result)
	if err != nil {
		return nil, fmt.Errorf("%s find: %s", utils.Use().GetStack(tm.FindBy), err.Error())
	}
	return result, nil
}

func (tm TarifManager) FindBy(param map[string]string) ([]*model.Tarif, error) {
	results := []*model.Tarif{}
	c := tm.session.DB(tm.dbName).C(tm.entity)
	err := c.Find(utils.Use().MapToBSON(param)).All(&results)
	if err != nil {
		return nil, fmt.Errorf("%s find: %s", utils.Use().GetStack(tm.FindBy), err.Error())
	}
	return results, nil
}

func (tm TarifManager) Create(tarif *model.Tarif) (*model.Tarif, error) {
	tarif.ObjectID = bson.NewObjectId()
	tarif.ID = bson.ObjectId.Hex(tarif.ObjectID)
	c := tm.session.DB(tm.dbName).C(tm.entity)
	_, err := time.Parse("2006-01-02T15:04:05Z", tarif.Date)
	if tarif.Date == "" || err != nil {
		tarif.Date = time.Now().Format("2006-01-02T15:04:05Z")
	}
	err = c.Insert(tarif)
	if err != nil {
		return nil, fmt.Errorf("%s Insert: %s", utils.Use().GetStack(tm.Create), err.Error())
	}
	return tarif, nil
}

func (tm TarifManager) RemoveAll() (int, error) {
	info := &mgo.ChangeInfo{}
	c := tm.session.
		DB(tm.dbName).
		C(tm.entity)
	info, err := c.RemoveAll(bson.M{})
	return info.Removed, err
}

func (tm TarifManager) RemoveBy(param map[string]string) (int, error) {
	info := &mgo.ChangeInfo{}
	c := tm.session.DB(tm.dbName).C(tm.entity)
	info, err := c.RemoveAll(utils.Use().MapToBSON(param))
	return info.Removed, err
}
