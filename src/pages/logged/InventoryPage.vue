<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Inventory
    q-btn(flat size="md" label="Back" @click="gotoMenu" icon="arrow_back").close-button

  div.buttons-area.full-width.row.justify-start
    doc-button(text="File Document" @click="gotoFileDoc").button
    doc-button(text="Update Outgoing Document").button

  div.flex.flex-center
    section(v-if="tabSelected === 'incoming'")
      span.tab__selected Filed Received Documents
      span(@click="switchTab('incoming')").tab Filed Outgoing Documents
    section(v-else)
      span(@click="switchTab('outgoing')").tab Filed Received Documents
      span.tab__selected Filed Outgoing Documents

    section(v-if="tabSelected === 'incoming'").dialog-content-table
      table.table
        thead
          tr
            th Folder
            th Page Number
            th Entry Code
            th Type of Communication
            th Source
            th Subject
            th Date Received
        tbody
          tr(v-for="(item, index) in fileReceivedDocList.result" :key="item").table-content-group
            td {{item}}
            td {{fileReceivedDocList.result2[index]}}
            td {{fileReceivedDocList.result3[index]}}
            td {{fileReceivedDocList.result4[index]}}
            td {{fileReceivedDocList.result5[index]}}
            td {{fileReceivedDocList.result6[index]}}
            td {{fileReceivedDocList.result7[index]}}

    section(v-else).dialog-content-table
      table.table
        thead
          tr
            th Folder
            th Page Number
            th Reference Code
            th Type of Communication
            th Respondent
            th Subject
            th Date Released
            th Date Received
        tbody
          tr(v-for="(item, index) in fileOutgoingDocList.result" :key="item").table-content-group
            td {{item}}
            td {{fileOutgoingDocList.result2[index]}}
            td {{fileOutgoingDocList.result3[index]}}
            td {{fileOutgoingDocList.result4[index]}}
            td {{fileOutgoingDocList.result5[index]}}
            td {{fileOutgoingDocList.result6[index]}}
            td {{fileOutgoingDocList.result7[index]}}
            td {{fileOutgoingDocList.result8[index]}}

</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { api } from 'boot/axios'
import { useCurrentPage } from 'stores/currentpage'
import { usePageWithTable } from 'stores/pagewithtable'

import docButton from 'components/docButton.vue'
import docLabel from 'components/docLabel.vue'

const router = useRouter()
const _currentpage = useCurrentPage()
const _pagewithtable = usePageWithTable()

let tabSelected = ref('incoming')

const switchTab = (value: string) => {
  if (value === 'incoming') tabSelected.value = 'outgoing'
  else tabSelected.value = 'incoming'
}

type FileReceived = {
  result: string
  result2: string
  result3: string
  result4: string
  result5: string
  result6: string
  result7: string
}
let fileReceivedDocList = ref({} as FileReceived)
let fileReceivedDocEmpty = ref(false)

const getFileReceivedDoc = async () => {
  try {
    const response = await api.get('/api/GetFileReceivedDoc')
    const data = response.data
    if (data.result.length > 0) {
      fileReceivedDocList.value = data
      fileReceivedDocEmpty.value = false
    } else fileReceivedDocEmpty.value = true
  } catch {
    fileReceivedDocEmpty.value = true
  }
}

type FileOutgoing = {
  result: string
  result2: string
  result3: string
  result4: string
  result5: string
  result6: string
  result7: string
  result8: string
}
let fileOutgoingDocList = ref({} as FileOutgoing)
let fileOutgoingDocEmpty = ref(false)

const getFileOutgoingDoc = async () => {
  try {
    const response = await api.get('/api/GetFileOutgoingDoc')
    const data = response.data
    if (data.result.length > 0) {
      fileOutgoingDocList.value = data
      fileOutgoingDocEmpty.value = false
    } else fileOutgoingDocEmpty.value = true
  } catch {
    fileOutgoingDocEmpty.value = true
  }
}

const gotoMenu = () => {
  _pagewithtable.pagewithtable = false
  _currentpage.currentpage = 'dashboard'
  router.push('/dashboard')
}

