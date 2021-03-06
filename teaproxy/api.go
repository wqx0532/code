package teaproxy

import (
	"github.com/TeaWeb/code/teaconfigs/api"
)

// 开启API功能
func EnableAPIServer(serverId string) {
	server, found := SERVERS[serverId]
	if found {
		server.API.On = true
	}
}

// 添加API
func AddAPI(serverId string, api *api.API) {
	server, found := SERVERS[serverId]
	if found {
		server.API.AddAPI(api)
	}
}

// 删除API
func DeleteAPI(serverId string, api *api.API) {
	server, found := SERVERS[serverId]
	if found {
		server.API.DeleteAPI(api)
	}
}

// 替换APIs
func ReplaceAPIs(serverId string, filenames []string) {
	server, found := SERVERS[serverId]
	if found {
		server.API.Files = filenames
		server.Validate()
	}
}
