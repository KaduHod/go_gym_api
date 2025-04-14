```bash
docker run -d --name gymdb --hostname gymdb --network gym -p 3306:3306 -e MYSQL_ROOT_PASSWORD=<senha> kaduhod/gymdb
docker run -p 6380:6380 --name gymredis --hostname redisgym --network gym -v /var/log/gymapi:/var/log/redis -d kaduhod/gymredis --port 6380
docker run -p 3005:3005 --name gymapi --hostname gymapi --network gym -v /var/log/gymapi:/app/logs -d kaduhod/gymapi
```
