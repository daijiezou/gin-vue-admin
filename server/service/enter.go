package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/worker"
	"github.com/flipped-aurora/gin-vue-admin/server/service/workerspace"
)

type ServiceGroup struct {
	SystemServiceGroup      system.ServiceGroup
	ExampleServiceGroup     example.ServiceGroup
	WorkerServiceGroup      worker.ServiceGroup
	WorkerspaceServiceGroup workerspace.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
