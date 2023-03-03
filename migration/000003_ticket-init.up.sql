CREATE TYPE ticket_status AS ENUM ('todo', 'on-progress', 'done');
CREATE TABLE "ticket" (
                          "id" uuid PRIMARY KEY,
                          "description" varchar null,
                          "metadata" varchar,
                          "issued_by" uuid not null,
                          "assigned" uuid null,
                          "status" ticket_status not null default 'todo'::ticket_status,
                          "created_at" timestamp with time zone not null ,
                            "updated_at" timestamp with time zone null
);

ALTER TABLE "ticket" ADD FOREIGN KEY ("issued_by") REFERENCES public.user ("id");

ALTER TABLE "ticket" ADD FOREIGN KEY ("assigned") REFERENCES public.user ("id");