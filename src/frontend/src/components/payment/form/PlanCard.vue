<template>
  <v-item id="plan-card" v-slot="{ active, toggle }" active-class="active-card">
    <v-card
      class="center-box card"
      :elevation="active? 9: 0"
      @click="toggle"
    >
      <v-card-title
        primary-title
        class="justify-center"
      >
        <h5 class="text-xs-center headline">{{ card.title }}</h5>
        <v-icon
          v-if="card.title.includes('Premium')"
          class="icon"
        >
          mdi-star-outline
        </v-icon>
      </v-card-title>
      <v-card-text class="pa-4">
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
    </v-card>
  </v-item>
</template>
<script>
export default {
  name: 'PlanCard',
  props: {
    card: {
      type: Object,
      required: true
    }

  }
}
</script>
