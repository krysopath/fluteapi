CREATE TABLE "flutes" (
  "id" bigserial PRIMARY KEY,
  "description" varchar NOT NULL,
  "available" boolean NOT NULL DEFAULT(false),
  "key" varchar NOT NULL,
  "name" varchar NOT NULL,
  "material" varchar NOT NULL,
  "holes" int NOT NULL DEFAULT(0),
  "scale" varchar NOT NULL,
  "pictures" varchar [],
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
