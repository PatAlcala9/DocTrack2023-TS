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
    section.section.fit.column.wrap.justify-start.items-center.content-center
      component(:is="docLabel" text="Code")
      span.information {{ code }}

    section.section.fit.column.wrap.justify-start.items-center.content-center
      component(:is="docLabel" text="Status")
      span.information {{ status }}

</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useCurrentPage } from 'stores/currentpage'
import { useQrValue } from 'stores/qrvalue'
import { useQrError } from 'stores/qrerror'
import { ref } from 'vue'

import docLabel from 'components/docLabel.vue'

let code = ref('99-9-9999')
let status = ref('SAMPLE STATUS')

const router = useRouter()
const _currentpage = useCurrentPage()
const _qrvalue = useQrValue()
const _qrerror = useQrError()

const qrPrefix = '**SCAN ME USING OCBO DOCTRACK** QrID::'

let validQR = ref(0)
const checkQR = () => {
  if ((_qrerror.qrerror === '' || _qrerror.qrerror == undefined) && _qrvalue.qrvalue.includes(qrPrefix) === true) validQR.value = 2
  else validQR.value = 1
}

const gotoHome = () => {
  _qrvalue.qrvalue = ''
  _currentpage.currentpage = '/'
  router.push('/')
}

;(async () => {
  checkQR()
})()
</script>

<style lang="sass">
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
  font-size: 2rem
  border: 1px solid #ffffff
  border-radius: 3rem
  padding: 0.5rem 1.5rem
  background-color: $background
</style>
