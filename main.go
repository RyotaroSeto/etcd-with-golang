package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	client, err := NewEtcdClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// resp, err := client.MemberList(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, m := range resp.Members {
	// 	fmt.Printf("%s\n", m.String())
	// }

	// _, err = client.Put(ctx, "/hoge", "fuga")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// res, err := client.Get(ctx, "/hoge")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if res.Count == 0 {
	// 	log.Fatal("/hoge not found")
	// }
	// fmt.Println(string(res.Kvs[0].Value))

	// _, err = client.Delete(ctx, "/hoge")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// client.Put(ctx, "/test/hoge3", "fuga3")
	// client.Put(ctx, "/test/hoge2", "fuga2")
	// client.Put(ctx, "/test/hoge1", "fuga1")
	// response, err := client.Get(ctx, "/test/",
	// 	clientv3.WithPrefix(), // キーに指定したプレフィックスから始まるキーをすべて取得(/test)
	// 	clientv3.WithSort(clientv3.SortByValue, clientv3.SortAscend), // 結果をソート(Valueの昇順)
	// 	clientv3.WithKeysOnly(), // キーのみを取得
	// 	// https://pkg.go.dev/github.com/coreos/etcd/clientv3#OpOption 他オプション
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, kv := range response.Kvs {
	// 	fmt.Printf("%s: %s\n", kv.Key, kv.Value)
	// }

	// fmt.Println("1----------------------")
	// client.Put(ctx, "/test/hoge", "123")
	// respo, _ := client.Get(ctx, "/test/hoge", clientv3.WithPrefix())
	// PrintResponse(respo)

	// fmt.Println("2----------------------")
	// client.Put(ctx, "/test/hoge", "456")
	// respo, _ = client.Get(ctx, "/test/hoge", clientv3.WithPrefix())
	// PrintResponse(respo)

	// fmt.Println("3----------------------")
	// client.Put(ctx, "/test/hage", "999")
	// respo, _ = client.Get(ctx, "/test/hage", clientv3.WithPrefix())
	// PrintResponse(respo)

	// fmt.Println("4----------------------")
	// respo, _ = client.Get(
	// 	ctx, "/test", clientv3.WithPrefix(),
	// 	clientv3.WithRev(respo.Kvs[0].CreateRevision),
	// )
	// PrintResponse(respo)

	// コンパクション:リビジョンを保存してあるデータの古い履歴を削除する
	client.Put(ctx, "/chapter3/compaction", "hoge")
	client.Put(ctx, "/chapter3/compaction", "fuga")
	client.Put(ctx, "/chapter3/compaction", "fuga")
	resp, err := client.Get(ctx, "/chapter3/compaction")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("--- prepared data: ")
	for i := resp.Kvs[0].CreateRevision; i <= resp.Kvs[0].ModRevision; i++ {
		r, err := client.Get(ctx, "/chapter3/compaction", clientv3.WithRev(i))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("rev: %d, value: %s\n", i, r.Kvs[0].Value)
	}
	_, err = client.Compact(ctx, resp.Kvs[0].ModRevision)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("--- compacted: %d\n", resp.Kvs[0].ModRevision)
	for i := resp.Kvs[0].CreateRevision; i <= resp.Kvs[0].ModRevision; i++ {
		r, err := client.Get(ctx, "/chapter3/compaction", clientv3.WithRev(i))
		if err != nil {
			fmt.Printf("failed to get: %v\n", err)
			continue
		}
		fmt.Printf("rev: %d, value: %s\n", i, r.Kvs[0].Value)
	}
}
