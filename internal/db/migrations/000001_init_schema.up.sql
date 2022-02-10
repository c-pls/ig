CREATE TABLE "users"
(
    "id"              bigserial,
    "user_id"         varchar UNIQUE              NOT NULL PRIMARY KEY,
    "username"        varchar                     NOT NULL,
    "salted_password" varchar                     NOT NULL,
    "first_name"      varchar                     NOT NULL,
    "last_name"       varchar                     NOT NULL,
    "bio"             varchar                     NOT NULL,
    "avatar_url"      varchar                     NOT NULL,
    "created_at"      timestamptz DEFAULT (now()) NOT NULL,
    "updated_at"      timestamptz DEFAULT (now()) NOT NULL
);

CREATE TABLE "posts"
(
    "id"         bigserial,
    "post_id"    varchar UNIQUE              NOT NULL PRIMARY KEY,
    "user_id"    varchar                     NOT NULL,
    "caption"    varchar                     NOT NULL,
    "longitude"  float                       NOT NULL,
    "latitude"   float                       NOT NULL,
    "created_at" timestamptz DEFAULT (now()) NOT NULL,
    "updated_at" timestamptz DEFAULT (now()) NOT NULL
);

CREATE TABLE "photos"
(
    "id"         bigserial,
    "photo_id"   varchar UNIQUE              NOT NULL PRIMARY KEY,
    "post_id"    varchar                     NOT NULL,
    "url"        varchar                     NOT NULL,
    "created_at" timestamptz DEFAULT (now()) NOT NULL
);

CREATE TABLE "comments"
(
    "id"         bigserial,
    "comment_id" varchar UNIQUE              NOT NULL PRIMARY KEY,
    "user_id"    varchar                     NOT NULL,
    "parent_id"  varchar                     NOT NULL,
    "content"    varchar                     NOT NULL,
    "type"       varchar                     NOT NULL,
    "created_at" timestamptz DEFAULT (now()) NOT NULL,
    "updated_at" timestamptz DEFAULT (now()) NOT NULL
);

CREATE TABLE "likes"
(
    "id"         bigserial,
    "like_id"    varchar UNIQUE PRIMARY KEY,
    "parent_id"  varchar                     NOT NULL,
    "user_id"    varchar                     NOT NULL,
    "created_at" timestamptz DEFAULT (now()) NOT NULL,
    "updated_at" timestamptz DEFAULT (now()) NOT NULL,
    "type"       varchar                     NOT NULL,
    "active"     boolean                     NOT NULL
);

CREATE TABLE "follows"
(
    "id"               bigserial,
    "follower_from_id" varchar                     NOT NULL,
    "follow_to_id"     varchar                     NOT NULL,
    "created_at"       timestamptz DEFAULT (now()) NOT NULL,
    "updated_at"       timestamptz DEFAULT (now()) NOT NULL,
    "active"           boolean                     NOT NULL
);

CREATE TABLE "hash_tags"
(
    "id"         bigserial,
    "hashtag_id" varchar PRIMARY KEY,
    "content"    varchar NOT NULL
);

CREATE TABLE "tags"
(

    "id"      bigserial,
    "tag_id"  varchar PRIMARY KEY NOT NULL,
    "post_id" varchar
);

-- CREATE TABLE "posts_photos"
-- (
--
--     "id"       bigserial,
--     "post_id"  varchar,
--     "photo_id" varchar,
--     PRIMARY KEY ("post_id", "photo_id")
-- );

-- CREATE TABLE "posts_comments"
-- (
--
--     "id"         bigserial,
--     "post_id"    varchar,
--     "comment_id" varchar,
--     PRIMARY KEY ("post_id", "comment_id")
-- );

CREATE TABLE "posts_hashtags"
(

    "id"         bigserial,
    "hashtag_id" varchar,
    "post_id"    varchar,
    PRIMARY KEY ("hashtag_id", "post_id")
);

ALTER TABLE "posts"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE;

ALTER TABLE "photos"
    ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id") ON DELETE CASCADE;

-- ALTER TABLE "posts_photos"
--     ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id") ON DELETE CASCADE;
--
-- ALTER TABLE "posts_photos"
--     ADD FOREIGN KEY ("photo_id") REFERENCES "photos" ("photo_id") ON DELETE CASCADE;

ALTER TABLE "posts_hashtags"
    ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id") ON DELETE CASCADE;

ALTER TABLE "posts_hashtags"
    ADD FOREIGN KEY ("hashtag_id") REFERENCES "hash_tags" ("hashtag_id") ON DELETE CASCADE;

ALTER TABLE "comments"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE;

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "likes" ("parent_id");

CREATE INDEX ON "tags" ("post_id");

-- Trigger function for updated_field
CREATE OR REPLACE FUNCTION update_modified_timestamp() RETURNS TRIGGER
    LANGUAGE plpgsql
AS
$$
BEGIN
    IF (NEW != OLD) THEN
        NEW.updated_at = CURRENT_TIMESTAMP;
        RETURN NEW;
    END IF;
    RETURN OLD;
END;
$$;

CREATE TRIGGER update_timestamp_user
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE update_modified_timestamp();

CREATE TRIGGER update_timestamp_post
    BEFORE UPDATE
    ON posts
    FOR EACH ROW
EXECUTE PROCEDURE update_modified_timestamp();

CREATE TRIGGER update_timestamp_comments
    BEFORE UPDATE
    ON comments
    FOR EACH ROW
EXECUTE PROCEDURE update_modified_timestamp();

CREATE TRIGGER update_timestamp_like
    BEFORE UPDATE
    ON likes
    FOR EACH ROW
EXECUTE PROCEDURE update_modified_timestamp();
