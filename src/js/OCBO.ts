const GenerateMD5 = require('md5')
const CryptoJS = require('crypto-js')

// declare module 'crypto-js'

export function hash(Password: string, salt: string, iteration = 1, bits: 128 | 256 | 512 | 1024 | 2048) {
  let expandedSalt = ''
  for (let i = 0; i < iteration; i++) {
    expandedSalt += salt
  }

  const md51 = GenerateMD5(expandedSalt + Password + expandedSalt)
  let vowelCount = 0
  let consCount = 0

  const p1 = md51.toString().substr(0, 8)
  const p2 = md51.toString().substr(8, 8)
  const p3 = md51.toString().substr(16, 8)
  const p4 = md51.toString().substr(24, 8)

  for (let x = 0; x < iteration; x++) {
    for (let i = 0; i < Password.length; i++) {
      if (Password.charAt(i) === 'a' || Password.charAt(i) === 'A' || Password.charAt(i) === 'e' || Password.charAt(i) === 'E' || Password.charAt(i) === 'i' || Password.charAt(i) === 'I' || Password.charAt(i) === 'o' || Password.charAt(i) === 'O' || Password.charAt(i) === 'u' || Password.charAt(i) === 'U') {
        vowelCount++
      } else {
        consCount++
      }
    }
  }

  const outputLength = bits / 128
  let result = ''

  const xtra = consCount.toString() + vowelCount.toString()
  const newConsCount = consCount * parseInt(xtra)
  const newVowelCount = vowelCount * parseInt(xtra)
  const newValue = md51 + newConsCount + newVowelCount
  const md52 = GenerateMD5(newValue)

  const p21 = md52.toString().substr(0, 8)
  const p22 = md52.toString().substr(8, 8)
  const p23 = md52.toString().substr(16, 8)
  const p24 = md52.toString().substr(24, 8)

  result = p2 + p22 + p1 + p21 + p4 + p24 + p3 + p23

  if (outputLength > 1) {
    const xtra = (consCount * outputLength).toString() + (vowelCount * outputLength).toString()
    const newConsCount = consCount * outputLength * parseInt(xtra)
    const newVowelCount = vowelCount * outputLength * parseInt(xtra)
    const newValue = result + newConsCount + newVowelCount
    const md53 = GenerateMD5(newValue)
    const md54 = GenerateMD5(newConsCount.toString() + newVowelCount.toString())

    const p31 = md53.toString().substr(0, 8)
    const p32 = md53.toString().substr(8, 8)
    const p33 = md53.toString().substr(16, 8)
    const p34 = md53.toString().substr(24, 8)

    const p31x = md54.toString().substr(0, 8)
    const p32x = md54.toString().substr(8, 8)
    const p33x = md54.toString().substr(16, 8)
    const p34x = md54.toString().substr(24, 8)

    result = result + p32x + p32 + p31x + p31 + p34x + p34 + p33x + p33
  }

  if (outputLength > 2) {
    const xtra = (consCount * outputLength).toString() + (vowelCount * outputLength).toString()
    const newConsCount = consCount * outputLength * parseInt(xtra)
    const newVowelCount = vowelCount * outputLength * parseInt(xtra)
    const newValue = result + newConsCount + newVowelCount
    const md55 = GenerateMD5(newValue)
    const md56 = GenerateMD5(newConsCount.toString() + newVowelCount.toString())

    const p41 = md55.toString().substr(0, 8)
    const p42 = md55.toString().substr(8, 8)
    const p43 = md55.toString().substr(16, 8)
    const p44 = md55.toString().substr(24, 8)

    const p41x = md56.toString().substr(0, 8)
    const p42x = md56.toString().substr(8, 8)
    const p43x = md56.toString().substr(16, 8)
    const p44x = md56.toString().substr(24, 8)

    result = result + p42x + p42 + p41x + p41 + p44x + p44 + p43x + p43
  }

  if (outputLength > 3) {
    const xtra = (consCount * outputLength).toString() + (vowelCount * outputLength).toString()
    const newConsCount = consCount * outputLength * parseInt(xtra)
    const newVowelCount = vowelCount * outputLength * parseInt(xtra)
    const newValue = result + newConsCount + newVowelCount
    const md57 = GenerateMD5(newValue)
    const md58 = GenerateMD5(newConsCount.toString() + newVowelCount.toString())

    const p51 = md57.toString().substr(0, 8)
    const p52 = md57.toString().substr(8, 8)
    const p53 = md57.toString().substr(16, 8)
    const p54 = md57.toString().substr(24, 8)

    const p51x = md58.toString().substr(0, 8)
    const p52x = md58.toString().substr(8, 8)
    const p53x = md58.toString().substr(16, 8)
    const p54x = md58.toString().substr(24, 8)

    result = result + p52x + p52 + p51x + p51 + p54x + p54 + p53x + p53
  }

  if (outputLength > 4) {
    const xtra = (consCount * outputLength).toString() + (vowelCount * outputLength).toString()
    const newConsCount = consCount * outputLength * parseInt(xtra)
    const newVowelCount = vowelCount * outputLength * parseInt(xtra)
    const newValue = result + newConsCount + newVowelCount
    const md59 = GenerateMD5(newValue)
    const md510 = GenerateMD5(newConsCount.toString() + newVowelCount.toString())

    const p61 = md59.toString().substr(0, 8)
    const p62 = md59.toString().substr(8, 8)
    const p63 = md59.toString().substr(16, 8)
    const p64 = md59.toString().substr(24, 8)

    const p61x = md510.toString().substr(0, 8)
    const p62x = md510.toString().substr(8, 8)
    const p63x = md510.toString().substr(16, 8)
    const p64x = md510.toString().substr(24, 8)

    result = result + p62x + p62 + p61x + p61 + p64x + p64 + p63x + p63
  }

  return result
}

export function comparePassword(dbPassword: string, strPassword: string, salt: string, iteration = 1, bits: 128 | 256 | 512 | 1024 | 2048) {
  let result
  const estrPassword = hash(strPassword, salt, iteration, bits)

  if (dbPassword === estrPassword) {
    result = true
  } else {
    result = false
  }
  return result
}

export function encrypt(text: string) {
  const key = CryptoJS.enc.Utf8.parse(hash('OCBODocTrack2023', 'AESKey', 4, 128))
  const iv = CryptoJS.enc.Utf8.parse(hash('OCBODocTrack2023', 'AESIV', 6, 128))
  const encrypted = CryptoJS.AES.encrypt(text, key, {
    mode: CryptoJS.mode.CTR,
    padding: CryptoJS.pad.Iso10126,
    iv: iv
  }).toString()

  const countEqual = encrypted.match(/=/g).length;
  const removeEqual = encrypted.slice(0, - countEqual)
  const newEncrypted = removeEqual + countEqual

  return newEncrypted
}

export function decrypt(ciphertext: string) {
  const key = CryptoJS.enc.Utf8.parse(hash('OCBODocTrack2023', 'AESKey', 4, 128))
  const iv = CryptoJS.enc.Utf8.parse(hash('OCBODocTrack2023', 'AESIV', 6, 128))

  const lastDigit = ciphertext.slice(-1)
  const removeLastDigit = ciphertext.slice(0, - 1)
  const equalString = lastDigit === '2' ? '==' : '='
  const newCiphertext = removeLastDigit + equalString

  const decrypted = CryptoJS.AES.decrypt(newCiphertext, key, {
    mode: CryptoJS.mode.CTR,
    padding: CryptoJS.pad.Iso10126,
    iv: iv
  }).toString(CryptoJS.enc.Utf8)

  return decrypted
}
