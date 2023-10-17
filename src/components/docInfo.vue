<template lang="pug">

div(v-if="left === false").section.fit.column.wrap.justify-start.items-center.content-center
  component(:is="docLabel" :text="props.label")
  span(v-if="wide === false").information {{ props.value }}
  span(v-else v-html="displayMultiline(props.value)").wide

div(v-else).section.fit.column.wrap.justify-start.items-start.content-center
  component(:is="docLabel" :text="props.label + ':'")
  span(v-if="wide === false").information {{ props.value }}
  span(v-else v-html="displayMultiline(props.value)").wide

</template>

<script setup lang="ts">
import docLabel from 'components/docLabel.vue'

export interface Props {
  label: string
  value: string
  wide: boolean
  left: boolean
}

const props = withDefaults(defineProps<Props>(), {
  label: 'Label',
  value: 'Value',
  wide: false,
  left: false
})
const emit = defineEmits(['update:value'])

const displayMultiline = (value: string) => {
  return value.replace(/\r?\n/g, '<br>');
}
</script>

<style lang="sass" scoped>
.information
  font-family: 'Inter'
  font-weight: 560
  font-size: 0.9rem
  border: 1px solid #ffffff
  border-radius: 0.4rem
  padding: 0.4rem 1rem
  background-color: rgba(0, 0, 0, 0.8)

.wide
  @extend .information
  font-weight: 450
  font-size: 0.9rem
  width: 50vw
  height: auto
  padding: 0.5rem 1rem
  text-align: justify
</style>
