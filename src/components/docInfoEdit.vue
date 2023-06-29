<template lang="pug">

div.section.fit.column.wrap.justify-start.items-center.content-center
  component(:is="docLabel" :text="props.label")
  input(v-if="wide === false" dark v-model="props.value" dense).information
  input(v-else dark v-html="displayMultiline(props.value)" dense type="multiline").wide

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

const displayMultiline = (value: string) => {
  return value.replace(/\r?\n/g, '<br>');
}
</script>

<style lang="sass" scoped>
.information
  font-family: 'Inter'
  font-weight: 560
  font-size: 1.4rem
  border: 1px solid #ffffff
  border-radius: 3rem
  padding: 0.4rem 0rem
  background-color: white
  text-align: center

.wide
  @extend .information
  font-weight: 450
  font-size: 1.2rem
  width: auto
  text-align: justify
  padding: 0.4rem 1.5rem
</style>
