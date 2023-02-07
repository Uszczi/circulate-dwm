package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	var log = logrus.New()
	log.Out = os.Stdout
}
