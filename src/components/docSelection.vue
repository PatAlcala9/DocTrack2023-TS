<template lang="pug">

div(v-if="$q.screen.width > 500")
  section.full-width.row.inline.wrap.justify-start.items-start.content-start.section
    section.column
      span.label {{text}}:
      q-select(filled dark :modelValue="props.modelValue" @update:modelValue="updateValue" :label="props.label" :options="options" style="width: 250px" color="blue-10" behavior="menu" label-color="blue-1" item-aligned)

div(v-else)
  section.fit.column.wrap.justify-center.items-center.content-center.section
    section.column.wrap.justify-center.items-center.content-center
      span.label {{text}}
      q-select(filled dark :modelValue="props.modelValue" @update:modelValue="updateValue" :label="props.label" :options="options" style="width: 250px" color="blue-10" behavior="dialog" label-color="blue-1" item-aligned)

</template>

<script setup lang="ts">
import { computed } from 'vue'

export interface Props {
  text: string
  label: string
  modelValue: string
  options: string[]

  width: number
  alignment: string
  transform: string
  mobileWidth: number
  type: string
}

const props = withDefaults(defineProps<Props>(), {
  text: 'Label',
  icon: '',

  value: '',
  width: 16,
  alignment: 'center',
  transform: 'uppercase',
  mobileWidth: 14,
  type: 'input'
})

const emit = defineEmits(['update:value'])

const updateValue = (value: string) => {
  emit('update:modelValue', value)
}

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
  font-weight: 400
  font-size: 1.1rem
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
  backdrop-filter: blur(1.6px) saturate(173%)
  // background-color: rgba(10, 10, 35, 0.52)
  // background-color: rgba(17, 25, 40, 0.82)
  background-color: rgba(17, 25, 40, 0.8)
  border-radius: 0.6rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  padding: 1.2rem 2rem

.icon
  padding: 1rem

.aa
  background-color: red

@media screen and (max-width: 500px)
  .label
    font-weight: 300
    font-size: 1.1rem

  .input
    font-size: 1.1rem
    width: var(--mobileWidth)

  .section
    padding: 1rem 1rem 2rem 1rem
</style>

