import FilterWarningsPlugin from 'webpack-filter-warnings-plugin'
import languages from './static/lang/languages'

export default {
  mode: 'universal',
  server: {
    port: 8080
  },
  /*
  ** Customize the progress-bar color
  */
  loading: {
    color: '#4996F2',
    height: '2px'
  },
  /*
   ** Headers of the page
   ** Doc: https://vue-meta.nuxtjs.org/api/#metainfo-properties
   */
  head: {
    htmlAttrs: {
      dir: 'ltr'
    },
    title: 'Proxy Cloud - skeletontagline',
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
        content: '#ffffff'
      },
      { name: 'msapplication-TileImage', content: '/favicons/ms-icon-144x144.png' },
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
      // Fonts and Icons
      // { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css?family=Montserrat:400,500,600&display=swap' }
      // { rel: 'stylesheet', href: 'https://fonts.googleapis.com/icon?family=Material+Icons' },
      // { rel: 'stylesheet', href: 'https://code.ionicframework.com/ionicons/2.0.1/css/ionicons.min.css' }
    ]
  },
  /*
   ** Global CSS
   ** Doc: https://nuxtjs.org/api/configuration-css
   */
  css: [
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
    },
    '@@/plugins/vue-fragment-config',
    '@@/plugins/vue-wow-config',
    { src: '@@/plugins/caroussel-config', ssr: false },
    { src: '@@/plugins/countup-config', ssr: false },
    { src: '@@/plugins/vue-scroll-nav', ssr: false }
  ],

  /*
   ** Nuxt.js modules
   ** Doc: https://nuxtjs.org/guide/modules
   */
  modules: [
    '@nuxtjs/eslint-module',
    '@nuxtjs/vuetify',
    ['@nuxtjs/html-minifier', { log: 'once', logHtml: true }],
    [
      'nuxt-mq',
      {
        // Default breakpoint for SSR
        defaultBreakpoint: 'default',
        breakpoints: {
          xsDown: 599,
          xsUp: 600,
          smDown: 959,
          smUp: 960,
          mdDown: 1279,
          mdUp: 1280,
          lgDown: 1919,
          lgUp: 1920,
          xl: Infinity
        }
      }
    ],
    [
      'nuxt-i18n',
      {
        // Options
        // to make it seo friendly remove below line and add baseUrl option to production domain
        seo: false,
        // baseUrl: 'https://my-nuxt-app.com',
        lazy: true,
        locales: languages,
        defaultLocale: 'en',
        vueI18n: {
          fallbackLocale: 'en'
        },
        detectBrowserLanguage: {
          useCookie: true,
          cookieKey: 'i18n_redirected',
          alwaysRedirect: true
        },
        langDir: 'static/lang/'
      }
    ]
  ],
  /*
    ** Render configuration
    */
  render: {
    bundleRenderer: {
      directives: {
        shouldPreload: (file, type) => {
          return ['script', 'style', 'font'].includes(type)
        },
        scroll (el, binding) {
          // const f = function (evt) {
          //   if (binding.value(evt, el)) {
          //     window.removeEventListener('scroll', f)
          //   }
          // }
          // window.addEventListener('scroll', f)
        }
      }
    }
  },
  // Doc: https://github.com/nuxt-community/vuetify-module
  vuetify: {
    customVariables: ['@@/assets/sass/dashboard_variables.scss', '@@/assets/sass/home_variables.sass'],
    optionsPath: './config/vuetify.options.js',
    treeShake: true
  },
  build: {
    /*
    ** You can extend webpack config here
    */
    // cssSourceMap: false,
    loaders: {
      vus: { cacheBusting: true },
      scss: { sourceMap: false }
    },
    extend (config, ctx) {
      config.plugins.push(
        new FilterWarningsPlugin({
          exclude: /Critical dependency: the request of a dependency is an expression/
        })
      )
      if (ctx.isDev && ctx.isClient) {
        config.module.rules.push({
          enforce: 'pre',
          test: /\.(js|vue)$/,
          loader: 'eslint-loader',
          exclude: /([node_modules, static])/,
          options: {
            fix: false
          }
        })
      }
    },
    extractCSS: {
      ignoreOrder: true
    },
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
    // scrollBehavior: () => ({ x: 0, y: 0 })
  },
  telemetry: false,
  env: {
    APIURL: process.env.API_URL || 'http://localhost:8081/api/v1'
  },
  /*
  ** Page Layout transition
  */
  layoutTransition: {
    name: 'layout',
    mode: 'out-in',
    beforeEnter (el) {
      console.log('Before enter...')
    },
    afterLeave (el) {
      console.log('afterLeave', el)
    }
  }
}
