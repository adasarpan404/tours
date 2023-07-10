package utils

import "log"

func CheckNilErr(err error) {
	log.Fatal(err)
}
