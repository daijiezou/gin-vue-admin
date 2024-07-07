package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/worker"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/workerspace"
)

type ApiGroup struct {
	SystemApiGroup      system.ApiGroup
	ExampleApiGroup     example.ApiGroup
	WorkerApiGroup      worker.ApiGroup
	WorkerspaceApiGroup workerspace.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
