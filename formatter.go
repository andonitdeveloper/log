package log

import (
	"github.com/sirupsen/logrus"
	"time"
	"encoding/json"
	"fmt"
	"strings"
)

type CommonFormatter struct {
}

func (this *CommonFormatter) Format(entry *logrus.Entry) ([]byte, error){

	data := make(logrus.Fields, len(entry.Data)+9)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	data["type"] = Identity
	data["version"] = version
	data["subversion"] = subversion
	data["ts"] = time.Now().UTC().Unix()
	data["pid"] = pid
	data["hostname"] = hostname
	data["appname"] = appname
	data["msg"] = entry.Message
	data["level"] = strings.ToUpper(entry.Level.String())
	serialized, err := json.Marshal(data)

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}

	return append(serialized, '\n'), nil

}
