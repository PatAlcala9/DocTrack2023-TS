<template lang="pug">

q-page(padding).column.fit.justify-center.content-center.items-center.text-center
  h2 Browser
  h4 {{ state.browserName }}
  component(:is="DocButton" text="Go Back" @click="gotoPage('dashboard')")

</template>

<script setup lang="ts">
import { Platform } from 'quasar'
import { reactive } from 'vue'
import DocButton from 'components/docButton.vue'
import { useCurrentPage } from 'stores/currentpage'
import { usePageWithTable } from 'stores/pagewithtable'
import { useRouter } from 'vue-router'

const router = useRouter()
const _currentpage = useCurrentPage()
const _pagewithtable = usePageWithTable()

const gotoPage = (page: string, table = false) => {
  _pagewithtable.pagewithtable = table
  _currentpage.currentpage = page
  router.push(page)
}


const state = reactive({
  browserName: '',
  // browserVersion: '' | 0,
  devicePlatform: ''
})

const detectBrowser = () => {
  state.browserName = Platform.is.name
  // state.browserVersion = Platform.is.version ?? 0
  state.devicePlatform = Platform.is.platform
}
</script>

