package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
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
	// client.Put(ctx, "/chapter3/compaction", "hoge")
	// client.Put(ctx, "/chapter3/compaction", "fuga")
	// client.Put(ctx, "/chapter3/compaction", "fuga")
	// resp, err := client.Get(ctx, "/chapter3/compaction")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("--- prepared data: ")
	// for i := resp.Kvs[0].CreateRevision; i <= resp.Kvs[0].ModRevision; i++ {
	// 	r, err := client.Get(ctx, "/chapter3/compaction", clientv3.WithRev(i))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("rev: %d, value: %s\n", i, r.Kvs[0].Value)
	// }
	// _, err = client.Compact(ctx, resp.Kvs[0].ModRevision)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("--- compacted: %d\n", resp.Kvs[0].ModRevision)
	// for i := resp.Kvs[0].CreateRevision; i <= resp.Kvs[0].ModRevision; i++ {
	// 	r, err := client.Get(ctx, "/chapter3/compaction", clientv3.WithRev(i))
	// 	if err != nil {
	// 		fmt.Printf("failed to get: %v\n", err)
	// 		continue
	// 	}
	// 	fmt.Printf("rev: %d, value: %s\n", i, r.Kvs[0].Value)
	// }

	// fmt.Println("------:Watch ")
	// // Watch
	// // キー・バリューの変更を通知するためにWatch APIを提供
	// // 他のプログラムが情報を更新したことを定期的にチェックしていたのでは効率が悪くなるため
	// ch := client.Watch(ctx, "/chapter3/watch/", clientv3.WithPrefix()) // /chapter3/watch/"から始まるすべてのキーを監視対象
	// for resp := range ch {
	// 	if resp.Err() != nil {
	// 		log.Fatal(resp.Err())
	// 	}
	// 	for _, ev := range resp.Events {
	// 		switch ev.Type {
	// 		case clientv3.EventTypePut:
	// 			switch {
	// 			case ev.IsCreate():
	// 				fmt.Printf("CREATE %q : %q\n", ev.Kv.Key, ev.Kv.Value)
	// 			case ev.IsModify():
	// 				fmt.Printf("MODIFY %q : %q\n", ev.Kv.Key, ev.Kv.Value)
	// 			}
	// 		case clientv3.EventTypeDelete:
	// 			fmt.Printf("DELETE %q : %q\n", ev.Kv.Key, ev.Kv.Value)
	// 		}
	// 	}
	// }

	// // 取りこぼしを防ぐ
	// // Watch APIを呼び出すと、呼び出した時点からの変更が通知される
	// // Watchを利用しているプログラムが停止している間にetcdに変更が加えられた場合、その変更を取りこぼすことになってしまう
	// // atch APIを呼び出すときにclientv3.WithRev()オプションを指定することで、特定のリビジョンからの変更をすべて受け取ることが可能
	// rev := nextRev()
	// fmt.Printf("loaded revision: %d\n", rev)
	// ch := client.Watch(ctx, "/chapter3/watch_file", clientv3.WithRev(rev))
	// for resp := range ch {
	// 	if resp.Err() != nil {
	// 		log.Fatal(resp.Err())
	// 	}
	// 	for _, ev := range resp.Events {
	// 		doSomething(ev)
	// 		err := saveRev(ev.Kv.ModRevision)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		fmt.Printf("saved: %d\n", ev.Kv.ModRevision)
	// 	}
	// }

	// 	// Watchとコンパクション
	// 	// リビジョンを指定してWatchを開始した場合、そのキーがすでにコンパクションされている可能性がある
	// 	// WatchResponseのCompactRevisionを利用すると、コンパクションされていない中で最も古いリビジョンが取得できるので、このリビジョンを使ってWatchを再開するなどの処理を行える
	// 	go func() {
	// 		for i := 0; i < 10; i++ {
	// 			client.Put(ctx, "/chapter3/watch_compact", strconv.Itoa(i))
	// 			time.Sleep(100 * time.Millisecond)
	// 		}
	// 	}()
	// 	time.Sleep(300 * time.Millisecond)
	// 	resp, err := client.Get(ctx, "/chapter3/watch_compact")
	// 	if err != nil || resp.Count == 0 {
	// 		log.Fatal(err)
	// 	}
	// 	rev := resp.Kvs[0].ModRevision + 1
	// 	time.Sleep(300 * time.Millisecond)
	// 	resp, err = client.Get(ctx, "/chapter3/watch_compact")
	// 	if err != nil || resp.Count == 0 {
	// 		log.Fatal(err)
	// 	}
	// 	_, err = client.Compact(ctx, resp.Kvs[0].ModRevision)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// RETRY:
	// 	fmt.Printf("watch from rev: %d\n", rev)
	// 	ch := client.Watch(ctx, "/chapter3/watch_compact", clientv3.WithRev(rev))
	// 	for resp := range ch {
	// 		if resp.Err() == rpctypes.ErrCompacted {
	// 			rev = resp.CompactRevision
	// 			goto RETRY
	// 		} else if resp.Err() != nil {
	// 			log.Fatal(resp.Err())
	// 		}
	// 		for _, ev := range resp.Events {
	// 			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
	// 		}
	// 	}

	// // Lease
	// // キー・バリューに有効期限を指定することができる。
	// // キーの登録時に有効期限を指定しておくと、その期限が過ぎたときに対象のキーは自動的に削除される
	// lease, err := client.Grant(ctx, 5) // 有効期限を秒単位で指定
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = client.Put(ctx, "/chapter3/lease", "value", clientv3.WithLease(lease.ID))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = client.KeepAliveOnce(ctx, lease.ID) // KeepAliveOnce()を呼び出すと、有効期限が最初に指定した時間分だけ延長される
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // _, err = client.KeepAlive(ctx, lease.ID) // KeepAliveOnce()を周期的に呼び出すための仕組みとして、KeepAlive()がある
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }
	// _, err = client.Revoke(ctx, lease.ID) // 指定した期限までまだ時間がある場合でも、そのキーを失効させたい場合、Revoke()を利用
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for {
	// 	resp, err := client.Get(ctx, "/chapter3/lease")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if resp.Count == 0 {
	// 		fmt.Println("'/chapter3/lease' disappeared")
	// 		break
	// 	}
	// 	fmt.Printf("[%v] %s\n", time.Now().Format("15:04:05"), resp.Kvs[0].Value)
	// 	time.Sleep(1 * time.Second)
	// }

	// // Namespace
	// // キーにはアプリケーションごとにプレフィックスをつけるのが一般的
	// // アプリケーションを開発する際にすべてのキーにプレフィックスを指定するのは少々めんどう
	// newClient := namespace.NewKV(client.KV, "/chapter3")
	// _, err = newClient.Put(ctx, "/ns/1", "hoge")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// resp, _ := client.Get(ctx, "/chapter3/ns/1")
	// fmt.Printf("%s: %s\n", resp.Kvs[0].Key, resp.Kvs[0].Value)

	// _, err = client.Put(ctx, "/chapter3/ns/2", "test")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// resp, _ = newClient.Get(ctx, "/ns/2")
	// fmt.Printf("%s: %s\n", resp.Kvs[0].Key, resp.Kvs[0].Value)

	// // namespace適用後にキーバリュー操作用のクライアントやWatch用のクライアントを個別に管理するのは面倒
	// // 次のように既存のクライアントの機能を上書きしてしまうのがおすすめ
	// client.KV = namespace.NewKV(client.KV, "/chapter3")
	// client.Watcher = namespace.NewWatcher(client.Watcher, "/chapter3") // namespaceをWatch関連の操作に適用する
	// client.Lease = namespace.NewLease(client.Lease, "/chapter3")       // Lease関連の操作に適用する

	// Transaction
	// etcdにアクセスするクライアントが常に1つしか存在しないのであれば何も問題はない
	// 現実には複数のクライアントが同時にetcdにデータを書き込んだり読み込んだりする
	// key := "/chapter4/conflict"
	// _, err = client.Put(ctx, key, "10")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// go addValue(client, key, 5)
	// go addValue(client, key, -3)
	// time.Sleep(1 * time.Second)
	// resp, _ := client.Get(ctx, key)
	// fmt.Println(string(resp.Kvs[0].Value)) // 12になってほしいところだが15になったり7になったり

	// // 実際のTransaction利用
	// key := "/chapter4/txn"
	// _, err = client.Put(ctx, key, "10")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// go addValueTxn(client, key, 5)
	// go addValueTxn(client, key, -3)
	// time.Sleep(1 * time.Second)
	// resp, _ := client.Get(context.TODO(), key)
	// fmt.Println(string(resp.Kvs[0].Value))

	// // Session
	// // 作成したキーにリース期間が設定されるため、プロセスが不意に終了してもキーの有効期限が過ぎたらロックは解除される
	// // プロセスが生きている間はリース期間を更新し続けるようにもなっている
	// session, err := concurrency.NewSession(client) // デフォルトでのリース期間は60秒に設定
	// // この時間を変更したい場合は、次のようにconcurrency.WithTTL()を利用することができる
	// // session, err := concurrency.NewSession(client, concurrency.WithTTL(180))
	// // ロックを取得したプロセスが何か処理をしているときに、etcdとの接続が切れてしまっていたりリースを失効していた場合は、処理を中止すべき
	// select {
	// case <-session.Done(): // セッションが切れた通知を受け取ることが可能
	// 	log.Fatal("session has been orphaned")
	// }

	// // Mutex
	// // etcdの提供するMutexでは異なるサーバ上のプロセス間での排他制御が可能
	session, err := concurrency.NewSession(client)
	if err != nil {
		log.Fatal(err)
	}
	// mutex := concurrency.NewMutex(session, "/chapter4/mutex")
	// err = mutex.Lock(context.TODO()) // すでに他のプロセスがロックを取得済みだった場合は、ロックが取得できるまでブロックされます。 ロックが取得できなかったときにタイムアウトさせたい場合は、タイムアウトを設定したcontextを渡します。
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("acquired lock")
	// time.Sleep(5 * time.Second)
	// err = mutex.Unlock(context.TODO()) // deferを利用してスコープを抜けたときに必ずロックを解放するのがおすすめ
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("released lock")
	// ロックした後に一度ネットワーク接続が切れていたり、Sessionのリース期間が終了するかもしれない
	// プログラムは自分がロックしたつもりで動作しているのに、実際にはロックされていないという状況に陥ってしまう
	// そこで、ロックを取ったあとにetcdのキーバリューを操作する際には、トランザクションのIf条件にMutex.IsOwner()を指定する。
	mutex := concurrency.NewMutex(session, "/chapter4/mutex_txn")
RETRY:
	select {
	case <-session.Done():
		log.Fatal("session has been orphaned")
	default:
	}
	err = mutex.Lock(ctx)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Txn(context.TODO()).
		If(mutex.IsOwner()).
		Then(clientv3.OpPut("/chapter4/mutex_owner", "test")).
		Commit()
	if err != nil {
		log.Fatal(err)
	}
	if !resp.Succeeded {
		fmt.Println("the lock was not acquired")
		mutex.Unlock(context.TODO())
		goto RETRY
	}
	err = mutex.Unlock(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

}

func addValueTxn(client *clientv3.Client, key string, d int) {
RETRY:
	resp, _ := client.Get(context.TODO(), key)
	rev := resp.Kvs[0].ModRevision
	value, _ := strconv.Atoi(string(resp.Kvs[0].Value))
	value += d
	tresp, err := client.Txn(context.TODO()).
		If(clientv3.Compare(clientv3.ModRevision(key), "=", rev)).
		Then(clientv3.OpPut(key, strconv.Itoa(value))).
		Commit()
	if err != nil { // トランザクションのIfの条件が成立しなくてもエラーにはならない
		log.Fatal(err)
	}
	if !tresp.Succeeded {
		goto RETRY // データが書き込めなかった場合にgotoを使って最初から処理をやり直し
	}
}

// トランザクション記法
// If(・・・).Then(・・・).Else(・・・).Commit()が使える
// If(
//
//	clientv3.Compare(clientv3.ModRevision(key1), "=", rev1),
//	clientv3.Compare(clientv3.ModRevision(key2), "=", rev2),
//
// ).
// Then(・・・).
// Commit()
// 比較演算子は"="の他に"!="、"<"、">"が利用できます。
// clientv3.Value()で値の比較、clientv3.Version()でバージョンの比較、clientv3.CreateRevision()でCreateRevisionの比較をすることが可能
// If(clientv3util.KeyExists(key))やIf(clientv3util.KeyMissing(key))を利用すれば、キーの有無によって条件分岐をすることが可能
func exampleTxn(client *clientv3.Client, key1, key2, value1 string, d int) error {
	tresp, err := client.Txn(context.TODO()).
		Then(
			clientv3.OpPut(key1, value1),
			clientv3.OpDelete(key2),
		).
		Commit()
	if err != nil {
		return err
	}
	presp := tresp.Responses[0].GetResponsePut()
	dresp := tresp.Responses[1].GetResponseDeleteRange()
	fmt.Println(presp, dresp)
	return nil
}

// etcdから現在の値を読み取り、そこに引数で指定した値を追加して保存する関数
func addValue(client *clientv3.Client, key string, d int) {
	resp, _ := client.Get(context.TODO(), key)
	value, _ := strconv.Atoi(string(resp.Kvs[0].Value))
	value += d
	client.Put(context.TODO(), key, strconv.Itoa(value))
}

func nextRev() int64 {
	p := "./last_revision"
	f, err := os.Open(p)
	if err != nil {
		os.Remove(p)
		return 0
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		os.Remove(p)
		return 0
	}
	rev, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		os.Remove(p)
		return 0
	}
	return rev + 1
}

func saveRev(rev int64) error {
	p := "./last_revision"
	f, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_SYNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(strconv.FormatInt(rev, 10))
	return err
}

func doSomething(ev *clientv3.Event) {
	fmt.Printf("[%d] %s %q : %q\n", ev.Kv.ModRevision, ev.Type, ev.Kv.Key, ev.Kv.Value)
}
