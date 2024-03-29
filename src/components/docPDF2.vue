<template lang="pug">

div.flex.flex-center
  component(:is="docButton" @click="createPDF" :text="props.text" :title="props.title" :type="props.type" :name="props.name" :address="props.address" :remarks="props.remarks" :employee="props.employee")
  component(:is="docQR" :text="preText + encText" :size=qrSize style="display: none;" id="qr")
  img(src="../assets/lungsod.png", alt="PNG Image" style="display: none" id="lungsod")
  img(src="../assets/ocbologo2.png", alt="PNG Image" style="display: none" id="ocbo")

</template>

<script setup lang="ts">
import { jsPDF } from 'jspdf'
import { decrypt, encrypt } from 'src/js/OCBO'
import docQR from 'components/docQR.vue'
import docButton from 'components/docButton.vue'
import { ref } from 'vue'
// const image = ref('../assets/lungsod.png')

import Lora from '../assets/fonts/Lora.js'
import LoraBold from '../assets/fonts/Lora-Bold.js'
import LoraItalic from '../assets/fonts/Lora-Italic.js'
import RedHatText from '../assets/fonts/RedHatText.js'
import RedHatTextBold from '../assets/fonts/RedHatText-Bold.js'

export interface Props {
  title: string
  code: string
  text: string
  structureOwner: string
  structureOwnerAddress: string
  lotOwner: string
  lotOwnerAddress: string
  phone: string
  locationOfConstruction: string
  useOfOccupancy: string
  noOfStorey: string
  content: string
  date: string
  type: string
  name: string
  address: string
  remarks: string
  employee: string
  sections: string[]
  sectionsNumber: string[]
}

const props = withDefaults(defineProps<Props>(), {
  title: 'Generated Document',
  code: '00-0-0000',
  text: 'Sample Text',
  content: 'Sample Content',
  type: 'First Notice of Violation',
  name: 'Juan Dela Cruz',
  address: 'Davao City',
  remarks: 'The quick brown fox jumps over the lazy dog.',
  employee: 'Juan Dela Xruz',
  structureOwner: 'Juan Dela Cruz',
  structureOwnerAddress: 'Davao City',
  lotOwner: 'Juan Dela Cruz',
  lotOwnerAddress: 'Davao City',
  phone: '0123456789',
  locationOfConstruction: 'Davao City',
  useOfOccupancy: 'Residential',
  noOfStorey: '0',
  sections: () => ['a', 'b', 'c'],
  sectionsNumber: () => ['1', '2', '3'],
})

const preText = '**SCAN ME USING DDMS** QrId::'
const encText = encrypt(props.code)
const qrSize = 200

