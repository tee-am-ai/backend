package config

import (
	"github.com/tee-am-ai/backend/helper"
)

var MongoString string = GetEnv("MONGO_STRING")

var mongoinfo = helper.DBInfo{
	DBString: MongoString,
	DBName:   "db_teeamaiii",
}

var Mongoconn, ErrorMongoconn = helper.MongoConnect(mongoinfo)