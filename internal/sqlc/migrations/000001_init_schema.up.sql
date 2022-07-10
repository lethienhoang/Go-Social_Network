CREATE TABLE "users"(
  "id" VARCHAR NOT NULL,
  "email" VARCHAR NOT NULL,
  "username" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now():::TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT now():::TIMESTAMP,
  CONSTRAINT users_pkey PRIMARY KEY (id ASC),
  UNIQUE INDEX users_email_key (email ASC),
  UNIQUE INDEX users_username_key (username ASC)
);


CREATE TABLE "accounts" (
  "id" INT8 NOT NULL DEFAULT unique_rowid(),
  "owner" VARCHAR NOT NULL,
  "balance" INT8 NOT NULL,
  "currency" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
  CONSTRAINT accounts_pkey PRIMARY KEY (id ASC)
);
