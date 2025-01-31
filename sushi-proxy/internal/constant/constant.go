package constant

// DB CONFIG
const DB_MODE = "db"
const DBLESS_MODE = "dbless"

// PORTS
const PORT_ADMIN_API = "8081"
const PORT_HTTP = "8008"
const PORT_HTTPS = "8443"

// PROTOCOLS
const UTF_8 = "UTF-8"
const HS_256 = "HS256"
const RSA_256 = "RS256"

var AVAILABLE_PROXY_PROTOCOLS = []string{"http", "https"}

// HTTP
var VALID_METHODS = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"}

// PLUGINS
var AVAILABLE_PLUGINS = []string{
	PLUGIN_BASIC_AUTH,
	PLUGIN_ACL,
	PLUGIN_BOT_PROTECTION,
	PLUGIN_RATE_LIMIT,
	PLUGIN_REQUEST_SIZE_LIMIT,
	PLUGIN_JWT,
	PLUGIN_KEY_AUTH,
	PLUGIN_MTLS,
	PLUGIN_HTTP_LOG,
	PLUGIN_CORS,
}

const PLUGIN_RESPONSE_HANDLER = "response_handler"
const PLUGIN_BASIC_AUTH = "basic_auth"
const PLUGIN_ACL = "acl"
const PLUGIN_CORS = "cors"
const PLUGIN_BOT_PROTECTION = "bot_protection"
const PLUGIN_RATE_LIMIT = "rate_limit"
const PLUGIN_REQUEST_SIZE_LIMIT = "request_size_limit"
const PLUGIN_JWT = "jwt"
const PLUGIN_KEY_AUTH = "key_auth"
const PLUGIN_MTLS = "mtls"
const PLUGIN_HTTP_LOG = "http_log"
