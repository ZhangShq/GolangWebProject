/*
 * @Descripttion: 
 * @version: 
 * @Author: zhangshq@corp.21cn.com
 * @Date: 2020-02-10 23:58:20
 * @LastEditors: zhangshq@corp.21cn.com
 * @LastEditTime: 2020-06-19 11:03:00
 */ 
package main

import (
	"github.com/kataras/iris"
)

func main()  {
	app := iris.New()
    app.Get("/", func(ctx iris.Context){})
    app.Run(iris.Addr(":8080"))
}

