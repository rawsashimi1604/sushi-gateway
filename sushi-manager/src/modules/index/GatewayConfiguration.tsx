import { useState } from "react";
import DashboardCard from "../../components/layout/DashboardCard";
import Header from "../../components/typography/Header";
import JsonView from "react18-json-view";

const data = {
  global: {
    name: "sushi-gateway-1",
    plugins: [
      {
        name: "http_log",
        enabled: true,
        data: {
          http_endpoint: "http://localhost:3000/v1/log",
          method: "POST",
          content_type: "application/json",
        },
      },
      {
        name: "mtls",
        enabled: false,
        data: {},
      },
      {
        name: "key_auth",
        enabled: false,
        data: {
          key: "123456",
        },
      },
      {
        name: "rate_limit",
        enabled: true,
        data: {
          limit_second: 10,
          limit_min: 10,
          limit_hour: 10,
        },
      },
      {
        name: "basic_auth",
        enabled: true,
        data: {
          username: "admin",
          password: "changeme",
        },
      },
      {
        name: "jwt",
        enabled: false,
        data: {
          alg: "HS256",
          iss: "someIssuerKey",
          secret: "123secret456",
        },
      },
      {
        name: "acl",
        enabled: true,
        data: {
          whitelist: ["127.0.0.1", "127.0.0.2"],
          blacklist: ["192.168.0.1"],
        },
      },
      {
        name: "bot_protection",
        enabled: true,
        data: {
          blacklist: ["googlebot", "bingbot", "yahoobot"],
        },
      },
      {
        name: "request_size_limit",
        enabled: true,
        data: {
          max_payload_size: 10,
        },
      },
      {
        name: "cors",
        enabled: true,
        data: {
          allow_origins: ["*"],
          allow_methods: ["GET", "POST"],
          allow_headers: ["Authorization", "Content-Type"],
          expose_headers: ["Authorization"],
          allow_credentials: true,
          allow_private_network: false,
          preflight_continue: false,
          max_age: 3600,
        },
      },
    ],
  },
  services: [
    {
      name: "sushi-svc",
      base_path: "/sushi-service",
      protocol: "http",
      load_balancing_strategy: "round_robin",
      upstreams: [
        { host: "localhost", port: 8001 },
        { host: "localhost", port: 8002 },
        { host: "localhost", port: 8003 },
      ],
      plugins: [],
      routes: [
        {
          name: "get-sushi",
          path: "/v1/sushi",
          methods: ["GET"],
          plugins: [
            {
              name: "rate_limit",
              enabled: true,
              data: {
                limit_second: 10,
                limit_min: 10,
                limit_hour: 100,
              },
            },
          ],
        },
        {
          name: "get-sushi-by-id",
          path: "/v1/sushi/{id}",
          methods: ["GET"],
          plugins: [],
        },
        {
          name: "get-sushi-restaurants",
          path: "/v1/sushi/restaurant",
          methods: ["GET"],
          plugins: [],
        },
        {
          name: "sushi-provision-jwt",
          path: "/v1/token",
          methods: ["GET"],
          plugins: [],
        },
      ],
    },
  ],
};

function GatewayConfiguration() {
  const [showConfig, setShowConfig] = useState(false);

  return (
    <DashboardCard className="flex flex-col gap-2 p-6 ">
      <Header text="gateway configuration" align="left" size="sm" />
      <button
        type="button"
        onClick={() => setShowConfig((prev) => !prev)}
        className="w-[80px] mt-2 text-xs py-1.5 px-2 pb-2 tlext-white bg-blue-500 shadow-md rounded-lg font-sans tracking-wider border-0 duration-300 transition-all hover:-translate-y-1 hover:shadow-lg text-white"
      >
        <span>{showConfig ? "hide" : "show"}</span>
      </button>
      {/* TODO: popup modal */}
      {showConfig && (
        <div className="overflow-y-scroll">
          <JsonView src={data} />
        </div>
      )}
    </DashboardCard>
  );
}

export default GatewayConfiguration;
