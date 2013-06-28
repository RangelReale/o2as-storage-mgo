package mgostorage

import (
	"github.com/RangelReale/o2aserver"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

// storage

type MongoDBStorage struct {
	session *mgo.Session
	db *mgo.Database
}

func NewMongoDBStorage(url string, database string) (*MongoDBStorage, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &MongoDBStorage{
		session: session,
		db: session.DB(database),
	}, nil
}

func (s *MongoDBStorage) GetClient(clientId string) *o2aserver.Client {
	log.Printf("GetClient: %s\n", clientId)

	if !bson.IsObjectIdHex(clientId) {
		log.Printf("%s is not a client id\n", clientId)
		return nil
	}

	result := MongoClient{}
	err := s.db.C("client").Find(bson.M{"_id": bson.ObjectIdHex(clientId), "enabled": true}).One(&result)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return nil
	}
	return result.exportData()
}

func (s *MongoDBStorage) SaveClient(client *o2aserver.Client) error {
	info, err := s.db.C("client").UpsertId(client.Id, mongoClientFromData(client))
	if err != nil {
		return err
	}
	if info.UpsertedId != nil {
		client.Id = info.UpsertedId.(string)
	}
	return nil
}

func (s *MongoDBStorage) GetAuthorize(code string) (*o2aserver.AuthorizationData, error) {
	log.Printf("LoadAuthorize: code:%s\n", code)

	result := MongoAuthorizationData{}
	err := s.db.C("authorization").Find(bson.M{"code": code}).One(&result)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return nil, err
	}
	return result.exportData(), nil
}

func (s *MongoDBStorage) SaveAuthorize(parameters o2aserver.AuthorizationData) error {
	log.Printf("SaveAuthorize: clientId:%s; code:%s; rediruri:%s; scope:%s\n", parameters.ClientId, parameters.Code, parameters.RedirectUri, parameters.Scope)

	_, err := s.db.C("authorization").Upsert(bson.M{"code": parameters.Code}, mongoAuthorizationDataFromData(&parameters))
	if err != nil {
		log.Printf("Error: %s\n", err)
		return err
	}
	return nil
}

func (s *MongoDBStorage) RemoveAuthorize(code string) error {
	log.Printf("RemoveAuthorize: code:%s\n", code)

	err := s.db.C("authorization").Remove(bson.M{"code": code})
	if err != nil {
		log.Printf("Error: %s\n", err)
		return err
	}
	return nil
}

func (s *MongoDBStorage) GetAccessToken(code string) (*o2aserver.AccessTokenData, error) {
	log.Printf("GetAccessToken: code:%s\n", code)

	result := MongoAccessTokenData{}
	err := s.db.C("accesstoken").Find(bson.M{"accesstoken": code}).One(&result)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return nil, err
	}
	return result.exportData(), nil
}

func (s *MongoDBStorage) SaveAccessToken(parameters o2aserver.AccessTokenData) error {
	log.Printf("SaveAccessToken: clientId:%s; accessToken:%s; rediruri:%s; scope:%s\n", parameters.ClientId, parameters.AccessToken, parameters.RedirectUri, parameters.Scope)

	_, err := s.db.C("accesstoken").Upsert(bson.M{"accesstoken": parameters.AccessToken}, mongoAccessTokenDataFromData(&parameters))
	if err != nil {
		log.Printf("Error: %s\n", err)
		return err
	}
	return nil
}
