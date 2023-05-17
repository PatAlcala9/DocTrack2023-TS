<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title QR Scanner
    q-btn(flat size="md" label="Back" @click="gotoHome" icon="arrow_back").close-button

  div.full-width.column.wrap.justify-center.items-center.content-center
    section.scanner
      component(:is="QrStream" @decode="handleDecode" @error="handleError")
    span {{ decodedText }}

</template>

<script setup>
import { QrStream } from 'vue3-qr-reader'

import { ref } from 'vue'
// import { api } from 'boot/axios'
import { useRouter } from 'vue-router'
import { useCurrentPage } from 'stores/currentpage'

let decodedText = ref('')

const router = useRouter()
const _currentpage = useCurrentPage()

const handleDecode = (text) => {
  decodedText.value = text
}

const handleError = () => {
  console.log('error')
}

const gotoHome = () => {
  _currentpage.currentpage = '/'
  router.push('/')
}
</script>

<style lang="sass">
.scanner
  margin-top: 2rem
  border: 1px solid rgba(255, 255, 255, 0.125)
</style>
