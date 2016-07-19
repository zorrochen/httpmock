package main

import (
	"fmt"
	"github.com/martini-contrib/render"
	"net/http"
)

func notFoundHandle(req *http.Request, r render.Render) {
	// 解析文件
	safeParse()
	ret := mockdata[req.URL.Path]
	r.JSON(200, ret)
}

type mock_t struct {
	ReqPath  string      `json:"reqPath"`
	RespData interface{} `json:"respData"`
}

func mockHandler(newMockData mock_t, w http.ResponseWriter) {
	if newMockData.ReqPath == "" {
		w.Write([]byte("reqPath error"))
		return
	}

	firstChar := string([]rune(newMockData.ReqPath)[0:1])
	if firstChar != "/" {
		w.Write([]byte("reqPath error"))
		return
	}

	mockdata[newMockData.ReqPath] = newMockData.RespData
	rewriteMockFile()

	w.Write([]byte("ok"))
}

func mockListHandler(w http.ResponseWriter) {
	ret := ""
	for reqPath, _ := range mockdata {
		ret += fmt.Sprintf("%s\n", reqPath)
	}

	w.Write([]byte(ret))
}
