<template>
  <div class="feature-more">
    <v-container>
      <u-animate-container>
        <scroll-wrap target="feature1">
          <v-row id="feature1">
            <v-col
              md="6"
              class="px-6"
              cols="12"
            >
              <div class="feature-item">
                <u-animate
                  name="fadeInLeft"
                  duration="0.6s"
                >
                  <h3 class="title-app display-2">Admin Dashboard</h3>
                </u-animate>
                <u-animate
                  name="fadeInLeft"
                  delay="0.3s"
                  duration="0.6s"
                >
                  <p class="text">
                    Fully featured admin dashboard based on Vuetify Material Dashboard.
                  </p>
                </u-animate>
              </div>
            </v-col>
            <v-col
              md="6"
              class="px-6"
              cols="12"
            >
              <div class="feature-item">
                <u-animate
                  name="zoomIn"
                  delay="0.3s"
                  duration="0.6s"
                >
                  <div class="deco1" />
                </u-animate>
                <u-animate
                  name="fadeInRight"
                  delay="0.5s"
                  duration="0.6s"
                >
                  <figure class="img">
                    <img
                      :src="require('@@/static/screenshots/4.jpg')"
                      alt="img"
                    >
                  </figure>
                </u-animate>
              </div>
            </v-col>
          </v-row>
        </scroll-wrap>
      </u-animate-container>
      <u-animate-container>
        <v-row>
          <v-col
            md="6"
            class="px-6"
            cols="12"
          >
            <div class="feature-item">
              <u-animate
                name="zoomIn"
                delay="0.3s"
                duration="0.6s"
              >
                <div class="deco1" />
              </u-animate>
              <u-animate
                name="fadeInLeft"
                delay="0.5s"
                duration="0.6s"
              >
                <figure class="img">
                  <img
                    :src="require('@@/static/screenshots/1.jpg')"
                    alt="img"
                  >
                </figure>
              </u-animate>
            </div>
          </v-col>
          <v-col
            md="6"
            class="px-6"
            cols="12"
          >
            <div class="feature-item">
              <u-animate
                name="fadeInLeft"
                duration="0.6s"
              >
                <h3 class="title-app display-2">Advanced Authentication</h3>
              </u-animate>
              <u-animate
                name="fadeInRight"
                delay="0.5s"
                duration="0.6s"
              >
                <p class="text">
                  JSON Web Tokens (JWT) based authentication. JSON REST API based on jsonapi.org standard. Social login using Facebook, LinkedIn, Google, Github.
                </p>
              </u-animate>
            </div>
          </v-col>
          <v-col
            md="12"
            class="px-12"
            cols="12"
          >
            <div class="feature-item">
              <u-animate
                name="fadeInRight"
                delay="0.5s"
                duration="0.6s"
              >
                <pricing />
              </u-animate>
            </div>
          </v-col>
        </v-row>
      </u-animate-container>
      <u-animate-container>
        <v-row>
          <v-col
            md="12"
            class="px-6"
            cols="12"
          >
            <div class="feature-item last text-center">
              <u-animate
                name="fadeInUp"
                duration="0.6s"
              >
                <h3 class="title-app display-2">And more...</h3>
              </u-animate>
              <u-animate
                name="fadeInUp"
                duration="0.6s"
                delay="0.3s"
              >
                <p class="text">
                  Check these quick screenshots or hit <v-btn color="primary" to="/auth/register/">register</v-btn>and start testing life demo.
                </p>
              </u-animate>
              <u-animate
                name="zoomIn"
                duration="0.6s"
                delay="0.3s"
              >
                <div class="deco2" />
              </u-animate>
              <u-animate
                name="fadeInUp"
                duration="0.6s"
                delay="0.5s"
              >
                <v-card
                  elevation="24"
                  class="mx-auto"
                  shaped
                >
                  <v-carousel hide-delimiters height="820" width="1440" cycle>
                    <v-carousel-item
                      v-for="n in 21"
                      :key="n"
                      :src="require(`@@/static/screenshots/${n}.jpg`)"
                      reverse-transition="fade-transition"
                      contain
                    />
                  </v-carousel>
                </v-card>
              </u-animate>
            </div>
          </v-col>
        </v-row>
      </u-animate-container>
    </v-container>
  </div>
</template>

<style scoped lang="sass">
@import './feature-styles'
</style>

<script>
import imgAPI from '@@/static/images/imgAPI'
import Pricing from '@@/components/home/Pricing/Pricing'
import ScrollWrap from '../ScrollWrap'

export default {
  name: 'MoreFeature',
  components: {
    Pricing,
    ScrollWrap
  },
  data () {
    return {
      imgAPI
    }
  }
}
</script>
