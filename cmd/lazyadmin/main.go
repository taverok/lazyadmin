package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/taverok/lazyadmin/pkg/admin"
	"github.com/taverok/lazyadmin/pkg/admin/config"
	"github.com/taverok/lazyadmin/pkg/logger"
)

func main() {
	logger.InitDefault()

	r := mux.NewRouter()
	app := admin.App{Router: r, Config: config.DefaultConfig()}
	err := app.Init()
	if err != nil {
		log.Fatalln(err.Error())
	}

	slog.Info(fmt.Sprintf("server started on port %d", app.Config.Port))
	err = http.ListenAndServe(fmt.Sprintf(":%d", app.Config.Port), r)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
