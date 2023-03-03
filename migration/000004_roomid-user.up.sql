CREATE TABLE "user_room" (
                             "user_id" uuid NOT NULL,
                             "room_id" uuid NOT NULL,
                             PRIMARY KEY ("user_id", "room_id"),
                             CONSTRAINT "user_room_user_fkey" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE,
                             CONSTRAINT "user_room_room_fkey" FOREIGN KEY ("room_id") REFERENCES "room" ("id") ON DELETE CASCADE
);
