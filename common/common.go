package common

import (
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
)

var FullLog string

func Logger(LogType int, log string) {
	FullLog = FullLog + "\n" + log

	switch LogType {
	case 0: // INFO
		color.White("[INFO] " + " " + log)

	case 1:
		color.Yellow("[WARN] " + " " + log)

	case 2:
		color.Red("[ERROR] " + log)

	default:
		color.White("[DEFAULT] " + " " + log)
	}

	ioutil.WriteFile("./fix.log", []byte(FullLog), 0777)
}

func HttpGet(Address string) ([]byte, error) {
	resp, err := http.Get(Address)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
