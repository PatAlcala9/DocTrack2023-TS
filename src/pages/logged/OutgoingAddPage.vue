<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Outgoing - New Entry
    q-btn(flat size="md" label="Back" @click="gotoMenu" icon="arrow_back").close-button

q-dialog(v-model="dialog" transition-show="flip-right" transition-hide="flip-left")
  q-card.dialog-card.text-white.flex.flex-center
    q-card-section.dialog-card__section
      div.dialog-title-area.column.justify-center.items-center
        span.dialog-card__title {{dialogMessage}}
        span.dialog-card__info {{dialogSubMessage}}
        component(:is="docButton" text="OK" v-close-popup)
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
let _currentpage = useCurrentPage()

let dialog = ref(false)
let dialogMessage = ref('')
let dialogSubMessage = ref('')

let showList = ref(false)
type Outgoing = {
  result: string
  result2: string
  result3: string
  result4: string
}
let outgoingList = ref({} as Outgoing)

const gotoMenu = () => {
  _currentpage.currentpage = 'dashboard'
  router.push('/dashboard')
}
</script>

<style lang="sass" scoped>
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
</style>
