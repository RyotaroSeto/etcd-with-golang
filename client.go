package main

import (
	"os"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewEtcdClient() (*clientv3.Client, error) {
	cfg := clientv3.Config{
		Endpoints:   []string{os.Getenv("ETCD_CLIENT_URL")},
		DialTimeout: 3 * time.Second,
	}

	client, err := clientv3.New(cfg)
	if err != nil {
		return nil, err
	}
	return client, nil
}
