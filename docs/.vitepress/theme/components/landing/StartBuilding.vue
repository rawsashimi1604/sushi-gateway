<template>
    <div class="hero__bgwrapper container">
        <!-- Sushi Chip -->
        <div class="sushi-chip">
            <img src="/images/Logo.png" alt="Sushi Gateway Logo" class="sushi-chip__logo" />
        </div>

        <h2>Try out Sushi Gateway</h2>
        <h3>
            Try it out now using our quick start guide.
        </h3>

        <div class="btn-container">
            <a href="/docs" class="btn btn--primary">Get started</a>
        </div>

    </div>
</template>

<script>
</script>

<style scoped>
.hero__bgwrapper {
    margin-top: 200px;
    width: 100%;
    position: relative;
    padding-top: 60px;
    text-align: center;
}

.sushi-chip {
    margin: 0 auto;
    margin-bottom: 30px;
    width: 100px;
    height: 100px;
    border-radius: 10px;
    background: linear-gradient(135deg, #2a78d4, #3fd9f0);
    display: flex;
    justify-content: center;
    align-items: center;
    box-shadow: 0 0 100px rgba(63, 217, 240, 0.6), 0 0 100px rgba(42, 120, 212, 0.5);
    z-index: 3;
    transition: box-shadow 0.3s ease-in-out, transform 0.3s ease-in-out;
}

.sushi-chip:hover {
    box-shadow: 0 0 30px rgba(63, 217, 240, 0.8), 0 0 15px rgba(42, 120, 212, 0.7);
    transform: scale(1.1);
}

.sushi-chip__logo {
    width: 80px;
    height: 80px;
    object-fit: contain;
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
    margin-bottom: 20px;

    @media (min-width: 768px) {
        font-size: 44px;
        letter-spacing: -0.88px;
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
    cursor: default;
    margin: 0 auto;
}

.btn-container {
    margin-top: 40px;
}

.btn {
    padding: 12px 24px;
    font-size: 16px;
    font-weight: 600;
    font-family: var(--font-secondary);
    border-radius: 8px;
    text-decoration: none;
    text-align: center;
    cursor: pointer;
    transition: all 0.3s ease;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
}

.btn--primary {
    color: #fff;
    background: linear-gradient(135deg, #2e88ee, #39c7dd);
    box-shadow: 0 4px 8px rgba(50, 194, 255, 0.3);
    border: none;

    &:hover {
        background: linear-gradient(90deg, #1a93cc, #32c2ff);
        box-shadow: 0 6px 12px rgba(50, 194, 255, 0.5);
        transform: translateY(-2px);
    }

    &:active {
        transform: translateY(1px);
        box-shadow: 0 3px 6px rgba(50, 194, 255, 0.2);
    }
}
</style>
