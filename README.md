## SSH tunnel
```bash
ssh -L 6379:INSTANCE.ruxlad.0001.sae1.cache.amazonaws.com:6379 -i ~/w/infrastructure/terraform/global/credentials/redis-bastion.pem ubuntu@18.230.22.135
```

## Find

Edit the following [index](https://github.com/devbytom/redis-test/blob/888dde7a9a4fe5525a86ad11667f0a39a325d54f/main.go#L27) parameter to some wishful value
```
val, err := rdb.Keys(ctx, "dex_descomplica_auth_*").Result()
```

## Run
Will print every index wich its value length is lesser than 1024 chars
```
go mod download
go run .
```