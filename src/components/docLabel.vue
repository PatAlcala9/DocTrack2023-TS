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
  font-weight: 500
  font-size: 1.1rem
  color: #ffffff
</style>
