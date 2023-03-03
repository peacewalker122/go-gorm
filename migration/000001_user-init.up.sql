CREATE TYPE user_status AS ENUM ('active', 'deactive');

CREATE TABLE "user" (
                        "id" uuid PRIMARY KEY,
                        "username" varchar NOT NULL,
                        "email" varchar NOT NULL,
                        "password" varchar NOT NULL,
                        "status" user_status NOT NULL,
                        "created_at" timestamp WITH TIME ZONE DEFAULT NOW(),
                        "deleted_at" timestamp WITH TIME ZONE NULL
);

ALTER TABLE "user" ADD CONSTRAINT "user_username_unique" UNIQUE ("username");
ALTER TABLE "user" ADD CONSTRAINT "user_email_unique" UNIQUE ("email");

ALTER TABLE "user" ADD CONSTRAINT "user_status_check" CHECK ("status" IN ('active', 'deactive'));
CREATE INDEX "user_status_index" ON "user" (created_at);

insert into "user" (id, username, email, password, status) values ('00000000-0000-0000-0000-000000000000', 'zero', 'admin@localhost', 'zero', 'active');
