<template lang="pug">

q-page(padding)
  div.grid
    section(v-if="$q.screen.width <=500").title.full-width.row.justify-between
      div.column
        transition(appear @before-enter="beforeEnterLogo" @enter="enterLogo")
          img(src="../assets/logo3.webp" alt="OCBO Logo").logo
        transition(appear @before-enter="beforeEnterTitle" @enter="enterTitle")
          section.name.fit.column.wrap.justify-start.items-start.content-start
            span.ocbo-title-mobile DDMS

    section(v-else).title.full-width.row.justify-between
      div.column
        transition(appear @before-enter="beforeEnterLogo" @enter="enterLogo")
          img(src="../assets/logo3.webp" alt="OCBO Logo").logo
        transition(appear @before-enter="beforeEnterTitle" @enter="enterTitle")
          section.name.fit.column.wrap.justify-start.items-start.content-start
            span.ocbo-title Digital Document
            span.ocbo-text Management System

    //- transition( appear @before-enter="beforeEnterForm" @enter="enterForm")
    div(v-if="inquiry === false").login#login
      div
        section.username-area.column.wrap.justify-center.items-center.content-center
          doc-label(text="Username")
          component(v-if="$q.screen.width <=500" :is="docInput" v-model:value="usernameEntry" :width=mobileWidth)
          doc-input(v-else v-model:value="usernameEntry")

        section.password-area.column.wrap.justify-center.items-center.content-center
          doc-label(text="Password")
          component(v-if="$q.screen.width <=500" :is="docInputPassword" v-model:value="passwordEntry" @keypress.enter="login")
          doc-input-password(v-else v-model:value="passwordEntry" @keypress.enter="login")

        section.button-area.column.wrap.justify-center.items-center.content-center
          doc-button(text="login" @click="login")

        section.register.column.wrap.justify-center.items-center.content-center
          span.register--text(@click="gotoRegister") Create New Account

    div(v-else).login#inquiry
      div
        section.username-area.column.wrap.justify-center.items-center.content-center
          span.inquiry-label Inquiry
        section.button-area--inquiry.column.wrap.justify-center.items-center.content-center
          doc-button(text="Incoming" @click="inquireReceivedTrigger")
        section.button-area--inquiry.column.wrap.justify-center.items-center.content-center
          doc-button(text="Outgoing" @click="inquireReleasedTrigger")
        section.button-area--inquiry.column.wrap.justify-center.items-center.content-center
          doc-button(text="QR" @click="qrScannerTrigger")

    transition(appear @before-enter="beforeEnterSwitch" @enter="enterSwitch")
      div.inquiry
        span(v-if="inquiry" @click="showInquiry").inquiry-text Login
        span(v-else @click="showLogin").inquiry-text Inquire
        //- component(:is="docPDF2" title="Sample Document PDF" text="aaaa" date="July 12, 2023")

    div.version
      section.version--section
        span(@click="versionDialog=true").version--text {{version}}
        span.version--text Developed By: Pat Alcala

q-dialog(v-model="error" transition-show="flip-right" transition-hide="flip-left" @keypress.enter="error=false").dialog
  q-card.dialog-card.text-white.flex.flex-center
    q-card-section.dialog-card__section
      div.dialog-title-area.column.justify-center.items-center
        span.dialog-card__title {{errorMessage}}
        span.dialog-card__info {{errorInformation}}
        doc-button(text="OK" @click="error=false")

q-dialog(v-model="versionDialog" full-width full-height transition-show="flip-right" transition-hide="flip-left" @keypress.enter="versionDialog=false").dialog
  q-card.dialog-card.text-white
    q-card-section.dialog-card__section
      div.dialog-title-area.column.justify-center.items-center
        component(:is="docVersions")
        doc-button(text="Close" @click="versionDialog=false").button-area
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { api } from 'boot/axios'
import { useRouter } from 'vue-router'
import { comparePassword, createAuth } from 'src/js/OCBO'
import { gsap } from 'gsap'
import { SessionStorage, Platform, useQuasar } from 'quasar'
import { checkConnection } from 'src/js/functions'
//- import vueQr from 'vue-qr/src/packages/vue-qr.vue'
//- import { PDFDocument, PDFImage, StandardFonts, PageSizes, PDFField, rgb } from 'pdf-lib'

