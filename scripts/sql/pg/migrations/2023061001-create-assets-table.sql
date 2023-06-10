CREATE TABLE ak.assets (
                           id UUID PRIMARY KEY,
                           user_id UUID NOT NULL REFERENCES ak.users (id),
                           asset_id UUID NOT NULL,
                           asset_type VARCHAR(20) NOT NULL,
                           name VARCHAR(255),
                           description TEXT,
                           UNIQUE (user_id, asset_id, asset_type)
);
