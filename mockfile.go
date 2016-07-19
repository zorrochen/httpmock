package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)


// 保证mockdata不为nil
func safeParse() {
	parsemockdata, err := parseData()
	if err != nil {
		mockdata = map[string]interface{}{}
	}
	mockdata = parsemockdata
}

func parseData() (map[string]interface{}, error) {
	bytes, err := ioutil.ReadFile(CfgData.MockFile)
	if err != nil {
		log.Printf("[parseData error] %v", err)
		return nil, err
	}

	dataInfo := map[string]interface{}{}
	if err = json.Unmarshal(bytes, &dataInfo); err != nil {
		log.Printf("[parseData error] %v", err)
		return nil, err
	}

	return dataInfo, nil
}

// 反写
func rewriteMockFile() error {
	mockdataJson, err := json.MarshalIndent(mockdata, "", "")
	if err != nil {
		log.Printf("[rewriteMockFile error] %v", err)
		return err
	}

	ioutil.WriteFile(CfgData.MockFile, mockdataJson, os.ModePerm)
	return nil
}
