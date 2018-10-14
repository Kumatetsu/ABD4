/*
 * File: user.go
 * Project: ABD4/VMD Escape Game
 * File Created: Thursday, 11th October 2018 3:32:09 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 14th October 2018 4:34:30 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package manager

import (
	"ABD4/API/model"
	"ABD4/API/utils"
	"fmt"
	"sort"
	"strconv"
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
type UserManager struct {
	session *mgo.Session
	dbName  string
	entity  string
}

func (um *UserManager) parseObjectIds(users []*model.User) {
	for _, user := range users {
		user.ID = bson.ObjectId.Hex(user.ObjectID)
	}
}

// IDataManager implementation
func (um *UserManager) Init(params map[string]string) error {
	mandatories := [2]string{"dbName", "entity"}
	for _, key := range mandatories {
		if params[key] == "" {
			return fmt.Errorf("%s missing mandatory: %s", utils.Use().GetStack(um.Init), key)
		}
	}
	um.SetDBName(params["dbName"])
	um.SetEntity(params["entity"])
	return nil
}

func (um *UserManager) SetDB(dbObject interface{}) error {
	var ok bool

	um.session, ok = dbObject.(*mgo.Session)
	if !ok {
		return fmt.Errorf("%s database object can't be casted in *mgo.Session", utils.Use().GetStack(um.SetDB))
	}
	return nil
}

func (um UserManager) GetDB() interface{} {
	return um.session
}

func (um *UserManager) SetEntity(entity string) {
	um.entity = entity
}

func (um UserManager) GetEntity() string {
	return um.entity
}

func (um *UserManager) SetDBName(dbName string) {
	um.dbName = dbName
}

func (um UserManager) GetDBName() string {
	return um.dbName
}

// IUserManager implementation

func (um UserManager) FindAll() ([]*model.User, error) {
	c := um.session.DB(um.dbName).C(um.entity)
	results := []*model.User{}
	err := c.Find(bson.M{}).All(&results)
	if err != nil {
		return nil, fmt.Errorf("%s find: %s", utils.Use().GetStack(um.FindAll), err.Error())
	}
	um.parseObjectIds(results)
	sort.Slice(results, func(i, j int) bool {
		return results[i].OrderByEmail(results, i, j)
	})
	return results, nil
}

func (um UserManager) FindOneBy(param map[string]string) (*model.User, error) {
	result := &model.User{}
	c := um.session.DB(um.dbName).C(um.entity)
	err := c.Find(utils.Use().MapToBSON(param)).One(result)
	if err != nil {
		return nil, fmt.Errorf("%s find: %s", utils.Use().GetStack(um.FindBy), err.Error())
	}
	result.ID = bson.ObjectId.Hex(result.ObjectID)
	return result, nil
}

func (um UserManager) FindBy(param map[string]string) ([]*model.User, error) {
	results := []*model.User{}
	c := um.session.DB(um.dbName).C(um.entity)
	err := c.Find(utils.Use().MapToBSON(param)).All(&results)
	if err != nil {
		return nil, fmt.Errorf("%s find: %s", utils.Use().GetStack(um.FindBy), err.Error())
	}
	um.parseObjectIds(results)
	sort.Slice(results, func(i, j int) bool {
		return results[i].OrderByEmail(results, i, j)
	})
	return results, nil
}

func (um UserManager) Create(user *model.User) (*model.User, error) {
	user.SetCreatedAt(time.Now())
	user.SetUpdatedAt(time.Now())
	user.ObjectID = bson.NewObjectId()
	c := um.session.DB(um.dbName).C(um.entity)
	err := c.Insert(user)
	if err != nil {
		return nil, fmt.Errorf("%s Insert: %s", utils.Use().GetStack(um.Create), err.Error())
	}
	user.ID = bson.ObjectId.Hex(user.ObjectID)
	return user, nil
}

func (um UserManager) RemoveAll() (int, error) {
	info := &mgo.ChangeInfo{}
	c := um.session.
		DB(um.dbName).
		C(um.entity)
	info, err := c.RemoveAll(bson.M{})
	return info.Removed, err
}

func (um UserManager) RemoveBy(param map[string]string) (int, error) {
	info := &mgo.ChangeInfo{}
	c := um.session.DB(um.dbName).C(um.entity)
	info, err := c.RemoveAll(utils.Use().MapToBSON(param))
	return info.Removed, err
}

// faker
func getDummyUsers(n int) []model.User {
	var it int
	users := []model.User{}
	for it = 0; it < n; it++ {
		name := "test_" + strconv.Itoa(it)
		u := model.User{
			Name:     name,
			Email:    name + "@etna-alternance.net",
			Password: "test",
		}
		users = append(users, u)
	}
	return users
}

func (um UserManager) FakePersist(n int) error {
	var iterator int
	var err error

	c := um.session.DB(um.dbName).C(um.entity)
	users := getDummyUsers(n)
	for iterator = 0; iterator < n; iterator++ {
		err = c.Insert(users[iterator])
		if err != nil {
			return fmt.Errorf("%s Insert: %s", utils.Use().GetStack(um.FakePersist), err.Error())
		}
	}
	return nil
}
