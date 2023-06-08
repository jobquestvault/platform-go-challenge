CREATE TABLE ak.insights (
                          id UUID PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          user_id UUID NOT NULL,
                          text TEXT NOT NULL,
                          topic VARCHAR(255) NOT NULL,
                          favorite BOOLEAN
);
