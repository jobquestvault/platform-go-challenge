CREATE SCHEMA ak;

CREATE USER ak WITH PASSWORD 'ak';

GRANT USAGE ON SCHEMA ak TO ak;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA ak TO ak;