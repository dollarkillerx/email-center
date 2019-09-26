/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:45 2019-09-25
 */
package controller

import (
	"email-center/config"
	"email-center/defs"
	"email-center/logic"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/kataras/iris"
)

func RegisterIris() *iris.Application {
	app := iris.New()

	router(app)

	return app
}

func router(app *iris.Application) {
	app.Post("/email/seed", Logic)
}

// 逻辑
func Logic(ctx iris.Context) {
	input := defs.SeedEmail{}

	err := ctx.ReadForm(&input)
	if err != nil {
		ctx.StatusCode(400)
		ctx.JSON(map[string]interface{}{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	// auth
	if input.Key != config.MyConfig.App.Key {
		ctx.StatusCode(401)
		ctx.JSON(map[string]interface{}{
			"code": 401,
			"msg":  "权限错误",
		})
		return
	}

	err = logic.Logic(&input)
	if err != nil {
		clog.PrintWa(err)
		ctx.StatusCode(500)
		ctx.JSON(map[string]interface{}{
			"code": 500,
			"msg":  "发送失败",
			"data": err.Error(),
		})
		return
	}
	ctx.StatusCode(200)
	ctx.JSON(map[string]interface{}{
		"code": 200,
		"msg":  "发送成功",
	})
}