import { useEmployeeName } from 'stores/employeename'
import { useUserID } from 'stores/userid'
import { usePageWithTable } from 'stores/pagewithtable'
import { useAccess } from 'stores/access'
import { useCurrentPage } from 'stores/currentpage'
import { useIsDemo } from 'stores/isdemo'
import { useIsLogged } from 'stores/islogged'
import { useAuthentication } from 'stores/authentication'
import { useEmployeeID } from 'stores/employeeid'

import docButton from 'components/docButton.vue'
import docInput from 'components/docInput.vue'
import docInputPassword from 'components/docInputPassword.vue'
import docLabel from 'components/docLabel.vue'
import docVersions from 'components/docVersions.vue'
import docPDF2 from 'components/docPDF2.vue'
import { consola, createConsola } from 'consola'

const _employeename = useEmployeeName()
const _userid = useUserID()
const _pagewithtable = usePageWithTable()
const _access = useAccess()
const _currentpage = useCurrentPage()
const _isdemo = useIsDemo()
const _islogged = useIsLogged()
const _authentication = useAuthentication()
const _employeeid = useEmployeeID()

let error = ref(false)
let errorMessage = ref('')
let errorInformation = ref('')

let versionDialog = ref(false)

const mobileWidth = ref(10)
const sampleMode = ref(false)

const router = useRouter()
const quasar = useQuasar()

let inquiry = ref(false)
let inquireReceived = ref(false)
let inquireReleased = ref(false)

const incomingList: any = ref([])
let incomingHeaderList = [
  { name: 'entryCodeNo', align: 'center', label: 'Entry Code', field: 'entryCodeNo', sortable: true },
  { name: 'receivedDate', label: 'Received Date', field: 'receivedDate', sortable: true },
  { name: 'sourceName', label: 'Source', field: 'sourceName' },
  { name: 'subjectInfo', label: 'Subject', field: 'subjectInfo' },
  { name: 'view', label: 'Details', field: 'view' },
]

let usernameEntry = ref('')
let passwordEntry = ref('')
let userid = 0
let employeeName = null
let loginSuccess = false

const version = ref('v 0.15')

const showInquiry = async () => {
  await exitInquiry()
  await enterLogin()
}

const showLogin = async () => {
  await exitLogin()
  await enterInquiry()
}

const inquireReceivedTrigger = async () => {
  _pagewithtable.pagewithtable = true
  SessionStorage.set('CurrentPage', 'incominginquire')
  _currentpage.currentpage = 'incominginquire'
  router.push('/incominginquire')
}

const inquireReleasedTrigger = async () => {
  _pagewithtable.pagewithtable = true
  SessionStorage.set('CurrentPage', 'outgoinginquire')
  _currentpage.currentpage = 'outgoinginquire'
  router.push('/outgoinginquire')
}

const qrScannerTrigger = async () => {
  _pagewithtable.pagewithtable = false
  SessionStorage.set('CurrentPage', 'qrscanner')
  _currentpage.currentpage = 'qrscanner'
  router.push('/qrscanner')
}

const getIncoming = async () => {
  const response = await api.get('/api/GetIncoming')
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) {
    incomingList.value = data
  }
}

const gotoRegister = async () => {
  quasar.loading.show()
  _pagewithtable.pagewithtable = false
  SessionStorage.set('CurrentPage', 'register')
  _currentpage.currentpage = 'register'
  router.push('/register')
}

let usernameAccepted = false
const checkUsername = async () => {
  try {
    const response = await api.get('/api/CheckUsername/' + usernameEntry.value.toUpperCase())
    const data = response.data
    const dataNum = parseInt(data.result)
    if (data !== undefined) usernameAccepted = dataNum > 0 ? true : false
  } catch {
    usernameAccepted = false
  }
}

let passwordAccepted = false
const checkPassword = async () => {
  try {
    const response = await api.get('/api/CheckPassword/' + usernameEntry.value.toUpperCase())
    const data = response.data
    if (data !== undefined && comparePassword(data.result, passwordEntry.value.toUpperCase(), 'doctrack2023', 3, 128)) passwordAccepted = true
  } catch {
    passwordAccepted = false
  }
}

