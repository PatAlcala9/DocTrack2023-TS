<template lang="pug">

q-page(padding)
  section.online
    q-icon(name="circle" size="1rem" :color="onlineColor")

  div.full-width.row.justify-between
    span.title Complaints
    q-btn(flat size="md" label="Back" @click="gotoMenu" icon="arrow_back").close-button


  div.button-group
    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      component(:is="docMenu" text="View List" icon="fact_check" @click="gotoComplaintInquire")

    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      component(:is="docMenu" text="Add New" icon="add_box" @click="gotoComplaintAdd")
      //- section.fit.column.wrap.justify-center.items-start.content-center.items-center
      //-     q-icon(name="add_box" size="3rem")
      //-     component.button(:is="docButton" text="Add New" )

  </template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useCurrentPage } from 'stores/currentpage'
import { usePageWithTable } from 'stores/pagewithtable'
import { useIsDemo } from 'stores/isdemo'
import { gsap } from 'gsap'

import docButton from 'components/docButton.vue'
import docMenu from 'components/docMenu.vue'

const router = useRouter()
const _currentpage = useCurrentPage()
const _pagewithtable = usePageWithTable()
const _isdemo = useIsDemo()

let onlineColor = ref('')

const checkOnline = () => {
  if (_isdemo.isdemo) onlineColor.value = 'red'
  else {
    onlineColor.value = 'green'
  }
}

const beforeEnterButton = (el: any) => {
  el.style.transform = 'translateY(40px)'
  el.style.opacity = 0
}
const enterButton = (el: any) => {
  gsap.to(el, { delay: 0.1, duration: 0.8, y: 0, opacity: 1 })
}

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

;(async () => {
  if (_currentpage.currentpage !== undefined) router.push(_currentpage.currentpage)
  else router.push('/dashboard')

  checkOnline()
})()
</script>

<style lang="sass" scoped>
.button-group
  display: flex
  flex-wrap: wrap
  justify-content: space-around
  align-items: center
  align-content: flex-end
  padding-top: 30%
  margin-top: 8.5rem

.button
  padding: 1.6rem
</style>
