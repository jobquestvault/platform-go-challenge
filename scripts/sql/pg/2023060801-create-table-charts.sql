CREATE TABLE ak.charts (
                        id UUID PRIMARY KEY,
                        title VARCHAR(255) NOT NULL,
                        x_axis_title VARCHAR(255) NOT NULL,
                        y_axis_title VARCHAR(255) NOT NULL,
                        data NUMERIC[],
                        favorite BOOLEAN
);
