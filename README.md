# TODO api-service

### Setup
Add .env file 
```
HTTP_PORT=8000

DB_PORT=5432
DB_HOST=postgresdb
DB_NAME={dbname}
DB_SSLMODE=disable
DB_DRIVER=postgres
DB_PASSWORD={password}
DB_USER={user}

MIGRATE_DIR=db/migration

POSTGRES_USER: {user}
POSTGRES_PASSWORD: {password}
POSTGRES_DB: {dbname}

JWT_SECRET={sercret}
JWT_EXPIRATION=86400
JWT_HEADER=Authorizatio
```
Run app
```
docker-compose up --build -d
```
