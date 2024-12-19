---
pageClass: landing dark

layout: home
aside: false
editLink: false
markdownStyles: false
# hero:
#   name: "Sushi Gateway"
#   tagline: Layer 7 Lightweight Modular Open Source API Gateway
#   actions:
#     - theme: brand
#       text: Quick Start
#       link: /getting-started/docker
#     - theme: alt
#       text: View Documentation
#       link: /docs-home

# features:
#   - icon: "🧩"
#     title: "Modular Plugins"
#     details: "Easily customize and extend functionality with configurable plugins."
#   - icon: "🔒"
#     title: "Secure & Reliable"
#     details: "Built with modern security protocols and reliable architecture."
#   - icon: "⚖️"
#     title: "Scalable API Management"
#     details: "Handle high-traffic APIs with dynamic routing and load balancing."
---

<script setup>
  import HeroSection from ".vitepress/theme/components/home/HeroSection.vue"
  import HeroDiagram from ".vitepress/theme/components/home/HeroDiagram.vue"
</script>

<div class="landing__wrapper">
  <HeroSection/>
  <HeroDiagram/>
</div>

<style scoped>

.landing__wrapper {
  position: relative;
  overflow: visible; 
}
</style>
