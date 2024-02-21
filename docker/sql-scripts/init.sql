\c sushi-db;

CREATE TABLE services (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    name TEXT UNIQUE,
    host TEXT,
    port BIGINT,
    protocol TEXT,
    tags TEXT[],
    enabled BOOLEAN DEFAULT true,
    health_check_enabled BOOLEAN,
    health BOOLEAN
);

CREATE TABLE routes (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    name TEXT,
    path TEXT,
    service_id UUID REFERENCES services(id),
    methods TEXT[],
    tags TEXT[]
);

DO $$
DECLARE
   service_1_uuid uuid = gen_random_uuid();
   service_2_uuid uuid = gen_random_uuid();
BEGIN
INSERT INTO services(id, created_at, updated_at, name, host, port, protocol) VALUES
    (service_1_uuid, current_timestamp, current_timestamp, 'sushi-test-service-1', 'localhost', 8001, 'http'),
    (service_2_uuid, current_timestamp, current_timestamp, 'sushi-test-service-2', 'localhost', 8002, 'http');

INSERT INTO routes(id, created_at, updated_at, name, path, service_id, methods) VALUES
    (gen_random_uuid(), current_timestamp, current_timestamp, 'List Sushi', '/v1/sushi', service_1_uuid, ARRAY['GET']),
    (gen_random_uuid(), current_timestamp, current_timestamp, 'List Sushi', '/v1/sushi', service_2_uuid, ARRAY['GET']),
    (gen_random_uuid(), current_timestamp, current_timestamp, 'List Sushi Restaurants', '/v1/sushi/restaurant', service_1_uuid, ARRAY['GET']),
    (gen_random_uuid(), current_timestamp, current_timestamp, 'List Sushi Restaurants', '/v1/sushi/restaurant', service_2_uuid, ARRAY['GET']);
END $$



