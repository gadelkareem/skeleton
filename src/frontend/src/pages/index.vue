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
