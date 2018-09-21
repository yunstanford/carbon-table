package main

import (
    "flag"
    "fmt"
    "go.uber.org/zap"
    "github.com/yunstanford/carbon-table/table"
    "github.com/yunstanford/carbon-table/api"
    "github.com/yunstanford/carbon-table/cfg"
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
    cfg, err := cfg.ParseConfigFile(*configFile)
    if err != nil {
        usage()
        return
    }

    
}
