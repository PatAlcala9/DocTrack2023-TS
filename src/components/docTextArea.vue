<template lang="pug">

div
  textarea(:value="props.value" @input="updateValue($event.target.value)" :style="widthComponent").textarea

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
.textarea
  font-family: 'OpenSans'
  font-size: 1.1rem
  border-radius: 0.6rem
  text-align: left
  text-transform: initial
  border: 1px solid $button
  color: #000000
  width: var(--width)

.textarea:focus
  outline: none
  border: 1px solid #ffffff
  background-color: #274c62
  color: #ffffff
</style>
