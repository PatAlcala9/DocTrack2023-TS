<template lang="pug">

div
  section(@click="('click')").full-width.column.inline.wrap.justify-start.items-center.content-start.section
    q-icon(:name="icon" :size="($q.screen.width < 500) ? 'md' : 'lg'").icon
    q-btn(outline padding="0.2rem 2rem" rounded color="white" :label="text" :size="($q.screen.width < 500) ? 'lg' : 'md'").button

</template>

<script setup lang="ts">
export interface Props {
  text: string
  icon: string
}

const props = withDefaults(defineProps<Props>(), {
  text: 'button'
})
</script>

<style lang="sass" scoped>
.button
  font-family: 'Inter'
  font-weight: 100 200

.section
  backdrop-filter: blur(1.6px) saturate(173%)
  background-color: rgba(17, 25, 40, 0.8)
  border-radius: 2rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  padding: 2rem
  cursor: pointer

.icon
  padding: 0 0 0.8rem 0
  cursor: pointer

@media screen and (max-width: 500px)
  .section
    padding: 1.2rem
    margin: 1.2rem
  </style>
