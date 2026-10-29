package main

import (
	"context"
	"fmt"
	"log"

	"go-resilient-ledger/internal/infra/config/envconfig"
	"go-resilient-ledger/internal/infra/fiberhttp"
	"go-resilient-ledger/internal/infra/gormdb"
	"go-resilient-ledger/internal/infra/validator"

	"go-resilient-ledger/internal/account"
)

func main() {
	val, err := validator.New()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	appCfg, err := envconfig.NewConfig(".env", val)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	gormDB, err := gormdb.NewDB(appCfg.PostgreConfig)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	fiberErrHandler := fiberhttp.NewErrHandler(val.Translator)
	fiberApp := fiberhttp.NewApp(fiberErrHandler, appCfg.AppConfig.TimeoutDur)

	accountCont := account.NewContainer(gormDB)

	accountWCtx := context.Background()
	accountCont.QWorker.StartWorker(accountWCtx)

	account.RegisterAPI(fiberApp, *accountCont.QProducer)

	addr := fmt.Sprintf(":%d", appCfg.AppConfig.Port)

	if err := fiberApp.Listen(addr); err != nil {
		log.Fatalf("[FATAL] Server failed to start: %v", err)
	}
}
