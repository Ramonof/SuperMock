create table projects (
    id serial PRIMARY KEY,
    name text NOT NULL UNIQUE,
    created_at text
);

create table reststubs (
    id serial PRIMARY KEY,
    name text NOT NULL UNIQUE,
    project_id int REFERENCES projects(id),
    created_at text,
    path text,
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
