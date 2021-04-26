<template>
  <div
    v-scroll="handleScroll"
    :class="{ 'active': active, 'pending': !active }"
    class="scroll-wrap"
  >
    <slot />
  </div>
</template>

<script>
export default {
  name: 'ScrollWrap',
  props: {
    target: {
      type: String,
      required: true
    }
  },
  data () {
    return {
      active: false
    }
  },
  computed: {
    offsetTop () {
      const elm = document.getElementById(this.target)
      return elm.getBoundingClientRect().y
    }
  },
  methods: {
    handleScroll () {
      const top = this.offsetTop - window.innerHeight
      if (window.scrollY > top) {
        return (this.active = true)
      }
      return false
    }
  }
}
</script>
