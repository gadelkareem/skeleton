<template>
  <v-card>
    <v-row>
      <v-col cols="12">
        <v-card>
          <v-card-title
            class="red darken-3 white--text subheading"
            primary-title
          >
            <span class="title">404 : Page Not Found.</span>
          </v-card-title>
          <v-card-text>
            Looks like you're in uncharted territory. We don't know about this page yet.
            Luckily, we know the way back.
          </v-card-text>
          <v-card-actions>
            <v-btn
              text
              class="red darken-3 white--text"
              to="/"
            >
              Back
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-card>
</template>

<script>

export default {
  name: 'Error'
}
</script>
