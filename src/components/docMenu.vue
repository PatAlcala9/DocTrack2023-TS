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
  background-color: rgba(12, 21, 42, 0.45)
  backdrop-filter: blur(9px) saturate(150%)
  border-radius: 2rem
  border: 3px solid rgba(255, 255, 255, 0.125)
  padding: 1.4rem
  cursor: pointer

.section:hover
  border: 3px solid rgba(255, 255, 255, 1)
  // margin-top: -1rem
  border-radius: 4rem

.icon
  padding: 0 0 0.8rem 0
  cursor: pointer

@media screen and (max-width: 500px)
  .section
    padding: 1.2rem
    margin: 1.2rem
  </style>
