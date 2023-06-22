<template lang="pug">

div.section.fit.column.wrap.justify-start.items-center.content-center
  component(:is="docLabel" :text="props.label")
  span(v-if="wide === false").information {{ props.value }}
  span(v-else).wide {{ props.value }}

</template>

<script setup lang="ts">
import docLabel from 'components/docLabel.vue'

export interface Props {
  label: string
  value: string
  wide: boolean
}

const props = withDefaults(defineProps<Props>(), {
  label: 'Label',
  value: 'Value',
  wide: false
})
const emit = defineEmits(['update:value'])

</script>

<style lang="sass" scoped>
.information
  font-family: 'Inter'
  font-weight: 560
  font-size: 1.4rem
  border: 1px solid #ffffff
  border-radius: 3rem
  padding: 0.4rem 1.5rem
  background-color: transparent

.wide
  @extend .information
  font-weight: 400
  font-size: 1.1rem
  width: 85vw
  text-align: justify
</style>
