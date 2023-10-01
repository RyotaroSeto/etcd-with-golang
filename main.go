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

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := client.MemberList(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, m := range resp.Members {
		fmt.Printf("%s\n", m.String())
	}

	_, err = client.Put(ctx, "/hoge", "fuga")
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Get(ctx, "/hoge")
	if err != nil {
		log.Fatal(err)
	}
	if res.Count == 0 {
		log.Fatal("/hoge not found")
	}
	fmt.Println(string(res.Kvs[0].Value))

	_, err = client.Delete(context.TODO(), "/hoge")
	if err != nil {
		log.Fatal(err)
	}

	client.Put(ctx, "/test/hoge3", "fuga3")
	client.Put(ctx, "/test/hoge2", "fuga2")
	client.Put(ctx, "/test/hoge1", "fuga1")
	response, err := client.Get(ctx, "/test/",
		clientv3.WithPrefix(), // キーに指定したプレフィックスから始まるキーをすべて取得(/test)
		clientv3.WithSort(clientv3.SortByValue, clientv3.SortAscend), // 結果をソート(Valueの昇順)
		clientv3.WithKeysOnly(), // キーのみを取得
		// https://pkg.go.dev/github.com/coreos/etcd/clientv3#OpOption 他オプション
	)
	if err != nil {
		log.Fatal(err)
	}
	for _, kv := range response.Kvs {
		fmt.Printf("%s: %s\n", kv.Key, kv.Value)
	}

	fmt.Println("----------------------")
	client.Put(ctx, "/test/hoge", "123")
	respo, _ := client.Get(ctx, "/test/hoge", clientv3.WithPrefix())
	printResponse(respo)

	fmt.Println("----------------------")
	client.Put(ctx, "/test/hoge", "456")
	respo, _ = client.Get(ctx, "/test/hoge", clientv3.WithPrefix())
	printResponse(respo)

	fmt.Println("----------------------")
	client.Put(ctx, "/test/hage", "999")
	respo, _ = client.Get(ctx, "/test/hage", clientv3.WithPrefix())
	printResponse(respo)
}

// revision:etcdのリビジョン番号。クラスタ全体で一つの値が利用。tcdに何らかの変更(キーの追加、変更、削除)が加えられると値が1増える
// create_revision:キーが作成されたときのリビジョン番号
// mod_revision:キーの内容が最後に変更されたときのリビジョン番号
// version:キーのバージョン。このキーに変更が加えられると値が1増える
func printResponse(resp *clientv3.GetResponse) {
	fmt.Println("header: " + resp.Header.String())
	for i, kv := range resp.Kvs {
		fmt.Printf("kv[%d]: %s\n", i, kv.String()) // RevisionとVersion確認
	}
}
