<template lang="pug">

div
  span.label {{text}}

</template>

<script setup lang="ts">
export interface Props {
  text: string
}

const props = defineProps<Props>()
</script>

<style lang="sass" scoped>
.label
  font-family: 'Inter'
  font-weight: bold
  font-size: 1.2rem
  color: #ffffff
</style>
