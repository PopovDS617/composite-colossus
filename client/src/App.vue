<script setup>
import Heading from './components/Heading/Heading.vue';
import Counter from './components/Counter/Counter.vue';
</script>

<template>
  <main>
    <div>
      <Heading />
      <Counter />
    </div>
    <div>
      <!-- text area -->
    </div>
  </main>
</template>

<style scoped>
main {
  font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
}
</style>
