```bash
docker run -d --name gymdb --hostname gymdb --network gym -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 kaduhod/gymdb
docker run -p 6379:6379 --hostname video-redis --name video-redis --hostname redisgym --network gym -d redis
docker run -p 3005:3005 --name gymapi --hostname gymapi --network gym kaduhod/gymapi
```
