// main.go
package main

import (
    "flag"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/binding"
    "github.com/martini-contrib/render"
    "log"
)

var (
    cfgPath = flag.String("cfgPath", "./config.yml", "config file")
)

var mockdata map[string]interface{}

func main() {
    log.Printf("httpmock start...\n")
    flag.Parse()
    ConfigInit(*cfgPath)
    // 解析文件
    safeParse()
    // 启动martini
    m := martini.Classic()
    m.Use(martini.Static("resource"))
    m.Use(render.Renderer(render.Options{}))

    m.Post("/mock", binding.Bind(mock_t{}), mockHandler)
    m.Get("/mockList", mockListHandler)
    m.NotFound(notFoundHandle)

    m.RunOnAddr(CfgData.SrvAddr)
}
