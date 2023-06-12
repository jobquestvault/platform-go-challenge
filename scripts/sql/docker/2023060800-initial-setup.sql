-- Check if the database exists
DO $$BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_database WHERE datname = 'pgc'
    ) THEN
        CREATE DATABASE pgc;
    END IF;
END$$;

\c pgc;

-- Check if the schema exists
DO $$BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_namespace WHERE nspname = 'ak'
    ) THEN
        CREATE SCHEMA ak;
    END IF;
END$$;

-- Check if the user exists
DO $$BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_user WHERE usename = 'ak'
    ) THEN
        CREATE USER ak WITH PASSWORD 'ak';
    END IF;
END$$;

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA ak TO ak;
GRANT USAGE ON SCHEMA ak TO ak;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
