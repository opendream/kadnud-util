package main

import (
	"os"
	"fmt"
	"log"
	"strings"
	"github.com/opendream/deeperror"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func elog(num int64, msg string, err error, domain string, oid string, mid string, email string) {
	msg = strings.Replace(msg, "\n", "\r", -1)
	derr := deeperror.New(num, msg, err)
	log.Print(fmt.Sprintf("%v | [domain: %v][oid: %v][mid: %v][email: %v]",
		derr, domain, oid, mid, email,
	))
}

func plog(msg string, domain string, oid string, mid string, email string) {
	msg = strings.Replace(msg, "\n", "\r", -1)
	fmt.Println(fmt.Sprintf("Message: %v | [domain: %v][oid: %v][mid: %v][email: %v]",
		msg, domain, oid, mid, email,
	))
}
