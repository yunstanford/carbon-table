package receiver

import (
    "bufio"
    "io"
    "net"
    "github.com/yunstanford/carbon-table/cfg"
    "github.com/yunstanford/carbon-table/table"
    m20 "github.com/metrics20/go-metrics20/carbon20"
)

type Receiver struct {
    tcpAddr   string
    table  *table.Table
    // Add more configs
}

// Receiver Handle
func (r *Receiver) Handle(c io.Reader) {
    scanner := bufio.NewScanner(c)
    for scanner.Scan() {
        // Note that everything in this loop should proceed as fast as it can
        // so we're not blocked and can keep processing
        // so the validation, the pipeline initiated via table.Dispatch(), etc
        // must never block.

        buf := scanner.Bytes()

        key, val, ts, err := m20.ValidatePacket(buf, p.config.Validation_level_legacy.Level, p.config.Validation_level_m20.Level)
        if err != nil {
            log.Debug("plain.go: Bad Line: %q", buf)
            p.bad.Add(key, buf, err)
            numInvalid.Inc(1)
            continue
        }

        // log.Debug("plain.go: Received Line: %q", buf)

        // Insert Into Table

    }
    if err := scanner.Err(); err != nil {
        // log.Error(err.Error())
    }
}

// listen
func listen(addr string, handler Handler) error {
    laddr, err := net.ResolveTCPAddr("tcp", addr)
    if err != nil {
        return err
    }
    l, err := net.ListenTCP("tcp", laddr)
    if err != nil {
        return err
    }
    go acceptTcp(l, handler)

    // TODO: ADD UDP Support

    return nil
}

func acceptTcp(l *net.TCPListener, handler Handler) {
    for {
        // wait for a tcp connection
        c, err := l.AcceptTCP()
        if err != nil {
            break
        }
        // handle the connection
        go acceptTcpConn(c, handler)
    }
}

func acceptTcpConn(c net.Conn, handler Handler) {
    defer c.Close()
    handler.Handle(c)
}

func NewReceiver(config *cfg.receiverConfig) {
    // New Receiver

    // Listen

}