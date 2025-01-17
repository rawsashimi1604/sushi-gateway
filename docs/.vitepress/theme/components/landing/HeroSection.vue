<template>
    <div class="hero">
        <div class="container">
            <!-- Promo Pill -->
            <a href="https://github.com/rawsashimi1604/sushi-gateway" class="hero__pill" target="_blank">
                <img src="/images/Logo.png" alt="Sushi Gateway logo" width="20" height="20" />
                <span>Sushi Gateway v0.3.0 Released</span>
            </a>

            <!-- Heading -->
            <h1>The API Gateway<br />for the Future</h1>

            <!-- Tagline -->
            <h3>
                Sushi Gateway delivers high-performance API management, empowering developers with effortless control,
                robust security, and dynamic adaptability.
            </h3>

            <!-- CTA Buttons -->
            <div class="hero__actions">
                <a href="/sushi-gateway/getting-started/docker" class="btn btn--primary">Get started</a>
                <a href="https://github.com/rawsashimi1604/sushi-gateway" target="_blank" class="btn btn--outline">
                    <img src="/github.svg" alt="GitHub logo" width="24" height="24" />
                    GitHub
                </a>
            </div>
        </div>

    </div>
</template>

<style scoped>
.hero {
    position: relative;
    z-index: 2;
    margin-bottom: 150px;

    @media (min-width: 768px) {
        margin-bottom: 60px;
    }

    @media(min-width: 1025px) {
        margin-bottom: 0;
    }
}

.container {
    display: flex;
    flex-direction: column;
    max-width: 900px;
    margin: 60px auto 0;
    align-items: center;
    position: relative;
    z-index: 2;

    @media (min-width: 768px) {
        margin: 95px auto 0;
    }
}

.hero__pill {
    color: #fff;
    font-family: var(--font-secondary);
    font-weight: 500;
    display: inline-flex;
    padding: 8px 18px;
    align-items: center;
    gap: 8px;
    border-radius: 100px;
    border: 1px solid rgba(50, 194, 255, 0.5);
    background: radial-gradient(1686.42% 113.39% at 83.25% 2.56%,
            rgba(50, 194, 255, 0.1) 0%,
            rgba(50, 194, 255, 0) 100%);
    box-shadow: none;
    margin-bottom: 40px;
    text-decoration: none;
    cursor: pointer;

    @media (min-width: 768px) {
        margin-bottom: 25px;
    }

    span {
        font-size: 16px;
        font-weight: 500;
        line-height: normal;
        letter-spacing: -0.28px;
    }

    &:hover {
        color: inherit;
        animation: hero-pill-glow 1400ms ease-out infinite alternate;
        text-decoration: none;
    }
}

@keyframes hero-pill-glow {
    0% {
        border-color: rgba(50, 194, 255, 0.5);
        box-shadow:
            0 0 5px rgba(50, 194, 255, 0.01),
            inset 0 0 5px rgba(50, 194, 255, 0.1);
    }

    100% {
        border-color: rgba(50, 194, 255, 0.6);
        box-shadow:
            0 0 30px rgba(50, 194, 255, 0.2),
            inset 0 0 10px rgba(50, 194, 255, 0.2);
    }
}

h1 {
    text-align: center;
    font-family: var(--font-primary);
    font-style: normal;
    font-weight: 600;
    background: linear-gradient(180deg,
            #fff 0%,
            rgba(255, 255, 255, 0.31) 100%);
    background-clip: text;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    text-wrap: balance;
    cursor: default;
    font-size: 44px;
    line-height: 120%;
    letter-spacing: -0.88px;
    padding: 0 20px;
    margin-bottom: 15px;

    @media (min-width: 768px) {
        font-size: 64px;
        line-height: 81px;
        letter-spacing: -1.28px;
    }

    @media (min-width: 1025px) {
        font-size: 72px;
        letter-spacing: -1.44px;
        padding-bottom: 8px;
        /* Fix for hanging descender on "g" in "tooling" */
    }
}

h3 {
    color: #a9a9a9;
    text-align: center;
    font-family: var(--font-secondary);
    font-size: 20px;
    font-style: normal;
    font-weight: 400;
    line-height: 150%;
    letter-spacing: -0.4px;
    max-width: 500px;
    text-wrap: balance;
    cursor: default;
    margin-bottom: 25px;
    padding: 0 20px;
}

.hero__actions {
    display: flex;
    flex-direction: row;
    gap: 20px;
    margin-top: 20px;

    @media (min-width: 768px) {
        margin-top: 0;
    }
}
</style>
