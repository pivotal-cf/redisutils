package redisconf

import (
	"fmt"
	"strings"
)

//Directive represents a single redis config, e.g. `save 900 1`
type Directive struct {
	Keyword string
	Args    []string
}

//Args is a convenience alias for NewDirectives
type Args []string

//NewDirective is the correct way to initialise a RedisConf
func NewDirective(keyword string, args ...string) (Directive, error) {
	return Directive{Keyword: keyword, Args: args}, nil
}

func (d Directive) String() string {
	return fmt.Sprintf("%s %s", d.Keyword, strings.Join(d.Args, " "))
}

//RedisConf represents a `redis.conf` file
type RedisConf []Directive

//New is the correct way to initialise a RedisConf
func New(directives ...Directive) ([]Directive, error) {
	return append([]Directive{}, directives...), nil
}

//Decode a `redis.conf` into a RedisConf
func Decode(config string) RedisConf {
	return RedisConf{}
}

//Encode a RedisConf to a `redis.conf`
func (redisConf RedisConf) Encode() (encoded string) {
	for _, c := range redisConf {
		encoded = encoded + fmt.Sprintln(c)
	}
	return
}

//Append a Conf to RedisConf
func (redisConf RedisConf) Append(directives ...Directive) RedisConf {
	return append(redisConf, directives...)
}

//Remove a Config from RedisConf
func (redisConf RedisConf) Remove(keyword string, args ...string) RedisConf {
	return redisConf
}
