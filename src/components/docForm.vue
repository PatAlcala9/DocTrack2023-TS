<template lang="pug">

div(v-if="$q.screen.width > 500")
  label(:for="text").full-width.row.inline.wrap.justify-start.items-start.content-start.section
    q-icon(:name="icon" size="3rem").icon
    section.column
      span.label {{text}}:
      input(:value="props.value" @input="updateValue(($event.target as HTMLInputElement)?.value)" :style="styleComponent" :id="text" :type="type").input

div(v-else)
  section.full-width.column.inline.wrap.justify-center.items-center.content-center.section
    q-icon(:name="icon" size="3rem").icon
    section.column.wrap.justify-center.items-center.content-center
      span.label {{text}}
      input(:value="props.value" @input="updateValue(($event.target as HTMLInputElement)?.value)" :style="styleComponent" :type="type").input
</template>

<script setup lang="ts">
import { computed } from 'vue'

export interface Props {
  text: string
  icon: string

  value: string
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
  mobileWidth: 16,
  type: 'input'
})

const emit = defineEmits(['update:value'])

const updateValue = (value: string) => {
  // const newValue = target.value
  emit('update:value', value)
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
  backdrop-filter: blur(16px) saturate(180%)
  background-color: rgba(10, 10, 35, 0.52)
  border-radius: 0.6rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  padding: 1.2rem 2rem 1.2rem 1.2rem

.icon
  padding: 1rem

@media screen and (max-width: 500px)
  .label
    font-weight: 300
    font-size: 1.4rem

  .input
    font-size: 1.2rem
    width: var(--mobileWidth)

  .section
    padding: 1rem 1rem 2rem 1rem
</style>
