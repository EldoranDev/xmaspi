<script setup lang="ts">
defineProps<{
    item: RenderElement
}>();
</script>

<template>
    <v-card
        class="mx-auto"
        max-width="344"
    >
        <v-card-item>
            <v-card-title v-if="item.displayName">{{ item.displayName }}</v-card-title>
            <v-card-title v-else>{{ item.name }}</v-card-title>
        </v-card-item>
        <v-card-text>
            {{ item.description }}
        </v-card-text>
        <v-card-actions>
            <v-btn
                variant="outlined"
                color="primary"
                @click="$emit('apply', item)"
            >
                Run
            </v-btn>
        </v-card-actions>
    </v-card>
</template>