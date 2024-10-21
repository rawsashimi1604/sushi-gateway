DROP TABLE IF EXISTS service_plugin;
DROP TABLE IF EXISTS route_plugin;
DROP TABLE IF EXISTS plugin;
DROP TABLE IF EXISTS route_methods;
DROP TABLE IF EXISTS route;
DROP TABLE IF EXISTS upstream;
DROP TABLE IF EXISTS service;

CREATE TABLE service (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    base_path TEXT NOT NULL,
    protocol TEXT NOT NULL,
    load_balancing_alg TEXT NOT NULL
);

CREATE TABLE upstream (
    id TEXT PRIMARY KEY,
    service_id TEXT REFERENCES service(id),
    host TEXT NOT NULL,
    port INTEGER NOT NULL,
    CONSTRAINT unique_service_host UNIQUE(service_id, host, port)
);

CREATE TABLE route (
   id TEXT PRIMARY KEY,
   service_id TEXT REFERENCES service(id),
   path TEXT NOT NULL
);

CREATE TABLE route_methods (
   route_id TEXT REFERENCES route(id),
   method TEXT NOT NULL,
   PRIMARY KEY (route_id, method)
);

CREATE TABLE plugin (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    config JSON NOT NULL,
    enabled BOOLEAN NOT NULL
);

CREATE TABLE service_plugin (
    service_id TEXT REFERENCES service(id),
    plugin_id TEXT REFERENCES plugin(id),
    PRIMARY KEY (service_id, plugin_id)
);

CREATE TABLE route_plugin (
    route_id TEXT REFERENCES route(id),
    plugin_id TEXT REFERENCES plugin(id),
    PRIMARY KEY (route_id, plugin_id)
);

