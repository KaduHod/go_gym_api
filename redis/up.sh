#/bin/bash
#sudo chown -R 999:999 /var/log/gymapi                                                                                                                                   130 ↵
#sudo chmod -R 755 /var/log/gymapi
docker run -p 6379:6379 \
    --name gymredis \
    --hostname redisgym \
    --network gym \
    -v /var/log/gymapi:/var/log/redis \
    -d kaduhod/gymredis


