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
---

<script setup>
  import HeroSection from ".vitepress/theme/components/landing/HeroSection.vue"
  import HeroDiagram from ".vitepress/theme/components/landing/HeroDiagram.vue"

  import Features from ".vitepress/theme/components/landing/Features.vue"
  import OpenSource from ".vitepress/theme/components/landing/OpenSource.vue"

  import StartBuilding from ".vitepress/theme/components/landing/StartBuilding.vue"
</script>

<div>
  <HeroSection/>
  <HeroDiagram/>

  <Features/>
  <OpenSource />

  <StartBuilding />
</div>
