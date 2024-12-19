<script setup>
import gsap from "gsap";
import { MotionPathPlugin } from "gsap/all";

gsap.registerPlugin(MotionPathPlugin);
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


    <div class="sushi-chip">
      <img src="/images/Logo.png" alt="Sushi Gateway Logo" class="sushi-chip__logo" />
    </div>
  </div>
</template>

<script>
export default {
  name: "HeroDiagram",
  mounted() {

    // Animate the dots along the paths
    const paths = document.querySelectorAll(".curved-lines path");
    paths.forEach((path, index) => {
      const dot = document.createElementNS("http://www.w3.org/2000/svg", "circle");
      dot.setAttribute("r", "3");
      dot.setAttribute("fill", "#c6caff");
      dot.setAttribute("cx", "0");
      dot.setAttribute("cy", "0");
      dot.classList.add("animated-dot");
      path.parentNode.appendChild(dot);

      gsap.to(dot, {
        motionPath: {
          path: path,
          align: path,
          alignOrigin: [0.5, 0.5],
        },
        duration: 3 + index * 0.5, // Vary the speed slightly
        repeat: -1, // Infinite repeat
        ease: "power2.inOut",
      });
    });

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
.curved-lines {
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  z-index: 1;
}

.hero__diagram {
  width: 100%;
  height: 400px;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
}

.sushi-chip {
  width: 134px;
  height: 134px;
  border-radius: 10px;
  margin-top: 180px;
  background: linear-gradient(135deg, #40a4ff, #ff57d9);
  /* Nice gradient background */
  display: flex;
  justify-content: center;
  align-items: center;
  box-shadow: 0 0 20px rgba(64, 115, 255, 0.6);
  z-index: 3;
  /* Simple glow effect */
  transition: box-shadow 0.3s ease-in-out, transform 0.3s ease-in-out;
}

.sushi-chip:hover {
  box-shadow: 0 0 30px rgba(64, 115, 255, 0.9);
  /* More intense glow on hover */
  transform: scale(1.05);
  /* Slight zoom-in on hover */
}

.sushi-chip__logo {
  width: 100px;
  height: 100px;
  object-fit: contain;
}
</style>
