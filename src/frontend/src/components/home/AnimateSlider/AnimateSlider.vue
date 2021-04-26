<template>
  <div class="hero-content">
    <transition-group
      :name="transition"
      tag="div"
      class="slider-wrapper"
    >
      <div
        v-for="number in [currentImg]"
        :key="number"
        :style="{ background: 'url(' + content[Math.abs(currentImg) % content.length].image + ') no-repeat center center' }"
        class="slide slider-content"
      >
        <div class="inner">
          <h1>{{ content[Math.abs(currentImg) % content.length].title }}</h1>
          <p>{{ content[Math.abs(currentImg) % content.length].description }}</p>
          <!--          <div class="hero-buttons d-flex justify-center">-->
          <!--            <v-btn-->
          <!--              color="primary"-->
          <!--              class="mx-2"-->
          <!--            >-->
          <!--              Main call to action-->
          <!--            </v-btn>-->
          <!--            <v-btn-->
          <!--              class="mx-2"-->
          <!--              outlined-->
          <!--              color="primary"-->
          <!--            >-->
          <!--              Secondary action-->
          <!--            </v-btn>-->
          <!--          </div>-->
          <section>
            <img
              :src="content[Math.abs(currentImg) % content.length].userProfile"
              :alt="content[Math.abs(currentImg) % content.length].user"
            >
            <span>
              Posted by&nbsp;
              <strong>
                {{ content[Math.abs(currentImg) % content.length].user }}
              </strong>
            </span>
          </section>
        </div>
      </div>
    </transition-group>
    <v-btn
      class="prev"
      fab
      small
      @click="prev"
    >
      <v-icon>
        mdi-arrow-left
      </v-icon>
    </v-btn>
    <v-btn
      class="next"
      fab
      small
      @click="next"
    >
      <v-icon>
        mdi-arrow-right
      </v-icon>
    </v-btn>
  </div>
</template>

<style scoped lang="sass">
@import './slider-styles'
</style>

<script>
import imgAPI from '@@/static/images/imgAPI'

export default {
  name: 'AnimateSlider',
  data () {
    return {
      show: true,
      transition: '',
      autoplay: null,
      content: [
        {
          title: 'Beautiful and tested Golang and NodeJS code',
          description: '',
          image: imgAPI.photo[0]
        },
        {
          title: 'Start your project with ready components',
          description: '',
          image: imgAPI.photo[1]
        },
        {
          title: 'Secured with JWT and MFA',
          description: '',
          image: imgAPI.photo[2]
        },
        {
          title: 'Beautiful material design using Vuetify',
          description: '',
          image: imgAPI.photo[4]
        },
        {
          title: 'Launch your project in minutes',
          description: '',
          image: imgAPI.photo[3]
        }
      ],
      currentImg: 0
    }
  },
  mounted () {
    this.playSlider()
  },
  beforeDestroy () {
    clearInterval(this.autoplay)
  },
  methods: {
    playSlider () {
      this.autoplay = setInterval(() => {
        this.currentImg = this.currentImg + 1
        this.transition = 'slide-right'
      }, 5000)
    },
    manualPlay () {
      clearInterval(this.autoplay)
      setTimeout(() => {
        this.playSlider()
      }, 1)
    },
    next () {
      this.manualPlay()
      this.currentImg += 1
      this.transition = 'slide-right'
    },
    prev () {
      this.manualPlay()
      this.currentImg -= 1
      this.transition = 'slide-left'
    }
  }
}
</script>
