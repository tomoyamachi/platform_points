package db

import (
	"platform_accounts/conf"

	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

func Init() *dbr.Session {
	session := getSession()
	return session
}

func getSessionByDbManager(db *conf.DbManager) string {
	return db.User + ":" + db.Password + "@tcp(" + db.Host + ":" + db.Port + ")/" + db.Dbname
}

func getSession() *dbr.Session {
	var master string
	master = getSessionByDbManager(conf.MasterDb)
	db, err := dbr.Open("mysql", master, nil)

	if err != nil {
		logrus.Error(err)
	} else {
		session := db.NewSession(nil)
		return session
	}
	return nil
}
