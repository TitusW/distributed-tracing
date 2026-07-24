export DATABASE_URL=postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable
export DBMATE_MIGRATIONS_DIR='./db/postgres/dbmate/migrations'
export DBMATE_SCHEMA_FILE='./db/postgres/dbmate/schema'
export DBMATE_NO_DUMP_SCHEMA=false
export DBMATE_WAIT=false
export DBMATE_WAIT_TIMEOUT='2m0s'
export DBMATE_MIGRATIONS_TABLE="dbmate_schema_migrations"