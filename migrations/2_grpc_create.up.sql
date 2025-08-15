create table grpcstubs (
    id serial PRIMARY KEY,
    name text NOT NULL UNIQUE,
    project_id text,
    created_at text,
    proto_file text,
    proto_method text,
    response_body text
);