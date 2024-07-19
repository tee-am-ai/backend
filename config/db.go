package config

import (
	"github.com/tee-am-ai/backend/helper"
)

var MongoString string = GetEnv("MONGOSTRING")

var mongoinfo = helper.DBInfo{
	DBString: MongoString,
	DBName:   "db_teeamai",
}

var Mongoconn, ErrorMongoconn = helper.MongoConnect(mongoinfo)