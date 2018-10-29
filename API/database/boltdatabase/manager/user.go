/*
 * File: user.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 16th October 2018 5:09:56 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package manager

import (
	"ABD4/API/database/boltdatabase"
	"ABD4/API/model"
	"ABD4/API/utils"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
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
	dbm        *boltdatabase.DBManager
	dbName     string
	entity     string
	bucket     string
	userPrefix string
}

// IDataManager implementation
func (um UserManager) SetDB(dbObject interface{}) error {
	var ok bool
	um.dbm, ok = dbObject.(*boltdatabase.DBManager)
	if !ok {
		return fmt.Errorf("%s database object can't be casted in *boltdatabase.DBManager", utils.Use().GetStack(um.SetDB))
	}
	return nil
}

func (um UserManager) GetDB() interface{} {
	return um.dbm
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

func (um *UserManager) Init(params map[string]string) error {
	mandatories := [3]string{"name", "fullpath", "secret"}
	for _, key := range mandatories {
		if params[key] == "" {
			return fmt.Errorf("%s missing mandatory: %s", utils.Use().GetStack(um.Init), key)
		}
	}
	name := params["name"]
	fullpath := params["fullpath"]
	secret := params["secret"]

	dbm, err := boltdatabase.NewDBManager(name, fullpath, secret, false, []string{"al-user"})
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(um.Init), err.Error())
	}
	if dbm != nil {
		um.dbm = dbm
		um.bucket = "al-user"
		um.userPrefix = "a-"
	}

	utils.Use().InitRand()
	return nil
}

func (um UserManager) GetDBM() *boltdatabase.DBManager {
	return um.dbm
}

func (um UserManager) NewUser(name string, email string, pass string, permission string, claim string) (*model.User, error) {
	user := &model.User{
		Name:       name,
		Password:   pass,
		Email:      email,
		Permission: permission,
		Claim:      claim,
	}

	user, err := um.Create(user)
	if err != nil {
		return nil, fmt.Errorf("<<<< %s %s", utils.Use().GetStack(um.NewUser), err.Error())
	}
	return user, nil
}

func (um UserManager) compareKeyValue(key, value string, user *model.User) (bool, error) {
	key = strings.Title(key)
	switch key {
	case "Name":
		return user.Name == value, nil
	case "Password":
		return user.Password == value, nil
	case "Email":
		return user.Email == value, nil
	case "Permission":
		return user.Permission == value, nil
	case "Claim":
		return user.Claim == value, nil
	case "CreatedAt":
		date, err := time.Parse("2006-12-24", value)
		if err != nil {
			return false, fmt.Errorf("%s %s", utils.Use().GetStack(um.compareKeyValue), err.Error())
		}
		return user.GetCreatedAt() == date, nil
	case "UpdatedAt":
		date, err := time.Parse("2006-12-24", value)
		if err != nil {
			return false, fmt.Errorf("%s %s", utils.Use().GetStack(um.compareKeyValue), err.Error())
		}
		return user.GetUpdatedAt() == date, nil
	default:
		return false, fmt.Errorf("%s %s: %s", utils.Use().GetStack(um.compareKeyValue), "unvalid key", key)
	}
}

func (um UserManager) seekBy(params map[string]string) ([]*model.User, error) {
	users, err := um.FindAll()
	if err != nil {
		return nil, fmt.Errorf("%s %s", utils.Use().GetStack(um.seekBy), err.Error())
	}
	rUsers := []*model.User{}
	for _, user := range users {
		check := false
		for key, value := range params {
			check, err = um.compareKeyValue(key, value, user)
			if err != nil {
				return rUsers, fmt.Errorf("%s %s", utils.Use().GetStack(um.seekBy), err.Error())
			}
			if check == false {
				break
			}
		}
		if check {
			rUsers = append(rUsers, user)
		}
	}
	return rUsers, nil
}

func (um UserManager) seekOneBy(params map[string]string) (*model.User, error) {
	users, err := um.FindAll()
	if err != nil {
		return nil, fmt.Errorf("%s %s", utils.Use().GetStack(um.seekOneBy), err.Error())
	}
	for _, user := range users {
		check := false
		fmt.Printf("\nuser email: %s, password: %s", user.Email, user.Password)
		for key, value := range params {
			fmt.Printf("\nseeking key: %s value: %s", key, value)
			check, err = um.compareKeyValue(key, value, user)
			if err != nil {
				return nil, fmt.Errorf("%s %s", utils.Use().GetStack(um.seekOneBy), err.Error())
			}
			fmt.Printf("\nValue of check: %v", check)
			if check == false {
				continue
			}
		}
		if check == true {
			return user, nil
		}
	}
	return nil, fmt.Errorf("%s %s", utils.Use().GetStack(um.seekOneBy), "No user found")
}

func (um UserManager) FindAll() (results []*model.User, err error) {
	if um.dbm == nil {
		return
	}
	results = make([]*model.User, 0)

	var byteResults [][]byte

	if byteResults, err = um.dbm.GetByPrefix(um.bucket, um.userPrefix); err != nil {
		return nil, fmt.Errorf("<<<< %s : %s", utils.Use().GetStack(um.FindAll), err.Error())
	}

	for _, iter := range byteResults {
		resNew := new(model.User)
		if err = json.Unmarshal([]byte(iter), resNew); err != nil {
			err = fmt.Errorf("<<<< %s : %s", utils.Use().GetStack(um.FindAll), err.Error())
			return nil, err
		} else {
			results = append(results, resNew)
		}
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].OrderByEmail(results, i, j)
	})
	return
}

func (um UserManager) FindBy(params map[string]string) (result []*model.User, err error) {
	users, err := um.seekBy(params)
	if err != nil {
		return nil, fmt.Errorf("%s %s", utils.Use().GetStack(um.FindBy), err.Error())
	}
	return users, nil
}

func (um UserManager) FindOneBy(params map[string]string) (*model.User, error) {
	user, err := um.seekOneBy(params)
	if err != nil {
		return nil, fmt.Errorf("%s %s %s", utils.Use().GetStack(um.RemoveBy), "failed to retrieve users", err.Error())
	}
	return user, nil
}

func (um UserManager) FindByID(id string) (result *model.User, err error) {
	if id == "" {
		return nil, fmt.Errorf("<<<< %s: id is empty", utils.Use().GetStack(um.FindByID))
	}

	var key string
	key = fmt.Sprintf("%s%s", um.userPrefix, id)

	var byt []byte
	byt, err = um.dbm.GetOne(um.bucket, key)

	result = new(model.User)
	if err = json.Unmarshal([]byte(byt), result); err != nil {
		return nil, fmt.Errorf("<<<< %s %s", utils.Use().GetStack(um.FindByID))
	}
	return
}

func (um UserManager) Create(record *model.User) (*model.User, error) {
	if record == nil {
		return nil, fmt.Errorf("<<<< %s record is nil", utils.Use().GetStack(um.Create))
	}
	if record.ID == "" || record.ID == "0" {
		record.ID = utils.Use().RandStringRunes(12)
	}

	record.SetCreatedAt(time.Now())
	record.SetUpdatedAt(time.Now())
	key := fmt.Sprintf("%s%s", um.userPrefix, record.ID)
	err := um.dbm.Save(um.bucket, key, record)
	if err != nil {
		return nil, fmt.Errorf("<<<< %s %s", utils.Use().GetStack(um.Create), err.Error())
	}
	return record, nil
}

func (um UserManager) Update(record *model.User) (*model.User, error) {
	if record == nil || record.ID == "0" {
		return nil, fmt.Errorf("<<<< %s record is nil", utils.Use().GetStack(um.Update))
	}

	var oldRecord *model.User
	var err error

	if oldRecord, err = um.FindByID(record.ID); err != nil {
		return nil, fmt.Errorf("<<<< %s record is nil", utils.Use().GetStack(um.Update))
	}
	//update the time values from the old record
	err = um.RemoveByID(record.ID)
	if err != nil {
		return record, fmt.Errorf("<<<< %s %s", utils.Use().GetStack(um.Update), err.Error())
	}
	record.SetCreatedAt(oldRecord.GetCreatedAt())
	record.SetUpdatedAt(time.Now())
	err = um.dbm.Save(um.bucket, fmt.Sprintf("%s%s", um.userPrefix, record.GetID()), record)
	if err != nil {
		return record, fmt.Errorf("<<<< %s record is nil", utils.Use().GetStack(um.Update))
	}
	return record, nil
}

func (um UserManager) RemoveByID(id string) error {
	if id == "" {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(um.RemoveByID), boltdatabase.ErrKeyInvalid.Error())
	}
	key := fmt.Sprintf("%s%s", um.userPrefix, id)
	err := um.dbm.Delete(um.bucket, key)
	if err != nil {
		return fmt.Errorf("<<<< %s record is nil", utils.Use().GetStack(um.Update))
	}
	return nil
}

func (um UserManager) RemoveBy(params map[string]string) (int, error) {
	deleted := 0
	users, err := um.seekBy(params)
	if err != nil {
		return deleted, fmt.Errorf("%s %s %s", utils.Use().GetStack(um.RemoveBy), "failed to retrieve users", err.Error())
	}
	for _, user := range users {
		um.RemoveByID(user.ID)
		if err != nil {
			return deleted, fmt.Errorf("%s %s %s", utils.Use().GetStack(um.RemoveAll), "failed to delete user", err.Error())
		}
		deleted++
	}
	return deleted, nil
}

func (um UserManager) RemoveAll() (int, error) {
	deleted := 0
	users, err := um.FindAll()
	if err != nil {
		return deleted, fmt.Errorf("%s %s %s", utils.Use().GetStack(um.RemoveAll), "failed to retrieve users", err.Error())
	}
	for _, user := range users {
		err = um.RemoveByID(user.ID)
		if err != nil {
			return deleted, fmt.Errorf("%s %s %s", utils.Use().GetStack(um.RemoveAll), "failed to delete user", err.Error())
		}
		deleted++
	}
	return deleted, nil
}

func (um UserManager) TestSaveUser() {
	var err error

	user, err := um.NewUser("test-name", "test-email", "test-pass", "test-permission", "test-claim")
	if err != nil {
		fmt.Printf("NewUser return err: %s", err)
	}

	if user.ID == "" {
		fmt.Print("NewUser Id is nil")
	}

	retRecord, err := um.FindByID(user.ID)
	if err != nil {
		fmt.Printf("GetByID return err: %s", err)
	} else {
		fmt.Printf("%+v", retRecord)
	}

	if user.Name != retRecord.Name {
		fmt.Print("GetByID returned values are not equal")
	}

	return
}

// fakers

func getDummyUsers(n int) []*model.User {
	var it int
	users := []*model.User{}
	for it = 0; it < n; it++ {
		name := "test_" + strconv.Itoa(it)
		u := &model.User{
			Name:     name,
			Email:    name + "@etna-alternance.net",
			Password: "test",
		}
		users = append(users, u)
	}
	return users
}

func (um UserManager) FakePersist(n int) error {
	users := getDummyUsers(n)
	for _, user := range users {
		_, err := um.Create(user)
		if err != nil {
			return fmt.Errorf("%s %s", utils.Use().GetStack(um.FakePersist), err.Error())
		}
	}
	return nil
}