const createPDF = async () => {
  const doc = new jsPDF({
    orientation: 'p',
    unit: 'mm',
    format: 'legal',
  })

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

  // const getSectionData = async (section: string): Promise<string> => {
  //   const response = await api.get('/api/GetSectionData/' + section)
  //   const data = response.data.length !== 0 ? response.data : 0
  //   return (data !== null) ? data.result : ''
  // }

  // const getAllSection = async () => {
  //   for (let item of props.sections) {
  //     const sectionItem = await getSectionData(item)
  //     props.sectionsNumber.push(sectionItem)
  //   }
  // }

  const pageWidth = doc.internal.pageSize.getWidth()

  const republicText = 'Republic of the Philippines'
  const republicTextWidth = getTextWidth(republicText, 12)
  const officeText = 'OFFICE OF THE CITY BUILDING OFFICIAL'
  const officeTextWidth = getTextWidth(officeText, 14)
  const cityText = 'City of Davao'
  const cityTextWidth = getTextWidth(cityText, 12)

  const workTextWidth = getTextWidth('WORK STOPPAGE ORDER', 14)
  const propsTypeWidth = getTextWidth(props.type, 14)
  const pdWidth = getTextWidth('National Building Code of the Philippines (PD 1096)', 12)

  const republicTextX = (pageWidth - republicTextWidth) / 2
  const officeTextX = (pageWidth - officeTextWidth) / 2
  const cityTextX = (pageWidth - cityTextWidth) / 2
  const workTextX = (pageWidth - workTextWidth) / 2
  const propsTypeX = (pageWidth - propsTypeWidth) / 2
  const pdX = (pageWidth - pdWidth) / 2

  const sentence =
    '     In view hereof, you are hereby directed to stop immediately the illegal construction activities of your structure and secure Necessary Permits for your construction within a period of fifteen (15) calendar days from receipt of this notice. Failure to comply with this Work Stoppage Order will prompt the Building Official to request the DCWD for disconnection of the water supply, request the DLPC for disconnection of the electricity supply, and/or order the closure of the structure. This letter shall serve as your formal notice, and it is incumbent on your part to show why no actions should be taken against you and your property.'
  const sentence2 = '     Your immediate attention and cooperation on this regard is earnestly sought.'

  // doc.addFont('../assets/fonts/PlaypenSans-Medium.ttf', 'lora', 'normal')
  // doc.addFileToVFS('PlaypenSans.ttf', PlaypenSans)
  // doc.addFont('PlaypenSans.ttf', 'PlaypenSans', 'normal')
  // doc.setFont('Lora')

  doc.setFont('Lora', 'normal')
  doc.setFontSize(12)
  doc.text(republicText, republicTextX, 10)
  doc.setFont('LoraBold', 'bold')
  doc.setFontSize(14)
  doc.text(officeText, officeTextX, 16)
  doc.setFont('Lora', 'normal')
  doc.setFontSize(12)
  doc.text(cityText, cityTextX, 22)

  doc.addImage(lungsodLink.href, 'PNG', 5, 2, 28, 28, 'lungsod', 'NONE', 0)
  doc.addImage(ocboLink.href, 'PNG', pageWidth - 35, 2, 28, 28, 'ocbo', 'NONE', 0)

  doc.setLineWidth(0.9)
  doc.line(10, 34, 200, 34)

  doc.setFont('LoraItalic', 'italic')
  doc.setFontSize(8)
  doc.text('Reference No. OCBO-2023-R', 10, 38)

  doc.setFont('LoraBold', 'bold')
  doc.setFontSize(15)
  doc.text('WORK STOPPAGE ORDER', workTextX, 46)

  doc.setFont('LoraBold', 'bold')
  doc.setFontSize(14)
  doc.text(props.type, propsTypeX, 56)

  doc.setFontSize(10)
  doc.setFont('Lora', 'normal')
  doc.text(props.date, pageWidth - props.date.length * 2.8, 38)

  doc.setFont('Lora', 'normal')
  doc.text('Name of Structure Owner:', 10, 72)
  doc.setFont('LoraBold', 'bold')
  doc.text(props.structureOwner, 55, 72)

  doc.setFont('Lora', 'normal')
  doc.text('Mailing Address of Structure Owner:', 10, 78)
  doc.setFont('LoraBold', 'bold')
  doc.text(props.structureOwnerAddress, 72, 78)

  doc.setFont('Lora', 'normal')
  doc.text('Name of Lot Owner:', 10, 84)
  doc.setFont('LoraBold', 'bold')
  doc.text(props.lotOwner, 45, 84)

  doc.setFont('Lora', 'normal')
  doc.text('Postal/Mailing Address of Lot Owner:', 10, 90)
  doc.setFont('LoraBold', 'bold')
  doc.text(props.lotOwnerAddress, getTextWidth('Postal/Mailing Address of Lot Owner:', 10) + 10, 90)

  doc.setFont('Lora', 'normal')
  // doc.setFontSize(10)
  doc.text('Telephone No./ Mobile Phone No.:', 10, 96)
  doc.setFont('LoraBold', 'bold')
  doc.text(props.phone, getTextWidth('Telephone No./ Mobile Phone No.:', 10) + 11, 96)

  doc.setFont('Lora', 'normal')
  doc.text('Location of Construction or Installation:', 10, 102)
  doc.setFont('LoraBold', 'bold')
  doc.text(props.locationOfConstruction, getTextWidth('Location of Construction or Installation:', 10) + 10, 102)

  doc.setFont('Lora', 'normal')
  doc.text('Use or Character of Occupancy:', 10, 108)
  doc.setFont('LoraBold', 'bold')
  doc.text(props.useOfOccupancy, getTextWidth('Use or Character of Occupancy:', 10) + 11, 108)

  doc.setFont('Lora', 'normal')
  doc.text('Number of Storey:', 10, 114)
  doc.setFont('LoraBold', 'bold')
  doc.text(props.noOfStorey, getTextWidth('Number of Storey:', 10) + 11, 114)

  doc.setFont('LoraBold', 'bold')
  doc.text('National Building Code of the Philippines (PD 1096)', pdX, 124)

  doc.setFontSize(10)
  for (let i = 0; i < props.sections.length; i++) {
    doc.text(props.sectionsNumber + ' - ' + props.sections[i], 20, 130 + i * 6)
  }
  // doc.text('SECTION 212. Administrative Fines', 20, 130)
  // doc.text('SECTION 301. Building Permits', 20, 136)
  // doc.text('SECTION 309. Certificate of Occupancy', 20, 142)
  // doc.text('SECTION 804. Sizes and Dimensions of Courts', 20, 148)
  // doc.text('SECTION 808. Windows Openings', 20, 154)
  // doc.text('SECTION 1106. Pedestrian Protection', 20, 160)
  // doc.text('SECTION 1202. Excavation, Foundations, and Retaining Walls', 20, 166)

  doc.setFont('RedHatText', 'normal')
  doc.setFontSize(10)
  doc.text('REMARKS:', 10, 176)
  doc.text(props.remarks, 10, 182, { maxWidth: 188, align: 'justify' })

  let remarksY = 182
  const sentenceY = remarksY + (props.remarks.length / 188) * 10 + 6
  const sentence2Y = sentenceY + (sentence.length / 188) * 10

  // for (let i = 0; i < props.sections.length; i++) {
  //   doc.text(props.sections[i], 20, 130 + (i * 6))
  // }
  // for (let j = 0; j < props.remarks.length; j++) {
  //   doc.text(props.remarks[j], 16, remarksY, { maxWidth: 190 })
  //   if (getTextWidth(props.remarks[j], 10) > 186) remarksY += 6
  //   remarksY += 5
  // }

  doc.setFontSize(10)
  doc.text(sentence, 10, sentenceY, { maxWidth: 188, align: 'justify' })
  doc.text(sentence2, 10, sentence2Y, { maxWidth: 188, align: 'justify' })

  doc.text('Very truly yours,', 140, sentence2Y + 18)
  doc.setFont('LoraBold', 'bold')
  doc.text('AR. KHASAHAYAR L. TOGHYANI', 120, sentence2Y + 30)
  doc.setFont('Lora', 'normal')
  doc.text('Officer-In-Charge', 140, sentence2Y + 36)

  doc.text('Received By    :', 10, sentence2Y + 54)
  doc.text('Date and Time :', 10, sentence2Y + 60)

  doc.setLineWidth(0.2)
  doc.line(40, sentence2Y + 54, 100, sentence2Y + 54)
  doc.line(40, sentence2Y + 60, 100, sentence2Y + 60)

  doc.text('Served By        :', 10, sentence2Y + 66)
  doc.setFont('LoraBold', 'bold')
  doc.text(props.employee, 39, sentence2Y + 66)

  const footerWidth = getTextWidth('Office of the City Building Official - Pichon Street (former Magallanes Street)', 9)
  const footerWidthX = (pageWidth - footerWidth) / 2
  const footer2Width = getTextWidth('fronting A. Bonifacio Monument Rotunda, Davao City', 9)
  const footer2WidthX = (pageWidth - footer2Width) / 2
  const footer3Width = getTextWidth('Tel. No. 291-6695 / email address: ocbo.davao@gmail.com', 9)
  const footer3WidthX = (pageWidth - footer3Width) / 2

  doc.setFont('LoraItalic', 'italic')
  doc.setFontSize(9)
  doc.text('Office of the City Building Official - Pichon Street (former Magallanes Street)', footerWidthX, sentence2Y + 84)
  doc.text('fronting A. Bonifacio Monument Rotunda, Davao City', footer2WidthX, sentence2Y + 89)
  doc.text('Tel. No. 291-6695 / email address: ocbo.davao@gmail.com', footer3WidthX, sentence2Y + 94)

  doc.addImage(qrLink.href, 'PNG', 180, sentence2Y + 48, 25, 25, 'qr', 'NONE', 0)

  doc.save(`${props.title} - ${props.code}`)
}
</script>