let detailsAllowed = false
const getUserDetails = async () => {
  try {
    const response = await api.get('/api/GetUserDetails/' + usernameEntry.value.toUpperCase())
    const data = response.data

    if (data !== undefined) {
      _userid.userid = data
      _employeename.employeename = data.result2

      if (data.result3 === '1') _access.access.push('incoming')
      if (data.result4 === '1') _access.access.push('outgoing')
      if (data.result5 === '1') _access.access.push('releasing')
      if (data.result6 === '1') _access.access.push('inventory')
      if (data.result7 === '1') _access.access.push('otherdocuments')
      if (data.result8 === '1') _access.access.push('complaint')
      if (data.result9 === '1') _access.access.push('inspector')
      detailsAllowed = true
    }
  } catch {
    detailsAllowed = false
  }
}

const getUserID = async () => {
  try {
    const response = await api.get('/api/GetUserID/' + usernameEntry.value.toUpperCase())
    const data = response.data
    _employeeid.employeeid = data !== undefined ? data.result : 0
  } catch {
    _employeeid.employeeid = 0
  }
}

const setDemo = async () => {
  SessionStorage.set('Demo', _isdemo.isdemo)
}

const login = async () => {
  if (usernameEntry.value.toUpperCase() === 'DEMO' && passwordEntry.value.toUpperCase() === 'DEMO') {
    _access.access.push('incoming')
    _access.access.push('outgoing')
    _access.access.push('releasing')
    _access.access.push('inventory')
    _access.access.push('otherdocuments')
    _access.access.push('complaint')
    _access.access.push('inspector')
    _isdemo.isdemo = true
    await setDemo()
    _islogged.islogged = true

    _pagewithtable.pagewithtable = false
    _employeename.employeename = 'DEMONSTRATION'
    SessionStorage.set('CurrentPage', 'dashboard')
    _currentpage.currentpage = 'dashboard'
    router.push('/dashboard')

    return
  }

  quasar.loading.show()
  if (await checkConnection()) {
    await checkUsername()
    if (usernameAccepted === false) {
      error.value = true
      errorMessage.value = 'Invalid Username'

      if (usernameEntry.value.length > 0) errorInformation.value = `${usernameEntry.value.toUpperCase()} does not exist`
      else errorInformation.value = 'Username is Empty'

      quasar.loading.hide()
      return
    }

    await checkPassword()
    if (passwordAccepted === false) {
      error.value = true
      errorMessage.value = 'Invalid Password'
      errorInformation.value = `Password does not match with ${usernameEntry.value.toUpperCase()}`

      if (passwordEntry.value.length > 0) errorInformation.value = `Password does not match with ${usernameEntry.value.toUpperCase()}`
      else errorInformation.value = 'Password is Empty'

      quasar.loading.hide()
      return
    }

    await getUserDetails()
    await getUserID()
    if (detailsAllowed === false) {
      error.value = true
      errorMessage.value = 'No Details Found'
      quasar.loading.hide()
      return
    }
    quasar.loading.hide()
    _authentication.authentication = createAuth(usernameEntry.value, passwordEntry.value)
    _isdemo.isdemo = false
    await setDemo()
    _islogged.islogged = true
    _pagewithtable.pagewithtable = false
    SessionStorage.set('CurrentPage', 'dashboard')
    _currentpage.currentpage = 'dashboard'

    router.push('/dashboard')
  } else {
    quasar.loading.hide()
    error.value = true
    errorMessage.value = 'Cannot Login'
    errorInformation.value = 'No Connection on Server'
  }
}

const exitLogin = async () => {
  await gsap.to('#login', { duration: 0.25, rotationY: 90, opacity: 0 })
  inquiry.value = true
}

const enterInquiry = async () => {
  gsap.from('#inquiry', { duration: 0.25, rotationY: 90, opacity: 0 })
}

const exitInquiry = async () => {
  await gsap.to('#inquiry', { duration: 0.25, rotationY: 90, opacity: 0 })
  inquiry.value = false
}

