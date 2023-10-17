<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title QR Scanner
    q-btn(flat size="md" label="Back" @click="gotoHome" icon="arrow_back").close-button

  div(v-if="validQR === 0").fit.column.wrap.justify-start.items-center.content-center
    div.text-negative
      q-icon(name="dangerous" size="10rem")
    span.invalid-message QR ERROR
    span.invalid-submessage {{ _qrerror.qrerror }}

  div(v-else-if="validQR === 1").fit.column.wrap.justify-start.items-center.content-center
    div.text-negative
      q-icon(name="gpp_bad" size="10rem")
    span.invalid-message Invalid
    span.invalid-message DocTrack QR

  div(v-else).fit.column.wrap.justify-start.items-center.content-center
    div(v-if="!invalid").section.fit.column.wrap.justify-start.items-center.content-center
      section.section.fit.column.wrap.justify-start.items-center.content-center
        component(:is="docLabel" text="Code")
        span.information {{ extractedData }}

      section.section.fit.column.wrap.justify-start.items-center.content-center
        component(:is="docLabel" text="Status")
        span.information {{ status }}

      section.section.fit.column.wrap.justify-start.items-center.content-center
        component(:is="docLabel" text="Transacted")
        span.information {{ transacted }}

      section.section.fit.column.wrap.justify-start.items-center.content-center
        component(:is="docLabel" text="Expiration")
        span.information {{ expiration }}

      section.section.fit.column.wrap.justify-start.items-center.content-center
        component(:is="docLabel" text="Complaint Type")
        span.information {{ type }}

      section.section.fit.column.wrap.justify-start.items-center.content-center
        component(:is="docLabel" text="Complaintant Name")
        span.information {{ complaintant }}

      section.section.fit.column.wrap.justify-start.items-center.content-center
        component(:is="docLabel" text="Respondent Name")
        span.information {{ respondent }}

      section.break

    div(v-else)
      div.text-negative
        q-icon(name="gpp_bad" size="10rem")
      span.invalid-message No Record
      span.invalid-message Found

</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useCurrentPage } from 'stores/currentpage'
import { useQrValue } from 'stores/qrvalue'
import { useQrError } from 'stores/qrerror'
import { ref } from 'vue'
import { decrypt } from 'src/js/OCBO'
import { date } from 'quasar'

import docLabel from 'components/docLabel.vue'
import { api } from 'src/boot/axios'

// let code = ref('99-9-9999')
let status = ref('SAMPLE STATUS')
let type = ref('')
let complaintant = ref('')
let respondent = ref('')
let transacted = ref('')
let expiration = ref('')

let invalid = ref(false)

const router = useRouter()
const _currentpage = useCurrentPage()
const _qrvalue = useQrValue()
const _qrerror = useQrError()

const qrPrefix = '**SCAN ME USING DDMS** QrId::'

let validQR = ref(0)
const checkQR = async () => {
  if ((_qrerror.qrerror === '' || _qrerror.qrerror == undefined) && _qrvalue.qrvalue.includes(qrPrefix) === true) validQR.value = 2
  else validQR.value = 1
}

// const loadSampleData = () => {
//   _qrvalue.qrvalue = '**SCAN ME USING DDMS** QrId::2dfg21fd23g23df12g31fd2g12fd12g3'
//   validQR.value = 2
// }

const getQRData = async () => {
  const response = await api.get('/api/GetQRData/' + extractedData.value)
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) {
    type.value = data.result
    complaintant.value = data.result2
    respondent.value = data.result3
    invalid.value = false
  } else invalid.value = true
}

const getQRStatus = async () => {
  const response = await api.get('/api/GetLatestStatusNameIndividual2/' + extractedData.value)
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) {
    status.value = data.result
    transacted.value = date.formatDate(data.result2, 'MMMM DD, YYYY')
    expiration.value = date.formatDate(data.result3, 'MMMM DD, YYYY')
  }
}

let extractedData = ref('')
const extractData = async () => {
  const contentLoc = _qrvalue.qrvalue.indexOf('::')
  const content = decrypt(_qrvalue.qrvalue.slice(contentLoc + 2))
  extractedData.value = content
}

const checkInvalid = async (): Promise<boolean> => {
  return !complaintant.value ? false : true
}

const gotoHome = () => {
  _qrvalue.qrvalue = ''
  _currentpage.currentpage = '/'
  router.push('/')
}

;(async () => {
  await checkQR()
  await extractData()
  await getQRData()

  if (await checkInvalid()) {
    getQRStatus()
  }
})()
</script>

<style lang="sass" scoped>
.section
  margin: 1.2rem

.invalid-message
  font-family: 'Inter'
  font-weight: 400
  font-size: 2.1rem

.invalid-submessage
  font-family: 'Inter'
  font-weight: 350
  font-size: 1.9rem

.information
  font-family: 'Inter'
  font-weight: 560
  font-size: 1.4rem
  border: 2px solid rgba(255, 255, 255, 0.5)
  background-color: rgba(12, 21, 42, 0.45)
  backdrop-filter: blur(9px) saturate(150%)
  border-radius: 3rem
  padding: 0.5rem 1.5rem
  //background-color: $background

.break
  padding: 0 0 2rem 0
</style>
