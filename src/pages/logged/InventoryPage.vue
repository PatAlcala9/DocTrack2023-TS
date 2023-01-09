<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Inventory
    q-btn(flat size="md" label="Back" @click="gotoMenu" icon="arrow_back").close-button

  div.buttons-area.full-width.row.justify-start
    doc-button(text="File Document").button
    doc-button(text="Update Outgoing Document").button

  div.flex.flex-center
    section(v-if="fileReceivedDocList.result !== ''").dialog-content-table
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
            //- td
              //- q-btn(color="button" icon="visibility" :ripple="false" @click="openDetails(item, incomingList.result2[index], incomingList.result3[index], incomingList.result4[index])").button-view
              //- q-btn(v-if="showText === false" v-else color="button" label="View" :ripple="false" @mouseleave="mouseLeave").button-view


</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { api } from 'boot/axios'
import { useCurrentPage } from 'stores/currentpage'

import docButton from 'components/docButton.vue'

const router = useRouter()
const _currentpage = useCurrentPage()

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
    }
    else fileReceivedDocEmpty.value = true
  } catch {
    fileReceivedDocEmpty.value = true
  }
}

const gotoMenu = () => {
  _currentpage.currentpage = 'dashboard'
  router.push('/dashboard')
}

;(async () => {
  await getFileReceivedDoc()
  if (_currentpage.currentpage !== undefined) router.push(_currentpage.currentpage)
  else router.push('/inventory')
})()
</script>

<style lang="sass" scoped>
// .calendar
//   background-color: rgba(8,25,35, 0.75)
//   border: 1px solid rgba(255, 255, 255, 0.525)
//   font-family: 'OpenSans'
//   font-size: 1.6rem
//   border-radius: 12px
//   ::v-deep .q-date__calendar-item
//     padding-top: 0.6rem
//     // padding-left: 0.5rem
//     // padding-right: 0.5rem
//   ::v-deep .block
//     font-size: 1.25rem

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
  padding: 1rem 4rem
  width: 20rem

.table tbody tr
  border-bottom: 1px solid #dddddd

.table tbody tr:nth-of-type(even)
  background-color: #1C4157

.table tbody tr:last-of-type
  border-bottom: 2px solid #000406
</style>
