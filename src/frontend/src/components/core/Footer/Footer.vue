<template>
  <v-footer
    class="footer"
    height="400"
    v-bind="$attrs"
  >
    <v-container>
      <v-row no-gutters>
        <v-col
          md="5"
          class="text-center"
        >
          <a class="logo" href="/">
            <v-img
              :src="logo"
              contain
              position="left"
            />
          </a>
          <p class="body-2">
            &copy;&nbsp; {{ brand.starter.footerText }}
          </p>
        </v-col>
        <v-col
          md="7"
        >
          <v-expansion-panels v-if="isMobile">
            <v-expansion-panel
              v-for="(footer, index) in footers"
              :key="index"
              class="accordion-content"
            >
              <v-expansion-panel-header>
                <h6 class="title mb-4">
                  {{ footer.title }}
                </h6>
              </v-expansion-panel-header>
              <v-expansion-panel-content>
                <ul>
                  <li
                    v-for="(item, index2) in footer.description"
                    :key="index2"
                  >
                    <a :href="footer.link[index2]">
                      {{ item }}
                    </a>
                  </li>
                </ul>
              </v-expansion-panel-content>
            </v-expansion-panel>
          </v-expansion-panels>
          <v-row
            v-if="isDesktop"
            justify="space-around"
          >
            <v-col
              v-for="(footer, index) in footers"
              :key="index"
              class="pa-4 site-map-item"
            >
              <h6 class="title mb-4">
                {{ footer.title }}
              </h6>
              <ul>
                <li
                  v-for="(item, index2) in footer.description"
                  :key="index2"
                >
                  <a :href="footer.link[index2]">
                    {{ item }}
                  </a>
                </li>
              </ul>
            </v-col>
          </v-row>
        </v-col>
        <!--        <v-col-->
        <!--          md="6"-->
        <!--          cols="12"-->
        <!--        >-->
        <!--          <div class="socmed">-->
        <!--            <v-btn-->
        <!--              text-->
        <!--              icon-->
        <!--              class="button"-->
        <!--            >-->
        <!--              <span class="ion-social-facebook icon" />-->
        <!--            </v-btn>-->
        <!--            <v-btn-->
        <!--              text-->
        <!--              icon-->
        <!--              class="button"-->
        <!--            >-->
        <!--              <span class="ion-social-twitter icon" />-->
        <!--            </v-btn>-->
        <!--            <v-btn-->
        <!--              text-->
        <!--              icon-->
        <!--              class="button"-->
        <!--            >-->
        <!--              <span class="ion-social-instagram icon" />-->
        <!--            </v-btn>-->
        <!--            <v-btn-->
        <!--              text-->
        <!--              icon-->
        <!--              class="button"-->
        <!--            >-->
        <!--              <span class="ion-social-linkedin icon" />-->
        <!--            </v-btn>-->
        <!--          </div>-->
        <!--        &lt;!&ndash;        <v-select&ndash;&gt;-->
        <!--        &lt;!&ndash;          v-model="lang"&ndash;&gt;-->
        <!--        &lt;!&ndash;          :items="langList"&ndash;&gt;-->
        <!--        &lt;!&ndash;          :value="lang"&ndash;&gt;-->
        <!--        &lt;!&ndash;          label=""&ndash;&gt;-->
        <!--        &lt;!&ndash;          outlined&ndash;&gt;-->
        <!--        &lt;!&ndash;          class="select-lang"&ndash;&gt;-->
        <!--        &lt;!&ndash;          @change="switchLang(lang)"&ndash;&gt;-->
        <!--        &lt;!&ndash;        />&ndash;&gt;-->
        <!--        </v-col>-->
      </v-row>
    </v-container>
  </v-footer>
</template>

<style scoped lang="sass">
@import './footer-style'
</style>

<script>
import brand from '@@/static/text/brand'

export default {
  data: () => ({
    brand,
    // lang: 'en',
    footers: [
      {
        title: 'Company',
        description: ['Contact us'],
        link: ['/home/contact/']
      },
      {
        title: 'Project Links',
        description: [
          'Github',
          'Gitlab'
        ],
        link: [
          'https://github.com/gadelkareem/skeleton',
          'https://gitlab.com/gadelkareem/skeleton'
        ]
      }
    ]
  }),
  computed: {
    // langList () {
    //   const list = []
    //   this.$i18n.locales.map((item) => {
    //     list.push({ text: this.$t('common.' + item.code), value: item.code })
    //   })
    //   return list
    // },
    isMobile () {
      const smDown = this.$store.state.breakpoints.smDown
      return smDown.includes(this.$mq)
    },
    isDesktop () {
      const mdUp = this.$store.state.breakpoints.mdUp
      return mdUp.includes(this.$mq)
    },
    logo () {
      return require(this.$vuetify.theme.dark ? '@@/static/logo-dark.svg' : '@@/static/logo.svg')
    }
  },
  mounted () {
    // this.lang = this.$i18n.locale
  },
  methods: {
    // switchLang (val) {
    //   this.$i18n.setLocale(val)
    // }
  }
}
</script>
