export MYSQL_DSN='golang:845924@tcp(127.0.0.1:3306)/MINIMAL_CHAT?charset=utf8mb4&parseTime=True&loc=Local'


{
    "username": "Alice",
    "password": "123456",
    "avatar_id": "1"
}

{
    "username": "bob",
    "password": "123456",
    "avatar_id": "2"
}

heartbeat

{"status":1,"data":{"uid":"1","room_id":"1","avatar_id":"1","username":"alice"}}

{"status":5,"data":{"to_uid":"2","content":"你好，Bob"}}

{"status":2,"data":{"room_id":"1"}}