<template>
  <v-card class="card">
    <p class="body1">
      {{ text }}
    </p>
    <div class="name">
      <v-avatar
        :size="30"
        class="avatar"
      >
        <img
          :src="avatar"
          :alt="name"
        >
      </v-avatar>
      <span class="caption">
        {{ name }}
      </span>
    </div>
  </v-card>
</template>

<script>
export default {
  name: 'TestCard',
  props: {
    text: {
      type: String,
      required: true
    },
    avatar: {
      type: String,
      required: true
    },
    name: {
      type: String,
      required: true
    }
  }
}
</script>
