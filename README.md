# 検証方法

## リポジトリをクローンする
```
https://github.com/leslesnoa/chatbox-grpc.git
```

## mysqlコンテナ生成、テストデータを作成
```
cd chatbox-grpc
docker-compose up
```

## gRPCサーバを起動する
```
cd server
go run server.go
```

## grpc-gatewayを起動する
```
cd gateway
go run gateway.go
```

## クライアントから検証する
```
cd client
open index.html
```