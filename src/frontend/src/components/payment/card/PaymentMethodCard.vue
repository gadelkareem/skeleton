<template>
  <v-item
    v-slot="{ active, toggle }"
    :value="pm.id"
    active-class="active-card"
    @change="$emit('setPaymentMethodID', pm.id)"
  >
    <v-card
      :elevation="active? 9: 0"
      class="center-box card"
      height="150"
      outlined
      shaped
      width="300"
      @click="toggle"
    >
      <confirm-modal ref="confirm" />
      <div>
        <img
          :alt="pm.card.brand"
          :src="`/images/icons/payment/`+ pm.card.brand +`.svg`"
          height="30"
          style="float:right"
          width="40"
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
          :elevation="active? 9: 0"
          :ripple="false"
          class="no-bg-btn"
          fab
          x-small
          @click="deletePaymentMethod"
        >
          <v-icon
            class="mx-1"
            color="error"
          >
            mdi-close
          </v-icon>
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-item>
</template>

<script>
import ConfirmModal from '@@/components/base/ConfirmModal'
export default {
  name: 'PaymentMethodCard',
  components: { ConfirmModal },
  props: {
    pm: {
      type: Object,
      required: true
    },
    showDelete: {
      type: Boolean,
      default: true
    }
  },
  methods: {
    deletePaymentMethod () {
      this.$refs.confirm.open(
        'Delete Payment Method',
        'Are you sure you want to delete your payment method?'
      ).then((confirm) => {
        if (confirm) {
          this.$emit('deletePaymentMethod', this.pm.id)
        }
      })
    }
  }
}
</script>
