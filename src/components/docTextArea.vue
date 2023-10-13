<template lang="pug">

div
  textarea(:value="props.value" @input="updateValue($event.target.value)" :style="widthComponent").textarea

</template>

<script setup lang="ts">
import { computed } from 'vue'

export interface Props {
  value: string
  width: number
  height: number
}

const props = defineProps<Props>()
const emit = defineEmits(['update:value'])

const updateValue = (value: string) => {
  emit('update:value', value)
}

const widthComponent = computed(() => {
  return {
    '--width': props.width + '%',
    '--height': props.height + 'rem',
  }
})
</script>

<style lang="sass" scoped>
.textarea
  font-family: 'OpenSans'
  font-size: 1rem
  border-radius: 0.6rem
  text-align: left
  text-transform: initial
  border: 1px solid $button
  color: #000000
  width: var(--width)
  height: var(--height)
  padding: 0.1rem 1rem 0.1rem 1rem

.textarea:focus
  outline: none
  border: 1px solid #ffffff
  //background-color: #274c62
  background-color: transparent
  color: #ffffff
</style>
