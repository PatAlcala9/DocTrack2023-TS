<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Outgoing Commmunications - New Entry
    q-btn(flat size="md" label="Back" @click="gotoMenu" icon="arrow_back").close-button

  div.form-area.fit.column.justify-center
    section.section
      //- span.inputs__label Respondent:
      //-   span.requiredWarning {{ missingRespondent }}
      //- component(:is="docInput" width="40" alignment="left" transform="initial" v-model:value="outRespondent")
      component(:is="docLabel" text="Respondent:").label--spaced
      component(:is="docInputEntry" v-model:value="missingRespondent" alignment="left" width=100)

    section.section
      //- span.inputs__label Address:
      //-   span.requiredWarning {{ missingAddress }}
      //- component(:is="docInput" width="40" alignment="left" transform="initial" v-model:value="outAddress")
      component(:is="docLabel" text="Address:").label--spaced
      component(:is="docInputEntry" v-model:value="missingAddress" alignment="left" width=100)

    section.section
      //- span.inputs__label From:
      //-   span.requiredWarning {{ missingFrom }}
      //- component(:is="docInput" width="40" v-model:value="outFrom" alignment="left")
      component(:is="docLabel" text="From:").label--spaced
      component(:is="docInputEntry" v-model:value="missingFrom" alignment="left" width=100)

    section.section
      //- span.inputs__label Attachments:
      //-   span.requiredWarning {{ missingAttachments }}
      //- component(:is="docTextArea" width="40" v-model:value="outAttachments")
      component(:is="docLabel" text="Attachments:").label--spaced
      component(:is="docInputEntry" v-model:value="missingAttachments" alignment="left" width=100)

    section.section
      //- span.inputs__label Subject:
      //-   span.requiredWarning {{ missingSubject }}
      //- component(:is="docInput" width="40" v-model:value="outSubject" alignment="left")
      component(:is="docLabel" text="Subject:").label--spaced
      component(:is="docInputEntry" v-model:value="missingSubject" alignment="left" width=100)

    section.section
      //- span.inputs__label Details:
      //-   span.requiredWarning {{ missingDetails }}
      //- component(:is="docTextArea" width="40" v-model:value="outDetails")
      component(:is="docLabel" text="Details:").label--spaced
      component(:is="docInputEntry" v-model:value="missingDetails" alignment="left" width=100)


    section.button-area.fit.row.items-center.justify-center
      doc-button(text="Save" type="submit" @click="showContents")

    //- section.list-area.column.text-center
    //-   span.list-area__span Show List of Outgoings

//- q-dialog(v-model="dialog" transition-show="flip-right" transition-hide="flip-left")
//-   q-card.dialog-card.text-white.flex.flex-center
//-     q-card-section.dialog-card__section
//-       div.dialog-title-area.column.justify-center.items-center
//-         span.dialog-card__title {{dialogMessage}}
//-         span.dialog-card__info {{dialogSubMessage}}
//-         component(:is="docButton" text="OK" v-close-popup)

</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { date } from 'quasar'
import { useRouter } from 'vue-router'
import { api } from 'boot/axios'

import { useCurrentPage } from 'stores/currentpage'

import docButton from 'components/docButton.vue'
import docTextArea from 'components/docTextArea.vue'
import docInput from 'components/docInput.vue'
import docInputEntry from 'components/docInputEntry.vue'
import docLabel from 'components/docLabel.vue'
import docSelection from 'components/docSelection.vue'
import docCalendar from 'components/docCalendar.vue'
import docList from 'components/docList.vue'

const router = useRouter()
const _currentpage = useCurrentPage()

let receivedDate = ref('')
let formattedReceivedDate = ref('')

let outRespondent = ref('')
let outAddress = ref('')
let outFrom = ref('')
let outAttachments = ref('')
let outSubject = ref('')
let outDetails = ref('')

let missingRespondent = ref('')
let missingAddress = ref('')
let missingFrom = ref('')
let missingAttachments = ref('')
let missingSubject = ref('')
let missingDetails = ref('')

const showContents = async () => {
  console.log('outRespondent', outRespondent.value)
  console.log('outAddress', outAddress.value)
  console.log('outFrom', outFrom.value)
  console.log('outAttachments', outAttachments.value)
  console.log('outSubject', outSubject.value)
  console.log('outDetails', outDetails.value)
}

const gotoMenu = () => {
  _currentpage.currentpage = 'dashboard'
  router.push('/dashboard')
}

;(async () => {
  if (_currentpage.currentpage !== undefined) router.push(_currentpage.currentpage)
  else router.push('/outgoing')
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
  // width: 5rem

.requiredWarning
  color: red

.section
  display: flex
  flex-direction: column
  background-color: rgba(12, 21, 42, 0.45)
  backdrop-filter: blur(9px) saturate(150%)
  border-radius: 0.6rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  padding: 1.2rem 2rem 1.2rem 1.2rem
  margin-bottom: 1rem

.section--calendar
  @extend .section
  background-color: rgba(128, 21, 21, 0.45)
  grid-area: section--calendar

.section--source, .section--subject, .section--details, .section--attachments
  @extend .section
  background-color: rgba(128, 21, 21, 0.45)
</style>



