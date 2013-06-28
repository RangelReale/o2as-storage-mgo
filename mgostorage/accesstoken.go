package mgostorage

import(
	"github.com/RangelReale/o2aserver"
	"labix.org/v2/mgo/bson"
	"time"
)

type MongoAccessTokenData struct {
	Id bson.ObjectId			`bson:"_id,omitempty"`
	ClientId bson.ObjectId
	AccessToken string
	RefreshToken string
	ExpiresIn int64
	Scope string
	RedirectUri string
	UserId string
	CreatedAt time.Time
}


func (c *MongoAccessTokenData) importData(accesstoken *o2aserver.AccessTokenData) *MongoAccessTokenData {
	c.ClientId = bson.ObjectIdHex(accesstoken.ClientId)
	c.AccessToken = accesstoken.AccessToken
	c.RefreshToken = accesstoken.RefreshToken
	c.ExpiresIn = accesstoken.ExpiresIn
	c.Scope = accesstoken.Scope
	c.RedirectUri = accesstoken.RedirectUri
	c.UserId = accesstoken.UserId
	c.CreatedAt = accesstoken.CreatedAt

	return c
}

func (c *MongoAccessTokenData) exportData() *o2aserver.AccessTokenData {
	return &o2aserver.AccessTokenData{
		ClientId: c.ClientId.Hex(),
		AccessToken: c.AccessToken,
		RefreshToken: c.RefreshToken,
		ExpiresIn: c.ExpiresIn,
		Scope: c.Scope,
		RedirectUri: c.RedirectUri,
		UserId: c.UserId,
		CreatedAt: c.CreatedAt,
	}
}

func mongoAccessTokenDataFromData(accesstoken *o2aserver.AccessTokenData) *MongoAccessTokenData {
	c := MongoAccessTokenData{}
	return c.importData(accesstoken)
}




