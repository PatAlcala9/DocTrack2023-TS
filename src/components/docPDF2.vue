<template lang="pug">

div.flex.flex-center
  component(:is="docButton" @click="createPDF" text="Create Sample PDF")
  component(:is="docQR" :text="preText + encText" :size=qrSize style="display: none;" id="qr")
  img(src="../assets/lungsod.png", alt="PNG Image" style="display: none" id="lungsod")
  img(src="../assets/ocbologobw.png", alt="PNG Image" style="display: none" id="ocbo")

</template>

<script setup lang="ts">
import { jsPDF } from 'jspdf'
import { decrypt, encrypt } from 'src/js/OCBO'
import docQR from 'components/docQR.vue'
import docButton from 'components/docButton.vue'
import { ref } from 'vue'
// const image = ref('../assets/lungsod.png')

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

  const ocboItem = document.getElementById('ocbo')
  const ocboSrc = ocboItem?.getAttribute('src')
  const ocboLink = document.createElement('a')
  ocboLink.href = ocboSrc ?? ''

  function getTextWidth(text: string, fontSize: number) {
    doc.setFontSize(fontSize)
    const textWidth = doc.getStringUnitWidth(text) * fontSize * 0.35 // Adjust the multiplier if needed
    return textWidth
  }

  const pageWidth = doc.internal.pageSize.getWidth()

  const republicText = 'Republic of the Philippines'
  const republicTextWidth = getTextWidth(republicText, 12)
  const officeText = 'OFFICE OF THE CITY BUILDING OFFICIAL'
  const officeTextWidth = getTextWidth(officeText, 12)
  const cityText = 'City of Davao'
  const cityTextWidth = getTextWidth(cityText, 12)

  const republicTextX = (pageWidth - republicTextWidth) / 2
  const officeTextX = (pageWidth - officeTextWidth) / 2
  const cityTextX = (pageWidth - cityTextWidth) / 2

  doc.addFont('../assets/fonts/lora.ttf', 'Lora', 'normal')
  doc.setFont('Lora')

  doc.text(republicText, republicTextX, 10)
  doc.text(officeText, officeTextX, 16)
  doc.text(cityText, cityTextX, 22)

  // doc.addImage(qrLink.href, 'PNG', 1, 10, 20, 20, 'qr', 'NONE', 0)
  doc.addImage(lungsodLink.href, 'PNG', 5, 2, 30, 30, 'lungsod', 'NONE', 0)
  doc.addImage(ocboLink.href, 'PNG', pageWidth - 35, 2, 30, 30, 'ocbo', 'NONE', 0)

  doc.line(10, 200, 1, 200, 'F')

  // const pdfData = doc.output('datauristring')
  // const fileName = 'sample.pdf'

  doc.save('sample.pdf')

  // const pdfData = doc.output()
  // const blob = new Blob([pdfData], { type: 'application/pdf' })
  // const url = URL.createObjectURL(blob)
  // window.open(url)
}
</script>
