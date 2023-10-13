<template lang="pug">

div(v-if="$q.screen.width > 500")
  section.full-width.row.inline.wrap.justify-start.items-start.content-start.section(:style="{ '--colorBackground': colorBackground }")
    q-icon(:name="icon" size="1.6rem").icon
    section.column
      span.label {{ text }}:
      q-option-group(dark :modelValue="props.modelValue" @update:modelValue="updateValue" :options="props.options" color="blue-10" type="checkbox" @keypress="keypressEvent").list

div(v-else)
  section.full-width.column.inline.wrap.justify-center.items-center.content-center.section(:style="{ '--colorBackground': colorBackground }")
    q-icon(:name="icon" size="1.8rem").icon
    section.column.wrap.justify-center.items-center.content-center.text-start
      span.label {{ text }}
      q-option-group(dark :modelValue="props.modelValue" @update:modelValue="updateValue" :options="props.options" color="blue-10" type="checkbox" @keypress="keypressEvent").list
</template>

<script setup lang="ts">
import { computed } from 'vue'

export interface Props {
  text: string
  icon: string
  modelValue: string[]
  options: Array<{ label: string; value: string }>
  alert: boolean
  keypressEvent: () => void
}

const props = withDefaults(defineProps<Props>(), {
  alert: false
})
const emit = defineEmits(['update:modelValue'])

const updateValue = (value: string[]) => {
  emit('update:modelValue', value)
}

const colorBackground = computed(() => {
  return props.alert ? 'rgba(128, 21, 21, 0.45)' : 'rgba(12, 21, 42, 0.45)'
})

// const styleComponent = computed(() => {
//   return {
//     '--colorBackground': colorBackground.value,
//   }
// })
</script>

<style lang="sass" scoped>
.label
  font-family: 'Inter'
  font-weight: 400
  font-size: 0.9rem
  color: #ffffff
  margin: 0 0 0 -2rem

.section
  // background-color: rgba(12, 21, 42, 0.45)
  background-color: var(--colorBackground)
  backdrop-filter: blur(9px) saturate(150%)
  border-radius: 0.6rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  padding: 1rem 1.6rem 1rem 0


.list
  width: 30rem
  font-size: 0.8rem
  margin: 0 0 0 -2rem
.icon
  padding: 1rem

@media screen and (max-width: 500px)
  .label
    font-weight: 300
    font-size: 1.2rem

  .section
    padding: 0.3rem 0.8rem 1.4rem 0.8rem

  .list
    width: 16rem
</style>
