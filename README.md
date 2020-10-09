## SSH tunnel
```bash
ssh -L 6379:INSTANCE_HOST:6379 -i ~/w/infrastructure/terraform/global/credentials/redis-bastion.pem ubuntu@BASTIONIP
```

## Find

Edit the following [index](https://github.com/devbytom/redis-test/blob/888dde7a9a4fe5525a86ad11667f0a39a325d54f/main.go#L27) parameter to some wishful value
```
val, err := rdb.Keys(ctx, "*").Result()
```

## Run
Will print every index wich its value length is lesser than 1024 chars
```
go mod download
go run .
```
