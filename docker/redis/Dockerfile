FROM redis

ENV REDIS_PASSWORD secret

CMD ["sh", "-c", "exec redis-server --requirepass \"$REDIS_PASSWORD\""]
