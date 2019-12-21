CREATE SEQUENCE IF NOT EXISTS users_id_seq;
CREATE TABLE "public"."users" (
                                "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
                                "username" varchar NOT NULL,
                                "email" varchar NOT NULL,
                                "password" varchar NOT NULL,
                                PRIMARY KEY ("id")
);

create unique index users_username_idx on users (username);
create unique index users_email_idx on users (email);

CREATE SEQUENCE IF NOT EXISTS tweets_id_seq;
CREATE TABLE "public"."tweets" (
                                 "id" int4 NOT NULL DEFAULT nextval('tweets_id_seq'::regclass),
                                 "message" text NOT NULL,
                                 "user_id" int4 NOT NULL,
                                 CONSTRAINT "tweets_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE ON UPDATE CASCADE,
                                 PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS user_subscriptions_id_seq;
CREATE TABLE "public"."user_subscriptions" (
                                             "id" int4 NOT NULL DEFAULT nextval('user_subscriptions_id_seq'::regclass),
                                             "user_id" int4 NOT NULL,
                                             "destination_user_id" int4 NOT NULL,
                                             CONSTRAINT "user_subscriptions_destination_user_id_fkey" FOREIGN KEY ("destination_user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE ON UPDATE CASCADE,
                                             CONSTRAINT "user_subscriptions_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE ON UPDATE CASCADE,
                                             PRIMARY KEY ("id")
);

