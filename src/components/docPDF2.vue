<template lang="pug">

div.flex.flex-center
  component(:is="docButton" @click="createPDF" text="Create Sample PDF")
  component(:is="docQR" :text="preText + encText" :size=qrSize style="display: none;" id="qr")
  img(src="../assets/lungsod.png", alt="PNG Image" style="display: none" id="lungsod")

</template>

<script setup lang="ts">
import { jsPDF } from 'jspdf'
import { decrypt, encrypt } from 'src/js/OCBO'
import docQR from 'components/docQR.vue'
import docButton from 'components/docButton.vue'
import { ref } from 'vue'
// const image = ref('../assets/lungsod.png')
// import imgUrl from '../assets/lungsod.png'

export interface Props {
  title: string
  text: string
  date: string
}

const props = withDefaults(defineProps<Props>(), {
  title: 'Generated Document',
  text: 'Sample Text',
})

const preText = '**SCAN ME USING OCBO DOCTRACK** QrId::'
const encText = encrypt(props.text)
const qrSize = 200

const createPDF = async () => {
  const doc = new jsPDF()

  const qrItem = document.getElementById('qr')
  const qrSrc = qrItem?.getAttribute('src')
  const qrLink = document.createElement('a')
  qrLink.href = qrSrc ?? ''

  const lungsodItem = document.getElementById('lungsod')
  const lungsodSrc = lungsodItem?.getAttribute('src')
  const lungsodLink = document.createElement('a')
  lungsodLink.href = lungsodSrc ?? ''

  const republicText = 'Republic of the Philippines'
  const officeText = 'OFFICE OF THE CITY BUILDING OFFICIAL'
  const cityText = 'City of Davao'

  doc.text(republicText, 20, 10)

  doc.addImage(qrLink.href, 'PNG', 10, 10, 20, 20, 'qr', 'NONE', 0)
  doc.addImage(lungsodLink.href, 'PNG', 20, 20, 50, 50, 'lungsod', 'NONE', 0)

  // const pdfData = doc.output('datauristring')
  // const fileName = 'sample.pdf'

  doc.save('sample.pdf')
}
</script>
