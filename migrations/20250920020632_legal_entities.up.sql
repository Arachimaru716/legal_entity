CREATE TABLE "public"."legal_entities" (
  "uuid" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" varchar(255) NOT NULL DEFAULT ''::character varying,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "deleted_at" timestamptz,
  CONSTRAINT "legal_entities_pkey" PRIMARY KEY ("uuid")
);

CREATE INDEX "legal_entities_uuid" ON "public"."legal_entities" ("uuid");
CREATE INDEX legal_entities_name_idx ON "public"."legal_entities" (lower("name"));
