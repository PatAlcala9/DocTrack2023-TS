<template lang="pug">

div.full-width.column.wrap.justify-start.items-start.content-start
  div(v-for="version in versions" :key="version")
    span.version {{ version }}
    span.info(v-for="(info, index) in infos[version]" :key="info" v-html="formatInfo(info, 0)")
</template>

<script setup lang="ts">
const versions = ['0.1', '0.2', '0.3']
const infos = {
  '0.1': ['1', '2'],
  '0.2': ['3', '4'],
  '0.3': ['5', '6'],
}
const formatInfo = (info: string, isFirst: number) => {
  if (isFirst) {
    return info
  } else {
    return '<br>- ' + info
  }
}
</script>

<style lang="sass">
.version
  font-size: 2rem

.info
  font-size: 1.2rem
</style>
