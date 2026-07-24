-- migrate:up
CREATE SCHEMA IF NOT EXISTS "notification";

CREATE TYPE IF NOT EXISTS "notification"."event_type" AS ENUM ('PAYMENT.INITIATED', 'PAYMENT.DEBITED', 'PAYMENT.CREDITED', 'PAYMENT.DEBIT.FAILED', 'PAYMENT.CREDIT.FAILED', 'PAYMENT.DEBIT.COMPENSATED')

CREATE TABLE IF NOT EXISTS "notification"."notifications" (
  "ksuid" varchar(255) PRIMARY KEY NOT NULL,
  "idempotency_key" varchar NOT NULL,
  "user_ksuid" varchar NOT NULL,
  "event_type" "notification"."event_type" NOT NULL,
  "transaction_ksuid" varchar NOT NULL,
  "context_data" JSONB NOT NULL,
  "inserted_at" timestamptz NOT NULL,
);

CREATE TYPE IF NOT EXISTS "notification"."channel" AS ENUM ('IN_APP', 'SMS', 'EMAIL', 'PUSH')

CREATE TYPE IF NOT EXISTS "notification"."status" AS ENUM ('PENDING', 'SENT', 'DELIVERED', 'FAILED')

CREATE TABLE IF NOT EXISTS "notification"."notification_dispatches" (
  "ksuid" varchar(255) PRIMARY KEY NOT NULL,
  "notification_ksuid" varchar NOT NULL,
  "channel" varchar NOT NULL,
  "recipient_target" varchar NOT NULL,
  "status" varchar NOT NULL,
  "provider_response" JSONB NOT NULL,
  "retry_count" integer NOT NULL,
  "sent_at" timestamptz NOT NULL,
  "inserted_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS "notification"."notifications";

DROP TYPE IF EXISTS "notification"."event_type";

DROP TYPE IF EXISTS "notification"."notification_dispatches";

DROP TYPE IF EXISTS "notification"."channel";

DROP SCHEMA IF EXISTS "notification";