<template>
  <div>
    <section id="home" />
    <div class="main-wrap">
      <main v-if="show" class="container-wrap">
        <section id="animated-container">
          <v-container>
            <animate-slider />
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
        <!--          id="pricing"-->
        <!--          class="space-top"-->
        <!--        >-->
        <!--          <pricing />-->
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

</style>

<script>
import Hidden from '@@/components/home/Hidden'
import AnimateSlider from '@@/components/home/AnimateSlider'
import Feature from '@@/components/home/Feature'
// import Counter from '@@/components/home/Counter'
// import Testimonials from '@@/components/home/Testimonials'
// import Pricing from '@@/components/home/Pricing'
// import Blog from '@@/components/home/Blog'
// import Subscribe from '@@/components/home/Subscribe'
import PageNav from '@@/components/home/PageNav'
import Notification from '@@/components/home/Notification'
import brand from '@@/static/text/brand'

export default {
  components: {
    AnimateSlider,
    Hidden,
    Feature,
    // Counter,
    // Testimonials,
    // Pricing,
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
