<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title QR Scanner
    q-btn(flat size="md" label="Back" @click="gotoHome" icon="arrow_back").close-button

  div.full-width.column.wrap.justify-center.items-center.content-center
    section.scanner
      component(:is="QrStream" @decode="handleDecode" @error="handleError")
    span.scanner-message Place the QR on the screen

</template>

<script setup lang="ts">
import { QrStream } from 'vue3-qr-reader'
import { useRouter } from 'vue-router'
import { useCurrentPage } from 'stores/currentpage'
import { useQrValue } from 'stores/qrvalue'
import { useQrError } from 'stores/qrerror'

const router = useRouter()
const _currentpage = useCurrentPage()
const _qrvalue = useQrValue()
const _qrerror = useQrError()

const handleDecode = (text: string) => {
  _qrerror.qrerror = ''
  _qrvalue.qrvalue = text
  gotoQrResult()
}

const handleError = (error: string) => {
  _qrerror.qrerror = error
  _qrvalue.qrvalue = ''
  gotoQrResult()
}

const gotoQrResult = () => {
  _currentpage.currentpage = 'qrresult'
  router.push('/qrresult')
}

const gotoHome = () => {
  _currentpage.currentpage = '/'
  router.push('/')
}
</script>

<style lang="sass" scoped>
.scanner
  margin-top: 2rem
  border: 1px solid rgba(255, 255, 255, 0.125)

.scanner-message
  font-family: 'Inter'
  font-weight: 400
  font-size: 1.4rem
  padding: 1.5rem
</style>
