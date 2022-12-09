<script setup lang="ts">
  
  const { getElements, apply } = useXmaspi();

  const elements = await getElements();
</script>

<template>
  <div class="container">
    <div class="row g-2">
        <Renderable v-for="element in elements" :item="element" @apply="apply(element)"/>
    </div>
  </div>
</template>
