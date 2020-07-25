<template>
  <span>
    <v-alert
      v-for="error in errors"
      :key="error.title"
      type="error"
    >
      <div v-for="s in split(error.title)" :key="s">
        {{ s }}
      </div>
      <ul v-if="error.meta">
        <li
          v-for="(msg, key) in error.meta"
          :key="key"
        >
          <strong>{{ key.split('.')[0] }}:</strong>
          <span v-for="s in split(msg)" :key="s">
            {{ s }}<br>
          </span>
        </li>
      </ul>
    </v-alert>
    <v-alert
      v-if="success"
      type="success"
    >
      <div v-for="s in split(successText)" :key="s">
        {{ s }}
      </div>
    </v-alert>
    <v-alert
      v-if="info"
      type="info"
    >
      <div v-for="s in split(info)" :key="s">
        {{ s }}
      </div>
    </v-alert>
  </span>
</template>

<script>

export default {
  name: 'Alert',
  props: {
    success: {
      type: Boolean,
      default: false
    },
    successText: {
      type: String,
      default: ''
    },
    info: {
      type: String,
      default: ''
    },
    errors: {
      type: Array,
      default: null
    }
  },
  watch: {
    errors () {
      if (this.errors && this.errors[0] && this.errors[0].title === 'Invalid authentication token.') {
        this.$router.push('/auth/logout/')
      }
    }
  },
  methods: {
    split: s => s.split('|')
  }
}
</script>
