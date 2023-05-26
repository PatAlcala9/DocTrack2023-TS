<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Complaints
    q-btn(flat size="md" label="Back" @click="gotoMenu" icon="arrow_back").close-button

  div.fit.row.wrap.justify-center.items-start.content-center.button-group
    section.fit.column.wrap.justify-center.items-start.content-center.items-center
      q-icon(name="fact_check" size="3rem")
      component.button(:is="docButton" text="View List" @click="gotoComplaintInquire")

    section.fit.column.wrap.justify-center.items-start.content-center.items-center
      q-icon(name="add_box" size="3rem")
      component.button(:is="docButton" text="Add New" @click="gotoComplaintAdd")

  </template>

  <script setup lang="ts">
import { useRouter } from 'vue-router'
import { useCurrentPage } from 'stores/currentpage'
import { usePageWithTable } from 'stores/pagewithtable';
import docButton from 'components/docButton.vue'

const router = useRouter()
let _currentpage = useCurrentPage()
let _pagewithtable = usePageWithTable()

const gotoComplaintInquire = () => {
  _currentpage.currentpage = 'complaintmain'
  router.push('/complaintmain')
}

const gotoComplaintAdd = () => {
  _pagewithtable.pagewithtable = true
  _currentpage.currentpage = 'complaintadd'
  router.push('/complaintadd')
}

const gotoMenu = () => {
  _currentpage.currentpage = 'dashboard'
  router.push('/dashboard')
}
</script>

<style lang="sass" scoped>
.button-group
  margin-top: 8.5rem

.button
  padding: 1.6rem
</style>
