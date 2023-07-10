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
  type: string
  name: string
  address: string
}

const props = withDefaults(defineProps<Props>(), {
  title: 'Generated Document',
  text: 'Sample Text',
  type: 'First Notice of Violation',
  name: 'Juan Dela Cruz',
  address: 'Davao City'
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
  const officeTextWidth = getTextWidth(officeText, 14)
  const cityText = 'City of Davao'
  const cityTextWidth = getTextWidth(cityText, 12)

  const workTextWidth = getTextWidth('WORK STOPPAGE ORDER', 14)
  const propsTypeWidth = getTextWidth(props.type, 14)

  const republicTextX = (pageWidth - republicTextWidth) / 2
  const officeTextX = (pageWidth - officeTextWidth) / 2
  const cityTextX = (pageWidth - cityTextWidth) / 2
  const workTextX = (pageWidth - workTextWidth) / 2
  const propsTypeX = (pageWidth - propsTypeWidth) / 2

  // doc.addFont('../assets/fonts/lora.ttf', 'lora', 'normal')
  // doc.setFont('lora')

  doc.setFont('times', 'normal')
  doc.text(republicText, republicTextX, 10)
  doc.setFont('times', 'bold')
  doc.text(officeText, officeTextX, 16)
  doc.setFont('times', 'normal')
  doc.text(cityText, cityTextX, 22)

  // doc.addImage(qrLink.href, 'PNG', 1, 10, 20, 20, 'qr', 'NONE', 0)
  doc.addImage(lungsodLink.href, 'PNG', 5, 2, 30, 30, 'lungsod', 'NONE', 0)
  doc.addImage(ocboLink.href, 'PNG', pageWidth - 35, 2, 30, 30, 'ocbo', 'NONE', 0)

  doc.setLineWidth(0.9)
  doc.line(10, 34, 200, 34)

  doc.setFont('times', 'italic')
  doc.setFontSize(9)
  doc.text('Reference No. OCBO-2023-R', 10, 38)

  doc.setFont('times', 'bold')
  doc.setFontSize(14)
  doc.text('WORK STOPPAGE ORDER', workTextX, 46)

  doc.setFont('times', 'bold')
  doc.setFontSize(14)
  doc.text(props.type, propsTypeX, 56)

  doc.setFontSize(12)
  doc.setFont('times', 'normal')
  doc.text(props.date, 156, 38)

  doc.setFont('times', 'normal')
  doc.text('Name of Structure Owner:', 10, 72)
  doc.setFont('times', 'bold')
  doc.text(props.name, 57, 72)

  doc.setFont('times', 'normal')
  doc.text('Mailing Address of Structure Owner:', 10, 78)
  doc.setFont('times', 'bold')
  doc.text(props.address, 75, 78)

  doc.setFont('times', 'normal')
  doc.text('Name of Lot Owner:', 10, 84)
  doc.setFont('times', 'bold')
  doc.text(props.address, 47, 84)

  doc.setFont('times', 'normal')
  doc.text('Postal/Mailing Address of Lot Owner:', 10, 90)
  doc.setFont('times', 'bold')
  doc.text(props.address, getTextWidth('Postal/Mailing Address of Lot Owner:', 12) + 8, 90)

  doc.setFont('times', 'normal')
  doc.text('Telephone No./ Mobile Phone No.:', 10, 96)
  doc.setFont('times', 'bold')
  doc.text(props.address, getTextWidth('Telephone No./ Mobile Phone No.:', 12) + 10, 96)

  doc.setFont('times', 'normal')
  doc.text('Location of Construction or Installation:', 10, 102)
  doc.setFont('times', 'bold')
  doc.text(props.address, getTextWidth('Location of Construction or Installation:', 12) + 10, 102)

  doc.setFont('times', 'normal')
  doc.text('Use or Character of Occupancy:', 10, 108)
  doc.setFont('times', 'bold')
  doc.text(props.address, getTextWidth('Use or Character of Occupancy:', 12) + 10, 108)

  doc.setFont('times', 'normal')
  doc.text('Number of Storey:', 10, 114)
  doc.setFont('times', 'bold')
  doc.text(props.address, getTextWidth('Number of Storey:', 12) + 10, 114)

  // const pdfData = doc.output('datauristring')
  // const fileName = 'sample.pdf'

  doc.save('sample.pdf')

  // const pdfData = doc.output()
  // const blob = new Blob([pdfData], { type: 'application/pdf' })
  // const url = URL.createObjectURL(blob)
  // window.open(url)
}
</script>
