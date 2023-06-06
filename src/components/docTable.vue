<template lang="pug">

div
  table.table
    thead
      tr
        th Complain Code
        th Type
        th Name
        th Status
        th Remaining Days
        th Details
    tbody
      tr(v-for="(item, index) in content" :key="item").table-content-group
        td {{item}}
        td {{content[index]}}
        td {{content[index]}}
        td {{content[index]}}
        td {{content[index]}}
        td
          q-btn(color="button" icon="visibility" :ripple="false" @click="('click')").button-view

</template>

<script setup lang="ts">
export interface Props {
  content: string[]
}

const props = withDefaults(defineProps<Props>(), {
  text: 'button'
})
</script>