const enterLogin = async () => {
  gsap.from('#login', { duration: 0.25, rotationY: 90, opacity: 0 })
}

const beforeEnterLogo = (el: any) => {
  el.style.scale = 0
}
const enterLogo = (el: any) => {
  gsap.to(el, { duration: 1, scale: 1, ease: 'elastic.out(1, 0.5)' })
}

const beforeEnterTitle = (el: any) => {
  el.style.transform = 'translateX(100px)'
  el.style.opacity = 0
}
const enterTitle = (el: any) => {
  gsap.to(el, { duration: 0.8, x: 1, opacity: 1 })
}

const beforeEnterForm = (el: any) => {
  el.style.scale = 0
}
const enterForm = (el: any) => {
  gsap.to(el, { duration: 0.6, scale: 1, ease: 'back.out(1.6)' })
}

const beforeEnterSwitch = (el: any) => {
  el.style.transform = 'translateX(-100px)'
  el.style.opacity = 0
}
const enterSwitch = (el: any) => {
  gsap.to(el, { duration: 0.8, x: 1, opacity: 1 })
}

interface State {
  browserName: string
  browserVersion: number
  devicePlatform: string
}

const state = reactive<State>({
  browserName: '',
  browserVersion: 0,
  devicePlatform: '',
})

const supportedBrowsers = [
  {
    name: 'firefox',
    minVersion: 93,
  },
  {
    name: 'chrome',
    minVersion: 85,
  },
  {
    name: 'safari',
    minVersion: 16.4,
  },
  {
    name: 'opera',
    minVersion: 71,
  },
]

const supportedMobile = [
  {
    name: 'firefox',
    minVersion: 93,
  },
  {
    name: 'chrome',
    minVersion: 116,
  },
  {
    name: 'safari',
    minVersion: 16,
  },
  {
    name: 'samsung',
    minVersion: 14,
  },
  {
    name: 'opera',
    minVersion: 73,
  },
  {
    name: 'operamini',
    minVersion: 0,
  },
]

const isBrowserSupported = () => {
  return supportedBrowsers.some((browser) => {
    return state.browserName === browser.name && state.browserVersion >= browser.minVersion
  })
}

const isMobileSupported = () => {
  return supportedMobile.some((browser) => {
    return state.browserName === browser.name && state.browserVersion >= browser.minVersion
  })
}

const detectBrowser = () => {
  state.browserName = Platform.is.name
  state.browserVersion = Number(Platform.is.version ?? 0)
  state.devicePlatform = Platform.is.platform

  const isMobile = Platform.is.mobile ? true : false

  if (!isBrowserSupported()) {
    _currentpage.currentpage = 'unsupported'
    router.push('/unsupported')
  }

  if (isMobile && !isMobileSupported()) {
    _currentpage.currentpage = 'unsupported'
    router.push('/unsupported')
  }
}

//- import { useQrValue } from 'stores/qrvalue'
//- import { useQrError } from 'stores/qrerror'
//- const _qrvalue = useQrValue()
//- const _qrerror = useQrError()

//- const testQRResult = () => {
//-   _qrerror.qrerror = ''
//-   _qrvalue.qrvalue = '**SCAN ME USING OCBO DOCTRACK** QrID::1as2d132as2d123sa123'
//-   _currentpage.currentpage = 'qrresult'
//-   router.push('/qrresult')
//- }

;(async () => {
  //- detectBrowser()

  if (_currentpage.currentpage !== undefined) router.push(_currentpage.currentpage)
  else router.push('/')
})()
</script>

<style lang="sass" scoped>
.ocbo-title
  font-size: 4rem
  font-family: 'RalewayBold'
  color: #ffffff

.ocbo-title-mobile
  font-size: 3rem
  font-family: 'RalewayBold'
  color: #ffffff

.ocbo-text
  font-size: 3.3rem
  font-family: 'Montserrat'
  margin-top: -2rem
  color: #ffffff

.grid
  display: grid
  grid-template-columns: 1fr 1fr 1fr
  grid-template-rows: 1fr 1fr 1fr
  gap: 0px 0px
  grid-auto-flow: row
  justify-items: end
  grid-template-areas: "title title title" ". login ." "inquiry . davao"
  height: calc(100vh - 2rem)

