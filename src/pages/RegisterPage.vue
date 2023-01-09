<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    transition(appear @before-enter="beforeEnterTitle" @enter="enterTitle")
      span.title Registration
    transition(appear @before-enter="beforeEnterClose" @enter="enterClose")
      q-btn(flat size="md" label="Back" @click="gotoHome" icon="arrow_back" ).close-button

  //- transition(appear @before-enter="beforeEnterPage" @enter="enterPage")
  section.login.column.items-center
    transition(appear @before-enter="beforeEnterInputs" @enter="enterInputs")
      div.column.items-center
        span.login__username--label Fullname
        component(:is="docInput" v-model:value="ifullname" width="30").login__username--input

        span.login__username--label Username
        component(:is="docInput" v-model:value="iusername").login__username--input

        span.login__username--label Password
        component(:is="docInputPassword" v-model:value="ipassword").login__username--input

        span.login__username--label Access
        q-option-group(dark v-model="accessList" :options="accessOption" color="indigo-9" type="checkbox").login__username--option
    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      doc-button(text="Register" @click="saveAccount")

q-dialog(v-model="dialog" transition-show="flip-right" transition-hide="flip-left")
  q-card.dialog-card.text-white
    q-card-section.dialog-card__section.flex.flex-center
      div.dialog-title-area.row.justify-between
        span.dialog-title {{dialogMessage}}
        component(:is="docButton" text="OK" v-close-popup)

</template>

<!-- <script lang="ts">
import { useCurrentPage } from 'stores/currentpage'

export default {
  preFetch({ redirect }) {
    let _currentpage = useCurrentPage()

    if (_currentpage.currentpage !== 'register') {
      redirect({ path: '/login' })
    }
  },
}
</script> -->

<script setup lang="ts">
import { api } from 'boot/axios'
import { ref } from 'vue'
import { encrypt } from 'src/js/OCBO'
import { useRouter } from 'vue-router'
import { gsap } from 'gsap'

import { usePageWithTable } from 'stores/pagewithtable'
import { useCurrentPage } from 'stores/currentpage'

import docInput from 'components/docInput.vue'
import docInputPassword from 'components/docInputPassword.vue'
import docButton from 'components/docButton.vue'

const router = useRouter()
const _pagewithtable = usePageWithTable()
const _currentpage = useCurrentPage()

const beforeEnterTitle = (el: any) => {
  el.style.transform = 'translateX(-100px)'
  el.style.opacity = 0
}
const enterTitle = (el: any) => {
  gsap.to(el, { duration: 0.8, x: 0, opacity: 1 })
}

const beforeEnterClose = (el: any) => {
  el.style.transform = 'translateX(100px)'
  el.style.opacity = 0
}
const enterClose = (el: any) => {
  gsap.to(el, { duration: 0.8, x: 0, opacity: 1 })
}

const beforeEnterInputs = (el: any) => {
  el.style.transform = 'translateY(-90px)'
  el.style.opacity = 0
}
const enterInputs = (el: any) => {
  gsap.to(el, { delay: 0.1, duration: 1, y: 0, opacity: 1 })
}

const beforeEnterButton = (el: any) => {
  el.style.transform = 'translateY(90px)'
  el.style.opacity = 0
}
const enterButton = (el: any) => {
  gsap.to(el, { duration: 1, y: 0, opacity: 1 })
}

let ifullname = ref('')
let iusername = ref('')
let ipassword = ref('')
let ipasswordEncrypted = ref('')
let accessList: any = ref([])
let accessOption = ref([
  {
    label: 'Incoming',
    value: 'is_incoming',
  },
  {
    label: 'Outgoing',
    value: 'is_outgoing',
  },
  {
    label: 'Releasing',
    value: 'is_releasing',
  },
  {
    label: 'Inventory',
    value: 'is_inventory',
  },
  {
    label: 'Other Documents',
    value: 'is_otherdocuments',
  },
])
let dialog = ref(false)
let dialogMessage = ref('')

const saveAccount = async () => {
  let ipasswordEncrypted = await encrypt(ipassword.value.toUpperCase())
  let iincoming = accessList.value.includes('is_incoming') ? 1 : 0
  let ioutgoing = accessList.value.includes('is_outgoing') ? 1 : 0
  let ireleasing = accessList.value.includes('is_releasing') ? 1 : 0
  let iinventory = accessList.value.includes('is_inventory') ? 1 : 0
  let iothers = accessList.value.includes('is_otherdocuments') ? 1 : 0
  const response = await api.post('/api/SaveAccount/' + ifullname.value.toUpperCase() + '/' + iusername.value.toUpperCase() + '/' + ipasswordEncrypted + '/' + iincoming + '/' + ioutgoing + '/' + ireleasing + '/' + iinventory + '/' + iothers)
  const data = response.data

  if (data.includes('Success')) {
    dialog.value = true
    dialogMessage.value = 'Successfully Registered'
  } else {
    dialog.value = true
    dialogMessage.value = 'Failed to Register'
  }
}

const sample = () => {
  dialog.value = true
  dialogMessage.value = 'Sample'
}

const gotoHome = () => {
  _pagewithtable.pagewithtable = false
  _currentpage.currentpage = '/'
  router.push('/')
}

;(async () => {
  if (_currentpage.currentpage !== undefined) router.push(_currentpage.currentpage)
  else router.push('/register')
})()
</script>

<style lang="sass" scoped>

.dialog
  font-family: 'Raleway'
  background-color: $button

.detail-dialog__info
  font-size: 2rem

.login
  margin: 1rem
  font-family: 'Montserrat'

.login__username--label
  font-family: 'Montserrat'
  font-size: 1.4rem

.login__username--input
  margin-bottom: 1.4rem

.login__username--option
  @extend .login__username--input
  text-transform: capitalize

.dialog-card
  font-family: "OpenSans"
  background-color: rgba(2,25,38, 0.7)
  width: 50%
  height: 50%

.dialog-card__section
  padding: 2rem
  font-size: 2rem
</style>
