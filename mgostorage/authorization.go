package mgostorage

import (
	"github.com/RangelReale/o2aserver"
	"labix.org/v2/mgo/bson"
	"time"
)

type MongoAuthorizationData struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	ClientId    bson.ObjectId
	Code        string
	ExpiresIn   int64
	Scope       string
	RedirectUri string
	UserId      string
	CreatedAt   time.Time
}

func (c *MongoAuthorizationData) importData(authorization *o2aserver.AuthorizationData) *MongoAuthorizationData {
	c.ClientId = bson.ObjectIdHex(authorization.ClientId)
	c.Code = authorization.Code
	c.ExpiresIn = authorization.ExpiresIn
	c.Scope = authorization.Scope
	c.RedirectUri = authorization.RedirectUri
	c.UserId = authorization.UserId
	c.CreatedAt = authorization.CreatedAt

	return c
}

func (c *MongoAuthorizationData) exportData() *o2aserver.AuthorizationData {
	return &o2aserver.AuthorizationData{
		ClientId:    c.ClientId.Hex(),
		Code:        c.Code,
		ExpiresIn:   c.ExpiresIn,
		Scope:       c.Scope,
		RedirectUri: c.RedirectUri,
		UserId:      c.UserId,
		CreatedAt:   c.CreatedAt,
	}
}

func mongoAuthorizationDataFromData(authorization *o2aserver.AuthorizationData) *MongoAuthorizationData {
	c := MongoAuthorizationData{}
	return c.importData(authorization)
}
