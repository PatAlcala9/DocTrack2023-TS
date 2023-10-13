<template lang="pug">

div
  input(:value="props.value" @input="updateValue($event.target.value)" :style="styleComponent").input

</template>

<script setup lang="ts">
import { computed } from 'vue'

export interface Props {
  value: string
  width: number
  alignment: string
  transform: string
  mobileWidth: number
}

const props = withDefaults(defineProps<Props>(), {
  value: '',
  width: 16,
  alignment: 'center',
  transform: 'uppercase',
  mobileWidth: 16,
})
const emit = defineEmits(['update:value'])

const updateValue = (value: string) => {
  emit('update:value', value)
}

const styleComponent = computed(() => {
  return {
    '--width': props.width + '%',
    '--alignment': props.alignment,
    '--transform': props.transform,
    '--mobileWidth': props.mobileWidth,
  }
})
</script>

<style lang="sass" scoped>
.input
  font-family: 'OpenSans'
  font-size: 1rem
  border-radius: 0.6rem
  text-align: var(--alignment)
  text-transform: var(--transform)
  border: 1px solid $button
  color: #000000
  width: var(--width)
  padding: 0.1rem 1rem 0.1rem 1rem

.input:focus
  outline: none
  border: 1px solid #ffffff
  //background-color: #274c62
  background-color: transparent
  color: #ffffff

@media screen and (max-width: 500px)
  .input
    font-size: 1.2rem
    width: var(--mobileWidth)
</style>
