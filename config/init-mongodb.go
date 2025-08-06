package config

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InitMongoDB() {
	err := mgm.SetDefaultConfig(nil, Env.DB_NAME, options.Client().ApplyURI(Env.MONGO_URI))
	if err != nil {
		log.Fatal("MongoDB connection error: ", err)
	}
}