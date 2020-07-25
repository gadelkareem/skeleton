<template>
  <v-list-group
    :group="group"
    :prepend-icon="item.icon"
    :sub-group="subGroup"
    append-icon="mdi-menu-down"
    :color="barColor !== 'rgba(255, 255, 255, 1), rgba(255, 255, 255, 0.7)' ? 'white' : 'grey darken-1'"
  >
    <template v-slot:activator>
      <v-list-item-icon
        v-if="text"
        class="v-list-item__icon--text"
        v-text="computedText"
      />

      <v-list-item-content>
        <v-list-item-title v-text="item.title" />
      </v-list-item-content>
    </template>

    <template v-for="(child, i) in children">
      <item-sub-group
        v-if="child.children"
        :key="`sub-group-${i}`"
        :item="child"
      />

      <item
        v-else
        :key="`item-${i}`"
        :item="child"
        text
      />
    </template>
  </v-list-group>
</template>

<script>

import ItemSubGroup from './ItemSubGroup'
import Item from './Item'
export default {
  name: 'ItemGroup',
  components: { Item, ItemSubGroup },
  inheritAttrs: false,

  props: {
    item: {
      type: Object,
      default: () => ({
        avatar: undefined,
        group: undefined,
        title: undefined,
        children: []
      })
    },
    subGroup: {
      type: Boolean,
      default: false
    },
    text: {
      type: Boolean,
      default: false
    }
  },

  computed: {
    barColor () {
      return this.$store.getters['dashboard/barColor']
    },
    children () {
      return this.item.children.map(item => ({
        ...item,
        to: item.to
      }))
    },
    computedText () {
      if (!this.item || !this.item.title) { return '' }

      let text = ''

      this.item.title.split(' ').forEach((val) => {
        text += val.substring(0, 1)
      })

      return text
    },
    group () {
      return this.genGroup(this.item.children)
    }
  },
  methods: {
    genGroup (children) {
      return children
        .filter(item => item.to)
        .map((item) => {
          return item.to
        }).join('|')
    }
  }
}
</script>

<style>
  .v-list-group__activator p {
    margin-bottom: 0;
  }
</style>
