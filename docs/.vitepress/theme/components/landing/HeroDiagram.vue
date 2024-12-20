<script setup type="module">
import { gsap } from "gsap/dist/gsap";
import MotionPathPlugin from "gsap/dist/MotionPathPlugin";
gsap.registerPlugin(MotionPathPlugin);
</script>

<template>
  <div class="hero__bgwrapper">
    <div class="hero__diagram container">
      <!-- Animated Text with Spinner -->
      <div class="api-processing">
        <div class="spinner"></div>
        <span class="processing-text">API request processing...</span>
      </div>

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
  </div>
</template>

<script>
export default {
  name: "HeroDiagram",
  mounted() {
    // Animate fade in and out for the API processing text
    gsap.fromTo(".processing-text", {
      opacity: 1,
      duration: 2,
      repeat: -1,
      ease: "linear",
    }, {
      opacity: 0.25,
      duration: 2,
      repeat: -1,
      ease: "linear",
    });

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
};
</script>

<style scoped>
.hero__bgwrapper {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: visible;
  min-width: 768px;
  margin: 0 auto;
}

.curved-lines {
  position: absolute;
  left: 0;
  top: 0;
  z-index: 1;
  overflow: hidden;
  width: 100%;

  @media (max-width: 768px) {
    height: 400px;
  }
}

.hero__diagram {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  margin-top: -100px;
  height: 500px;
  display: none;


  @media (min-width: 768px) {
    display: flex;
  }

  @media (min-width: 1280px) {
    height: 700px;
  }

}

.sushi-chip {
  width: 80px;
  height: 80px;
  border-radius: 10px;
  position: absolute;
  background: linear-gradient(135deg, #2a78d4, #3fd9f0);


  box-shadow: 0 0 20px rgba(63, 217, 240, 0.6), 0 0 10px rgba(42, 120, 212, 0.5);
  z-index: 3;
  transition: box-shadow 0.3s ease-in-out, transform 0.3s ease-in-out;

  display: flex;
  justify-content: center;
  align-items: center;
  left: 270px;
  top: 120px;

  @media (min-width: 768px) {
    display: flex;
    width: 100px;
    height: 100px;
    left: 335px;
    top: 150px;
  }

  @media (min-width: 1024px) {
    display: flex;
    width: 105px;
    height: 105px;
    left: 460px;
    top: 220px;
  }

  @media (min-width: 1280px) {
    display: flex;
    width: 120px;
    height: 120px;
    left: 580px;
    top: 280px;
  }

  @media (min-width: 1536px) {
    display: flex;
    width: 134px;
    height: 134px;
    left: 690px;
    top: 330px;
  }

}

.sushi-chip:hover {
  box-shadow: 0 0 30px rgba(63, 217, 240, 0.8), 0 0 15px rgba(42, 120, 212, 0.7);
  /* Slightly darker glow on hover */
  transform: scale(1.1);
  /* Slight zoom-in on hover */
}

.sushi-chip__logo {
  width: 110px;
  height: 110px;
  object-fit: contain;
}

.api-processing {
  display: none;
  position: absolute;
  top: 280px;
  left: 645px;
  /* Adjust position above the sushi-chip */
  align-items: center;
  justify-content: center;
  font-family: Arial, sans-serif;
  font-size: 16px;
  font-weight: bold;
  text-shadow: 0 0 3px rgba(255, 255, 255, 0.8);
  color: #ffffff;
  z-index: 4;

  @media (min-width: 1280px) {
    display: flex;
    top: 230px;
    left: 535px;
  }

  @media (min-width: 1536px) {
    display: flex;
    top: 280px;
    left: 645px;
  }
}

.spinner {
  width: 16px;
  height: 16px;
  border: 3px solid rgba(255, 255, 255, 0.6);
  border-top-color: #ffffff;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-right: 8px;
  /* Space between spinner and text */
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}
</style>
