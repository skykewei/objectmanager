package common

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var session *mgo.Seesion

func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{AppConfig.MongoDBHost},
			Username: AppConfig.DBUser,
			Password: AppConfig.DBPwd,
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.Fatalf("[GetSession]:%s\n", err)
		}
	}
	return session
}

func createDbSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConfig.MongoDBHost},
		Username: AppConfig.DBUser,
		Password: AppConfig.DBPwd,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[GetSession]:%s\n", err)
	}
}

//Add indexes into MongoDB
func addIndexes() {
	var err error
	ownerIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
	logicalobjectIndex := mgo.Index{
		Key:        []string{"createdby"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}
	objectownedIndex := mgo.Index{
		Key:        []string{"logicalobjectid"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}

	//Add indexs into MongoDB
	session := GetSession().Copy()
	defer session.Close()

	ownerCol := session.DB(AppConfig.Database).C("owners")
	logicalobjectCol := session.DB(AppConfig.Database).C("logicalobjects")
	objectownedCol := session.DB(AppConfig.Database).C("objectowned")
	err = ownerCol.EnsureIndex(ownerIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
	err = logicalobjectCol.EnsureIndex(logicalobjectIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
	err = objectownedCol.EnsureIndex(ownerIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
}
