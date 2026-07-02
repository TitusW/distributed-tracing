-- migrate:up
CREATE SCHEMA IF NOT EXISTS "account";

CREATE TABLE IF NOT EXISTS "account"."financial_accounts" (
  "ksuid" varchar(255) PRIMARY KEY NOT NULL,
  "current_amount" numeric,
  "inserted_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS "account"."users";

DROP SCHEMA IF EXISTS "account";