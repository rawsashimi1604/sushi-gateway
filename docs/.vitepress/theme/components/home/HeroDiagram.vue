<script setup>
import gsap from "gsap";
</script>

<template>
  <div class="hero__diagram">
    <!-- SVG for curved lines -->
    <svg class="curved-lines" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 800 400">

      <defs>
        <!-- Define Gradient for the Lines -->
        <linearGradient id="line-gradient" x1="0%" y1="0%" x2="100%" y2="0%">
          <stop offset="0%" stop-color="#c6caff" stop-opacity="0.1" />
          <stop offset="30%" stop-color="#c6caff" stop-opacity="0.5" />
          <stop offset="50%" stop-color="#ffffff" stop-opacity="0.9" />
          <stop offset="70%" stop-color="#c6caff" stop-opacity="0.5" />
          <stop offset="100%" stop-color="#c6caff" stop-opacity="0.1" />
        </linearGradient>
      </defs>

      <!-- Incoming Lines -->
      <path d="M0,80 C200,80 300,180 370,190" stroke="url(#line-gradient)" fill="transparent" stroke-width="1" />
      <path d="M0,130 C150,130 300,200 370,200" stroke="url(#line-gradient)" fill="transparent" stroke-width="1" />
      <path d="M0,200 C150,210 300,210 370,210" stroke="url(#line-gradient)" fill="transparent" stroke-width="1" />
      <path d="M0,280 C240,270 300,240 370,220" stroke="url(#line-gradient)" fill="transparent" stroke-width="1" />
      <path d="M0,340 C240,310 300,260 370,230" stroke="url(#line-gradient)" fill="transparent" stroke-width="1" />

      <!-- Outgoing Lines -->
      <path d="M420,190 C750,150 850,10 1200,200" stroke="url(#line-gradient)" fill="transparent" stroke-width="1" />
      <path d="M420,210 C750,200 850,200 1200,200" stroke="url(#line-gradient)" fill="transparent" stroke-width="1" />
      <path d="M420,230 C750,250 850,360 1200,200" stroke="url(#line-gradient)" fill="transparent" stroke-width="1" />

    </svg>

    <!-- Sushi Chip -->
    <div class="sushi-chip">
      <img src="/images/Logo.png" alt="Sushi Gateway Logo" />
    </div>
  </div>
</template>

<script>
export default {
  name: "HeroDiagram",
  mounted() {
    // Animate the paths using GSAP
    gsap.from(".curved-lines path", {
      duration: 1.5,
      ease: "power2.out",
    });
  },
  props: {
    text: {
      type: String,
      required: true,
      default: "Default text",
    },
  },
};
</script>

<style scoped>
.hero__diagram {
  width: 100%;
  height: 400px;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
}

.curved-lines {
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
}

.sushi-chip {
  width: 134px;
  height: 134px;
  border-radius: 10px;
  margin-top: 200px;
  overflow: hidden;
  transition: all 0.6s ease-out;
  transform: translate3d(0, 0, 0) scale(0.85);
  background: white;
}

.sushi-chip img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}
</style>
