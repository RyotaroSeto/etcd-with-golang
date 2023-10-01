package main

import (
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// revision:etcdのリビジョン番号。クラスタ全体で一つの値が利用。tcdに何らかの変更(キーの追加、変更、削除)が加えられると値が1増える
// create_revision:キーが作成されたときのリビジョン番号
// mod_revision:キーの内容が最後に変更されたときのリビジョン番号
// version:キーのバージョン。このキーに変更が加えられると値が1増える
func PrintResponse(resp *clientv3.GetResponse) {
	fmt.Println("header: " + resp.Header.String())
	for i, kv := range resp.Kvs {
		fmt.Printf("kv[%d]: %s\n", i, kv.String()) // RevisionとVersion確認
	}
}
