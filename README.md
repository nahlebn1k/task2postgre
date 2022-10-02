1. Open project and run docker-compose command
```
docker-compose up --build
```
2. Use next commands for DB migrations  

UP:
```
migrate -database postgresql://user:pass@localhost:5430/testdb_task?sslmode=disable -path ./migrations up
```
DOWN:
```
migrate -database postgresql://user:pass@localhost:5430/testdb_task?sslmode=disable -path ./migrations down
```
3. To test app go to `http://localhost:8000/get`
