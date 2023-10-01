# etcd-with-golang

## コマンド

### 書き込み

- curl -X POST -d '{"key": "hoge","value": "hogehoge"}' http://localhost:12379/v3/kv/put

### 読み込み

- curl -L http://localhost:12379/v3/kv/range -X POST -d '{"key": "hoge"}'
