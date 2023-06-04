<template lang="pug">

q-page(padding)
  component(:is="docOnline" :state="online" )

  section(v-if="quasar.screen.width > 500").full-width.column.items-start.content-center
    img(src="../assets/ocbologobw.avif" alt="OCBO Logo").logo

  section(v-else).full-width.column.content-center.items-center
    img(src="../assets/ocbologobw.avif" alt="OCBO Logo").logo

  div.fit.column.text-center.main
    transition(appear @before-enter="beforeEnterTitle" @enter="enterTitle")
      span.name {{_employeename.employeename}}

    transition(appear @before-enter="beforeEnterLogout" @enter="enterLogout")
      q-btn(flat size="md" label="Logout" @click="logout" icon="logout" ).logout-button

  div.button-area
    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      component(:is="docMenu" text="Incoming" icon="description" v-if="_access.access.includes('incoming')" @click="gotoPage('incoming', false)" @mouseover="showDescription")

    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      component(:is="docMenu" text="Outgoing" icon="upload_file" v-if="_access.access.includes('outgoing')" @click="gotoPage('outgoing', false)")

    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      component(:is="docMenu" text="Releasing" icon="start" v-if="_access.access.includes('releasing')" @click="gotoPage('releasing', false)")

    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      component(:is="docMenu" text="Inventory" icon="summarize" v-if="_access.access.includes('inventory')" @click="gotoPage('inventory', false)")

    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      component(:is="docMenu" text="Complaint" icon="gavel" v-if="_access.access.includes('complaint')" @click="gotoPage('complaint', false)")

  </template>

<script setup lang="ts">
import { ref } from 'vue'
import { useQuasar, SessionStorage } from 'quasar'
import { gsap } from 'gsap'
import { useRouter } from 'vue-router'

import docOnline from 'components/docOnline.vue'
import docMenu from 'components/docMenu.vue'

import { useEmployeeName } from 'stores/employeename'
import { useAccess } from 'stores/access'
import { useCurrentPage } from 'stores/currentpage'
import { usePageWithTable } from 'stores/pagewithtable'
import { useIsLogged } from 'stores/islogged'
import { useIsDemo } from 'stores/isdemo'

const router = useRouter()
const quasar = useQuasar()
const _employeename = useEmployeeName()
const _access = useAccess()
const _currentpage = useCurrentPage()
const _pagewithtable = usePageWithTable()
const _islogged = useIsLogged()
const _isdemo = useIsDemo()

let onlineColor = ref('')
let online = ref('OFFLINE')

const setName = async () => {
  if (_employeename.employeename !== '') SessionStorage.set('EmployeeName', _employeename.employeename)
}
const getName = async () => {
  const session: string | null = SessionStorage.getItem('EmployeeName')
  _employeename.employeename = session ?? _employeename.employeename
}

const setAccess = async () => {
  if (_access.access.length > 0) SessionStorage.set('Access', _access.access)
}
const getAccess = async () => {
  const session: string | null = SessionStorage.getItem('Access')
  _access.access = session ? [...session] : []
}

const setDemo = async () => {
  if (_isdemo.isdemo) SessionStorage.set('Demo', _isdemo.isdemo)
}
const getDemo = async () => {
  const session: string | null = SessionStorage.getItem('Demo')
  _isdemo.isdemo = session ? true : false
}

const showDescription = () => {
  console.log('yah')
}

const beforeEnterTitle = (el: any) => {
  el.style.transform = 'translateY(-50px)'
  el.style.opacity = 0
}
const enterTitle = (el: any) => {
  gsap.to(el, { duration: 0.8, y: 0, opacity: 1 })
}

const beforeEnterMenu = (el: any) => {
  el.style.transform = 'translateY(-50px)'
  el.style.opacity = 0
}
const enterMenu = (el: any) => {
  gsap.to(el, { delay: 0.3, duration: 0.8, y: 0, opacity: 1 })
}

const beforeEnterLogout = (el: any) => {
  el.style.transform = 'translateX(30px)'
  el.style.opacity = 0
}
const enterLogout = (el: any) => {
  gsap.to(el, { delay: 0.1, duration: 0.8, x: 0, opacity: 1 })
}

const beforeEnterButton = (el: any) => {
  el.style.transform = 'translateY(40px)'
  el.style.opacity = 0
}
const enterButton = (el: any) => {
  gsap.to(el, { delay: 0.1, duration: 0.8, y: 0, opacity: 1 })
}

const logout = async () => {
  _islogged.islogged = false
  _pagewithtable.pagewithtable = false
  quasar.sessionStorage.remove('EmployeeName')
  _access.access = []
  gotoPage('/')
}

const gotoPage = (page: string, table = false) => {
  _pagewithtable.pagewithtable = table
  _currentpage.currentpage = page
  router.push(page)
}

const checkOnline = () => {
  if (_isdemo.isdemo) {
    onlineColor.value = 'red'
    online.value = 'OFFLINE'
  }
  else {
    onlineColor.value = 'green'
    online.value = 'ONLINE'
  }
}

;(async () => {
  await setName()
  await getName()

  await setAccess()
  await getAccess()

  await setDemo()
  await getDemo()

  if (_currentpage.currentpage !== undefined) router.push(_currentpage.currentpage)
  else router.push('/dashboard')

  checkOnline()
})()
</script>

<style lang="sass" scoped>
.main
  margin-top: 2rem
.logo
  width: 16rem
  height: auto
  position: fixed
  opacity: 0.1
  // filter: invert(100)

.name
  font-size: 3.2rem
  font-family: 'RalewayBold'
  color: #ffffff

.button-area
  display: flex
  flex-wrap: wrap
  justify-content: space-around
  align-items: center
  align-content: flex-end
  padding-top: 30%

.logout-button
  font-family: 'Inter'
  font-weight: bold
  padding: 1.6rem
  position: absolute
  top: 0
  right: 0
  /* margin-top: -5.5rem */


@media screen and (max-width: 500px)
  .name
    font-size: 1.9rem
    margin-top: 2rem
    padding-bottom: 1.2rem

  .button-area
    padding-top: 0
</style>
