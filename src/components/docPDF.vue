<template lang="pug">

div.flex.flex-center
  component(:is="docButton" @click="createPDF" text="Create Sample PDF")
  component(:is="docQR" :text="preText + encText" :size=qrSize style="display: none;" id="qr")
  img(src="../assets/lungsod.png", alt="PNG Image" style="display: none" id="lungsod")

</template>

<script setup lang="ts">
import { PDFDocument, StandardFonts, PageSizes, rgb } from 'pdf-lib'
import { decrypt, encrypt } from 'src/js/OCBO'
import docQR from 'components/docQR.vue'
import docButton from 'components/docButton.vue'
import { ref } from 'vue'
import fs from 'fs'
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

const randomData1 = encrypt(Math.random().toString()).substring(4, 10)
const randomData2 = encrypt('some random string').substring(25)
const qrSize = 200

const createPDF = async () => {
  const pdfDoc = await PDFDocument.create()
  const page = pdfDoc.addPage(PageSizes.Legal)

  const qrItem = document.getElementById('qr')
  const qrSrc = qrItem?.getAttribute('src')
  const qrLink = document.createElement('a')
  qrLink.href = qrSrc ?? ''
  const qrImage = await pdfDoc.embedPng(qrLink.href)
  const qrImageDims = qrImage.scale(0.5)

  const lungsodItem = document.getElementById('lungsod')
  const lungsodSrc = lungsodItem?.getAttribute('src')
  const lungsodLink = document.createElement('a')
  lungsodLink.href = lungsodSrc ?? ''

  const response = await fetch(lungsodLink.href)
  const blob = await response.blob()
  let img = ''

  const reader = new FileReader()
  reader.readAsDataURL(blob)
  reader.onload = async () => {
    if (typeof reader.result === 'string') {
      img = reader.result
      const lungsodImage = await pdfDoc.embedPng(img)
      const lungsodImageDims = lungsodImage.scale(0.5)

      page.drawImage(lungsodImage, {
        x: 250,
        y: 600,
        width: lungsodImageDims.width,
        height: lungsodImageDims.height,
      })
    } else {
      img = 'yeah'
      console.log(new Error('Failed to convert image to data URL'))
    }
  }

  const timesRoman = await pdfDoc.embedFont(StandardFonts.TimesRoman)
  const timesRomanBold = await pdfDoc.embedFont(StandardFonts.TimesRomanBold)
  const timesRomanItalic = await pdfDoc.embedFont(StandardFonts.TimesRomanItalic)
  const textSize = 14

  const republicText = 'Republic of the Philippines'
  const officeText = 'OFFICE OF THE CITY BUILDING OFFICIAL'
  const cityText = 'City of Davao'
  page.drawText(republicText, { x: page.getWidth() / 2 - republicText.length * 3, y: page.getHeight() / 1 - 25, size: textSize, font: timesRoman })
  page.drawText(officeText, { x: page.getWidth() / 2 - officeText.length * 4, y: page.getHeight() / 1 - 42, size: textSize, font: timesRomanBold })
  page.drawText(cityText, { x: page.getWidth() / 2 - cityText.length * 3, y: page.getHeight() / 1 - 59, size: textSize, font: timesRoman })

  page.drawText('Reference No. OCBO-2023-R', { x: 50, y: page.getHeight() / 1 - 93, size: 8, font: timesRomanItalic })
  page.drawText(props.date, { x: page.getWidth() - 80, y: page.getHeight() / 1 - 93, size: 9, font: timesRoman })

  page.drawText('WORK STOPPAGE ORDER', { x: page.getWidth() / 2 - 98, y: page.getHeight() / 1 - 127, size: textSize, font: timesRomanBold })

  // page.drawImage(qrImage, {
  //   x: page.getWidth() / 20 - qrImageDims.width / 2 + 500,
  //   y: page.getHeight() / 1 - qrImageDims.height,
  //   width: qrImageDims.width,
  //   height: qrImageDims.height,
  // })

  pdfDoc.setTitle(props.title)

  const pdfBytes = await pdfDoc.save()
  const pdfBlob = new Blob([pdfBytes], { type: 'application/pdf' })
  const link2 = document.createElement('a')
  link2.href = URL.createObjectURL(pdfBlob)
  window.open(link2.href)
}
</script>
