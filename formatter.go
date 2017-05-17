package log

import (
	"github.com/sirupsen/logrus"
	"time"
)

type CommonFormatter struct {
	jsonFormatter logrus.JSONFormatter
}

func (this *CommonFormatter) Format(entry *logrus.Entry) ([]byte, error){
	entry.Data["type"] = Identity
	entry.Data["version"] = version
	entry.Data["subversion"] = subversion
	entry.Data["ts"] = time.Now().UTC().Unix()
	entry.Data["pid"] = pid
	entry.Data["hostname"] = hostname
	entry.Data["appname"] = appname

	return this.jsonFormatter.Format(entry)
}
