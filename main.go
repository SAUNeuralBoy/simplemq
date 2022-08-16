package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
)

type Args struct {
	req string
	src string
	dst string
	msg string
	cnt int
}

func main() {
	var args Args
	flag.StringVar(&args.req, "r", "len", "")
	flag.StringVar(&args.src, "s", "", "")
	flag.StringVar(&args.dst, "t", "", "")
	flag.StringVar(&args.msg, "m", "hello", "")
	flag.IntVar(&args.cnt, "n", -1, "")
	flag.Parse()
	src := md5.Sum([]byte(args.src))
	dst := md5.Sum([]byte(args.dst))
	var backend Backend
	backend = newRedisBackend(redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // 指定
		Password: "",
		DB:       0, // redis一共16个库，指定其中一个库即可
	}))
	switch args.req {
	case "len":
		println(backend.Len(src))
	case "get":
		get, err := backend.Get(src, CNT(args.cnt))
		for i := 0; i < len(get); i++ {
			fmt.Printf("%s:%s\n", string(get[i].Src[:]), string(get[i].Content))
		}
		if err != nil {
			print(err)
		}
	case "write":
		println(backend.Write(dst, &MSG{Src: src, Content: []byte(args.msg)}))
	case "skip":
		println(backend.Skip(src, CNT(args.cnt)))
	}
}
