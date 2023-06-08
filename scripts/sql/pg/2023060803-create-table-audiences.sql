CREATE TABLE ak.audiences (
                           id UUID PRIMARY KEY,
                           gender VARCHAR(255) NOT NULL,
                           birth_country VARCHAR(255) NOT NULL,
                           age_group VARCHAR(255) NOT NULL,
                           hours_spent_on_social INT,
                           num_purchases_last_month INT,
                           favorite BOOLEAN
);
