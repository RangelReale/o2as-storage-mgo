package mgostorage

import(
	"github.com/RangelReale/o2aserver"
	"labix.org/v2/mgo/bson"
)

type MongoAuthorizationData struct {
	Id bson.ObjectId			`bson:"_id,omitempty"`
	ClientId bson.ObjectId
	Code string
	Scope string
	RedirectUri string
	UserId string
}

func (c *MongoAuthorizationData) importData(authorization *o2aserver.AuthorizationData) *MongoAuthorizationData {
	c.ClientId = bson.ObjectIdHex(authorization.ClientId)
	c.Code = authorization.Code
	c.Scope = authorization.Scope
	c.RedirectUri = authorization.RedirectUri
	c.UserId = authorization.UserId

	return c
}

func (c *MongoAuthorizationData) exportData() *o2aserver.AuthorizationData {
	return &o2aserver.AuthorizationData{
		ClientId: c.ClientId.Hex(),
		Code: c.Code,
		Scope: c.Scope,
		RedirectUri: c.RedirectUri,
		UserId: c.UserId,
	}
}

func mongoAuthorizationDataFromData(authorization *o2aserver.AuthorizationData) *MongoAuthorizationData {
	c := MongoAuthorizationData{}
	return c.importData(authorization)
}


