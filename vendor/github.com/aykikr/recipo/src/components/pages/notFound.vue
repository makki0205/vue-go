<template>
  <section>
    <h1>404 Not Found.</h1>
  </section>
</template>

<script>
  export default {
    name: "notFound"
  }
</script>

<style lang="scss" scoped>
:scope {
  display: block;

  section {
    width: 100vw;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: rgb(241, 243, 246);
  }

  h1 {
    font-size: 5rem;
    color: gray;
  }
}
</style>

