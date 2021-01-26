CREATE TABLE "users" (
  "id" uuid VARCHAR(36) PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "password" VARCHAR(255) NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp
);