-- Insert the default service (sushi-svc)
INSERT INTO service (id, name, base_path, protocol, load_balancing_alg)
VALUES
    ('1', 'sushi-svc', '/sushi-service', 'http', 'round_robin');

-- Insert the upstreams for sushi-svc
INSERT INTO upstream (id, service_id, host, port)
VALUES
    ('1', '1', 'localhost', 8001),
    ('2', '1', 'localhost', 8002),
    ('3', '1', 'localhost', 8003);

-- Insert the plugins for the global scope with enabled status
INSERT INTO plugin (id, name, config, enabled)
VALUES
    ('1', 'http_log', '{"http_endpoint": "http://localhost:3000/v1/log", "method": "POST", "content_type": "application/json"}', true),
    ('2', 'mtls', '{}', false),
    ('3', 'key_auth', '{"key": "123456"}', false),
    ('4', 'rate_limit', '{"limit_second": 10, "limit_min": 10, "limit_hour": 10}', true),
    ('5', 'basic_auth', '{"username": "admin", "password": "changeme"}', true),
    ('6', 'jwt', '{"alg": "HS256", "iss": "someIssuerKey", "secret": "123secret456"}', false),
    ('7', 'acl', '{"whitelist": ["127.0.0.1", "127.0.0.2"], "blacklist": ["192.168.0.1"]}', true),
    ('8', 'bot_protection', '{"blacklist": ["googlebot", "bingbot", "yahoobot"]}', true),
    ('9', 'request_size_limit', '{"max_payload_size": 10}', true),
    ('10', 'cors', '{"allow_origins": ["*"], "allow_methods": ["GET", "POST"], "allow_headers": ["Authorization", "Content-Type"], "expose_headers": ["Authorization"], "allow_credentials": true, "allow_private_network": false, "preflight_continue": false, "max_age": 3600}', true);

-- Global plugins to service_plugin (assuming plugins applied globally to service)
INSERT INTO service_plugin (service_id, plugin_id)
VALUES
    ('1', '4');  -- Example, only rate_limit is enabled

-- Insert routes for the service
INSERT INTO route (id, service_id, path)
VALUES
    ('1', '1', '/v1/sushi'),
    ('2', '1', '/v1/sushi/{id}'),
    ('3', '1', '/v1/sushi/restaurant'),
    ('4', '1', '/v1/token');

-- Insert route methods for each route
INSERT INTO route_methods (route_id, method)
VALUES
    ('1', 'GET'),   -- for /v1/sushi
    ('2', 'GET'),   -- for /v1/sushi/{id}
    ('3', 'GET'),   -- for /v1/sushi/restaurant
    ('4', 'GET');   -- for /v1/token

-- Insert the plugins for each route with enabled status
INSERT INTO plugin (id, name, config, enabled)
VALUES
    ('11', 'rate_limit', '{"limit_second": 10, "limit_min": 10, "limit_hour": 100}', true);

-- Insert route_plugin mappings (plugins specific to routes)
INSERT INTO route_plugin (route_id, plugin_id)
VALUES
    ('1', '11');  -- rate_limit for /v1/sushi
