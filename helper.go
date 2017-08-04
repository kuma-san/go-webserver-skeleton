package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"reflect"

	"github.com/hashicorp/go-uuid"
	"github.com/kuma-san/kengo/conf"
	"gopkg.in/mailgun/mailgun-go.v1"
)

func toHash(password string) string {
	converted := sha512.Sum512([]byte(password))
	return hex.EncodeToString(converted[:])
}

func getUUID() string {
	u, _ := uuid.GenerateUUID()
	return u
}

func sendEmail(to string, title string, html string) (string, error) {
	mg := mailgun.NewMailgun(conf.MailgunDomain, conf.MailgunAPIKey, "")
	m := mg.NewMessage(
		conf.MailgunFromAddress,
		title,
		"",
		to,
	)
	m.SetHtml(html)
	_, id, err := mg.Send(m)

	return id, err
}

func structToStringMap(data interface{}) map[string]string {
	result := make(map[string]string)
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := fmt.Sprint(elem.Field(i).Interface())
		result[field] = value
	}

	return result
}
