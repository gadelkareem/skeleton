<template>
  <v-card class="center-box price-card">
    <v-card-title
      primary-title
      class="price-header justify-center"
    >
      <h5 class="text-xs-center headline">{{ card.title }}</h5>
      <div class="small-title text-xs-center body-1"> {{ card.subheader }}</div>
      <v-icon
        v-if="card.title === 'Pro'"
        class="icon"
      >
        mdi-star-outline
      </v-icon>
    </v-card-title>
    <v-card-text class="pa-4">
      <div class="card-pricing">
        <h2 class="price text-xs-center display-2">
          $ {{ card.price }}
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
        :outlined="card.buttonVariant === 'outlined'"
        block
        color="primary"
        :disabled="card.disabled"
        :to="card.link"
      >
        {{ card.buttonText }}
      </v-btn>
    </div>
  </v-card>
</template>

<style scoped lang="sass">
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
