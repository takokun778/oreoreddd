package main

import (
	"github.com/takokun778/oreoreddd/internal/adapter/controller"
	"github.com/takokun778/oreoreddd/internal/adapter/gateway"
	"github.com/takokun778/oreoreddd/internal/driver/config"
	"github.com/takokun778/oreoreddd/internal/driver/server"
	"github.com/takokun778/oreoreddd/internal/usecase/interactor/sample"
)

func main() {
	config.Init()

	sr := gateway.NewSample()

	sci := sample.NewCreateInteractor(sr)

	sri := sample.NewReadInteractor(sr)

	sui := sample.NewUpdateInteractor(sr)

	sample := controller.NewSample(sci, sri, sui)

	server.NewHTTPServer(sample).Run()
}
