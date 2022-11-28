const GenerateMD5 = require('md5')

export function encrypt(Password: string) {
  const salt = 'ocboSalt@2023'
  const md51 = GenerateMD5(salt + Password + salt)
  let vowelCount = 0
  let consCount = 0

  const p1 = md51.substr(0, 8)
  const p2 = md51.substr(8, 8)
  const p3 = md51.substr(16, 8)
  const p4 = md51.substr(24, 8)

  for (let i = 0; i < Password.length; i++) {
    if (Password.charAt(i) === 'a' || Password.charAt(i) === 'A' || Password.charAt(i) === 'e' || Password.charAt(i) === 'E' || Password.charAt(i) === 'i' || Password.charAt(i) === 'I' || Password.charAt(i) === 'o' || Password.charAt(i) === 'O' || Password.charAt(i) === 'u' || Password.charAt(i) === 'U') {
      vowelCount++
    } else {
      consCount++
    }
  }

  const xtra = consCount.toString() + vowelCount.toString()
  const newConsCount = consCount * parseInt(xtra)
  const newVowelCount = vowelCount * parseInt(xtra)
  const newValue = Password + newConsCount + newVowelCount
  const md52 = GenerateMD5(newValue)

  const p21 = md52.substr(0, 8)
  const p22 = md52.substr(8, 8)
  const p23 = md52.substr(16, 8)
  const p24 = md52.substr(24, 8)

  return p2 + p22 + p1 + p21 + p4 + p24 + p3 + p23
}

export function comparePassword(dbPassword: string, strPassword: string) {
  let result
  const estrPassword = encrypt(strPassword)

  if (dbPassword === estrPassword) {
    result = true
  } else {
    result = false
  }
  return result
}
