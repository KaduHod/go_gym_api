FROM redis
RUN mkdir -p /var/log/redis && \
    chown redis:redis /var/log/redis && \
    chmod 755 /var/log/redis
COPY redis.conf /usr/local/etc/redis/redis.conf
CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]
