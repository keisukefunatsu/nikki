CREATE TABLE "notes" (
"id" UUID NOT NULL,
PRIMARY KEY("id"),
"message" text NOT NULL,
"created_at" timestamp NOT NULL,
"updated_at" timestamp NOT NULL
);