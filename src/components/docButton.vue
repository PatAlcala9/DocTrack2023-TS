<template lang="pug">

div
  q-btn(outline padding="0.2rem 2rem" rounded color="white" :label="text" :size="($q.screen.width < 500) ? 'lg' : 'md'" @click="('click')" ).button

</template>

<script setup lang="ts">
export interface Props {
  text: string
}

const props = withDefaults(defineProps<Props>(), {
  text: 'button'
})
</script>

<style lang="sass" scoped>
.button
  font-family: 'Inter'
  font-weight: 100 200
</style>
