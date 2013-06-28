package mgostorage

import(
	"github.com/RangelReale/o2aserver"
	"labix.org/v2/mgo/bson"
)

// o2aserver.Client wrapper
type MongoClient struct {
	Id			bson.ObjectId		`bson:"_id,omitempty"`
	Secret		string
	Name		string
	RedirectUri	string
	Enabled		bool
}

func (c *MongoClient) importData(client *o2aserver.Client) *MongoClient {
	c.Id = bson.ObjectIdHex(client.Id)
	c.Secret = client.Secret
	c.Name = client.Name
	c.RedirectUri = client.RedirectUri
	c.Enabled = client.Enabled

	return c
}

func (c *MongoClient) exportData() *o2aserver.Client {
	return &o2aserver.Client{
		Id: c.Id.Hex(),
		Secret: c.Secret,
		Name: c.Name,
		RedirectUri: c.RedirectUri,
		Enabled: c.Enabled,
	}
}

func mongoClientFromData(client *o2aserver.Client) *MongoClient {
	c := MongoClient{}
	return c.importData(client)
}


