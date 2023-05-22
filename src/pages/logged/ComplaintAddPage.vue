<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Complaint - Add New
    q-btn(flat size="md" label="Back" @click="gotoComplaintDashboard" icon="arrow_back").close-button


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

const router = useRouter()
let _currentpage = useCurrentPage()

const gotoMenu = () => {
  _currentpage.currentpage = 'dashboard'
  router.push('/dashboard')
}

const gotoComplaintDashboard = () => {
  _currentpage.currentpage = 'complaint'
  router.push('/complaint')
}
</script>
