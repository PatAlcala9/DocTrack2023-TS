<template lang="pug">

div.section.fit.column.wrap.justify-start.items-center.content-center
  component(:is="docLabel" :text="props.label")
  input(v-if="wide === false" dark :value="props.value" @input="updateValue(($event.target as HTMLInputElement)?.value)" dense  @blur="$emit('blur')").information
  textarea(v-else dark :value="props.value" @input="updateValue(($event.target as HTMLInputElement)?.value)"  dense type="multiline" @blur="$emit('blur')").wide(v-html="displayMultiline(props.value)")

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
  wide: false,
})

const emit = defineEmits(['update:value', 'blur'])

const updateValue = (value: string) => {
  emit('update:value', value)
}

const displayMultiline = (value: string) => {
  return value.replace(/\r?\n/g, '<br>')
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
  background-color: white
  text-align: center

.wide
  @extend .information
  font-weight: 450
  font-size: 0.9rem
  width: 50vw
  padding: 1rem 2rem
  height: 30vh
  text-align: justify
  padding: 0.4rem 1rem
</style>
