package headers

import (
	"github.com/TeaWeb/code/teaconfigs"
	"github.com/TeaWeb/code/teaweb/actions/default/proxy/global"
	"github.com/iwind/TeaGo/actions"
)

type AddIgnoreAction actions.Action

func (this *AddIgnoreAction) Run(params struct {
	Filename      string
	LocationIndex int
	Name          string
	Must          *actions.Must
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入Name")

	proxy, err := teaconfigs.NewServerConfigFromFile(params.Filename)
	if err != nil {
		this.Fail(err.Error())
	}

	location := proxy.LocationAtIndex(params.LocationIndex)
	if location == nil {
		this.Fail("找不到要修改的路径规则")
	}

	location.AddIgnoreHeader(params.Name)
	proxy.Save()

	global.NotifyChange()

	this.Refresh().Success()
}
