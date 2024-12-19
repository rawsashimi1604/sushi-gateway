import{_ as i,c as a,a2 as t,o as e}from"./chunks/framework.B1nutZSU.js";const u=JSON.parse('{"title":"Basic Authentication Plugin","description":"","frontmatter":{},"headers":[],"relativePath":"plugins/basic-auth.md","filePath":"plugins/basic-auth.md"}'),n={name:"plugins/basic-auth.md"};function l(o,s,h,r,p,c){return e(),a("div",null,s[0]||(s[0]=[t(`<h1 id="basic-authentication-plugin" tabindex="-1">Basic Authentication Plugin <a class="header-anchor" href="#basic-authentication-plugin" aria-label="Permalink to &quot;Basic Authentication Plugin&quot;">​</a></h1><p>The Basic Authentication (<code>basic_auth</code>) plugin secures APIs by requiring users to provide a username and password. This plugin is compliant with <strong><a href="https://datatracker.ietf.org/doc/html/rfc7617" target="_blank" rel="noreferrer">RFC 7617</a></strong> and ensures that only authorized users can access your services.</p><h2 id="how-it-works" tabindex="-1">How It Works <a class="header-anchor" href="#how-it-works" aria-label="Permalink to &quot;How It Works&quot;">​</a></h2><p>The Basic Authentication plugin validates incoming requests by checking the provided credentials against a predefined username and password. Requests with invalid or missing credentials are rejected with a <strong>401 Unauthorized</strong> response.</p><h3 id="key-features" tabindex="-1">Key Features <a class="header-anchor" href="#key-features" aria-label="Permalink to &quot;Key Features&quot;">​</a></h3><ul><li>Simple and lightweight authentication method.</li><li>Configurable at global, service, or route levels.</li></ul><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>Learn how to integrate this plugin into your setup in the <strong><a href="./../plugins/">Plugins Overview</a></strong>.</p></div><h2 id="configuration-fields" tabindex="-1">Configuration Fields <a class="header-anchor" href="#configuration-fields" aria-label="Permalink to &quot;Configuration Fields&quot;">​</a></h2><table tabindex="0"><thead><tr><th>Field</th><th>Type</th><th>Description</th><th>Example Value</th></tr></thead><tbody><tr><td><code>username</code></td><td>String</td><td>The username required for authentication.</td><td><code>admin</code></td></tr><tr><td><code>password</code></td><td>String</td><td>The password associated with the username.</td><td><code>adminpass</code></td></tr></tbody></table><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>Ensure passwords are strong and securely stored to enhance security.</p></div><h2 id="example-configuration" tabindex="-1">Example Configuration <a class="header-anchor" href="#example-configuration" aria-label="Permalink to &quot;Example Configuration&quot;">​</a></h2><p>Below is an example of configuring the Basic Authentication plugin for a route:</p><div class="language-json vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">json</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">{</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">  &quot;name&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;basic_auth&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">  &quot;enabled&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">true</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">  &quot;config&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: {</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">    &quot;username&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;admin&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">    &quot;password&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;adminpass&quot;</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">  }</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">}</span></span></code></pre></div><h3 id="explanation" tabindex="-1">Explanation <a class="header-anchor" href="#explanation" aria-label="Permalink to &quot;Explanation&quot;">​</a></h3><ul><li><strong><code>username</code></strong>: The username required for accessing the route or service.</li><li><strong><code>password</code></strong>: The password associated with the username.</li></ul><h2 id="applying-the-plugin" tabindex="-1">Applying the Plugin <a class="header-anchor" href="#applying-the-plugin" aria-label="Permalink to &quot;Applying the Plugin&quot;">​</a></h2><p>The Basic Authentication plugin can be applied at various levels:</p><ol><li><strong>Global Level</strong>: Secures all incoming requests to the gateway.</li><li><strong>Service Level</strong>: Secures all routes within a specific service.</li><li><strong>Route Level</strong>: Secures individual routes for precise control.</li></ol><p>Example of applying the plugin to a specific route:</p><div class="language-json vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">json</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">{</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">  &quot;name&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;example-route&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">  &quot;path&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;/v1/sushi&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">  &quot;methods&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: [</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;GET&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">],</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">  &quot;plugins&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: [</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    {</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">      &quot;name&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;basic_auth&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">      &quot;enabled&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">true</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">      &quot;config&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: {</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">        &quot;username&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;admin&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">        &quot;password&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;adminpass&quot;</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">      }</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    }</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">  ]</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">}</span></span></code></pre></div><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>Use route-level Basic Authentication for APIs that require granular access control.</p></div><h2 id="connecting-to-the-api" tabindex="-1">Connecting to the API <a class="header-anchor" href="#connecting-to-the-api" aria-label="Permalink to &quot;Connecting to the API&quot;">​</a></h2><p>To authenticate requests using Basic Authentication:</p><ol><li>Base64 encode the <code>username:password</code> pair.</li><li>Include the encoded value in the <code>Authorization</code> header as:<div class="language-http vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">http</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#22863A;--shiki-dark:#85E89D;">Authorization</span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;"> Basic &lt;base64-encoded-credentials&gt;</span></span></code></pre></div></li></ol><h3 id="example" tabindex="-1">Example: <a class="header-anchor" href="#example" aria-label="Permalink to &quot;Example:&quot;">​</a></h3><p>If the username is <code>admin</code> and the password is <code>adminpass</code>, the header would look like:</p><div class="language-http vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">http</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#22863A;--shiki-dark:#85E89D;">Authorization</span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;"> Basic YWRtaW46YWRtaW5wYXNz</span></span></code></pre></div><div class="warning custom-block"><p class="custom-block-title">WARNING</p><p><strong>Basic Authentication is insufficient to protect production environments.</strong></p><p>Reasons:</p><ul><li><strong>Lack of Encryption</strong>: Credentials are transmitted in plaintext unless HTTPS is used, making them vulnerable to interception.</li><li><strong>No Token Revocation</strong>: Unlike modern authentication mechanisms (e.g., OAuth), credentials cannot be easily revoked or rotated.</li><li><strong>Prone to Brute Force Attacks</strong>: Without additional security measures like rate limiting, attackers can guess credentials.</li><li><strong>Static Credentials</strong>: Managing static username-password pairs at scale is cumbersome and insecure.</li></ul><p>Consider using advanced authentication methods like JWT for production-grade security.</p></div><h2 id="use-cases" tabindex="-1">Use Cases <a class="header-anchor" href="#use-cases" aria-label="Permalink to &quot;Use Cases&quot;">​</a></h2><ol><li><strong>Restrict Access</strong>: Secure endpoints by limiting access to authorized users.</li><li><strong>Protect Development APIs</strong>: Add a simple authentication layer for staging or development environments.</li></ol><h2 id="tips-for-using-the-basic-authentication-plugin" tabindex="-1">Tips for Using the Basic Authentication Plugin <a class="header-anchor" href="#tips-for-using-the-basic-authentication-plugin" aria-label="Permalink to &quot;Tips for Using the Basic Authentication Plugin&quot;">​</a></h2><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>Use HTTPS to encrypt traffic and protect credentials from being intercepted.</p></div><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>Combine Basic Authentication with plugins like Rate Limiting to prevent brute-force attacks.</p></div><div class="tip custom-block"><p class="custom-block-title">TIP</p><p>Store and manage credentials securely using environment variables or secret management tools.</p></div><p>For more plugins, visit the <strong><a href="./../plugins/">Plugins Overview</a></strong>.</p>`,35)]))}const k=i(n,[["render",l]]);export{u as __pageData,k as default};