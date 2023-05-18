<template lang="pug">

q-page
  div.full-width.row.justify-between
    span.title QR Scanner
    q-btn(flat size="md" label="Back" @click="gotoHome" icon="arrow_back").close-button

  h1 {{ _qrvalue.qrvalue }}

</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useCurrentPage } from 'stores/currentpage'
import { useQrValue } from 'stores/qrvalue'

const router = useRouter()
const _currentpage = useCurrentPage()
const _qrvalue = useQrValue()

const gotoHome = () => {
  _currentpage.currentpage = '/'
  router.push('/')
}
</script>
