CREATE TABLE IF NOT EXISTS "users"(
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "email" VARCHAR NOT NULL UNIQUE,
    "username" VARCHAR NOT NULL UNIQUE,
    "created_at" TIMESTAMP NOT NULL DEFAULT now(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);


CREATE TABLE IF NOT EXISTS "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
