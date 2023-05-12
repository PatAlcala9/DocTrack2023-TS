<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Complaint Tracking
    q-btn(flat size="md" label="Back" @click="gotoMenu" icon="arrow_back").close-button


</template>

<script setup lang="ts">
import { ref } from 'vue'
import { date } from 'quasar'
import { useRouter } from 'vue-router'
import { api } from 'boot/axios'

import { useCurrentPage } from 'stores/currentpage'

import docButton from 'components/docButton.vue'
import docTextArea from 'components/docTextArea.vue'
import docInput from 'components/docInput.vue'
import docLabel from 'components/docLabel.vue'

let entryCode = ref('')
let dataExist = ref(false)

let sourceName = ref('')
let dateReceived = ref('')
let subject = ref('')
let details = ref('')
let attachments = ref('')
let notes = ref('')
let docLogsList = []
let actionMadeList = []

const router = useRouter()
let _currentpage = useCurrentPage()

const gotoMenu = () => {
  _currentpage.currentpage = 'dashboard'
  router.push('/dashboard')
}

</script>
