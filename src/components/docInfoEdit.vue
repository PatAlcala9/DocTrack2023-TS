<template lang="pug">

div.section.fit.column.wrap.justify-start.items-center.content-center
  component(:is="docLabel" :text="props.label")
  q-input(v-if="wide === false" dark v-model="props.value" dense).information
  q-input(v-else dark v-model="props.value" dense).wide

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
