<template>
  <v-item v-slot="{ active, toggle }" active-class="active-card">
    <v-card
      class="center-box price-card text-wrap"
      :elevation="active? 9: 0"
      min-width="50"
      @click="toggle"
    >
      <v-card-title
        primary-title
        class="justify-center "
      >
        <h5 class="text-xs-center headline text-wrap">{{ title }}</h5>
      </v-card-title>
      <v-card-text v-if="card.unit_amount" class="pa-4">
        <div class="card-pricing">
          <span class=" text-xs-center display-2">
            € {{ price }}
          </span>
          <span class="title interval">
            / {{ interval }}
          </span>
        </div>
      </v-card-text>
    </v-card>
  </v-item>
</template>
<style scoped lang="sass">
@import './card-styles'
</style>
<script>
export default {
  name: 'TermCard',
  props: {
    card: {
      type: Object,
      required: true
    }
  },
  computed: {
    title () {
      if (!this.card.unit_amount) {
        return 'Free'
      }
      return this.interval === 'month' ? 'Monthly' : 'Yearly'
    },
    interval () {
      return this.card.recurring.interval
    },
    price () {
      return this.card.unit_amount / 100
    }
  }
}
</script>
