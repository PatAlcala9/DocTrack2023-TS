<template lang="pug">

div(v-if="$q.screen.width > 500")
  section.full-width.row.inline.wrap.justify-start.items-start.content-start.section(:style="{ '--colorBackground': colorBackground }")
    section.column
      span.label {{text}}:
      q-select(filled dark :modelValue="props.modelValue" @update:modelValue="updateValue" :options="options" style="width: 260px; margin-left: -1rem" color="blue-10" behavior="menu" label-color="blue-1" item-aligned)

div(v-else)
  section.fit.column.wrap.justify-center.items-center.content-center.section(:style="{ '--colorBackground': colorBackground }")
    section.column.wrap.justify-center.items-center.content-center
      span.label {{text}}
      q-select(filled dark :modelValue="props.modelValue" @update:modelValue="updateValue" :options="options" style="width: 250px" color="blue-10" behavior="dialog" label-color="blue-1" item-aligned)

</template>

<script setup lang="ts">
import { computed } from 'vue'

export interface Props {
  text: string
  modelValue: string
  options: string[]
  alert: boolean
}

const props = withDefaults(defineProps<Props>(), {
  text: 'Label',
  modelValue: '',
  alert: false
})

const emit = defineEmits(['update:value'])

const updateValue = (modelValue: string) => {
  emit('update:modelValue', modelValue)
}

const colorBackground = computed(() => {
  return props.alert ? 'rgba(128, 21, 21, 0.45)' : 'rgba(12, 21, 42, 0.45)'
})

const styleComponent = computed(() => {
  return {
    '--width': props.width + 'rem',
    '--alignment': props.alignment,
    '--transform': props.transform,
    '--mobileWidth': props.mobileWidth,
  }
})
</script>

<style lang="sass" scoped>
.label
  font-family: 'Inter'
  font-weight: 500
  font-size: 0.9rem
  color: #ffffff

.input
  font-family: 'Inter'
  font-size: 1.3rem
  border-radius: 0.6rem
  text-align: var(--alignment)
  text-transform: var(--transform)
  border: 1px solid $button
  color: #000000
  width: var(--width)

.input:focus
  outline: none
  border: 1px solid #ffffff
  background-color: transparent
  color: #ffffff

.section
  backdrop-filter: blur(9px) saturate(150%)
  // background-color: rgba(12, 21, 42, 0.45)
  background-color: var(--colorBackground)
  border-radius: 0.6rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  padding: 1.2rem 2rem

@media screen and (max-width: 500px)
  .label
    margin: 0 0 0.5rem 0

  .input
    font-size: 1.1rem
    width: var(--mobileWidth)

  .section
    // width: 90%
    // padding: 1rem 1rem 2rem 1rem
</style>
