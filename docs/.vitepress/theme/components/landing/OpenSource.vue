<template>
    <div class="hero__bgwrapper container">


        <img class="icon-heart" src="/heart.svg"
            alt="Sushi Gateway is MIT Licensed and will always be free and open source." width="58" height="55" />
        <h2>Free &amp; open source</h2>
        <h3>
            Sushi Gateway is MIT Licensed and will always be free and open source.
        </h3>
    </div>
</template>

<script>
</script>

<style scoped>
.icon-heart {
    margin: 0 auto;
    margin-bottom: 15px;
}

.hero__bgwrapper {
    margin-top: 200px;
    width: 100%;
    position: relative;
    padding-top: 60px;
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
</style>
