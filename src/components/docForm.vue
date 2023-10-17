<template lang="pug">

div(v-if="$q.screen.width > 500")
  label(:for="text").full-width.row.inline.wrap.justify-start.items-start.content-start.section(:style="{ '--colorBackground': colorBackground }")
    q-icon(:name="icon" size="1.6rem").icon
    section.column
      span.label {{text}}:
      input(:value="props.value" @input="updateValue(($event.target as HTMLInputElement)?.value)" :style="styleComponent" :id="text" :type="type" @keypress="keypressEvent").input

div(v-else)
  section.fit.column.wrap.justify-center.items-center.content-center.section(:style="{ '--colorBackground': colorBackground }")
    q-icon(:name="icon" size="1.8rem").icon
    section.column.wrap.justify-center.items-center.content-center
      span.label {{text}}
      input(:value="props.value" @input="updateValue(($event.target as HTMLInputElement)?.value)" :style="styleComponent" :type="type" @keypress="keypressEvent").input

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
  alert: boolean
  keypressEvent: () => void
}

const props = withDefaults(defineProps<Props>(), {
  text: 'Label',
  icon: '',

  value: '',
  width: 16,
  alignment: 'center',
  transform: 'uppercase',
  mobileWidth: 14,
  type: 'input',
  alert: false,
})

const emit = defineEmits(['update:value'])

const updateValue = (value: string) => {
  emit('update:value', value)
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
    '--colorBackground': colorBackground.value,
  }
})
</script>

<style lang="sass" scoped>
.label
  font-family: 'Inter'
  font-weight: 300
  font-size: 0.9rem
  color: #ffffff

.input
  font-family: 'Inter'
  font-size: 1.1rem
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
  // backdrop-filter: blur(9px) saturate(173%)
  // background-color: rgba(10, 10, 35, 0.52)
  // background-color: rgba(17, 25, 40, 0.82)
  // background-color: rgba(17, 25, 40, 0.8)
  backdrop-filter: blur(9px) saturate(150%)
  // background-color: rgba(12, 21, 42, 0.45)
  background-color: var(--colorBackground)
  border-radius: 0.6rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  padding: 0.8rem 1.4rem 0.9rem 0.1rem

.icon
  padding: 0.9rem

@media screen and (max-width: 500px)
  .label
    font-weight: 300
    font-size: 1rem

  .input
    font-size: 1.1rem
    width: var(--mobileWidth)

  .section
    padding: 0.1rem 0.8rem 1.1rem 0.8rem
</style>
