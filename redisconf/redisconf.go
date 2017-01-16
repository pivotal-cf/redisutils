package redisconf

import (
	"fmt"
	"strings"
)

type Conf struct {
	Keyword string
	Args    []string
}

func (conf Conf) String() string {
	return fmt.Sprintf("%s %s", conf.Keyword, strings.Join(conf.Args, " "))
}

func NewConf(keyword string, args ...string) Conf {
	return Conf{Keyword: keyword, Args: args}
}

type RedisConf []Conf

func Decode(config string) RedisConf {
	return RedisConf{}
}

func (rc *RedisConf) Encode() (encoded string) {
	for _, c := range *rc {
		encoded = encoded + fmt.Sprintln(c)
	}
	return
}

func (rc *RedisConf) Append(confs ...Conf) error {
	*rc = append(*rc, confs...)
	return nil
}

func (rc *RedisConf) Remove(keyword string, args ...string) {}
