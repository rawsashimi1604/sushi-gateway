import { defineConfig } from "vitepress";

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Sushi Gateway",
  description: "Lightweight Layer 7 Open Source Gateway",
  base: "/sushi-gateway/",
  themeConfig: {
    search: {
      provider: "local",
    },
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: "Home", link: "/" },
      { text: "Docs", link: "/docs-home" },
      {
        text: "Get Started",
        link: "/getting-started/docker",
      },
    ],

    sidebar: [
      {
        text: "Getting Started",
        collapsed: false,
        items: [
          { text: "Quick Start with Docker", link: "/getting-started/docker" },
          {
            text: "Install Sushi Gateway with Postgres",
            link: "/getting-started/postgres",
          },
        ],
      },
      {
        text: "Concepts",
        collapsed: false,
        items: [
          {
            text: "What is an API Gateway?",
            link: "/concepts/what-is-api-gateway",
          },
          {
            text: "Sushi Gateway Architecture",
            link: "/concepts/architecture",
          },
          {
            text: "Entities",
            link: "/concepts/entities/",
            items: [
              { text: "Service", link: "/concepts/entities/service" },
              { text: "Upstream", link: "/concepts/entities/upstream" },
              { text: "Route", link: "/concepts/entities/route" },
              { text: "Plugin", link: "/concepts/entities/plugin" },
            ],
          },
          { text: "Routing", link: "/concepts/routing" },
          { text: "Load Balancing", link: "/concepts/load-balancing" },
          {
            text: "Data Persistence Modes",
            link: "/concepts/data-persistence",
          },
          {
            text: "Configuration",
            link: "/concepts/configuration/",
            items: [
              {
                text: "Environment Variables",
                link: "/concepts/configuration/environment",
              },
              {
                text: "Declarative Configuration File",
                link: "/concepts/configuration/files",
              },
            ],
          },
        ],
      },
      {
        text: "Plugins",
        collapsed: true,
        link: "/plugins/",
        items: [
          {
            text: "Basic Authentication",
            link: "/plugins/basic-auth",
          },
          { text: "JSON Web Token (JWT)", link: "/plugins/jwt" },
          {
            text: "API Key Authentication",
            link: "/plugins/key-auth",
          },
          {
            text: "Mutual Transport Layer Security (MTLS)",
            link: "/plugins/mtls",
          },
          { text: "Bot Protection", link: "/plugins/bot-protection" },
          { text: "CORS", link: "/plugins/cors" },
          { text: "Access Control List", link: "/plugins/acl" },
          { text: "Rate Limiting", link: "/plugins/rate-limit" },
          {
            text: "Request Size Limit",
            link: "/plugins/request-size-limit",
          },
          { text: "HTTP Log", link: "/plugins/http-log" },
        ],
      },
      {
        text: "Admin REST API",
        collapsed: true,
        link: "/api/",
        items: [{ text: "Endpoints", link: "/api/endpoints" }],
      },
    ],

    socialLinks: [
      {
        icon: "github",
        link: "https://github.com/rawsashimi1604/sushi-gateway",
      },
    ],
  },
});
