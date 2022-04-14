<template>
  <div class="page-section">
    <v-row>
      <v-col class="px-6" md="4">
        <div class="feature-list">
          <v-icon color="primary" class="icon">mdi-receipt</v-icon>
          <h5 class="headline">
            Subscription Management & Recurring Billing
          </h5>
          <p class="body-1">
            Skeleton uses <a href="https://stripe.com/" rel="nofollow noreferrer noopener" target="_blank">Stripe</a> to provide a subscription management platform.
          </p>
        </div>
      </v-col>
      <v-col class="px-6" md="4">
        <div class="feature-list">
          <v-icon color="primary" class="icon">mdi-open-source-initiative</v-icon>
          <h5 class="headline">
            Free and Open Source
          </h5>
          <p class="body-1">
            Skeleton is totally open source and available on both <a href="https://github.com/gadelkareem/skeleton">Github</a> and <a href="https://gitlab.com/gadelkareem/skeleton">Gitlab</a>.
          </p>
        </div>
      </v-col>
      <v-col class="px-6" md="4">
        <div class="feature-list">
          <v-icon color="primary" class="icon">mdi-robot-excited-outline</v-icon>
          <h5 class="headline">
            Dev env & CI/CD included
          </h5>
          <p class="body-1">
            Automated development initialization using Docker compose and init file. Complete CI/CD pipelines including tests using GitLab .gitlab-ci.yml file. Deploy to Heroku using few easy steps.
          </p>
        </div>
      </v-col>
      <v-col class="px-6" md="4" />
      <v-col class="px-6" md="4">
        <div class="feature-list">
          <v-icon color="primary" class="icon">mdi-code-tags</v-icon>
          <h5 class="headline">
            Based on the best frameworks
          </h5>
          <p class="body-1">
            Backend written in <a href="https://golang.org/" rel="nofollow noreferrer noopener" target="_blank">Golang</a> using <a href="https://beego.me/" rel="nofollow noreferrer noopener" target="_blank">Beego framework</a>. Frontend written in <a href="https://nodejs.org/en/" rel="nofollow noreferrer noopener" target="_blank">NodeJS</a> using <a href="https://nuxtjs.org/" rel="nofollow noreferrer noopener" target="_blank">NUXT</a> and <a href="https://vuejs.org/" rel="nofollow noreferrer noopener" target="_blank">Vue.js</a> frameworks. Material design using <a href="https://vuetifyjs.com/" rel="nofollow noreferrer noopener" target="_blank">Vuetify</a>.
          </p>
        </div>
      </v-col>
    </v-row>
  </div>
</template>

<style scoped lang="sass">
@import './feature-styles'
</style>

<script>
export default {
  name: 'MainFeature'
}
</script>
