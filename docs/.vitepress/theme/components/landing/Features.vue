<script setup>
</script>

<template>
    <div class="hero__bgwrapper container">

        <div class="header-circle"></div>

        <!-- Header Text -->
        <h2>Redefining API proxying</h2>

        <!-- Feature Grid -->
        <div class="feature-grid">
            <div class="feature-card">
                <div class="feature-icon">🧩</div>
                <h3 class="feature-title">Modular Plugins</h3>
                <p class="feature-details">Easily customize and extend functionality with configurable plugins.</p>
            </div>
            <div class="feature-card">
                <div class="feature-icon">🔒</div>
                <h3 class="feature-title">Secure & Reliable</h3>
                <p class="feature-details">Built with modern security protocols and reliable architecture.</p>
            </div>
            <div class="feature-card">
                <div class="feature-icon">⚖️</div>
                <h3 class="feature-title">Scalable API Management</h3>
                <p class="feature-details">Handle high-traffic APIs with dynamic routing and load balancing.</p>
            </div>
            <div class="feature-card">
                <div class="feature-icon">⚖️</div>
                <h3 class="feature-title">Scalable API Management</h3>
                <p class="feature-details">Handle high-traffic APIs with dynamic routing and load balancing.</p>
            </div>
        </div>
    </div>
</template>

<script>
</script>

<style scoped>
.header-circle {
    width: 20px;
    height: 20px;
    background: linear-gradient(135deg, #2a78d4, #3fd9f0);
    border-radius: 100%;
    margin: 0 auto;
    margin-bottom: 40px;
    box-shadow: 0 0 10px rgba(63, 217, 240, 0.7),
        0 0 20px rgba(63, 217, 240, 0.5),
        0 0 30px rgba(63, 217, 240, 0.3);
    transition: box-shadow 0.3s ease-in-out;
}

.header-circle:hover {
    box-shadow: 0 0 20px rgba(63, 217, 240, 0.9),
        0 0 40px rgba(63, 217, 240, 0.7),
        0 0 60px rgba(63, 217, 240, 0.5);
}


.hero__bgwrapper {
    width: 100%;
    position: relative;
    padding-top: px;
    text-align: center;
}


h2 {
    position: relative;
    display: block;
    z-index: 2;
    font-family: var(--font-primary);
    font-size: 32px;
    font-style: normal;
    font-weight: 600;
    line-height: 120%;
    letter-spacing: -0.64px;
    cursor: default;
    text-wrap: balance;
    padding: 0 20px;

    @media (min-width: 768px) {
        font-size: 44px;
        letter-spacing: -0.88px;
    }
}

.feature-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 20px;
    margin-top: 40px;
}

.feature-card {
    background: #1c1c1c;
    padding: 40px;
    border-radius: 10px;
    text-align: center;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
    transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.feature-card:hover {
    transform: translateY(-10px);
    box-shadow: 0 6px 15px rgba(0, 0, 0, 0.4);
}

.feature-icon {
    font-size: 36px;
    margin-bottom: 15px;
}

.feature-title {
    font-size: 20px;
    font-weight: 600;
    margin-bottom: 10px;
    color: #ffffff;
}

.feature-details {
    font-size: 16px;
    color: #a3a3a3;
}
</style>
