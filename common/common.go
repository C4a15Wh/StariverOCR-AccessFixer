package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var FullLog string

func Logger(LogType int, log string) {
	var Header string

	FullLog = FullLog + "\n" + log

	switch LogType {
	case 0: // INFO
		Header = "[INFO] "

	case 1:
		Header = "[WARN]"

	case 2:
		Header = "[ERROR]"

	default:
		Header = "[DEFAULT]"
	}

	fmt.Println(Header + " " + log)
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
