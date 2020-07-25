<template>
  <item-group
    :item="item"
    text
    sub-group
  />
</template>

<script>
import ItemGroup from './ItemGroup'
export default {
  name: 'ItemSubGroup',
  components: { ItemGroup },
  props: {
    item: {
      type: Object,
      default: () => ({
        avatar: undefined,
        group: undefined,
        title: undefined,
        children: []
      })
    }
  }
}
</script>
