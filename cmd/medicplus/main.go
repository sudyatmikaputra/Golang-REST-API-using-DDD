package main

import (
	"fmt"
	netHttp "net/http"

	"github.com/go-chi/chi/v5"
	"github.com/medicplus-inc/medicplus-feedback/cmd/container"
	"github.com/medicplus-inc/medicplus-feedback/cmd/http"
	"github.com/medicplus-inc/medicplus-feedback/config"
	"github.com/medicplus-inc/medicplus-feedback/database"
	"github.com/oklog/oklog/pkg/group"
	"github.com/sirupsen/logrus"
)

func main() {
	var g = group.Group{}
	var logger = config.GetLogger()

	container.NewIOC()

	runMigration()
	injectSeed()
	runHTTP(&g, *logger)
	runHTTPProfiler(&g, *logger)

	logger.Fatal("exit", g.Run())
}

func runMigration() {
	db := config.DB()

	err := database.Migrate(db)
	if nil != err {
		panic(fmt.Sprintf("Error on migrating database: %v", err))
	}
}

func injectSeed() {
	db := config.DB()

	_ = database.Seed(db)
}

func runHTTP(
	g *group.Group,
	logger logrus.Logger,
) {
	port := config.GetValue(config.HTTP_PORT)

	if len(port) < 1 {
		panic(fmt.Sprintf("Environment Missing!\n*%s* is required", port))
	}

	var router *chi.Mux
	router = chi.NewRouter()

	router.Mount("/api", http.CompileRoute(router))

	server := &netHttp.Server{
		Addr:    port,
		Handler: router,
	}

	fmtLog := logger.WithFields(logrus.Fields{
		"transport": "debug/HTTP",
		"addr":      port,
	})

	g.Add(
		func() error {
			fmtLog.Info("HTTP transport run at ", port)
			return server.ListenAndServe()
		},
		func(err error) {
			if nil != err {
				fmtLog.Warn("Error Occurred ", err.Error())
				panic(err)
			}
		},
	)
}

func runHTTPProfiler(
	g *group.Group,
	logger logrus.Logger,
) {
	profilerPort := config.GetValue(config.HTTP_PROFILER_PORT)

	if len(profilerPort) < 1 {
		panic(fmt.Sprintf("Environment Missing!\n*%s* is required", profilerPort))
	}

	var router *chi.Mux
	router = chi.NewRouter()

	router.Mount("/profiler", http.CompileProfilingRoute(router))

	server := &netHttp.Server{
		Addr:    profilerPort,
		Handler: router,
	}

	fmtLog := logger.WithFields(logrus.Fields{
		"transport": "debug/HTTP",
		"addr":      profilerPort,
	})

	g.Add(
		func() error {
			fmtLog.Info("HTTP Profiler transport run at ", profilerPort)
			return server.ListenAndServe()
		},
		func(err error) {
			if nil != err {
				fmtLog.Warn("Error Occurred ", err.Error())
				panic(err)
			}
		},
	)
}
