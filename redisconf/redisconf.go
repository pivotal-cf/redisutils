package redisconf

import (
	"fmt"
	"strings"
)

//Conf represents a single redis config, e.g. `save 900 1`
type Conf struct {
	Keyword string
	Args    []string
}

//NewConf is the correct way to initialise a Conf
func NewConf(keyword string, args ...string) Conf {
	return Conf{Keyword: keyword, Args: args}
}

func (conf Conf) String() string {
	return fmt.Sprintf("%s %s", conf.Keyword, strings.Join(conf.Args, " "))
}

//RedisConf represents a `redis.conf` file
type RedisConf []Conf

//Decode a `redis.conf` into a RedisConf
func Decode(config string) RedisConf {
	return RedisConf{}
}

//Encode a RedisConf to a `redis.conf`
func (rc *RedisConf) Encode() (encoded string) {
	for _, c := range *rc {
		encoded = encoded + fmt.Sprintln(c)
	}
	return
}

//Append a Conf to RedisConf
func (rc *RedisConf) Append(confs ...Conf) error {
	*rc = append(*rc, confs...)
	return nil
}

//Remove a Config from RedisConf
func (rc *RedisConf) Remove(keyword string, args ...string) {}
