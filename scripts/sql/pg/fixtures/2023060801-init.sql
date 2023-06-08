-- Install UUID-OSSP extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Creating ten rows for the 'charts' table
INSERT INTO ak.charts (id, title, x_axis_title, y_axis_title, data, favorite)
SELECT
    uuid_generate_v4(),
    'Chart ' || i,
    'X-Axis Title ' || i,
    'Y-Axis Title ' || i,
    ARRAY[i * 1.0, i * 2.0, i * 3.0],
    i % 2 = 0

FROM generate_series(1, 10) AS i;

-- Creating ten rows for the 'insights' table
INSERT INTO ak.insights (id, text, topic, favorite)
SELECT
    uuid_generate_v4(),
    'Insight ' || i,
    'Topic ' || i,
    i % 2 = 0

FROM generate_series(1, 10) AS i;

-- Creating ten rows for the 'audience' table
INSERT INTO ak.audiences (id, gender, birth_country, age_group, hours_spent_on_social, num_purchases_last_month, favorite)
SELECT
    uuid_generate_v4(),
    'Gender ' || i,
    'Country ' || i,
    'Age Group ' || i,
    i * 10,
    i,
    i % 2 = 0

FROM generate_series(1, 10) AS i;
