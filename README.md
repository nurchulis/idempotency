# idempotency

### Setup Environment
```
APP_NAME=Idempotency
APP_PORT=3000

DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=db_idempotency
DB_DRIVER=mysql
DB_AUTO_MIGRATE=true
DB_AUTO_SEED=true

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
```

### Run App
```
go run main.go
```

### Note

**Auto Migrate & Auto Seed setting by env config**
