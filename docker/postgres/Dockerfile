FROM postgres:14-alpine3.15

COPY scripts/create-databases.sh /docker-entrypoint-initdb.d/

RUN chmod +x /docker-entrypoint-initdb.d/create-databases.sh
