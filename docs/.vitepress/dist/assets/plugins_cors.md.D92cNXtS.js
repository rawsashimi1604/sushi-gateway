import{_ as i,c as t,a2 as e,o as a}from"./chunks/framework.CJ3Fe3Yg.js";const k=JSON.parse('{"title":"CORS Plugin","description":"","frontmatter":{},"headers":[],"relativePath":"plugins/cors.md","filePath":"plugins/cors.md"}'),n={name:"plugins/cors.md"};function o(l,s,r,h,p,d){return a(),t("div",null,s[0]||(s[0]=[e(`<h1 id="cors-plugin" tabindex="-1">CORS Plugin <a class="header-anchor" href="#cors-plugin" aria-label="Permalink to &quot;CORS Plugin&quot;">​</a></h1><p>The CORS (Cross-Origin Resource Sharing) (<code>cors</code>) plugin enables APIs to manage requests from different origins, ensuring compliance with <strong><a href="https://datatracker.ietf.org/doc/html/rfc6454" target="_blank" rel="noreferrer">RFC 6454</a></strong>. This is essential for enabling secure cross-origin requests, particularly in browser-based applications.</p><h2 id="how-it-works" tabindex="-1">How It Works <a class="header-anchor" href="#how-it-works" aria-label="Permalink to &quot;How It Works&quot;">​</a></h2><p>The CORS plugin inspects incoming requests for the <code>Origin</code> header and validates them against the configured rules. If the request meets the specified criteria, the plugin adds appropriate CORS headers to the response.</p><p>Requests that do not meet the criteria are either denied or processed without CORS headers, depending on the configuration.</p><h3 id="key-features" tabindex="-1">Key Features <a class="header-anchor" href="#key-features" aria-label="Permalink to &quot;Key Features&quot;">​</a></h3><ul><li>Supports granular configuration for origins, methods, headers, and credentials.</li><li>Helps APIs comply with modern web security standards.</li></ul><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>Learn how to integrate this plugin into your setup in the <strong><a href="./../plugins/">Plugins Overview</a></strong>.</p></div><h2 id="configuration-fields" tabindex="-1">Configuration Fields <a class="header-anchor" href="#configuration-fields" aria-label="Permalink to &quot;Configuration Fields&quot;">​</a></h2><table tabindex="0"><thead><tr><th>Field</th><th>Type</th><th>Description</th><th>Example Value</th></tr></thead><tbody><tr><td><code>allow_origins</code></td><td>Array</td><td>List of allowed origins. Use <code>&quot;*&quot;</code> to allow all origins.</td><td><code>[&quot;*&quot;]</code></td></tr><tr><td><code>allow_methods</code></td><td>Array</td><td>List of HTTP methods allowed for cross-origin requests.</td><td><code>[&quot;GET&quot;, &quot;POST&quot;]</code></td></tr><tr><td><code>allow_headers</code></td><td>Array</td><td>List of headers allowed in cross-origin requests.</td><td><code>[&quot;Authorization&quot;, &quot;Content-Type&quot;]</code></td></tr><tr><td><code>expose_headers</code></td><td>Array</td><td>List of headers exposed in the response.</td><td><code>[&quot;Authorization&quot;]</code></td></tr><tr><td><code>allow_credentials</code></td><td>Boolean</td><td>Whether credentials (cookies, authorization headers) are allowed.</td><td><code>true</code></td></tr><tr><td><code>allow_private_network</code></td><td>Boolean</td><td>Whether private network requests (e.g., local IPs) are allowed.</td><td><code>false</code></td></tr><tr><td><code>preflight_continue</code></td><td>Boolean</td><td>Whether the gateway should continue processing preflight requests.</td><td><code>false</code></td></tr><tr><td><code>max_age</code></td><td>Integer</td><td>Maximum time (in seconds) for which the preflight response is cached.</td><td><code>3600</code></td></tr></tbody></table><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>For maximum security, use specific origins instead of <code>&quot;*&quot;</code> in production environments.</p></div><h2 id="example-configuration" tabindex="-1">Example Configuration <a class="header-anchor" href="#example-configuration" aria-label="Permalink to &quot;Example Configuration&quot;">​</a></h2><p>Below is an example of configuring the CORS plugin:</p><div class="language-json vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">json</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">{</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">  &quot;name&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;cors&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">  &quot;enabled&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">true</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">  &quot;config&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: {</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">    &quot;allow_origins&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: [</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;*&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">],</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">    &quot;allow_methods&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: [</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;GET&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;POST&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">],</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">    &quot;allow_headers&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: [</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;Authorization&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;Content-Type&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">],</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">    &quot;expose_headers&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: [</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;Authorization&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">],</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">    &quot;allow_credentials&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">true</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">    &quot;allow_private_network&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">false</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">    &quot;preflight_continue&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">false</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">    &quot;max_age&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">3600</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">  }</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">}</span></span></code></pre></div><h3 id="explanation" tabindex="-1">Explanation <a class="header-anchor" href="#explanation" aria-label="Permalink to &quot;Explanation&quot;">​</a></h3><ul><li><strong><code>allow_origins</code></strong>: Allows all origins (<code>&quot;*&quot;</code>). Replace with specific origins for stricter control.</li><li><strong><code>allow_methods</code></strong>: Permits <code>GET</code> and <code>POST</code> requests from cross-origin sources.</li><li><strong><code>allow_headers</code></strong>: Enables <code>Authorization</code> and <code>Content-Type</code> headers in requests.</li><li><strong><code>expose_headers</code></strong>: Exposes <code>Authorization</code> in the response.</li><li><strong><code>allow_credentials</code></strong>: Allows cookies and authorization headers in requests.</li><li><strong><code>allow_private_network</code></strong>: Blocks private network requests for security.</li><li><strong><code>preflight_continue</code></strong>: Stops further processing of preflight requests.</li><li><strong><code>max_age</code></strong>: Sets a 1-hour cache duration for preflight responses.</li></ul><h2 id="applying-the-plugin" tabindex="-1">Applying the Plugin <a class="header-anchor" href="#applying-the-plugin" aria-label="Permalink to &quot;Applying the Plugin&quot;">​</a></h2><p>The CORS plugin can be applied at various levels:</p><ol><li><strong>Global Level</strong>: Applies CORS rules to all services and routes.</li><li><strong>Service Level</strong>: Applies CORS rules to all routes within a service.</li><li><strong>Route Level</strong>: Customizes CORS rules for specific routes.</li></ol><p>Example of applying the plugin globally:</p><div class="language-json vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">json</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">{</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">  &quot;name&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;cors&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">  &quot;enabled&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">true</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">  &quot;config&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: {</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">    &quot;allow_origins&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: [</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;https://example.com&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">],</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">    &quot;allow_methods&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: [</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;GET&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">],</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">    &quot;allow_headers&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: [</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;Content-Type&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">  }</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">}</span></span></code></pre></div><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>Use route-level configurations for APIs with varying CORS requirements.</p></div><div class="warning custom-block"><p class="custom-block-title">WARNING</p><p>Misconfigured CORS policies can expose your API to vulnerabilities. Ensure <code>allow_origins</code> is restrictive in production environments.</p></div><h2 id="use-cases" tabindex="-1">Use Cases <a class="header-anchor" href="#use-cases" aria-label="Permalink to &quot;Use Cases&quot;">​</a></h2><ol><li><strong>Enable Cross-Origin Requests</strong>: Allow secure communication between frontend applications and APIs.</li><li><strong>Control Sensitive Endpoints</strong>: Restrict origins, methods, and headers for specific endpoints.</li><li><strong>Optimize Performance</strong>: Cache preflight responses to reduce latency.</li></ol><h2 id="tips-for-using-the-cors-plugin" tabindex="-1">Tips for Using the CORS Plugin <a class="header-anchor" href="#tips-for-using-the-cors-plugin" aria-label="Permalink to &quot;Tips for Using the CORS Plugin&quot;">​</a></h2><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>Combine the CORS plugin with authentication plugins like JWT to secure cross-origin requests.</p></div><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>Regularly audit allowed origins and headers to align with your security policies.</p></div><p>For more plugins, visit the <strong><a href="./../plugins/">Plugins Overview</a></strong>.</p>`,29)]))}const u=i(n,[["render",o]]);export{k as __pageData,u as default};
