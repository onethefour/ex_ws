package utils

import(
	"github.com/go-redis/redis"
)

//
var Redis *redis.Client

func init(){
	addr:="47.106.136.96:6379"
	Redis =redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_,err := Redis.Ping().Result()
	if err != nil {
		panic("redis connect faild ")
	}


}
