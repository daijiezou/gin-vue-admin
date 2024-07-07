package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/worker"
	"github.com/flipped-aurora/gin-vue-admin/server/router/workerspace"
)

type RouterGroup struct {
	System      system.RouterGroup
	Example     example.RouterGroup
	Worker      worker.RouterGroup
	Workerspace workerspace.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
