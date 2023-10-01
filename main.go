package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cfg := clientv3.Config{
		Endpoints:   []string{os.Getenv("ETCD_CLIENT_URL")},
		DialTimeout: 3 * time.Second,
	}

	client, err := clientv3.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	resp, err := client.MemberList(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	for _, m := range resp.Members {
		fmt.Printf("%s\n", m.String())
	}
}
