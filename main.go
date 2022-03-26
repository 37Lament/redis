package main

import (
	"fmt"
	"github.com/go-redis/redis"
)
var redisDB *redis.Client

func initClient() (err error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err = redisDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

//单个订阅subscribe
func subscribe(channel string) {
	pubsub := redisDB.Subscribe(channel)
	defer pubsub.Close()
	for msg := range pubsub.Channel() {
		fmt.Printf("channel=%s message=%s\n", msg.Channel, msg.Payload)
	}
}
//模式订阅psubscribe
func psubscribe(channel string) {
	pubsub := redisDB.PSubscribe(channel+"*")
	defer pubsub.Close()
	for msg := range pubsub.Channel() {
		fmt.Printf("channel=%s message=%s\n", msg.Channel, msg.Payload)
	}

}
//频道发布publish
func publish(channel string,message string)  {
	n, err := redisDB.Publish(channel, message).Result()
	if err != nil{
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf("%d clients received the message\n", n)
}

func main()  {
	err:=initClient()
	x:="test1"
	message:="hello"
	subscribe(x)
	psubscribe(x)
	publish(x,message)
	if err != nil {
		fmt.Println("err is",err)
	}else{
		fmt.Println("OK")
	}
}
