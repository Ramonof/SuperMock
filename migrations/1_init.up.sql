create table reststubs (
    id serial PRIMARY KEY,
    name text NOT NULL UNIQUE,
    project_id text,
    created_at text,
    path text,
    response_body text
);