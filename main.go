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
	"go-resilient-ledger/internal/payment"
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
	paymentCont := payment.NewContainer(gormDB, accountCont.Service)

	accountWCtx := context.Background()
	accountCont.QWorker.StartWorker(accountWCtx)

	paymentWCtx := context.Background()
	paymentCont.QWorker.StartWorker(paymentWCtx)

	account.RegisterAPI(fiberApp, *accountCont.QProducer)
	payment.RegisterAPI(fiberApp, *paymentCont.QProducer)

	addr := fmt.Sprintf(":%d", appCfg.AppConfig.Port)

	if err := fiberApp.Listen(addr); err != nil {
		log.Fatalf("[FATAL] Server failed to start: %v", err)
	}
}
