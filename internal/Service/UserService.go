package Service

import (
	"context"
)

type UserService struct {
}

func (u UserService) StoreUser() { // map[string]string
	ctx := context.Background()
	redis := getRedis()
	err := redis.Set(ctx, "user:"+"2"+":username", "erfan", 0).Err()
	if err != nil {
		panic(err)
	}

}

func (u UserService) GetUser(userId string) string {
	ctx := context.Background()
	redis := getRedis()
	username, err := redis.Get(ctx, "user:"+userId+":username").Result()
	if err != nil {
		panic(err)
	}
	return username
}

func (u UserService) StoreUsers() map[string]string {
	ctx := context.Background()
	client := getRedis()

	pip := client.Pipeline()

	session := map[string]string{
		"name":    "John",
		"surname": "Smith",
		"company": "Redis",
		"age":     "29",
	}
	for k, v := range session {
		err := pip.HSet(ctx, "user-session:123", k, v).Err()
		if err != nil {
			panic(err)
		}
	}

	_, err := pip.Exec(ctx)
	if err != nil {
		panic(err)
	}

	return client.HGetAll(ctx, "user-session:123").Val()
	//fmt.Println(userSession)
}
