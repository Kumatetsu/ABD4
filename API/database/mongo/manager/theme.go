/*
 * File: themes.go
 * Project: ABD4/VMD Escape Game
 * File Created: Tuesday, 30th October 2018 6:43:36 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 30th October 2018 9:09:21 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package manager

import (
	"ABD4/API/model"
	"ABD4/API/utils"
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ThemeManager struct {
	session *mgo.Session
	dbName  string
	entity  string
}

// IDataManager implementation
func (tm *ThemeManager) Init(params map[string]string) error {
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

func (tm *ThemeManager) SetDB(dbObject interface{}) error {
	var ok bool

	tm.session, ok = dbObject.(*mgo.Session)
	if !ok {
		return fmt.Errorf("%s database object can't be casted in *mgo.Session", utils.Use().GetStack(tm.SetDB))
	}
	return nil
}

func (tm ThemeManager) GetDB() interface{} {
	return tm.session
}

func (tm *ThemeManager) SetEntity(entity string) {
	tm.entity = entity
}

func (tm ThemeManager) GetEntity() string {
	return tm.entity
}

func (tm *ThemeManager) SetDBName(dbName string) {
	tm.dbName = dbName
}

func (tm ThemeManager) GetDBName() string {
	return tm.dbName
}

func (tm ThemeManager) FindAll() ([]*model.Theme, error) {
	c := tm.session.DB(tm.dbName).C(tm.entity)
	results := []*model.Theme{}
	err := c.Find(bson.M{}).All(&results)
	if err != nil {
		return nil, fmt.Errorf("%s find: %s", utils.Use().GetStack(tm.FindAll), err.Error())
	}
	return results, nil
}

func (tm ThemeManager) FindOneBy(param map[string]string) (*model.Theme, error) {
	result := &model.Theme{}
	c := tm.session.DB(tm.dbName).C(tm.entity)
	err := c.Find(utils.Use().MapToBSON(param)).One(result)
	if err != nil {
		return nil, fmt.Errorf("%s find: %s", utils.Use().GetStack(tm.FindBy), err.Error())
	}
	return result, nil
}

func (tm ThemeManager) FindBy(param map[string]string) ([]*model.Theme, error) {
	results := []*model.Theme{}
	c := tm.session.DB(tm.dbName).C(tm.entity)
	err := c.Find(utils.Use().MapToBSON(param)).All(&results)
	if err != nil {
		return nil, fmt.Errorf("%s find: %s", utils.Use().GetStack(tm.FindBy), err.Error())
	}
	return results, nil
}

func (tm ThemeManager) RemoveAll() (int, error) {
	info := &mgo.ChangeInfo{}
	c := tm.session.
		DB(tm.dbName).
		C(tm.entity)
	info, err := c.RemoveAll(bson.M{})
	return info.Removed, err
}

func (tm ThemeManager) RemoveBy(param map[string]string) (int, error) {
	info := &mgo.ChangeInfo{}
	c := tm.session.DB(tm.dbName).C(tm.entity)
	info, err := c.RemoveAll(utils.Use().MapToBSON(param))
	return info.Removed, err
}

func (tm ThemeManager) Create(theme *model.Theme) (*model.Theme, error) {
	theme.ObjectID = bson.NewObjectId()
	theme.ID = bson.ObjectId.Hex(theme.ObjectID)
	c := tm.session.DB(tm.dbName).C(tm.entity)
	err := c.Insert(theme)
	if err != nil {
		return nil, fmt.Errorf("%s Insert: %s", utils.Use().GetStack(tm.Create), err.Error())
	}
	return theme, nil
}
