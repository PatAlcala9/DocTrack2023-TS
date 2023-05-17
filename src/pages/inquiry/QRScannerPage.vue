<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title QR Scanner
    q-btn(flat size="md" label="Back" @click="gotoHome" icon="arrow_back").close-button

  div.flex.flex-center
    component(:is="VueQrReader" @code-scanned="handleDecode" @error-captured="handleError" videoWidth="300" videoHeight="300" :responsive="true")

</template>

<script setup lang="ts">
// import { QrcodeStream } from 'vue-qrcode-reader'
import VueQrReader from 'vue-qr-reader/dist/lib/vue-qr-reader.umd.js'
import { ref } from 'vue'
import { api } from 'boot/axios'
import { useRouter } from 'vue-router'
import { useCurrentPage}  from 'stores/currentpage'

import docButton from 'components/docButton.vue'

const router = useRouter()
const _currentpage = useCurrentPage()



const handleDecode = (text: string) => {
  console.log(text)
}

const handleError = () => {
  console.log('error')
}

const gotoHome = () => {
  _currentpage.currentpage = '/'
  router.push('/')
}
</script>
