package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
	log "github.com/sirupsen/logrus"
)

type Redis struct {
	clients map[string]*redis.Pool
}

func NewRedis() *Redis {
	r := new(Redis)
	r.clients = make(map[string]*redis.Pool)

	return r
}

func (r *Redis) Register(key string, opt ConnOptions) {
	opts := redis.DialOption{}
	if opt.Timeout != 0 {
		opts = redis.DialConnectTimeout(opt.Timeout)
	}

	p := &redis.Pool{
		MaxIdle:     3,
		MaxActive:   50,
		IdleTimeout: opt.Timeout,
		Dial:        func() (redis.Conn, error) { return redis.Dial(connectionTypeTCP, opt.Address+":"+opt.Port, opts) },
	}
	r.clients[key] = p
}

func (r *Redis) RegisterMock(key string, conn *redigomock.Conn) {
	p := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return conn, nil
		},
	}
	r.clients[key] = p
}

func (r *Redis) GetCon(key string) redis.Conn {
	return r.clients[key].Get()
}

func (c *Redis) Cmd(name, cmd string, args ...interface{}) (interface{}, error) {
	if p, ok := c.clients[name]; ok && p != nil {
		c := p.Get()
		defer c.Close()

		i, err := c.Do(cmd, args...)
		if err != nil {
			log.Printf("[Cmd][%s] Do [%s][%s]\n", name, cmd, err.Error())
		}
		return i, err
	}
	return nil, ErrConnNotFound
}

// Get execute command "GET" and return data as string. If key return `redis.ErrNil`, it will ignore the error and
// return empty string instead.
func (c *Redis) Get(name, key string) (string, error) {
	if p, ok := c.clients[name]; ok && p != nil {
		c := p.Get()
		defer c.Close()

		i, err := redis.String(c.Do(CmdGet, key))
		if err != nil && err != redis.ErrNil {
			log.Printf("[Cmd][%s] Do [%s][%s]\n", name, CmdGet, err.Error())
			return i, err
		}
		return i, nil
	}
	return "", ErrConnNotFound
}

func (c *Redis) MultiExec(name string, commands ...Command) (interface{}, error) {
	if p, ok := c.clients[name]; ok && p != nil {
		c := p.Get()
		defer c.Close()

		c.Send(CmdMulti)
		for _, cmd := range commands {
			c.Send(cmd.Name, cmd.Args...)
		}
		i, err := c.Do(CmdExec)
		if err != nil {
			log.Printf("[Cmd][%s] Do [%s][%s]\n", name, fmt.Sprintf("%s.%s", CmdMulti, CmdExec), err.Error())
			return i, err
		}
		return i, nil
	}
	return nil, ErrConnNotFound
}
