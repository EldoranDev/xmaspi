// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    ssr: false,
    app: {
        head: {
            link: [
                { rel: 'stylesheet', 'href': "https://cdn.jsdelivr.net/npm/@mdi/font@5.x/css/materialdesignicons.min.css" },
            ]
        }
    },
    build: {
        transpile: [
            'vuetify',
        ]
    },
    css: [
        'vuetify/lib/styles/main.sass',
    ],
})
