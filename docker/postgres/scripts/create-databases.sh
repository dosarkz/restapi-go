#!/usr/bin/env bash

set -e
set -u

function create_user_and_database() {
	local database=$1
	echo "  Creating user and database '$database'"
	psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
	    CREATE USER $database;
	    CREATE DATABASE $database;
	    GRANT ALL PRIVILEGES ON DATABASE $database TO $database;
EOSQL
}

if [ -n "$POSTGRES_MAIN_DB" ]; then
	echo "Main database creation requested: $POSTGRES_MAIN_DB"
  create_user_and_database $(echo $POSTGRES_MAIN_DB)
	echo "Main databases created"
fi

if [ -n "$POSTGRES_TEST_DB" ]; then
	echo "Test database creation requested: $POSTGRES_TEST_DB"
  create_user_and_database $(echo $POSTGRES_TEST_DB)
	echo "Test databases created"
fi