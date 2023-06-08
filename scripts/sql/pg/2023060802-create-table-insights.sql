CREATE TABLE ak.insights (
                          id UUID PRIMARY KEY,
                          text TEXT NOT NULL,
                          topic VARCHAR(255) NOT NULL,
                          favorite BOOLEAN
);
