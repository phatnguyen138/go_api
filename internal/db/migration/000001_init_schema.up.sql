CREATE TABLE todo (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date DATE DEFAULT CURRENT_DATE,
    completed BOOLEAN DEFAULT FALSE
);