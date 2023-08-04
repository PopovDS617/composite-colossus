<script setup>
import './Counter.css';
import Display from './Display/Display.vue';
import Btn from './Button/Btn.vue';
import { ref } from 'vue';

const count = ref(0);

const plus = () => count.value++;
const minus = () => count.value--;
</script>

<template>
  <div class="counter-container">
    <Display count="count" />
    <div class="controls"><Btn onClick="minus" /> <Btn onClick="plus" /></div>
  </div>
</template>
