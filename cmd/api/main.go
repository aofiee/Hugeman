package main

// @title APIs
// @version 1.0
// @description This is a sample server server.
// @termsOfService https://www.aofiee.dev/

// @contact.name API Support
// @contact.url https://www.aofiee.dev/
// @contact.email aofiee@aofiee.dev

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9089
// @BasePath /
// @schemes http
import (
	_ "hugeman/docs"
	protocol "hugeman/protocal"

	_ "github.com/arsmn/fiber-swagger/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	err := protocol.ServeHTTP()
	if err != nil {
		logrus.Println(err)
	}
}
