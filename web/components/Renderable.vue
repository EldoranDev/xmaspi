<script setup lang="ts">
defineProps<{
    item: RenderElement
}>();
</script>

<template>
    <div class="card" style="width: 18rem">
        <div class="card-body">
            <h5 v-if="item.displayName" class="card-title">{{ item.displayName }}</h5>
            <h5 v-else class="card-title">{{ item.name }}</h5>
            
            <p v-if="item.description" class="card-text">
                {{ item.description }}
            </p>

            <button class="btn btn-primary" @click="$emit('apply', item)">
                Run
            </button>
        </div>
    </div>
</template>