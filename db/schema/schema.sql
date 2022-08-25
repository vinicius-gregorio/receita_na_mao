-- CREATE TABLE "category" (
--   "id" bigserial PRIMARY KEY,
--   "name" varchar UNIQUE NOT NULL,
--   "created_at" timestamptz NOT NULL DEFAULT (now())
-- );

-- CREATE TABLE "recipes" (
--   "id" bigserial PRIMARY KEY,
--   "name" varchar NOT NULL,
--   "categories" varchar[] NOT NULL,
--   "description" varchar NOT NULL,
--   "prepare_method" varchar NOT NULL,
--   "ingredients" varchar[] NOT NULL,
--   "rating" integer DEFAULT 1,
--   "preparation_time" timestamptz NOT NULL DEFAULT (now()),
--   "created_at" timestamptz NOT NULL DEFAULT (now())
-- );

-- CREATE TABLE "users" (
--   "id" bigserial PRIMARY KEY,
--   "email" varchar UNIQUE NOT NULL,
--   "hashed_password" varchar NOT NULL,
--   "nick_name" varchar NOT NULL,
--   "updated_at" timestamptz NOT NULL DEFAULT (now()),
--   "created_at" timestamptz NOT NULL DEFAULT (now())
-- );
-- -- ALTER TABLE recipes 
-- --     ALTER COLUMN categories SET DEFAULT array[]::varchar[];

CREATE TABLE "category" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "recipes" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "categories" varchar[],
  "description" varchar NOT NULL,
  "prepare_method" varchar NOT NULL,
  "ingredients" varchar[] NOT NULL,
  "rating" integer DEFAULT 1,
  "preparation_time" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "nick_name" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
