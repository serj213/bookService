

CREATE TABLE authors(
    id serial PRIMARY KEY NOT NULL,
    author_name text NOT NULL,
    created_at 		timestamp with time zone 	DEFAULT now() NOT NULL,
    updated_at 		timestamp with time zone
);

CREATE TABLE categories(
    id serial PRIMARY KEY NOT NULL,
    category_name text NOT NULL,
    created_at 		timestamp with time zone 	DEFAULT now() NOT NULL,
    updated_at 		timestamp with time zone
);

CREATE TABLE users(
    id serial PRIMARY KEY NOT NULL,
    username text NOT NULL UNIQUE,
    email text NOT NULL UNIQUE,
    hash_password text NOT NULL,
    created_at 	timestamp with time zone 	DEFAULT now() NOT NULL,
    updated_at timestamp with time zone
);

CREATE TABLE books(
    id serial PRIMARY KEY NOT NULL,
    title text NOT NULL,
    author_id integer,
    categories_id integer,
    price integer NOT NULL,
    created_at timestamp with time zone    DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,

    FOREIGN KEY(author_id) REFERENCES authors(id),
    FOREIGN KEY(categories_id) REFERENCES categories(id)
);
