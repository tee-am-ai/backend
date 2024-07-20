package config

import (
	"os"

	"github.com/tee-am-ai/backend/helper"
)

var mongoinfo = helper.DBInfo{
	DBString: os.Getenv("GO_MONGO_STRING"),
	DBName:   "db_teeamai",
}

var Mongoconn, ErrorMongoconn = helper.MongoConnect(mongoinfo)
