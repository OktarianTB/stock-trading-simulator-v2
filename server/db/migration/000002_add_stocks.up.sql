CREATE TABLE "stocks" (
  "username" varchar,
  "ticker" varchar,
  "quantity" bigint NOT NULL DEFAULT 0,
  PRIMARY KEY ("username", "ticker")
);

CREATE INDEX ON "stocks" ("username");

CREATE INDEX ON "stocks" ("ticker");

CREATE INDEX ON "stocks" ("username", "ticker");

ALTER TABLE "stocks" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");