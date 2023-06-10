-- Inserting the user fixture
INSERT INTO ak.users (id, username, email, password, name)
VALUES
    ('c03dc326-7160-4b63-ac36-7105a4c96fa3', 'owner', 'owner@localhost.com', '5d2cGFzc3dvcmQ=', 'Admin'),
    ('efd8cec6-3e45-4fb1-b0d7-3a1be9cfae2c', 'johndoe', 'john.dow@localhost.com', 'QW5a1kcmVhMzI1=', 'John Doe'),
    ('bc59a3bd-3291-4832-adf3-611e104cc846', 'emilysmith', 'emily.smith@localhost.com', 'bG9yZW56c29tZA=', 'Emily Smith');

-- Creating ten rows for the 'charts' table
INSERT INTO ak.charts (id, name, owner_id, title, x_axis_title, y_axis_title, data)
SELECT
    uuid_generate_v4(),
    'Chart ' || i,
    'c03dc326-7160-4b63-ac36-7105a4c96fa3',
    'Title ' || i,
    'X-Axis Title ' || i,
    'Y-Axis Title ' || i,
    ARRAY[i * 1.0, i * 2.0, i * 3.0]
FROM generate_series(1, 10) AS i;

-- Creating ten rows for the 'insights' table
INSERT INTO ak.insights (id, name, owner_id, text, topic)
SELECT
    uuid_generate_v4(),
    'Insight ' || i,
    'c03dc326-7160-4b63-ac36-7105a4c96fa3',
    'Text ' || i,
    'Topic ' || i
FROM generate_series(1, 10) AS i;

-- Creating ten rows for the 'audiences' table
INSERT INTO ak.audiences (id, name, owner_id, gender, birth_country, age_group, hours_spent_on_social, num_purchases_last_month)
SELECT
    uuid_generate_v4(),
    'Audience ' || i,
    'c03dc326-7160-4b63-ac36-7105a4c96fa3',
    'Gender ' || i,
    'Country ' || i,
    'Age Group ' || i,
    i * 10,
    i
FROM generate_series(1, 10) AS i;
