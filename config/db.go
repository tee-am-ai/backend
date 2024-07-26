package config

import (
	"github.com/tee-am-ai/backend/helper"
)

var mongoinfo = helper.DBInfo{
	// DBString: os.Getenv("GO_MONGO_STRING"),
	DBString: "mongodb+srv://fdirga63:-3aN!r*5-r8sNju@cluster0.xhievrg.mongodb.net",
	DBName:   "db_teeamai",
}

var Mongoconn, ErrorMongoconn = helper.MongoConnect(mongoinfo)
