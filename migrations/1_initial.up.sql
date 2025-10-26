create table projects (
    id serial PRIMARY KEY,
    name text NOT NULL UNIQUE,
    created_at text
);

create table users (
    id serial PRIMARY KEY
);

create table role_models (
    id serial PRIMARY KEY,
    user_id int NOT NULL REFERENCES users(id),
    project_id int NOT NULL REFERENCES projects(id),
    role text NOT NULL
);

create table reststubs (
    id serial PRIMARY KEY,
    name text NOT NULL UNIQUE,
    project_id int REFERENCES projects(id),
    created_at text,
    path text,
    method text,
    type text,
    response_body text
);

create table grpc_protos (
    id serial PRIMARY KEY,
    service_name text NOT NULL UNIQUE,
    project_id int REFERENCES projects(id),
    created_at text,
    proto_file text
);

create table grpc_stubs (
    id serial PRIMARY KEY,
    name text NOT NULL UNIQUE,
    project_id int REFERENCES projects(id),
    created_at text,
    path text,
    response_body text
);
