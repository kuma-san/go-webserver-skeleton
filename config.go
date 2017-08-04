package main

import "os"

var (
	RedisServer = os.Getenv("REDIS_URL")
	RedisPort   = os.Getenv("REDIS_PORT")

	MySQLServer   = os.Getenv("MYSQL_SERVER")
	MySQLPort     = os.Getenv("MYSQL_PORT")
	MySQLUser     = os.Getenv("MYSQL_USER")
	MySQLPassword = os.Getenv("MYSQL_PASSWORD")
	MySQLDatabase = os.Getenv("MYSQL_DATABASE")
	MySQLIsProxy  = os.Getenv("MYSQL_PROXY")

	MailgunAPIKey      = os.Getenv("MAILGUN_APIKEY")
	MailgunDomain      = os.Getenv("MAILGUN_DOMAIN")
	MailgunFromAddress = os.Getenv("MAILGUN_FROM_ADDRESS")

	ServerTimezone = os.Getenv("SERVER_TIMEZONE")
	NudgeHour      = os.Getenv("NUDGE_HOUR")
)

const (
	LoginExpire        = 86400
	TimeFormat         = "2006-01-02"
	TimeDateTimeFormat = "2006-01-02 15:04:05"
)
