<template lang="pug">

div.flex.flex-center
  component(:is="docButton" @click="createPDF" text="Create Sample PDF")
  component(:is="docQR" :text="preText + encText" :size=qrSize style="display: none;" id="qr")
  img(src="../assets/lungsod.png", alt="PNG Image" style="display: none" id="lungsod")

</template>

<script setup lang="ts">
import { PDFDocument, StandardFonts, PageSizes, PDFField } from 'pdf-lib'
import { decryptAES, encryptAES } from 'src/js/OCBO'
import docQR from 'components/docQR.vue'
import docButton from 'components/docButton.vue'
import { ref } from 'vue'
import fs from 'fs'

export interface Props {
  title: string
  text: string
}

const props = withDefaults(defineProps<Props>(), {
  title: 'Generated Document',
  text: 'Sample Text',
})

const preText = '**SCAN ME USING OCBO DOCTRACK** QrId::'
const encText = encryptAES(props.text)

const randomData1 = encryptAES(Math.random().toString()).substring(4, 10)
const randomData2 = encryptAES('some random string').substring(25)
const sampleQR = ref(`**SCAN ME USING OCBO DOCTRACK** QrId::${encryptAES(randomData1)}??${encryptAES(randomData2)}`)
const qrSize = 200

const createPDF = async () => {
  const pdfDoc = await PDFDocument.create()

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

  const response = await fetch(lungsodLink.href);
  const blob = await response.blob()

  const reader = new FileReader();
    reader.onloadend = async () => {
      if (typeof reader.result === 'string') {
        const lungsodImage = await pdfDoc.embedPng(reader.result)
        const lungsodImageDims = lungsodImage.scale(1)
        page.drawImage(lungsodImage, {
          x: 250,
          y: 600,
          width: lungsodImageDims.width,
          height: lungsodImageDims.height,
        })
        console.log(reader.result)
      } else {
        console.log(new Error('Failed to convert image to data URL'));
      }
    }
    reader.readAsDataURL(blob);

  const page = pdfDoc.addPage(PageSizes.Legal)

  const font = await pdfDoc.embedFont(StandardFonts.Helvetica)
  const textSize = 14
  page.drawText('Republic of the Philippines', { x: page.getWidth() / 2 - 80, y: page.getHeight() / 1 - 25, size: textSize, font })
  page.drawText('OFFICE OF THE CITY BUILDING OFFICIAL', { x: page.getWidth() / 2 - 140, y: page.getHeight() / 1 - 42, size: textSize, font })
  page.drawText('City of Davao', { x: page.getWidth() / 2 - 50, y: page.getHeight() / 1 - 59, size: textSize, font })
  // page.drawText('Space Launch System', { x: 340, y: 500, size: textSize, font })



  page.drawImage(qrImage, {
    x: page.getWidth() / 20 - qrImageDims.width / 2 + 55,
    y: page.getHeight() / 1 - qrImageDims.height,
    width: qrImageDims.width,
    height: qrImageDims.height,
  })



  pdfDoc.setTitle(props.title)

  const pdfBytes = await pdfDoc.save()
  const pdfBlob = new Blob([pdfBytes], { type: 'application/pdf' })
  const link2 = document.createElement('a')
  link2.href = URL.createObjectURL(pdfBlob)
  window.open(link2.href)
}
</script>
