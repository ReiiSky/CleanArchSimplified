package infrastructure

import (
	"github.com/Kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var defaultMongoClient *mongo.Client

// SetupMongoConnection this code just a shorthand to setup default mongo client
func SetupMongoConnection(dbname, mongoURL string) error {
	mgm.SetDefaultConfig(nil, dbname, options.Client().ApplyURI(mongoURL))
	_, c, _, err := mgm.DefaultConfigs()
	if err == nil {
		SetDefaultClient(c)
	}
	return err
}

// SetDefaultClient set default mongo client variable
func SetDefaultClient(client *mongo.Client) {
	defaultMongoClient = client
}

// MongoClientPing tryng to ping connection
func MongoClientPing(client *mongo.Client) error {
	return client.Ping(mgm.Ctx(), nil)
}
