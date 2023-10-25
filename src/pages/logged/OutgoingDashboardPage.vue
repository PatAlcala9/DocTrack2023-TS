<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Outgoing
    q-btn(flat size="md" label="Back" @click="gotoMenu" icon="arrow_back").close-button

  //- div.fit.row.wrap.justify-center.items-start.content-center.button-group
  //-   section.fit.column.wrap.justify-center.items-start.content-center.items-center
  //-     q-icon(name="fact_check" size="3rem")
  //-     component.button(:is="docButton" text="View List" @click="gotoIncomingInquire")

  //-   section.fit.column.wrap.justify-center.items-start.content-center.items-center
  //-     q-icon(name="add_box" size="3rem")
  //-     component.button(:is="docButton" text="Add New" @click="gotoIncomingAdd")

  div.button-group
    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      component(:is="docMenu" text="View List" icon="fact_check" @click="gotoOutgoingInquire")

    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      component(:is="docMenu" text="Add New" icon="add_box" @click="gotoOutgoingAdd")

</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useCurrentPage } from 'stores/currentpage'
import { ref } from 'vue'
import { usePageWithTable } from 'stores/pagewithtable'
import { useIsDemo } from 'stores/isdemo'
import { gsap } from 'gsap'

import docButton from 'components/docButton.vue'
import docMenu from 'components/docMenu.vue'

const router = useRouter()
const _currentpage = useCurrentPage()
const _pagewithtable = usePageWithTable()
const _isdemo = useIsDemo()

const gotoOutgoingInquire = () => {
  _currentpage.currentpage = 'outgoinginquire'
  router.push('/outgoinginquire')
}

const gotoOutgoingAdd = () => {
  _currentpage.currentpage = 'outgoingadd'
  router.push('/outgoingadd')
}

const gotoMenu = () => {
  _currentpage.currentpage = 'dashboard'
  router.push('/dashboard')
}

const beforeEnterButton = (el: any) => {
  el.style.transform = 'translateY(40px)'
  el.style.opacity = 0
}
const enterButton = (el: any) => {
  gsap.to(el, { delay: 0.1, duration: 0.8, y: 0, opacity: 1 })
}
</script>

<style lang="sass" scoped>
.button-group
  display: flex
  flex-wrap: wrap
  justify-content: space-around
  align-items: center
  align-content: flex-end
  height: calc(100vh - 10rem)

</style>