const gotoFileDoc = () => {
  _pagewithtable.pagewithtable = true
  _currentpage.currentpage = 'inventoryfile'
  router.push('/inventoryfile')
}

;(async () => {
  await getFileReceivedDoc()
  await getFileOutgoingDoc()
  if (_currentpage.currentpage !== undefined) router.push(_currentpage.currentpage)
  else router.push('/inventory')
})()
</script>

<style lang="sass" scoped>


.title
  grid-area: title
  justify-self: start
  align-self: stretch
  padding: 0 0 0 1rem
  margin: 0 0 -2rem 0

.logo
  width: 12rem
  height: auto
  opacity: 0.9
  margin-right: 2rem

.name
  grid-area: name
  justify-self: start
  align-self: center

.login
  grid-area: login
  justify-self: center
  align-self: center
  padding: 1.4rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  border-radius: 2rem
  background-color: #021926
  margin-top: -10rem
  // backdrop-filter: blur(16px) saturate(180%)

.inquiry
  grid-area: inquiry
  justify-self: start
  align-self: center

.inquiry-text
  font-family: 'Raleway'
  font-size: 1.2rem
  text-decoration: underline
  padding: 2rem
  cursor: pointer
  color: #ffffff

.username-area
  margin: 1rem

.username-label
  font-family: 'Raleway'
  font-size: 1.4rem
  color: #ffffff

.username-input
  font-family: 'Raleway'
  font-size: 1.3rem
  border-radius: 0.6rem
  text-align: center
  text-transform: uppercase

.password-area
  margin: 1rem

.password-label
  @extend .username-label

.password-input
  @extend .username-input

.button-area
  padding-top: 1rem

.davao
  grid-area: davao
  justify-self: end
  align-self: center

.davaologo
  width: 18rem
  height: auto
  opacity: 0.4

.inquiry-label
  font-family: 'Raleway'
  font-size: 2.1rem
  color: #ffffff

.dialog
  width: 100vw

.dialog-title-area
  padding: 0.1rem 0  1rem 0.1rem

.dialog-title
  font-family: 'Raleway'
  font-size: 1.4rem

.dialog-subtitle
  font-family: 'Raleway'
  font-size: 1.2rem

.inputs
  margin-left: 2rem

.inputs__label
  font-size: 1.2rem

.inputs__input
  font-family: 'OpenSans'
  font-size: 1.3rem
  border-radius: 0.6rem
  // text-align: center
  margin-bottom: 1rem
  width: 50vw

.inputs__label--date
  @extend .inputs__label
  font-family: 'OpenSans'
  margin-top: 0.5rem
  font-size: 1.4rem

.button-area
  margin-top: 2rem

.list-area
  padding-top: 2rem
  font-size: 0.5rem

.list-area__span
  cursor: pointer
  text-decoration: underline
  font-size: 1.2rem

.form-area
  margin-top: 2rem

.requiredWarning
  color: red

.buttons-area
  margin: 2rem 0

.button
  margin-left: 1rem

.table
  font-family: 'Montserrat'
  font-size: 0.8rem
  text-transform: uppercase
  border-collapse: collapse
  margin: 2rem 0
  min-width: 25rem
  border-radius: 1rem 1rem 0 0
  overflow: hidden
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.2)

.table thead tr
  background-color: #000406
  color: #ffffff
  text-align: left
  font-weight: bold

.table th, .table td
  padding: 1rem 1rem
  width: 20rem

.table tbody tr
  border-bottom: 1px solid #dddddd

.table tbody tr:nth-of-type(even)
  background-color: #1C4157

.table tbody tr:last-of-type
  border-bottom: 2px solid #000406

.tab-area
  margin-top: 2rem

.tab
  font-family: 'Raleway'
  font-size: 1.2rem
  padding: 1rem
  border: 1px solid white
  cursor: pointer

.tab__selected
  @extend .tab
  font-family: 'RalewayBold'
  background-color: rgba(255, 255, 255, 0.1)
</style>
