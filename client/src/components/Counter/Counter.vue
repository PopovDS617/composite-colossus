<script setup>
import './Counter.css';
import Display from './Display/Display.vue';
import Btn from './Button/Btn.vue';
import { ref } from 'vue';

const count = ref(0);
</script>

<template>
  <div class="counter-container">
    <Display count="count" />
    <div class="controls"><Btn /> <Btn /></div>
  </div>
</template>
