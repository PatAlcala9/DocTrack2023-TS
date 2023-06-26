<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    transition(appear @before-enter="beforeEnterTitle" @enter="enterTitle")
      span.title Registration
    transition(appear @before-enter="beforeEnterClose" @enter="enterClose")
      q-btn(flat size="md" label="Back" @click="gotoHome" icon="arrow_back" ).close-button

  section.login.column.items-center
    transition(appear @before-enter="beforeEnterInputs" @enter="enterInputs")
      div.column.items-center
        component(:is="docForm" text="Fullname" v-model:value="ifullname" :width=30 :mobileWidth=14 icon="badge" :alert="redFullname" :keypressEvent="checkCompleteFullname").login__username--input
        component(:is="docForm" text="Username" v-model:value="iusername" :width=30 :mobileWidth=14 icon="person" :alert="redUsername" :keypressEvent="checkCompleteUsername").login__username--input
        component(:is="docForm" text="Password" v-model:value="ipassword" :width=30 :mobileWidth=14 icon="lock" type="password" :alert="redPassword" :keypressEvent="checkCompletePassword").login__username--input
        component(:is="docForm" text="Confirm Password" v-model:value="icpassword" :width=30 :mobileWidth=14 icon="lock" type="password" :alert="redCPassword" :keypressEvent="checkCompleteCPassword").login__username--input
        component(:is="docList" text="Access" v-model:modelValue="accessList" icon="format_list_bulleted" :options="accessOption" :alert="redAccess" :keypressEvent="checkCompleteAccess").login__username--input

        //- span.login__username--label Access
        //- q-option-group(dark v-model="accessList" :options="accessOption" color="indigo-9" type="checkbox").login__username--option
    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      doc-button.register-button(text="Register" @click="saveAccount")

q-dialog(v-model="dialog" transition-show="flip-right" transition-hide="flip-left").dialog
  q-card.dialog-card.text-white.flex.flex-center
    q-card-section.dialog-card__section
      div.dialog-title-area.column.justify-center.items-center
        span.dialog-card__title {{dialogTitle}}
        span.dialog-card__info {{dialogMessage}}
        component(:is="docButton" text="OK" v-close-popup)

</template>

<script setup lang="ts">
import { api } from 'boot/axios'
import { ref } from 'vue'
import { encrypt } from 'src/js/OCBO'
import { useRouter } from 'vue-router'
import { gsap } from 'gsap'
import { checkConnection } from 'src/js/functions'

import { usePageWithTable } from 'stores/pagewithtable'
import { useCurrentPage } from 'stores/currentpage'

import docInput from 'components/docInput.vue'
import docInputPassword from 'components/docInputPassword.vue'
import docButton from 'components/docButton.vue'
import docForm from 'components/docForm.vue'
import docList from 'components/docList.vue'

const sample = () => {
  console.log('yeah')
}
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

let redFullname = ref(false)
let redUsername = ref(false)
let redPassword = ref(false)
let redCPassword = ref(false)
let redAccess = ref(false)

let ifullname = ref('')
let iusername = ref('')
let ipassword = ref('')
let icpassword = ref('')
let accessList = ref<string[]>([])
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
  {
    label: 'Complaint',
    value: 'is_complaint',
  },
])
let dialog = ref(false)
let dialogTitle = ref('')
let dialogMessage = ref('')

let missingDetails: string[] = []
const checkComplete = async () => {
  missingDetails = []

  if (!ifullname.value) {
    missingDetails.push('fullname')
    redFullname.value = true
  } else redFullname.value = false

  if (!iusername.value) {
    missingDetails.push('username')
    redUsername.value = true
  } else redUsername.value = false

  if (!ipassword.value) {
    missingDetails.push('password')
    redPassword.value = true
  } else redPassword.value = false

  if (!icpassword.value) {
    missingDetails.push('confirm password')
    redCPassword.value = true
  } else redCPassword.value = false

  if (accessList.value.length === 0) {
    missingDetails.push('access')
    redAccess.value = true
  } else redAccess.value = false

  if (missingDetails.length > 0) return true
  else return false
}

const checkCompleteFullname = () => {
  if (!ifullname.value) redFullname.value = true
  else redFullname.value = false
}
const checkCompleteUsername = () => {
  if (!iusername.value) redUsername.value = true
  else redUsername.value = false
}
const checkCompletePassword = () => {
  if (!ipassword.value) redPassword.value = true
  else redPassword.value = false
}
const checkCompleteCPassword = () => {
  if (!icpassword.value) redCPassword.value = true
  else redCPassword.value = false
}
const checkCompleteAccess = () => {
  if  (accessList.value.length === 0) redAccess.value = true
  else redAccess.value = false
}

const passwordConfirm = async () => {
  if (ipassword.value === icpassword.value) return true
  else false
}

// const checkConnection = async (): Promise<boolean> => {
//   try {
//     const response = await api.get('/api/CheckConnection')
//     const data = response.data
//     if (data !== undefined && data.result === '1') return true
//   } catch {
//     return false
//   }
//   return false
// }

const saveAccount = async () => {
  if (await checkConnection()) {
    if ((await checkComplete()) === false) {
      if (await passwordConfirm()) {
        let ipasswordEncrypted = encrypt(ipassword.value.toUpperCase(), 'doctrack2023', 3, 128)
        let iincoming = accessList.value.includes('is_incoming') ? 1 : 0
        let ioutgoing = accessList.value.includes('is_outgoing') ? 1 : 0
        let ireleasing = accessList.value.includes('is_releasing') ? 1 : 0
        let iinventory = accessList.value.includes('is_inventory') ? 1 : 0
        let iothers = accessList.value.includes('is_otherdocuments') ? 1 : 0
        let icomplaint = accessList.value.includes('is_complaint') ? 1 : 0

        const response = await api.post('/api/PostAccount', {
          data: ifullname.value.toUpperCase(),
          data2: iusername.value.toUpperCase(),
          data3: ipasswordEncrypted,
          data4: iincoming,
          data5: ioutgoing,
          data6: ireleasing,
          data7: iinventory,
          data8: iothers,
          data9: icomplaint,
        })

        const data = response.data

        if (data.includes('Success')) {
          showDialog('Successfully Register', `You can now login as ${iusername.value.toUpperCase()}`)
        } else {
          showDialog('Failed to Register', 'Something went wrong. Please try again')
        }
      } else {
        showDialog('Cannot Register', 'Password does not match')
      }
    } else {
      showDialog('Cannot Register', `Missing ${missingDetails.toString().toUpperCase()}`)
    }
  } else {
    showDialog('Cannot Register', 'No Connection on Server')
  }
}

const showDialog = (title: string, message: string) => {
  dialog.value = true
  dialogTitle.value = title
  dialogMessage.value = message
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
  width: 100vw

.dialog-title-area
  padding: 0.1rem 0  1rem 0.1rem

.dialog-title
  font-family: 'Raleway'
  font-size: 1.4rem

// .dialog-card
//   background-color: #021926

.login
  margin: 1rem
  font-family: 'Inter'
  font-weight: 310

.login__username--label
  font-family: 'Inter'
  font-weight: 400
  font-size: 1.2rem

.login__username--input
  margin-bottom: 1.4rem

.login__username--option
  @extend .login__username--input
  text-transform: capitalize

// .dialog-card
//   font-family: "OpenSans"
//   background-color: rgba(2,25,38, 0.7)
//   width: 50%
//   height: 50%

// .dialog-card__section
//   padding: 2rem
//   font-size: 2rem

.register-button
  padding: 3rem 0rem 5rem 0rem
</style>
