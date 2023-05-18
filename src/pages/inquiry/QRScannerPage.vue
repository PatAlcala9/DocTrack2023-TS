<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title QR Scanner
    q-btn(flat size="md" label="Back" @click="gotoHome" icon="arrow_back").close-button

  div.full-width.column.wrap.justify-center.items-center.content-center
    section.scanner
      component(:is="QrStream" @decode="handleDecode" @error="handleError")
    span.scanner-message Place the QR on the screen
    section(v-if="qr !== undefined")
      component(:is="vueQr" text="qr")

</template>

<script setup lang="ts">
import { QrStream } from 'vue3-qr-reader'

import { ref } from 'vue'
// import { api } from 'boot/axios'
import { useRouter } from 'vue-router'
import { useCurrentPage } from 'stores/currentpage'
import { useQrValue } from 'stores/qrvalue'

import vueQr from 'vue-qr/src/packages/vue-qr.vue'
import { encryptAES } from 'src/js/OCBO'

const router = useRouter()
const _currentpage = useCurrentPage()
const _qrvalue = useQrValue()

let qr = ref('')

const handleDecode = (text: string) => {
  _qrvalue.qrvalue = text
  gotoQrResult()
}

const handleError = () => {
  console.log('error')
}

const gotoQrResult = () => {
  _currentpage.currentpage = 'qrresult'
  router.push('/qrresult')
}

const gotoHome = () => {
  _currentpage.currentpage = '/'
  router.push('/')
}

const createSampleQR = (text: string) => {
  const prefix = '**SCAN ME USING OCBO DOCTRACK** QrID::'
  const qrtext = encryptAES(text)
  qr.value = prefix + qrtext
}
</script>

<style lang="sass">
.scanner
  margin-top: 2rem
  border: 1px solid rgba(255, 255, 255, 0.125)

.scanner-message
  font-family: 'Inter'
  font-weight: 300
  font-size: 1.2rem
  padding: 1.5rem
</style>
