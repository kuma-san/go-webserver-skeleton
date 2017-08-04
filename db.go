package main

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"gopkg.in/redis.v4"
)

func dbInit() *dbr.Session {

	DbSession := getDbSession()

	return DbSession
}

func getDbSession() *dbr.Session {

	db, err := dbr.Open("mysql",
		MySQLUser+":"+MySQLPassword+"@tcp("+MySQLServer+":"+MySQLPort+")/"+MySQLDatabase,
		nil)

	if err != nil {
		logrus.Error(err)
	} else {
		session := db.NewSession(nil)
		return session
	}
	return nil
}

func redisInit() *redis.Client {

	RedisSession := getRedisSession()

	return RedisSession
}

func getRedisSession() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     RedisServer + ":" + RedisPort,
		Password: "",
		DB:       0,
	})

	return client
}
