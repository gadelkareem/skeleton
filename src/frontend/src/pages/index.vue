<template>
  <span>
    <v-carousel hide-delimiters height="30vh">
      <v-carousel-item
        v-for="(item,i) in items"
        :key="i"
        reverse-transition="fade-transition"
        transition="fade-transition"
        :src="item.src"
        style="background-blend-mode:darken;"
      >
        <v-row
          class="fill-height"
          align="center"
          justify="center"
        >
          <div class="display-4 white--text px-4 rounded" style="background-color: rgba(0,0,0,0.1);font-size: 2vw !important;">
            {{ item.text }}
          </div>
        </v-row>
      </v-carousel-item>
    </v-carousel>
    <material-section>
      <v-container>
        <v-row
          align="center"
          justify="center"
        >
          <v-col
            cols="3"
          >
            <v-icon size="100" color="#e0e0e0">mdi-image-multiple</v-icon>
            <div class="display-1 ml-2" style="color: #4e656d">Screenshots</div>
          </v-col>

          <v-col
            cols="8"
          >
            <v-card
              elevation="24"
              class="mx-auto"
              shaped
            >
              <v-carousel hide-delimiters height="800px">
                <v-carousel-item
                  v-for="n in 15"
                  :key="n"
                  :src="require(`@@/static/screenshots/${n}.png`)"
                  reverse-transition="fade-transition"
                  transition="fade-transition"
                  position="top"
                />
              </v-carousel>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </material-section>
    <material-section>
      <v-container>
        <v-row>
          <ul dir="auto">
            <li>Backend written in <a href="https://golang.org/" rel="nofollow noreferrer noopener" target="_blank">Golang</a> using <a href="https://beego.me/" rel="nofollow noreferrer noopener" target="_blank">Beego framework</a>.</li>
            <li>Frontend written in <a href="https://nodejs.org/en/" rel="nofollow noreferrer noopener" target="_blank">NodeJS</a> using <a href="https://nuxtjs.org/" rel="nofollow noreferrer noopener" target="_blank">NUXT</a> and <a href="https://vuejs.org/" rel="nofollow noreferrer noopener" target="_blank">Vue.js</a> frameworks.</li>
            <li>Material design using <a href="https://vuetifyjs.com/" rel="nofollow noreferrer noopener" target="_blank">Vuetify</a>.</li>
            <li>JSON REST API based on <a href="https://jsonapi.org/" rel="nofollow noreferrer noopener" target="_blank">jsonapi.org</a> standard.</li>
            <li>Fully featured user registration, login, password reminder, profile update, 2FA authentication, SMS based mobile confirmation, ... etc.</li>
            <li>
              <a href="https://jwt.io/" rel="nofollow noreferrer noopener" target="_blank">JSON Web Tokens (JWT)</a> based authentication.</li>
            <li>
              <a href="https://gitlab.com/gadelkareem/skeleton/-/blob/master/src/backend/services/SocialAuthService.go">Social login</a> using Facebook, LinkedIn, Google, Github.</li>
            <li>
              <a href="https://en.wikipedia.org/wiki/Search_engine_optimization" rel="nofollow noreferrer noopener" target="_blank">SEO friendly</a> thanks to <a href="https://nuxtjs.org/guide/#static-generated-pre-rendering-" rel="nofollow noreferrer noopener" target="_blank">the NUXT Static Generated (Pre Rendering)</a>.</li>
            <li>
              <a href="https://gitlab.com/gadelkareem/skeleton/-/tree/master/src/backend/migrations/sql">SQL migration</a> using <a href="https://github.com/rubenv/sql-migrate" rel="nofollow noreferrer noopener" target="_blank">sql-migrate</a>.</li>
            <li>
              <a href="https://gitlab.com/gadelkareem/skeleton/-/tree/master/src/backend/limiter">Rate limiter</a> for API routes to easily set a rate limit per IP for one or more routes.</li>
            <li>
              <a href="https://gitlab.com/gadelkareem/skeleton/-/tree/master/src/backend/rbac">Role-based access control (RBAC)</a> for API routes and permissions.</li>
            <li>
              <a href="https://gitlab.com/gadelkareem/skeleton/-/tree/master/src/backend/utils/paginator">Pagination</a> implementation for API and frontend with caching.</li>
            <li>
              <a href="https://gitlab.com/gadelkareem/skeleton/-/blob/master/src/backend/services/AuthenticatorService.go">Multi-factor authentication</a> using <a href="https://github.com/pquerna/otp" rel="nofollow noreferrer noopener" target="_blank">One Time Passwords</a> and mobile SMS code.</li>
            <li>
              <a href="https://gitlab.com/gadelkareem/skeleton/-/blob/master/src/backend/services/EmailService.go">Nice Email templates</a> using <a href="https://github.com/matcornic/hermes" rel="nofollow noreferrer noopener" target="_blank">Hermes</a>.</li>
            <li>
              <a href="https://gitlab.com/gadelkareem/skeleton/-/tree/master/src/frontend/src/pages/dashboard">Fully featured admin dashboard</a> based on <a href="https://demos.creative-tim.com/vuetify-material-dashboard/?partner=116160&amp;ref=vuetifyjs.com#/" rel="nofollow noreferrer noopener" target="_blank">Vuetify Material Dashboard</a>.</li>
            <li>
              <a href="https://gitlab.com/gadelkareem/skeleton/-/pipelines">Complete CI/CD pipelines</a> including tests using <a href="https://gitlab.com/gadelkareem/skeleton/-/blob/master/.gitlab-ci.yml">GitLab .gitlab-ci.yml</a> file.</li>
            <li>
              <a href="#deploy-to-heroku">Deploy to Heroku</a> using few easy steps.</li>
            <li>Automated development initialization using <a href="https://gitlab.com/gadelkareem/skeleton/-/blob/master/docker-compose.yml">Docker compose</a> and <a href="https://gitlab.com/gadelkareem/skeleton/-/blob/master/init.sh">init file</a>.</li>
            <li>Application Cache using <a href="https://github.com/gadelkareem/cachita" rel="nofollow noreferrer noopener" target="_blank">Cachita</a> with support for memory, Redis, database and file cache.</li>
            <li>
              <a href="https://gitlab.com/gadelkareem/skeleton/-/blob/master/src/backend/di/Container.go">Dependency injection</a>.</li>
            <li>Backend API integration and unit tests.</li>
            <li>Frontend <a href="https://github.com/facebook/jest" rel="nofollow noreferrer noopener" target="_blank">Jest</a> tests.</li>
          </ul>
        </v-row>
      </v-container>
    </material-section>
  </span>
</template>

<script>
import MaterialSection from '@@/components/base/MaterialSection'
export default {
  name: 'Home',
  components: { MaterialSection },
  layout: 'Home',
  data () {
    return {
      items: [
        {
          text: 'Beautiful and tested Golang and NodeJS code',
          src: require('@@/static/bg/code-projected-over-woman-3861969.jpg')
        },
        {
          text: 'Start your project with ready components',
          src: require('@@/static/bg/abstract-art-circle-clockwork-414579.jpg')
        },
        {
          text: 'Beautiful material design using Vuetify',
          src: require('@@/static/bg/office-working-app-computer-97077.jpg')
        },
        {
          text: 'Launch your project in minutes',
          src: require('@@/static/bg/space-rocket-launching-73871.jpg')
        },
        {
          text: 'Secured with JWT and MFA',
          src: require('@@/static/bg/strong-lock-locked-padlock-39624.jpg')
        }
      ]
    }
  }
}
</script>
