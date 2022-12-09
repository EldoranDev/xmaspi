<script lang="ts" setup>
    let drawer = ref(false);
    const router = useRouter();

    const navigationItems = [
        { icon: 'mdi-pine-tree', title: 'Tree', target: 'index'},
        { icon: 'mdi-cog', title: 'Settings', target: 'settings'}
    ]

    function navigate(item: any): void {
        router.push({
            name: item.target,
        });
    }
</script>

<template>
    <v-app>
        <v-app-bar
            color="primary"
            prominent
        >
            <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
            <v-toolbar-title>XMAS-Pi</v-toolbar-title>
        </v-app-bar>

        <v-navigation-drawer v-model="drawer">
            <v-list
                :lines="false"
                nav
                @click:select="navigate"
            >
                <v-list-item
                    v-for="(item, i) in navigationItems"
                    active-color="primary"
                    :key="i"
                    :value="item.target"
                    
                >
                    <template #prepend>
                        <v-icon :icon="item.icon"></v-icon>
                        <v-list-item-title>{{ item.title }}</v-list-item-title>
                    </template>
                </v-list-item>
            </v-list>
        </v-navigation-drawer>

        <v-main>
            <slot />
        </v-main>
    </v-app>
</template>