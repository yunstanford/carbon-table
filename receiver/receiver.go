package receiver

import (
    "bufio"
    "io"
    "net"
    "go.uber.org/zap"
    "github.com/yunstanford/carbon-table/cfg"
    "github.com/yunstanford/carbon-table/table"
    m20 "github.com/metrics20/go-metrics20/carbon20"
)

type Receiver struct {
    TcpAddr   string
    table     *table.Table
    // Add more configs
    Logger    *zap.Logger
}

type Handler interface {
    Handle(io.Reader)
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

        // We're using medium validation, for now.
        // TODO: Make validation level
        key, _, _, err := m20.ValidatePacket(buf, m20.MediumLegacy, m20.MediumM20)
        if err != nil {
            r.Logger.Debug("receiver.go: Bad Line", zap.String("line", string(buf)))
            continue
        }

        r.Logger.Debug("receiver.go: Received Line", zap.String("line", string(buf)))

        // Insert Into Table
        r.table.Insert(string(key))

    }
    if err := scanner.Err(); err != nil {
        // log.Error(err.Error())
    }
}

// listen
func Listen(addr string, handler Handler) error {
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

// NewReceiver
func NewReceiver(c *cfg.ReceiverConfig, t *table.Table, l *zap.Logger) *Receiver{
    // New Receiver
    rec := &Receiver {
        TcpAddr: c.TcpAddr,
        table:   t,
        Logger:  l,
    }

    return rec
}
