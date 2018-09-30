package main

import (
    "flag"
    "fmt"
    "os"
    "go.uber.org/zap"
    "github.com/yunstanford/carbon-table/table"
    "github.com/yunstanford/carbon-table/api"
    "github.com/yunstanford/carbon-table/cfg"
    "github.com/yunstanford/carbon-table/receiver"
)

const Version = "0.1.0"

func init() {

}

func usage() {
    fmt.Fprintln(
        os.Stderr,
        "Usage: carbon-table -config=<path-to-config-file>",
    )
}

func main() {
    // Command line flags
    configFile := flag.String("config", "", "config filename")
    flag.Parse()

    // Parse Config File...
    config, err := cfg.ParseConfigFile(*configFile)
    if err != nil {
        usage()
        return
    }

    // Logger
    logger, _ := zap.NewProduction()
    defer logger.Sync() // flushes buffer, if any

    logger.Info("Starting carbon-table")

    // New Table
    tbl := table.NewTable(config.Table)

    // New Receiver
    logger.Info("Starting Receiver")
    rec := receiver.NewReceiver(config.Receiver, tbl, logger)
    receiver.Listen(rec.TcpAddr , rec)

    // New API
    logger.Info("Starting Web API")
    webApi := api.NewApi(config.Api, tbl, logger)
    webApi.Start()
}
