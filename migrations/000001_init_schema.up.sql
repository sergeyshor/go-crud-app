CREATE TABLE IF NOT EXISTS "users" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT(now())
);

CREATE TABLE IF NOT EXISTS "tweets" (
    "id" bigserial PRIMARY KEY,
    "author_id" bigint NOT NULL ,
    "content" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT(now())
);

ALTER TABLE "tweets" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id");