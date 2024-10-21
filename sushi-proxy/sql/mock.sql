-- Insert mock data into the service table
INSERT INTO service (name, base_path, protocol, load_balancing_alg)
VALUES
    ('sushi-svc', '/sushi-service', 'http', 'round_robin');

-- Insert mock data into the upstream table
INSERT INTO upstream (id, service_name, host, port)
VALUES
    ('1', 'sushi-svc', 'localhost', 8001),
    ('2', 'sushi-svc', 'localhost', 8002),
    ('3', 'sushi-svc', 'localhost', 8003);

-- Insert mock data into the route table
INSERT INTO route (name, service_name, path)
VALUES
    ('get-sushi', 'sushi-svc', '/v1/sushi'),
    ('get-sushi-by-id', 'sushi-svc', '/v1/sushi/{id}'),
    ('get-sushi-restaurants', 'sushi-svc', '/v1/sushi/restaurant'),
    ('sushi-provision-jwt', 'sushi-svc', '/v1/token');

-- Insert mock data into the route_methods table
INSERT INTO route_methods (route_name, method)
VALUES
    ('get-sushi', 'GET'),
    ('get-sushi-by-id', 'GET'),
    ('get-sushi-restaurants', 'GET'),
    ('sushi-provision-jwt', 'GET');

-- Insert mock data into the plugin table
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

-- Insert mock data into the service_plugin table (associating plugins with the service)
INSERT INTO service_plugin (service_name, plugin_id)
VALUES
    ('sushi-svc', '4'), -- rate_limit plugin for sushi-svc
    ('sushi-svc', '5'); -- basic_auth plugin for sushi-svc

-- Insert mock data into the route_plugin table (associating plugins with specific routes)
INSERT INTO route_plugin (route_name, plugin_id)
VALUES
    ('get-sushi', '4'); -- rate_limit plugin for the get-sushi route