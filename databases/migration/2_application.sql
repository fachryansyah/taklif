CREATE TABLE "application"
(
    "id"         uuid VARCHAR(36) PRIMARY KEY,
    "pid"        INTEGER,
    "name"       VARCHAR(255) NOT NULL,
    "runtime"    INTEGER      NOT NULL,
    "path"       TEXT         NOT NULL,
    "created_at" timestamp,
    "updated_at" timestamp
);