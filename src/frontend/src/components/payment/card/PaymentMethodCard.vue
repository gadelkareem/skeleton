<template>
  <v-item
    v-slot="{ active, toggle }"
    :value="pm.id"
    active-class="active-card"
    @change="$emit('setPaymentMethodID', pm.id)"
  >
    <v-card
      outlined
      shaped
      class="center-box card"
      :elevation="active? 9: 0"
      height="150"
      width="300"
      @click="toggle"
    >
      <div>
        <img
          :src="`/images/icons/payment/`+ pm.card.brand +`.svg`"
          width="40"
          height="30"
          style="float:right"
          :alt="pm.card.brand"
        >
      </div>
      <v-card-title>{{ pm.billing_details.name }}</v-card-title>
      <v-card-text>
        **** **** **** {{ pm.card.last4 }}<br>
        {{ pm.card.exp_month }}/{{ pm.card.exp_year }}<br>
        {{ pm.is_default ? 'PRIMARY' : '' }}
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn
          v-if="!pm.is_default && showDelete"
          x-small
          fab
          :ripple="false"
          class="no-bg-btn"
          :elevation="active? 9: 0"
          @click="$emit('deletePaymentMethod',pm.id)"
        >
          <v-icon
            color="error"
            class="mx-1"
          >
            mdi-close
          </v-icon>
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-item>
</template>

<script>

export default {
  name: 'PaymentMethodCard',
  props: {
    pm: {
      type: Object,
      required: true
    },
    showDelete: {
      type: Boolean,
      default: true
    }
  }
}
</script>
