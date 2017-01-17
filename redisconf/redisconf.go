package redisconf

import (
	"fmt"
	"reflect"
	"strings"
)

//Directive represents a single redis config, e.g. `save 900 1`
type Directive struct {
	Keyword string
	Args    []string
}

//NewDirective is the correct way to initialise a RedisConf
func NewDirective(keyword string, args ...string) Directive {
	return Directive{Keyword: keyword, Args: args}
}

//Args is a convenience alias for initialising Directives
type Args []string

func (d Directive) String() string {
	return fmt.Sprintf("%s %s", d.Keyword, strings.Join(d.Args, " "))
}

//RedisConf represents a `redis.conf` file
type RedisConf []Directive

//New is the preferred way to initialise an empty RedisConf
func New(directives ...Directive) RedisConf {
	return RedisConf(directives)
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
	for _, directive := range directives {
		if !redisConf.contains(directive) {
			redisConf = append(redisConf, directive)
		}
	}
	return redisConf
}

//Remove a Config from RedisConf
func (redisConf RedisConf) Remove(keyword string, args ...string) RedisConf {
	return redisConf
}

func (redisConf RedisConf) contains(assertion Directive) bool {
	for _, directive := range redisConf {
		if reflect.DeepEqual(directive, assertion) {
			return true
		}
	}
	return false
}
