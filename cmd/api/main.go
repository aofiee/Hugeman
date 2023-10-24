package main

import (
	protocol "hugeman/protocal"

	"github.com/sirupsen/logrus"
)

func main() {
	err := protocol.ServeHTTP()
	if err != nil {
		logrus.Println(err)
	}
}
