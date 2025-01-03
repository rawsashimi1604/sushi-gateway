DROP TABLE IF EXISTS gateway;
DROP TABLE IF EXISTS plugin_scope;
DROP TABLE IF EXISTS service_plugin;
DROP TABLE IF EXISTS route_plugin;
DROP TABLE IF EXISTS plugin;
DROP TABLE IF EXISTS route_methods;
DROP TABLE IF EXISTS route;
DROP TABLE IF EXISTS upstream;
DROP TABLE IF EXISTS service;

CREATE TABLE gateway (
    name TEXT PRIMARY KEY
);

CREATE TABLE service (
    name TEXT PRIMARY KEY,
    base_path TEXT NOT NULL,
    protocol TEXT NOT NULL,
    load_balancing_alg TEXT NOT NULL
);

CREATE TABLE upstream (
    id TEXT PRIMARY KEY,
    service_name TEXT REFERENCES service(name) ON DELETE CASCADE,
    host TEXT NOT NULL,
    port INTEGER,
    CONSTRAINT unique_service_host UNIQUE(service_name, host, port)
);

CREATE TABLE route (
   name TEXT PRIMARY KEY,
   service_name TEXT REFERENCES service(name) ON DELETE CASCADE,
   path TEXT NOT NULL
);

CREATE TABLE route_methods (
   route_name TEXT REFERENCES route(name) ON DELETE CASCADE,
   method TEXT NOT NULL,
   PRIMARY KEY (route_name, method)
);

CREATE TABLE plugin (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    config JSON NOT NULL,
    enabled BOOLEAN NOT NULL,
    scope TEXT CHECK (scope IN ('global', 'service', 'route')) NOT NULL
);

CREATE TABLE service_plugin (
    service_name TEXT REFERENCES service(name) ON DELETE CASCADE,
    plugin_id TEXT REFERENCES plugin(id) ON DELETE CASCADE,
    PRIMARY KEY (service_name, plugin_id)
);

CREATE TABLE route_plugin (
    route_name TEXT REFERENCES route(name) ON DELETE CASCADE,
    plugin_id TEXT REFERENCES plugin(id) ON DELETE CASCADE,
    PRIMARY KEY (route_name, plugin_id)
);

