package cache

import (
	"github.com/go-redis/redis/v8"
)

func RedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// func ExampleClient() {
// 	var ctx = context.Background()
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})
// 	// c := repository.CategoryRepo().GetAll()
// 	// cc, _ := json.Marshal(c)
// 	fmt.Println()
// 	// err := rdb.Set(ctx, "key1", cc, 10*time.Second).Err()
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	val, err := rdb.Get(ctx, "key").Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("key", val)

// 	val2, err := rdb.Get(ctx, "key2").Result()
// 	if err == redis.Nil {
// 		fmt.Println("key2 does not exist")
// 	} else if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("key2", val2)
// 	}
// }
