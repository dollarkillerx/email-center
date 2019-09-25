/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:45 2019-09-25
 */
package main

import (
	"email-center/config"
	"email-center/controller"
	"github.com/kataras/iris"
)

func main() {
	app := controller.RegisterIris()

	app.Run(iris.Addr(config.MyConfig.App.Host), iris.WithCharset("UTF-8"))
}
