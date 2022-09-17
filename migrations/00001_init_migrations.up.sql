CREATE EXTENSION IF NOT EXISTS "uuid-ossp" SCHEMA public;

CREATE TYPE GENDER_TYPE AS ENUM ('male', 'female');

CREATE TABLE "user"
(
    id          UUID        NOT NULL DEFAULT public.uuid_generate_v4(),
    first_name  TEXT        NOT NULL,
    last_name   TEXT        NOT NULL,
    gender      GENDER_TYPE NOT NULL,
    birthday    DATE        NOT NULL,
    created     TIMESTAMP   NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX idx_unique_user ON "user" (first_name, last_name, birthday);

CREATE TYPE DESTINATION_TYPE AS ENUM ('Mars', 'Moon', 'Pluto', 'Asteroid Belt', 'Europa', 'Titan', 'Ganymede');

CREATE TABLE "launch"
(
    ID              UUID                NOT NULL DEFAULT public.uuid_generate_v4(),
    launchpad_id    TEXT                NOT NULL,
    destination     DESTINATION_TYPE    NOT NULL,
    launch_date     DATE                NOT NULL,
    user_id         UUID                NOT NULL,
    created         TIMESTAMP           NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES "user" (id)
);

CREATE UNIQUE INDEX idx_unique_launch ON "launch" (launchpad_id, launch_date);
