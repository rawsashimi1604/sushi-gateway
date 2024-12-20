import{_ as i,c as e,a2 as a,o as s}from"./chunks/framework.B1nutZSU.js";const c=JSON.parse('{"title":"Plugins in Sushi Gateway","description":"","frontmatter":{},"headers":[],"relativePath":"plugins/index.md","filePath":"plugins/index.md"}'),n={name:"plugins/index.md"};function l(r,t,o,d,h,p){return s(),e("div",null,t[0]||(t[0]=[a(`<h1 id="plugins-in-sushi-gateway" tabindex="-1">Plugins in Sushi Gateway <a class="header-anchor" href="#plugins-in-sushi-gateway" aria-label="Permalink to &quot;Plugins in Sushi Gateway&quot;">​</a></h1><p>Plugins are modular extensions that enhance the gateway&#39;s functionality. They can be used for tasks such as authentication, rate limiting, transformations, and more. Each plugin operates within a middleware chain, allowing precise control over how requests and responses are processed.</p><h2 id="what-are-plugins" tabindex="-1">What are Plugins? <a class="header-anchor" href="#what-are-plugins" aria-label="Permalink to &quot;What are Plugins?&quot;">​</a></h2><p>Plugins are:</p><ul><li>Reusable components that add features to services and routes.</li><li>Configurable to meet specific API requirements.</li><li>Applied at different scopes (global, service, route) for fine-grained control.</li></ul><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>Learn about plugin fields and configurations in the <strong><a href="./../concepts/entities/plugin.html">Plugin Entity Documentation</a></strong>.</p></div><h2 id="plugin-middleware-chain" tabindex="-1">Plugin Middleware Chain <a class="header-anchor" href="#plugin-middleware-chain" aria-label="Permalink to &quot;Plugin Middleware Chain&quot;">​</a></h2><p>Plugins in Sushi Gateway operate in a defined middleware chain:</p><ol><li><strong>Global Plugins</strong>: Applied to all services and routes.</li><li><strong>Service-Level Plugins</strong>: Applied to all routes within a specific service.</li><li><strong>Route-Level Plugins</strong>: Applied to individual routes, overriding service and global plugins if applicable.</li></ol><h3 id="plugin-priority" tabindex="-1">Plugin Priority <a class="header-anchor" href="#plugin-priority" aria-label="Permalink to &quot;Plugin Priority&quot;">​</a></h3><p>The table below illustrates the priority of specific plugins in Sushi Gateway. Plugins with higher priority values are executed earlier in the middleware chain.</p><table tabindex="0"><thead><tr><th>Priority</th><th>Plugin</th></tr></thead><tbody><tr><td>2500</td><td>Bot Protection</td></tr><tr><td>2000</td><td>Cross Origin Resource Sharing (RFC 6454)</td></tr><tr><td>1600</td><td>Mutual Transport Layer Security (RFC 8705)</td></tr><tr><td>1450</td><td>JSON Web Token (RFC 7519)</td></tr><tr><td>1250</td><td>API Key Authentication</td></tr><tr><td>1100</td><td>Basic Authentication (RFC 7617)</td></tr><tr><td>951</td><td>Request Size Limit</td></tr><tr><td>950</td><td>Access Control List</td></tr><tr><td>910</td><td>Rate Limit</td></tr><tr><td>12</td><td>HTTP Log</td></tr></tbody></table><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>Use route-level plugins for the highest level of specificity and ensure priority alignment with your gateway logic.</p></div><h2 id="available-plugins" tabindex="-1">Available Plugins <a class="header-anchor" href="#available-plugins" aria-label="Permalink to &quot;Available Plugins&quot;">​</a></h2><p>Sushi Gateway supports several plugins. Currently, there are <strong>10 plugins</strong> available. The table below provides an overview:</p><table tabindex="0"><thead><tr><th>Plugin Name</th><th>Description</th><th>Documentation</th></tr></thead><tbody><tr><td><code>bot_protection</code></td><td>Protects against automated bots.</td><td><a href="./../plugins/bot-protection.html">Bot Protection Plugin</a></td></tr><tr><td><code>cors</code></td><td>Manages CORS policies for APIs.</td><td><a href="./../plugins/cors.html">CORS Plugin</a></td></tr><tr><td><code>mtls</code></td><td>Implements mutual TLS authentication.</td><td><a href="./../plugins/mtls.html">mTLS Plugin</a></td></tr><tr><td><code>jwt</code></td><td>Validates JSON Web Tokens (JWT).</td><td><a href="./../plugins/jwt.html">JWT Plugin</a></td></tr><tr><td><code>key_auth</code></td><td>Secures APIs using API Key Authentication.</td><td><a href="./../plugins/key-auth.html">API Key Plugin</a></td></tr><tr><td><code>basic_auth</code></td><td>Secures routes with basic authentication.</td><td><a href="./../plugins/basic-auth.html">Basic Auth Plugin</a></td></tr><tr><td><code>request_size_limit</code></td><td>Limits the size of incoming requests.</td><td><a href="./../plugins/request-size-limit.html">Request Size Limit Plugin</a></td></tr><tr><td><code>acl</code></td><td>Manages access control lists for API consumers.</td><td><a href="./../plugins/acl.html">Access Control List Plugin</a></td></tr><tr><td><code>rate_limit</code></td><td>Controls request rates for clients.</td><td><a href="./../plugins/rate-limit.html">Rate Limiting Plugin</a></td></tr><tr><td><code>http_log</code></td><td>Logs HTTP requests and responses for monitoring purposes.</td><td><a href="./../plugins/http-log.html">HTTP Log Plugin</a></td></tr></tbody></table><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>Click on a plugin name to learn more about its configuration and use cases.</p></div><h2 id="example-plugin-configuration" tabindex="-1">Example Plugin Configuration <a class="header-anchor" href="#example-plugin-configuration" aria-label="Permalink to &quot;Example Plugin Configuration&quot;">​</a></h2><p>Here’s how to configure a <code>rate_limit</code> plugin:</p><div class="language-json vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">json</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">{</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">  &quot;name&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;rate_limit&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">  &quot;enabled&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">true</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">  &quot;config&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: {</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">    &quot;limit_second&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">10</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">    &quot;limit_min&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">100</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">    &quot;limit_hour&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">1000</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">  }</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">}</span></span></code></pre></div><h3 id="explanation" tabindex="-1">Explanation <a class="header-anchor" href="#explanation" aria-label="Permalink to &quot;Explanation&quot;">​</a></h3><ul><li><strong><code>name</code></strong>: The plugin type (e.g., <code>rate_limit</code>).</li><li><strong><code>enabled</code></strong>: Toggles the plugin on or off.</li><li><strong><code>config</code></strong>: Plugin-specific settings.</li></ul><h2 id="tips-for-using-plugins" tabindex="-1">Tips for Using Plugins <a class="header-anchor" href="#tips-for-using-plugins" aria-label="Permalink to &quot;Tips for Using Plugins&quot;">​</a></h2><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>Combine multiple plugins at the route level to customize behavior for specific APIs.</p></div>`,24)]))}const g=i(n,[["render",l]]);export{c as __pageData,g as default};