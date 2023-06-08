-- Install UUID-OSSP extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Inserting the user fixture
INSERT INTO ak.users (id, username, email, password, name)
VALUES ('c03dc326-7160-4b63-ac36-7105a4c96fa3', 'username', 'username@localhost.com', 'cGFzc3dvcmQ=', 'John Doe');

-- Creating ten rows for the 'charts' table
INSERT INTO ak.charts (id, name, user_id, title, x_axis_title, y_axis_title, data, favorite)
SELECT
    uuid_generate_v4(),
    'Chart ' || i,
    'c03dc326-7160-4b63-ac36-7105a4c96fa3',
    'Title ' || i,
    'X-Axis Title ' || i,
    'Y-Axis Title ' || i,
    ARRAY[i * 1.0, i * 2.0, i * 3.0],
    i % 2 = 0
FROM generate_series(1, 10) AS i;

-- Creating ten rows for the 'insights' table
INSERT INTO ak.insights (id, name, user_id, text, topic, favorite)
SELECT
    uuid_generate_v4(),
    'Insight ' || i,
    'c03dc326-7160-4b63-ac36-7105a4c96fa3',
    'Text ' || i,
    'Topic ' || i,
    i % 2 = 0
FROM generate_series(1, 10) AS i;

-- Creating ten rows for the 'audiences' table
INSERT INTO ak.audiences (id, name, user_id, gender, birth_country, age_group, hours_spent_on_social, num_purchases_last_month, favorite)
SELECT
    uuid_generate_v4(),
    'Audience ' || i,
    'c03dc326-7160-4b63-ac36-7105a4c96fa3',
    'Gender ' || i,
    'Country ' || i,
    'Age Group ' || i,
    i * 10,
    i,
    i % 2 = 0
FROM generate_series(1, 10) AS i;
