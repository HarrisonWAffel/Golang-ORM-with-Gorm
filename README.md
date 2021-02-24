#dbTrain
An example MVC-like micro-service to serve as an example as for how to possibly use gorm with a psql database in golang




 
This repository's main goal was practicing gorm, not 
designing a perfect micro-service. 


list of service packages 
+ `/posts`
+ `/user` 
+ `/userPosts`

all database models are located in the `models` folder 

all repository types are located in the `repositories` package

the API handlers are located in `handlers/handlers.go`

the ORM used for this project is Gorm

This project uses a simple database as an example of creating models, services, and repositories within golang using gorm.


# to start, run `docker-compose up`

database schema
```sql
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
```