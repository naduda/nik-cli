package scheduler

import (
	"encoding/json"
	"io/ioutil"
	"nik-cli/lms/scheduler/model"
	"os"
)

func getConfig() (model.LmsConfig, error) {
	var conf model.LmsConfig
	jsonFile, err := os.Open("config.json")
	if err != nil {
		return conf, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &conf)
	return conf, err
}
