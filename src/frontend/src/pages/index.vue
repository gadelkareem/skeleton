<template>
  <div>
    <section id="home" />
    <div class="main-wrap">
      <main v-if="show" class="container-wrap">
        <section id="animated-container">
          <v-container>
            <v-container fluid fill-height class="home-hero">
              <v-layout justify-center align-center column pa-5 mt-10>
                <img :src="require($vuetify.theme.dark ? '@@/static/logo-dark.svg' : '@@/static/logo.svg')" width="500">
              </v-layout>
            </v-container>
          </v-container>
        </section>
        <section
          id="feature"
          class="space-top space-bottom"
        >
          <v-container>
            <feature />
          </v-container>
        </section>
        <!--        <section-->
        <!--          class="page-section"-->
        <!--        >-->
        <!--          <counter />-->
        <!--        </section>-->
        <!--        <section-->
        <!--          id="testimonials"-->
        <!--          class="space-top"-->
        <!--        >-->
        <!--          <Testimonials />-->
        <!--        </section>-->
        <!--        <section-->
        <!--          id="blog"-->
        <!--          class="space-top space-bottom"-->
        <!--        >-->
        <!--          <blog />-->
        <!--        </section>-->
        <!--        <section id="subscribe">-->
        <!--          <subscribe />-->
        <!--        </section>-->
      </main>
      <hidden point="mdDown">
        <page-nav />
      </hidden>
      <hidden point="mdDown">
        <notification />
      </hidden>
    </div>
  </div>
</template>

<style scoped lang="sass">
@function section-margin($margin)
  @return $margin * 15

.main-wrap
  position: relative
  width: 100%
  overflow: hidden

.space-bottom
  margin-bottom: section-margin($spacing1)

.space-top
  padding-top: section-margin($spacing1)

.container-wrap
  & > section
    position: relative
.home-hero
  background-color: #eaeaea
.home-hero::after
  content: ""
  position: absolute
  height: 90px
  width: 100%
  background-color: #eaeaea
  bottom: 0
  left: 0
  transform: translateY(50%)
  border-radius: 0 0 100% 100%
.theme--dark
  .home-hero
    background-color: #1e1e1e
  .home-hero::after
    background-color: #1e1e1e

</style>

<script>
import Hidden from '@@/components/home/Hidden'
// import AnimateSlider from '@@/components/home/AnimateSlider'
import Feature from '@@/components/home/Feature'
// import Counter from '@@/components/home/Counter'
// import Testimonials from '@@/components/home/Testimonials'

// import Blog from '@@/components/home/Blog'
// import Subscribe from '@@/components/home/Subscribe'
import PageNav from '@@/components/home/PageNav'
import Notification from '@@/components/home/Notification'
import brand from '@@/static/text/brand'
export default {
  components: {
    // AnimateSlider,
    Hidden,
    Feature,
    // Counter,
    // Testimonials,
    // Blog,
    // Subscribe,
    PageNav,
    Notification
  },
  data: () => ({
    show: false
  }),
  mounted () {
    this.show = true
  },
  layout: 'Home',
  head () {
    return {
      title: brand.starter.name + ' - Home Page'
    }
  }
}
</script>
