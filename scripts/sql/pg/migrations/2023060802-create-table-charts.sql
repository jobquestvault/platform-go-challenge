CREATE TABLE ak.charts (
                        id UUID PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        owner_id UUID NOT NULL,
                        title VARCHAR(255) NOT NULL,
                        x_axis_title VARCHAR(255) NOT NULL,
                        y_axis_title VARCHAR(255) NOT NULL,
                        data FLOAT[]
);
