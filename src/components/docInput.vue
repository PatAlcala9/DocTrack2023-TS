<template lang="pug">

div
  input(:value="props.value" @input="updateValue($event.target.value)" :style="widthComponent").input

</template>

<script setup lang="ts">
import { computed } from 'vue'

export interface Props {
  value: string
  width: number
}

const props = defineProps<Props>()
const emit = defineEmits(['update:value'])

const updateValue = (value: string) => {
  emit('update:value', value)
}

const widthComponent = computed(() => {
  return {
    '--width': props.width + 'rem'
  }
})
</script>

<style lang="sass" scoped>
.input
  font-family: 'Raleway'
  font-size: 1.3rem
  border-radius: 0.6rem
  text-align: center
  text-transform: uppercase
  border: 1px solid $button
  color: #000000
  width: var(--width)

.input:focus
  outline: none
  border: 1px solid #ffffff
  background-color: #274c62
  color: #ffffff
</style>
