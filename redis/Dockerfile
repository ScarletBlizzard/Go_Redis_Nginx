FROM debian
RUN apt-get update && \
    apt-get install -y lsb-release curl gpg && \
    curl -fsSL https://packages.redis.io/gpg | gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg && \
    echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" > /etc/apt/sources.list.d/redis.list && \
    apt-get update && \
    apt-get install -y redis && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*
EXPOSE 6379
ENV CONF_PATH /etc/redis/redis.conf
COPY redis.conf $CONF_PATH
COPY *.crt redis.key ./
CMD redis-server $CONF_PATH --daemonize no --requirepass $REDIS_PASSWORD