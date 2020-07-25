export default {
  /*
   ** Rendering mode
   ** Doc: https://nuxtjs.org/api/configuration-mode
   */
  mode: 'universal',
  server: {
    port: 8080
  },
  loading: {
    color: '#4996F2',
    height: '2px'
  },
  /*
   ** Headers of the page
   ** Doc: https://vue-meta.nuxtjs.org/api/#metainfo-properties
   */
  head: {
    title: 'Skeleton - skeletontagline',
    titleTemplate: '%s - Skeleton',
    meta: [
      {
        charset: 'utf-8'
      },
      {
        name: 'viewport',
        content: 'width=device-width, initial-scale=1'
      },
      {
        hid: 'description',
        name: 'description',
        content: 'skeletontagline'
      },
      {
        hid: 'msapplication-TileColor',
        name: 'msapplication-TileColor',
        content: '#da532c'
      },
      {
        hid: 'theme-color',
        name: 'theme-color',
        content: '#ffffff'
      }
    ],
    link: [
      { rel: 'apple-touch-icon', href: '/apple-touch-icon.png', sizes: '180x180' },
      { rel: 'icon', type: 'image/png', href: '/favicon-32x32.png', sizes: '32x32' },
      { rel: 'icon', type: 'image/png', href: '/favicon-16x16.png', sizes: '16x16' },
      { rel: 'mask-icon', href: '/safari-pinned-tab.svg', color: '#5bbad5' },
      { rel: 'manifest', href: '/site.webmanifest' }
    ]
  },
  /*
   ** Global CSS
   ** Doc: https://nuxtjs.org/api/configuration-css
   */
  css: [
    '@@/assets/sass/overrides.sass'
  ],

  /*
   ** Plugins to load before mounting the App
   ** Doc: https://nuxtjs.org/guide/plugins
   */
  plugins: [
    {
      src: '@@/plugins/chartist'
    },
    {
      src: '@@/plugins/mixin'
    },
    {
      src: '@@/plugins/i18n'
    }
  ],

  /*
   ** Nuxt.js modules
   ** Doc: https://nuxtjs.org/guide/modules
   */
  modules: [
    '@nuxtjs/eslint-module',
    '@nuxtjs/vuetify'
  ],

  // Doc: https://github.com/nuxt-community/vuetify-module
  vuetify: {
    customVariables: ['@@/assets/sass/variables.scss'],
    optionsPath: './vue.config.js'
  },
  build: {
    extractCSS: true,
    transpile: ['vuetify/lib'],
    babel: {
      configFile: './babel.config.js'
    }
  },
  axios: {
    credentials: true,
    headers: {
      common: {
        'X-Requested-With': 'XMLHttpRequest'
      }
    }
  },
  router: {
    trailingSlash: true
  },
  telemetry: false,
  env: {
    APIURL: process.env.API_URL || 'http://localhost:8081/api/v1'
  }
}
