<template lang="pug">

q-page(padding).column.fit.justify-center.content-center.items-center
  h2 This page is still in development
  component(:is="DocButton" text="Go Back" @click="gotoPage('dashboard')")

</template>

<script setup lang="ts">
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
</script>
