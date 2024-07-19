package config

import (
	"github.com/tee-am-ai/backend/helper" // Package yang mungkin berisi fungsi-fungsi bantuan terkait konfigurasi atau pengaturan
)

var MongoString string = GetEnv("MONGOSTRING")

var mongoinfo = helper.DBInfo{
	DBString: MongoString,
	DBName:   "db_teamai",
}
var Mongoconn, ErrorMongoconn = helper.MongoConnect(mongoinfo)
