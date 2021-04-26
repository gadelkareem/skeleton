<template>
  <div class="counter-wrap">
    <div id="watched_counter" />
    <v-container>
      <v-row
        align="center"
        justify="center"
        class="root"
      >
        <v-col sm="4">
          <div class="counter-item">
            <span class="icon ion-leaf" />
            <div class="text">
              <h4 v-if="loaded" class="display-1">
                <span
                  v-countUp:onWindowScroll.once="{
                    watchedElId: 'watched_counter',
                    startValue: 0,
                    endValue: 123,
                    options: { duration: 1 }
                  }"
                />
              </h4>
              <h6 class="title">Lorem</h6>
            </div>
          </div>
        </v-col>
        <v-col sm="4">
          <div class="counter-item">
            <span class="icon ion-ios-lightbulb-outline" />
            <div class="text">
              <h4 v-if="loaded" class="display-1">
                <span
                  v-countUp:onWindowScroll.once="{
                    watchedElId: 'watched_counter',
                    startValue: 0,
                    endValue: 456,
                    options: { duration: 1 }
                  }"
                />
              </h4>
              <h6 class="title">Lorem</h6>
            </div>
          </div>
        </v-col>
        <v-col sm="4">
          <div class="counter-item">
            <span class="icon ion-ios-world-outline" />
            <div class="text">
              <h4 v-if="loaded" class="display-1">
                <span
                  v-countUp:onWindowScroll.once="{
                    watchedElId: 'watched_counter',
                    startValue: 0,
                    endValue: 789,
                    options: { duration: 1 }
                  }"
                />
              </h4>
              <h6 class="title">Lorem</h6>
            </div>
          </div>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<style scoped lang="sass">
@import './counter-style'
</style>

<script>
export default {
  name: 'Counter',
  data: () => ({
    loaded: false
  }),
  mounted () {
    this.loaded = true
  }
}
</script>
