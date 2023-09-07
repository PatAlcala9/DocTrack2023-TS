<template lang="pug">

q-page(padding).column.fit.justify-center.content-center.items-center.text-center.ns-page
  span.main Browser Not Supported
  span.sub--main Doctrack was built using the latest web technologies which your current browser does not fully support.
  span.sub--main Please consider upgrading to a newer browser version for the best experience.

  span.info Browser Info: {{ state.browserName }} version {{ state.browserVersion }}

  section.button
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

const capitalize = (str: string) => {
  return str[0].toUpperCase() + str.slice(1)
}

interface State {
  browserName: string
  browserVersion: number
  devicePlatform: string
}

const state = reactive<State>({
  browserName: '',
  browserVersion: 0,
  devicePlatform: '',
})

const detectBrowser = () => {
  state.browserName = capitalize(Platform.is.name)
  state.browserVersion = Number(Platform.is.version ?? 0)
  state.devicePlatform = Platform.is.platform
}

detectBrowser()
</script>

<style lang="sass" scoped>
.main
  font-size: 4rem

.sub--main
  font-size: 1.6rem

.ns-page
  background-color: black
  font-family: 'Inter'

.info
  margin-top: 3rem
  font-size: 1.2rem

.button
  padding-top: 2rem
</style>
