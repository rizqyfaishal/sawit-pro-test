/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

ALTER DATABASE database SET timezone TO 'Asia/Jakarta';

CREATE TABLE users
(
    id                  BIGSERIAL PRIMARY KEY,
    phone_number        VARCHAR(13)  NOT NULL UNIQUE,
    full_name           VARCHAR(60)  NOT NULL,
    password            VARCHAR(255) NOT NULL,
    login_success_count BIGINT    DEFAULT 0,
    created_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX phone_number_unique_index ON users (phone_number);

CREATE FUNCTION update_updated_at_users_task()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at
BEFORE UPDATE
ON
   users
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_users_task();