package log

import (
	"github.com/TeaWeb/code/teaweb/configs"
	"github.com/TeaWeb/code/teaweb/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			EndAll().
			Helper(&helpers.UserMustAuth{
				Grant: configs.AdminGrantLog,
			}).
			Helper(new(Helper)).
			Prefix("/log").
			Get("", new(IndexAction)).
			Get("/get", new(GetAction)).
			Get("/responseHeader/:logId", new(ResponseHeaderAction)).
			Get("/requestHeader/:logId", new(RequestHeaderAction)).
			Get("/cookies/:logId", new(CookiesAction)).
			GetPost("/runtime", new(RuntimeAction)).
			EndAll()
	})
}
