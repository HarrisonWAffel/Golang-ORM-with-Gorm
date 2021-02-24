CREATE DATABASE learn;
\c learn;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
                       id uuid DEFAULT uuid_generate_v4 () UNIQUE NOT NULL,
                       user_name VARCHAR ( 50 ) UNIQUE NOT NULL,
                       password VARCHAR ( 50 ) NOT NULL,
                       email VARCHAR ( 255 ) UNIQUE NOT NULL,
                       last_login TIMESTAMP,
                       created_at TIMESTAMP,
                       updated_at TIMESTAMP,
                       deleted_at TIMESTAMP
);

CREATE TABLE posts (
                       id uuid DEFAULT uuid_generate_v4 () UNIQUE NOT NULL,
                       content VARCHAR(9999),
                       private BOOL DEFAULT 'f',
                       created_at TIMESTAMP,
                       updated_at TIMESTAMP,
                       deleted_at TIMESTAMP
);

CREATE TABLE user_posts (
                            id uuid DEFAULT uuid_generate_v4 () UNIQUE NOT NULL,
                            user_id uuid DEFAULT uuid_generate_v4()  NOT NULL,
                            post_id uuid DEFAULT uuid_generate_v4()  NOT NULL,
                            private BOOL DEFAULT 'f',
                            created_at TIMESTAMP,
                            updated_at TIMESTAMP,
                            deleted_at TIMESTAMP,
                            PRIMARY KEY (user_id, post_id),
                            FOREIGN KEY (post_id)
                                REFERENCES posts (id),
                            FOREIGN KEY (user_id)
                                REFERENCES users (id)
);