CREATE TABLE  IF NOT EXISTS categories (
    "id" bigserial primary key,
    "name" character varying(250) NOT NULL DEFAULT '',
    "img" character varying(250) NOT NULL DEFAULT '',
    "updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')  NOT NULL,
    "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')  NOT NULL
    );

CREATE TABLE  IF NOT EXISTS products (
    "id" bigserial primary key,
    "cat_id" int,
    "name" character varying(250) NOT NULL DEFAULT '',
    "img" character varying(250) NOT NULL ,
    "item" json NOT NULL DEFAULT '{}'::json,
    "updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')  NOT NULL,
    "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')  NOT NULL,
    CONSTRAINT sub_cat_id_fk
        FOREIGN KEY ("cat_id")
            REFERENCES categories("id")
            ON UPDATE CASCADE ON DELETE CASCADE
    );

CREATE TABLE  IF NOT EXISTS banner (
    "id" bigserial primary key,
    "img" character varying(250) NOT NULL ,
    "updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')  NOT NULL,
    "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')  NOT NULL
    );

CREATE TABLE  IF NOT EXISTS news (
    "id" bigserial primary key,
    "title" character varying(250) NOT NULL,
    "content" text NOT NULL DEFAULT '',
    "img" character varying(250) NOT NULL,
    "date" character varying(250) NOT NULL,
    "updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')  NOT NULL,
    "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')  NOT NULL
    );