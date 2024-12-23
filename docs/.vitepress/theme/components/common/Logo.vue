<script setup>
</script>

<template>
    <div class="logo__container">
        <!-- <span>logo</span> -->

        <div class="sushi-chip">
            <img src="/images/Logo.png" alt="Sushi Gateway Logo" class="sushi-chip__logo" />
        </div>

        <h1>
            Sushi Gateway
        </h1>
    </div>
</template>

<style scoped>
.logo__container {
    padding: 36px;
    background: white;
    color: black;

    display: flex;
    justify-content: center;
    align-items: center;
}

.sushi-chip {
    width: 80px;
    height: 80px;
    border-radius: 10px;

    background: linear-gradient(135deg, #2a78d4, #3fd9f0);


    box-shadow: 0 0 20px rgba(63, 217, 240, 0.6), 0 0 10px rgba(42, 120, 212, 0.5);
    z-index: 3;
    transition: box-shadow 0.3s ease-in-out, transform 0.3s ease-in-out;

    display: flex;
    justify-content: center;
    align-items: center;
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

h1 {
    text-align: center;
    font-family: var(--font-primary);
    font-style: normal;
    background: linear-gradient(180deg,
            #fff 0%,
            rgba(255, 255, 255, 0.31) 100%);
    background-clip: text;
    text-wrap: balance;
    cursor: default;
    font-size: 32px;
    line-height: 120%;
    letter-spacing: -1px;
    padding: 0 10px;
}
</style>
