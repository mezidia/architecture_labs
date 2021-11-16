-- Create tables.
DROP TABLE IF EXISTS "channels";
CREATE TABLE "channels"
(
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL UNIQUE
);

-- Insert demo data.
INSERT INTO "channels" (name) VALUES ('lectures');
INSERT INTO "channels" (name) VALUES ('practice');
