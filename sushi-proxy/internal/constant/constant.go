package constant

// PROTOCOLS
const UTF_8 = "UTF-8"
const HS_256 = "HS256"

// PLUGINS
var AVAILABLE_PLUGINS = []string{
	PLUGIN_BASIC_AUTH,
	PLUGIN_ACL,
	PLUGIN_ANALYTICS,
	PLUGIN_BOT_PROTECTION,
	PLUGIN_RATE_LIMIT,
	PLUGIN_REQUEST_SIZE_LIMIT,
	PLUGIN_JWT,
	PLUGIN_KEY_AUTH,
}

const PLUGIN_BASIC_AUTH = "basic_auth"
const PLUGIN_ACL = "acl"
const PLUGIN_ANALYTICS = "analytics"

// TODO: add cors
const PLUGIN_CORS = "cors"
const PLUGIN_BOT_PROTECTION = "bot_protection"
const PLUGIN_RATE_LIMIT = "rate_limit"
const PLUGIN_REQUEST_SIZE_LIMIT = "request_size_limit"
const PLUGIN_JWT = "jwt"
const PLUGIN_KEY_AUTH = "key_auth"
