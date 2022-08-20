CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "password_changed_at" timestamptz NOT NULL DEFAULT (now()),
  "balance" float NOT NULL DEFAULT 100000
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "ticker" varchar NOT NULL,
  "quantity" bigint NOT NULL,
  "price" float NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "transactions" ("username");

COMMENT ON COLUMN "transactions"."quantity" IS 'can be positive or negative';

ALTER TABLE "transactions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
