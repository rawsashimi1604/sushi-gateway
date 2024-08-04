import React from 'react'
import Header from '../../components/typography/Header'
import ReactJson from 'react-json-view'

function Json() {
    return (
        <div>
            <Header text="Json" align="left" size="md" />
            <div className='bg-neutral-200 px-4 py-4 rounded-lg shadow-sm w-[80%]'>
                <ReactJson src={{
                    "global": {
                        "name": "sushi-gateway-1",
                        "plugins": [
                            {
                                "name": "http_log",
                                "enabled": true,
                                "data": {
                                    "http_endpoint": "http://localhost:8003/v1/log",
                                    "method": "POST",
                                    "content_type": "application/json"
                                }
                            },
                            {
                                "name": "mtls",
                                "enabled": false,
                                "data": {
                                }
                            },
                            {
                                "name": "rate_limit",
                                "enabled": true,
                                "data": {
                                    "limit_second": 100,
                                    "limit_min": 100,
                                    "limit_hour": 100
                                }
                            },
                            {
                                "name": "basic_auth",
                                "enabled": false,
                                "data": {
                                    "username": "admin",
                                    "password": "changeme"
                                }
                            },
                            {
                                "name": "jwt",
                                "enabled": false,
                                "data": {
                                    "alg": "HS256",
                                    "iss": "someIssuerKey",
                                    "secret": "123secret456"
                                }
                            },
                            {
                                "name": "acl",
                                "enabled": true,
                                "data": {
                                    "whitelist": ["127.0.0.1", "127.0.0.2"],
                                    "blacklist": ["192.168.0.1"]
                                }
                            },
                            {
                                "name": "bot_protection",
                                "enabled": true,
                                "data": {
                                    "blacklist": ["googlebot", "bingbot", "yahoobot"]
                                }
                            },
                            {
                                "name": "request_size_limit",
                                "enabled": true,
                                "data": {
                                    "max_payload_size": 10
                                }
                            },
                            {
                                "name": "cors",
                                "enabled": true,
                                "data": {
                                    "allow_origins": ["*"],
                                    "allow_methods": ["GET", "POST"],
                                    "allow_headers": ["Authorization", "Content-Type"],
                                    "expose_headers": ["Authorization"],
                                    "allow_credentials": true,
                                    "allow_private_network": false,
                                    "preflight_continue": false,
                                    "max_age": 3600
                                }
                            }
                        ]
                    },
                    "services": [
                        {
                            "name": "sushi-svc",
                            "base_path": "/sushi-service",
                            "protocol": "http",
                            "upstreams": [
                                { "host": "localhost", "port": 8001 },
                                { "host": "localhost", "port": 8002 }
                            ],
                            "plugins": [],
                            "routes": [
                                {
                                    "name": "get-sushi",
                                    "path": "/v1/sushi",
                                    "methods": ["GET"],
                                    "plugins": [
                                        {
                                            "name": "rate_limit",
                                            "enabled": true,
                                            "data": {
                                                "limit_second": 10,
                                                "limit_min": 10,
                                                "limit_hour": 100
                                            }
                                        }
                                    ]
                                },
                                {
                                    "name": "get-sushi-restaurants",
                                    "path": "/v1/sushi/restaurant",
                                    "methods": ["GET"],
                                    "plugins": []
                                },
                                {
                                    "name": "sushi-provision-jwt",
                                    "path": "/v1/token",
                                    "methods": ["GET"],
                                    "plugins": []
                                }
                            ]
                        }
                    ]
                }} />
            </div>
        </div>
    )
}

export default Json