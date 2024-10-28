## 概要
- `https://yukizemi.online` `https://yukizemi.online/spoke`へのアクセスを監視
- Cloud Run, Cloud SQLへデプロイ
- [React製クライアント](https://github.com/hiroshijp/hcce-observer-client)あり(認証機能を導入してから追いついていない)

## 技術スタック
- echo
- jwt

## 手順(開発環境)
```
$ git clone https://github.com/hiroshijp/hcce-observer.git
$ touch .env 　　　#docker-compose.yaml参考
$ docker compose up -d 
$ docker exec -it {{ container_name }} bash
$ cd /code/app
$ go mod tidy && go mod tidy
```
## API
### public
`localhost:8080/visited`: yukizemi.onlineクライアントのjsファイルに仕込む。
```
curl -X POST localhost:8080/visited \
    -H "Content-Type: application/json" \
    -d '{
            "visitor": {
                "mail": "{{ mail }}"
            },
            "visited_from": "{{ url }}"
        }'    
```
`localhost:8080/signin`: jwtが返る。adminの場合はカスタムクレイムのadminがtrueになる。
```
curl -X POST localhost:8080/signin -u {{ name }}:{{ password }} 
```

### private
`localhost:8080/api/history`: 訪問履歴を取得する。
```
curl -X GET localhost:8080/api/history \
    -H "Authorization: Bearer {{ jwt }}"
```
`localhost:8080/api/history/tx`: トランザクションを張って訪問履歴を取得する。
```
curl -X GET localhost:8080/api/history/tx \
    -H "Authorization: Bearer {{ jwt }}"
```
### admin only
`localhost:8080/api/user`: adminのみ新規ユーザーを追加できる。
```
curl -X POST localhost:8080/api/user \
    -H "Authorization: Bearer {{ jwt }}" \
    -H "Content-Type: application/json" \
    -d '{
            "name": "{{ name }}",
            "password": "{{ password}}"
        }'       
```

## TODO
- 今後の機能追加に備えgo-migrateを導入
- GET /api/userを追加
- GET /api/visitorを追加




