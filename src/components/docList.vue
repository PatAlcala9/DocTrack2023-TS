<template lang="pug">

div(v-if="$q.screen.width > 500")
  section.full-width.row.inline.wrap.justify-start.items-start.content-start.section
    q-icon(:name="icon" size="2.5rem").icon
    section.column
      span.label {{ text }}:
      q-option-group(dark :modelValue="props.modelValue" @update:modelValue="updateValue" :options="props.options" color="blue-10" type="checkbox").list

div(v-else)
  section.full-width.column.inline.wrap.justify-center.items-center.content-center.section
    q-icon(:name="icon" size="3rem").icon
    section.column.wrap.justify-center.items-center.content-center.text-start
      span.label {{ text }}
      q-option-group(dark :modelValue="props.modelValue" @update:modelValue="updateValue" :options="props.options" color="blue-10" type="checkbox").list
</template>

<script setup lang="ts">
import { computed } from 'vue'

export interface Props {
  text: string
  icon: string
  modelValue: string[]
  options: Array<{ label: string; value: string }>

  // width: number
  // alignment: string
  // mobileWidth: number
}

const props = defineProps<Props>()
const emit = defineEmits(['update:modelValue'])

const updateValue = (value: string[]) => {
  emit('update:modelValue', value)
}

// const styleComponent = computed(() => {
//   return {
//     '--width': props.width + 'rem',
//     '--alignment': props.alignment,
//     '--mobileWidth': props.mobileWidth,
//   }
// })
</script>

<style lang="sass" scoped>
.label
  font-family: 'Inter'
  font-weight: 400
  font-size: 1.1rem
  color: #ffffff

.section
  backdrop-filter: blur(16px) saturate(173%)
  background-color: rgba(17, 25, 40, 0.65)
  border-radius: 0.6rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  padding: 1.2rem 2rem 1.2rem 1.2rem


.list
  width: 30rem
.icon
  padding: 1rem

@media screen and (max-width: 500px)
  .label
    font-weight: 300
    font-size: 1.4rem

  .section
    padding: 1rem 1rem 2rem 1rem

  .list
    width: 16rem
</style>
