#!/usr/bin/env bash

set -euo pipefail


if [ "${PG_PASSWORD-}" == "" ]; then
    PG_PASSWORD=dev_awTf9d2GceKRNzhkCb4H5B8nfmq
fi

until psql -c "select 1" > /dev/null 2>&1; do
  echo "Waiting for postgres server..."
  sleep 1
done

psql -v ON_ERROR_STOP=1  --dbname=template1  <<- EOSQL
    CREATE DATABASE skeleton_backend WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';

    CREATE USER skeleton_backend WITH PASSWORD '$PG_PASSWORD';
    GRANT ALL PRIVILEGES ON DATABASE skeleton_backend TO skeleton_backend;
    GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO skeleton_backend;
    GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO skeleton_backend;
    GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public TO skeleton_backend;
    ALTER DATABASE skeleton_backend OWNER TO skeleton_backend;
    ALTER ROLE skeleton_backend SET statement_timeout TO '5s';

    \connect skeleton_backend
    CREATE EXTENSION pg_trgm;
    SELECT * FROM pg_extension;
EOSQL
