package main

import (
    "flag"
    "fmt"
    "os"
    // "go.uber.org/zap"
    "github.com/yunstanford/carbon-table/table"
    "github.com/yunstanford/carbon-table/api"
    "github.com/yunstanford/carbon-table/cfg"
    "github.com/yunstanford/carbon-table/receiver"
)

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

    // New Table
    tbl := table.NewTable(config.Table)

    // New Receiver
    rec := receiver.NewReceiver(config.Receiver, tbl)
    receiver.Listen(rec.TcpAddr , rec)

    // New API
    webApi := api.NewApi(config.Api, tbl)
    webApi.Start()
}
