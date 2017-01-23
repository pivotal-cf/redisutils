package redisconf

import "fmt"

//Directive describes the methods for `redis.conf` directives
type Directive interface {
	String() string
	Validate() error
}

//RedisConf is a representation of `redis.conf`
type RedisConf []Directive

//New initialises a RedisConf from Directives and validates each
func New(directives ...Directive) (redisConf RedisConf, err error) {
	err = validateDirectives(directives...)
	if err == nil {
		redisConf = append(RedisConf{}, directives...)
	}
	return
}

func validateDirectives(directives ...Directive) error {
	for _, directive := range directives {
		err := directive.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (conf RedisConf) String() (confString string) {
	for _, directive := range conf {
		confString = confString + fmt.Sprintln(directive)
	}
	return
}
