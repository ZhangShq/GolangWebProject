/*
 * @Descripttion: 
 * @version: 
 * @Author: zhangshq@corp.21cn.com
 * @Date: 2020-02-10 23:58:20
 * @LastEditors: zhangshq@corp.21cn.com
 * @LastEditTime: 2020-06-22 15:29:29
 */ 
package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func TestHelloWorld(ctx iris.Context) int {
	fmt.Println("Tst")
	ctx.WriteString("hello world")
	fmt.Println("Tes")
	return 0;
}
func main(){
	app := iris.New()
    app.Get("/", func(ctx iris.Context){
		fmt.Println("new connect")
		go TestHelloWorld(ctx)
	})
	app.Run(iris.Addr(":8080"))
}
