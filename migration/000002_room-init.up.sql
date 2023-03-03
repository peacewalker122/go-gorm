
CREATE TABLE "room" (
                        "id" uuid PRIMARY KEY,
                        "created_by" uuid,
                        "name" varchar,
                        "description" varchar,
                        "max_member" int,
                        "created_at" timestamp with time zone
);

ALTER TABLE "room" ADD FOREIGN KEY ("created_by") REFERENCES public.user (id);
