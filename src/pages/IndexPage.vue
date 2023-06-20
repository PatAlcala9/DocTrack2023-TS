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

  div(v-if="$q.screen.width > 500").message-area
    span {{ menuMessage }}

  div.button-area.flex.flex-center
    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      component(:is="docMenu" text="Incoming" icon="description" v-if="_access.access.includes('incoming')" @click="gotoPage('incoming', false)" @mouseover="setMessage('incoming')" @mouseout="returnDefault")

    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      component(:is="docMenu" text="Outgoing" icon="upload_file" v-if="_access.access.includes('outgoing')" @click="gotoPage('outgoing', false)" @mouseover="setMessage('outgoing')" @mouseout="returnDefault")

    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      component(:is="docMenu" text="Releasing" icon="start" v-if="_access.access.includes('releasing')" @click="gotoPage('releasing', false)" @mouseover="setMessage('releasing')" @mouseout="returnDefault")

    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      component(:is="docMenu" text="Inventory" icon="summarize" v-if="_access.access.includes('inventory')" @click="gotoPage('inventory', false)" @mouseover="setMessage('inventory')" @mouseout="returnDefault")

    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      component(:is="docMenu" text="Complaint" icon="gavel" v-if="_access.access.includes('complaint')" @click="gotoPage('complaint', false)" @mouseover="setMessage('complaint')" @mouseout="returnDefault")

  </template>

<script setup lang="ts">
import { ref } from 'vue'
import { useQuasar, SessionStorage } from 'quasar'
import { gsap } from 'gsap'
import { useRouter } from 'vue-router'
import { checkConnection } from 'src/js/functions'

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

let timer: NodeJS.Timeout

let onlineColor = ref('')
let online = ref('OFFLINE')
let menuMessage = ref('')
const defaultMessage = 'Please make a selection from the menu'

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
  SessionStorage.set('Demo', _isdemo.isdemo)
}
const getDemo = async () => {
  const session: string | null = SessionStorage.getItem('Demo')
  _isdemo.isdemo = session ? true : false
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
  clearInterval(timer)
  gotoPage('/')
}

const gotoPage = (page: string, table = false) => {
  _pagewithtable.pagewithtable = table
  _currentpage.currentpage = page
  router.push(page)
}

const checkOnline = async () => {
  await getDemo()

  if (_isdemo.isdemo) {
    // onlineColor.value = 'white'
    online.value = 'DEMO MODE'
  } else {
    if (await checkConnection()) {
      // onlineColor.value = 'green'
      online.value = 'ONLINE'
    } else {
      // onlineColor.value = 'red'
      online.value = 'OFFLINE'
    }

  }
}

const setDefault = async () => {
  menuMessage.value = defaultMessage
}

const setMessage = (link: string) => {
  switch (link) {
    case 'incoming':
      menuMessage.value = "Government's internal document record"
      break
    case 'outgoing':
      menuMessage.value = 'I have no idea what this is'
      break
    case 'releasing':
      menuMessage.value = 'Release the Kraken,'
      break
    case 'inventory':
      menuMessage.value = 'A list of Inventions, I guess'
      break
    case 'complaint':
      menuMessage.value = 'Listahan sa mga reklamador'
      break
    default:
      break
  }
  rotateMessage()
}
const returnDefault = () => {
  setDefault()
  rotateMessageBack()
}

const rotateMessage = () => {
  gsap.fromTo('.message-area', { rotationY: 0 }, { rotationY: 360, duration: 0.9 })
}
const rotateMessageBack = () => {
  gsap.fromTo('.message-area', { rotationY: 0 }, { rotationY: -360, duration: 0.9 })
}

;(async () => {
  await setDefault()

  await setName()
  await getName()

  await setAccess()
  await getAccess()

  // await setDemo()
  // await getDemo()

  if (_currentpage.currentpage !== undefined) router.push(_currentpage.currentpage)
  else router.push('/dashboard')

  await checkOnline()
  timer = setInterval(checkOnline, 5000)
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
  height: calc(100vh - 10rem)

.logout-button
  font-family: 'Inter'
  font-weight: bold
  padding: 1.6rem
  position: absolute
  top: 0
  right: 0
  /* margin-top: -5.5rem */

.message-area
  position: absolute
  top: 50%
  left: 50%
  transform: translate(-50%, -50%)
  height: 100vh
  display: flex
  justify-content: center
  align-items: center

.message-area span
  font-family: 'Inter'
  font-weight: 300
  font-size: 1.6rem
  padding: 5rem
  background-color: rgba(12, 21, 42, 0.45)
  backdrop-filter: blur(9px) saturate(150%)
  border-radius: 2rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  display: inline-block
  white-space: nowrap

@media screen and (max-width: 500px)
  .name
    font-size: 1.9rem
    margin-top: 2rem
    padding-bottom: 1.2rem

  .button-area
    padding-top: 0
</style>
