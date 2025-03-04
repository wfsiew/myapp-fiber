package database

import (
	"strconv"
	"database/sql"
	"app/utils"
    "app/config"
	"github.com/sijms/go-ora/v2"
)

var DB *sql.DB

func SetDb(db *sql.DB) {
	DB = db
}

func GetDb() *sql.DB {
	if DB == nil {
		ConnectDB()
	}

	return DB
}

func ConnectDB() {
	p := config.Config("DB_PORT")
    port, _ := strconv.Atoi(p)
	connStr := go_ora.BuildUrl(config.Config("DB_SERVER"), port, config.Config("DB_SERVICE"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), nil)
	db, err := sql.Open("oracle", connStr)
	// dsn := fmt.Sprintf(`user="%s" password="%s" connectString="%s" heterogeneousPool=false standaloneConnection=false`, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_URL"))
	// fmt.Println(dsn)
	// DB, err := sql.Open("godror", dsn)

	if err != nil {
		utils.Logger.Err(err).Msg(err.Error())
	} else {
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(5)
		SetDb(db)
		utils.ILogger.Info().Msg("Connection Opened to Database")
	}
}