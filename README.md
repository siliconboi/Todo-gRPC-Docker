### Simple Golang gRPC CRUD app

### Technologies used:
* Go
* gRPC
* Postgres
* GORM
* Docker

### Setup
1. Write the DB_URL in the db.go CreateDB func
2. Build the Docker images by running the Dockerfiles

```
docker build -t grpc-first -f server.Dockerfile .
docker build -t rest-first-wrapper -f client.Dockerfile .
```
3. Run the docker compose file

```
docker compose up --build
```
4. Congo!!!!!!!