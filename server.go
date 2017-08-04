package main

import "github.com/Sirupsen/logrus"

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {

	e := initServer()

	e.Logger.Fatal(e.Start(":1323"))
}
