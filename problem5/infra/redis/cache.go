package redis

import (
	"errors"
	"time"
)

const (
	CmdSet       string = "SET"
	CmdNotExist  string = "NX"
	CmdExpired   string = "EX"
	CmdTTL       string = "TTL"
	CmdSIsMember string = "SISMEMBER"
	CmdSMembers  string = "SMEMBERS"
	CmdSAdd      string = "SADD"
	CmdSRem      string = "SREM"
	CmdDel       string = "DEL"
	CmdGet       string = "GET"
	CmdMulti     string = "MULTI"
	CmdExec      string = "EXEC"
	CmdPing      string = "PING"
	CmdKeys      string = "KEYS"
	CmdIncr      string = "INCR"

	TTLForever int = -1
	TTL30s     int = 30
	TTL15m     int = 60 * 15
	TTL1h      int = 60 * 60
	TTL1d      int = 60 * 60 * 24
	TTL1w      int = 60 * 60 * 24 * 7

	ClientRedis       = "redis"
	connectionTypeTCP = "tcp"
)

var ErrConnNotFound = errors.New("cache connection not found")

type ConnOptions struct {
	Address string
	Port    string
	Timeout time.Duration
}

type Command struct {
	Name string
	Args []interface{}
}

type Client interface {
	Register(key string, opt ConnOptions)
	Cmd(name, command string, args ...interface{}) (interface{}, error)
	Get(name, key string) (string, error)
	MultiExec(name string, commands ...Command) (interface{}, error)
}

var c Client

func NewClient(clientType string) Client {
	if ClientRedis == clientType {
		c = NewRedis()
	}

	return c
}

func GetClient() Client {
	return c
}

func SetClient(cl Client) {
	c = cl
}
