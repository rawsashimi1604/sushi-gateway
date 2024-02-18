\c sushi-db;

CREATE TABLE apis (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    name TEXT,
    host TEXT,
    port BIGINT,
    tags TEXT[],
    enabled BOOLEAN,
    health_check_enabled BOOLEAN,
    health TEXT
);

CREATE TABLE routes (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    name TEXT,
    service_id UUID REFERENCES apis(id),
    methods TEXT[],
    tags TEXT[]
);


