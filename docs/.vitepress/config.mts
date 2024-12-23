import { defineConfig } from "vitepress";

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Sushi Gateway",
  description: "Lightweight Layer 7 Open Source Gateway",
  base: "/sushi-gateway/",
  head: [["link", { rel: "icon", href: "/sushi-gateway/favicon.ico" }]],
  themeConfig: {
    search: {
      provider: "local",
    },
    logo: "/images/Logo_gradient.png",
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: "Home", link: "/" },
      { text: "Docs", link: "/docs-home" },
      {
        text: "Get Started",
        items: [
          { text: "Docker", link: "/getting-started/docker" },
          { text: "Postgres", link: "/getting-started/postgres" },
        ],
      },
      {
        text: "Plugins",
        items: [
          { text: "Overview", link: "/plugins/" },
          { text: "Basic Auth", link: "/plugins/basic-auth" },
          { text: "JWT", link: "/plugins/jwt" },
          { text: "Key Auth", link: "/plugins/key-auth" },
          { text: "MTLS", link: "/plugins/mtls" },
          { text: "Bot Protection", link: "/plugins/bot-protection" },
          { text: "CORS", link: "/plugins/cors" },
          { text: "ACL", link: "/plugins/acl" },
          { text: "Rate Limit", link: "/plugins/rate-limit" },
          { text: "Request Size Limit", link: "/plugins/request-size-limit" },
          { text: "HTTP Log", link: "/plugins/http-log" },
        ],
      },
    ],
    footer: {
      message: `Released under the MIT License.`,
      copyright: "Copyright © 2024-present Sushi Gateway.",
    },
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
      {
        icon: "discord",
        link: "https://discord.gg/aPv4QhQ6",
      },
      {
        icon: "docker",
        link: "https://hub.docker.com/repository/docker/rawsashimi/sushi-proxy",
      },
    ],
  },
});
