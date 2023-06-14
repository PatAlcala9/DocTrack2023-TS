<template lang="pug">

div.flex.flex-center
  component(:is="docButton" @click="createPDF" text="Create Sample PDF")
  component(:is="docQR" :text="preText + encText" :size=qrSize style="display: none;" id="qr")

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
  text: 'Sample Text'
})

const preText = '**SCAN ME USING OCBO DOCTRACK** QrId::'
const encText = encryptAES(props.text)

const randomData1 = encryptAES(Math.random().toString()).substring(4, 10)
const randomData2 = encryptAES('some random string').substring(25)
const sampleQR = ref(`**SCAN ME USING OCBO DOCTRACK** QrId::${encryptAES(randomData1)}??${encryptAES(randomData2)}`)
const qrSize = 200

const createPDF = async () => {
  // const filePath = '~/assets/wso.pdf'
  // const existingPdfBytes = await fetch(filePath).then(res => res.arrayBuffer())
  // const bytes = new Uint8Array(existingPdfBytes)
  // const pdfDoc = await PDFDocument.load(bytes)






  const item = document.getElementById('qr')
  const image = item?.getAttribute('src')
  const link = document.createElement('a')
  link.href = image ?? ''

  const pdfDoc = await PDFDocument.create()
  const pngImage = await pdfDoc.embedPng(link.href)
  const pngDims = pngImage.scale(0.5)
  const page = pdfDoc.addPage(PageSizes.Letter)

  const font = await pdfDoc.embedFont(StandardFonts.Helvetica)
  const textSize = 14
  // page.drawText('Falcon Heavy', { x: 120, y: 560, size: textSize, font })
  // page.drawText('Saturn IV', { x: 120, y: 500, size: textSize, font })
  // page.drawText('Delta IV Heavy', { x: 340, y: 560, size: textSize, font })
  // page.drawText('Space Launch System', { x: 340, y: 500, size: textSize, font })

  page.drawImage(pngImage, {
    x: page.getWidth() / 4 - pngDims.width / 2 + 75,
    y: page.getHeight() / 1 - pngDims.height,
    width: pngDims.width,
    height: pngDims.height,
  })

  pdfDoc.setTitle(props.title)

  const pdfBytes = await pdfDoc.save()
  const pdfBlob = new Blob([pdfBytes], { type: 'application/pdf' })
  const link2 = document.createElement('a')
  link2.href = URL.createObjectURL(pdfBlob)
  window.open(link2.href)
}
</script>
