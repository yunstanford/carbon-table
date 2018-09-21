package cfg

import (
    "github.com/BurntSushi/toml"
)

// Table Config
// ttl: Number of seconds to rebuild trie according to current traffic. This is designed for dropping pretty old metrics.
type tableConfig struct {
    Ttl  int
}

// api Config
// addr: Http Query Addr
type apiConfig struct {
    ApiAddr  string
}

// tcp Config
// addr: Tcp Listen Addr
type receiverConfig struct {
    TcpAddr  string
}

type Config struct {
    Table      tableConfig
    Api        apiConfig
    Receiver   receiverConfig
    // Put any other configs here...
}

// NewConfig
// Provdes default Values
func NewConfig() *Config {
    cfg := &Config{
        Table: tableConfig{
            Ttl: 3600 * 12, // 12 hours
        },
        Api: apiConfig{
            ApiAddr: "127.0.0.1:8080",
        },
        Receiver: receiverConfig{
            TcpAddr: ":3000"
        }
    }
    return cfg
}

func ParseConfigFile(file string) (*Config, error) {
    cfg := NewConfig()

    if file != "" {
        bytes, err := ioutil.ReadFile(file)
        if err != nil {
            return nil, err
        }
        body := string(bytes)

        if _, err := toml.Decode(body, cfg); err != nil {
            return nil, err
        }
    }
    return cfg, nil
}
