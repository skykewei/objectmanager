package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	Owner struct {
		Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName    string        `json:"firstname"`
		LastName     string        `json:"lastname"`
		Email        string        `json:"email"`
		Password     string        `json:"password,omitempty"`
		HashPassword []byte        `json:"hashpassword,omitempty"`
	}
	LogicalObject struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"createdby"`
		BucketName  string        `json:"bucketname"`
		ObjectKey   string        `json:"objectkey"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Due         time.Time     `json:"due"`
		Version     string        `json:"version,omitempty"`
		Status      string        `json:"status,omitempty"`
		Tags        []string      `json:"tags,omitempty"`
	}
	ObjectOwned struct {
		Id              bson.ObjectId `bson:"_id,omitempty" json:"id"`
		LogicalObjectId bson.ObjectId `json:"logicalobjectid"`
		Description     string        `json:"description"`
		CreatedOn       time.Time     `json:"createdon,omitempty"`
	}
)
