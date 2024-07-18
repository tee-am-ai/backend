package config

import (
	"os"

	"github.com/tee-am-ai/backend/helper"
)

var MongoString string = os.Getenv("MONGOSTRING")

var mongoinfo = helper.DBInfo{
	DBString: MongoString,
	DBName:   "db_chatteeamai",
}

var Mongoconn, ErrorMongoconn = helper.MongoConnect(mongoinfo)