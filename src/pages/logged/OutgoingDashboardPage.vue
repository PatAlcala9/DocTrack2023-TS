<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Outgoing
    q-btn(flat size="md" label="Back" @click="gotoMenu" icon="arrow_back").close-button

  div.fit.row.wrap.justify-center.items-start.content-center.button-group
    section.fit.column.wrap.justify-center.items-start.content-center.items-center
      q-icon(name="fact_check" size="3rem")
      component.button(:is="docButton" text="View List" @click="gotoIncomingInquire")

    section.fit.column.wrap.justify-center.items-start.content-center.items-center
      q-icon(name="add_box" size="3rem")
      component.button(:is="docButton" text="Add New" @click="gotoIncomingAdd")

</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useCurrentPage } from 'stores/currentpage'
import docButton from 'components/docButton.vue'

const router = useRouter()
let _currentpage = useCurrentPage()

const gotoIncomingInquire = () => {
  _currentpage.currentpage = 'incominginquire'
  router.push('/incominginquire')
}

const gotoIncomingAdd = () => {
  _currentpage.currentpage = 'incomingadd'
  router.push('/incomingadd')
}


const gotoMenu = () => {
  _currentpage.currentpage = 'dashboard'
  router.push('/dashboard')
}
</script>
