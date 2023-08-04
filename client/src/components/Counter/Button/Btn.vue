<script setup>
import './Btn.css';

defineProps({ onClick: { type: Function } });
</script>

<template>
  <button class="btn" @click="onClick">+</button>
</template>