.title
  grid-area: title
  justify-self: start
  align-self: stretch
  padding: 0 0 10rem 1rem
  margin: 0 0 2rem 0
  overflow: hidden

.logo
  width: 10rem
  height: auto
  opacity: 0.6
  margin-right: 2rem
  // filter: invert(100%)

.name
  grid-area: name
  justify-self: start
  align-self: center

.main-name
  backdrop-filter: blur(9px) saturate(150%)
  background-color: rgba(12, 21, 42, 0.45)
  border-radius: 0.6rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  padding: 0rem 1rem 0.9rem 1rem

.login
  grid-area: login
  justify-self: center
  align-self: center
  padding: 1.4rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  border-radius: 2rem
  // background-color: #021926
  margin-top: -1rem
  width: 22rem
  height: 22rem
  backdrop-filter: blur(9px) saturate(150%)
  background-color: rgba(12, 21, 42, 0.45)

.inquiry
  grid-area: inquiry
  justify-self: left
  align-self: end
  padding: 1.1rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  border-radius: 2rem
  backdrop-filter: blur(9px) saturate(150%)
  background-color: rgba(12, 21, 42, 0.45)

.inquiry-text
  font-family: 'Inter'
  font-weight: 300
  font-size: 1.1rem
  text-decoration: underline
  padding: 2rem
  cursor: pointer
  color: #ffffff

.username-area
  margin: 1rem

.password-area
  margin: 1rem

.button-area
  padding-top: 1rem

.button-area--inquiry
  padding-top: 1.6rem

.davao
  grid-area: davao
  justify-self: end
  align-self: center

.davaologo
  width: 18rem
  height: auto
  opacity: 0.4

.inquiry-label
  font-family: 'Raleway'
  font-size: 2.1rem
  color: #ffffff

.dialog
  width: 100vw

.dialog-title-area
  padding: 0.1rem 0  1rem 0.1rem

.dialog-title
  font-family: 'Raleway'
  font-size: 1.4rem
.table
  width: 93vw

.table-header-group
  font-family: 'Montserrat'
  font-size: 0.9em

.table-header
  padding: 0.2rem
  background-color: #ffffff
  color: #021926
  text-transform: uppercase

.table-content-group
  font-family: 'Montserrat'
  font-size: 0.9rem

.table-content
  padding: 0.2rem
  color: #ffffff

.register
  padding-top: 1.5rem
  cursor: pointer

.register:hover
  text-decoration: underline

.register--text
  font-family: 'Inter'
  font-family: 0.6rem

.version
  position: absolute
  bottom: 0
  left: 50%
  transform: translateX(-50%)
  text-align: center
  padding-bottom: 2rem

.version--text
  font-family: 'Inter'
  font-size: 0.9rem
  color: #ffffff
  cursor: pointer

.version--section
  display: flex
  flex-direction: column

.no-scrolling
  overflow: hidden

@media screen and (max-width: 900px)
  .grid
    grid-template-areas: "title title title" "login login login" "inquiry inquiry ."

  .logo
    width: 9rem

  .ocbo-title
    font-size: 3.2rem

  .ocbo-text
    font-size: 2.5rem
    margin-top: -1rem

  .register--text
    font-size: 1rem
  .inquiry-text
    font-size: 1.1rem



@media screen and (max-width: 500px)
  .grid
    grid-template-areas: "title title title" "login login login" "inquiry inquiry inquiry"
  .logo
    width: 5rem
    opacity: 0.6

  .ocbo-title
    font-size: 2.2rem

  .ocbo-text
    font-size: 1.2rem

  .login
    margin-top: -5rem
    width: 90%
    height: 22rem

  .register--text
    font-size: 0.9rem

  .inquiry
    margin-top: 1.2rem
    justify-self: center
    align-self: center

  .inquiry-text
    font-family: 'Inter'
    font-weight: 300
    font-size: 1.1rem
    text-decoration: underline
    padding: 2rem
    cursor: pointer
    color: #ffffff
    font-size: 1.1rem

  .version
    padding-bottom: 2%
</style>
