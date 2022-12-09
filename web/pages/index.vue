<script setup lang="ts">
  const { getElements, apply } = useXmaspi();

  const loading = ref(true);
  const elements = ref([] as Array<RenderElement>);
    
  getElements()
  .then((el) => {
    unref(elements).push(...el)
  });

</script>

<template>
  <v-container>
    <v-row dense>
        <v-col
          cols="12"
          xs="12"
          md="4"
          v-for="element in elements"
        >
          <Renderable :item="element" @apply="apply(element)"/>
        </v-col>
      </v-row>
  </v-container>
</template>
