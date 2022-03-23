<template>
  <v-card class="center-box price-card">
    <v-card-title
      class="price-header justify-center"
      primary-title
    >
      <h5 class="text-xs-center headline">{{ card.title }}</h5>
      <div class="small-title text-xs-center body-1"> {{ card.subheader }}</div>
      <v-icon
        v-if="card.title.includes('Premium')"
        class="icon"
      >
        mdi-star-outline
      </v-icon>
    </v-card-title>
    <v-card-text class="pa-4">
      <div class="card-pricing">
        <h2 class="price text-xs-center display-2">
          {{ formatMoney(card.price) }}
        </h2>
        <h6 class="title">
          / mo
        </h6>
      </div>
      <ul>
        <li
          v-for="(line, index) in card.description"
          :key="index"
          class="text-center subtitle-1"
        >
          {{ line }}
        </li>
      </ul>
    </v-card-text>
    <div class="text-center pa-4">
      <v-btn
        :disabled="card.disabled"
        :outlined="card.buttonVariant === 'outlined'"
        :to="`/dashboard/subscription/plans/?plan=`+card.priceID"
        block
        color="primary"
      >
        {{ card.buttonText }}
      </v-btn>
    </div>
  </v-card>
</template>

<style lang="sass" scoped>
@import './card-styles'
</style>

<script>
export default {
  name: 'PricingCard',
  props: {
    card: {
      type: Object,
      required: true
    }
  }
}
</script>
