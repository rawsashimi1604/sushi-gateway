import{_ as e,c as a,a2 as n,o}from"./chunks/framework.CJ3Fe3Yg.js";const g=JSON.parse('{"title":"Configuration in Sushi Gateway","description":"","frontmatter":{},"headers":[],"relativePath":"concepts/configuration/index.md","filePath":"concepts/configuration/index.md"}'),t={name:"concepts/configuration/index.md"};function r(s,i,l,c,u,f){return o(),a("div",null,i[0]||(i[0]=[n('<h1 id="configuration-in-sushi-gateway" tabindex="-1">Configuration in Sushi Gateway <a class="header-anchor" href="#configuration-in-sushi-gateway" aria-label="Permalink to &quot;Configuration in Sushi Gateway&quot;">​</a></h1><p>Sushi Gateway supports two primary types of configurations: <strong>Declarative Configurations</strong> using <code>config.json</code> files and <strong>Environment Variable Configurations</strong>. Each type serves specific use cases, providing flexibility for different deployment models.</p><h2 id="declarative-configuration-config-json" tabindex="-1">Declarative Configuration (<code>config.json</code>) <a class="header-anchor" href="#declarative-configuration-config-json" aria-label="Permalink to &quot;Declarative Configuration (`config.json`)&quot;">​</a></h2><p>Declarative configurations allow you to define the gateway&#39;s behavior and settings in a JSON file. This approach is ideal for:</p><ul><li>Stateless deployments.</li><li>GitOps workflows where configuration is managed through version control.</li></ul><p>When using declarative configuration:</p><ul><li>All entities, such as services, routes, upstreams, and plugins, are defined in a single file.</li><li>Easy to replicate and manage across environments.</li></ul><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>To learn more about declarative configurations, visit the <strong><a href="./files.html">Declarative Configuration Guide</a></strong>.</p></div><h2 id="environment-variable-configuration" tabindex="-1">Environment Variable Configuration <a class="header-anchor" href="#environment-variable-configuration" aria-label="Permalink to &quot;Environment Variable Configuration&quot;">​</a></h2><p>Environment variable configurations allow you to dynamically adjust gateway behavior on startup. Simply specify it at docker runtime as an environment variable to utilize them.</p><ul><li>Define database connections, persistence modes, and other runtime options.</li></ul><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>For a detailed list of supported environment variables, see the <strong><a href="./../configuration/environment.html">Environment Variable Configuration Guide</a></strong>.</p></div>',12)]))}const p=e(t,[["render",r]]);export{g as __pageData,p as default};
