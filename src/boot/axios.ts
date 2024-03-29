import { boot } from 'quasar/wrappers'
import axios, { AxiosInstance } from 'axios'
import axiosRetry from 'axios-retry'

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $axios: AxiosInstance
  }
}

// Be careful when using SSR for cross-request state pollution
// due to creating a Singleton instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)

const api = axios.create({
  baseURL: '',
  headers: {
    'X-Requested-With': 'XMLHttpRequest',
  },
})

// const api = axios.create({
//   baseURL: 'http://192.168.7.160:8080/doctrack-server/',
//   headers: {
//     'X-Requested-With': 'XMLHttpRequest',
//   },
// })

// const api = axios.create({
//   baseURL: 'https://ncyhljvhdyblusgeqzre.supabase.co/rest/v1',
//   headers: {
//     'X-Requested-With': 'XMLHttpRequest',
//     Authorization: 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Im5jeWhsanZoZHlibHVzZ2VxenJlIiwicm9sZSI6ImFub24iLCJpYXQiOjE2OTYzNDAyNDgsImV4cCI6MjAxMTkxNjI0OH0.1ENfnJGp10NYpkxpy15PXD--cyvNJUQJbk6hwzvvum0',
//     apikey: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Im5jeWhsanZoZHlibHVzZ2VxenJlIiwicm9sZSI6ImFub24iLCJpYXQiOjE2OTYzNDAyNDgsImV4cCI6MjAxMTkxNjI0OH0.1ENfnJGp10NYpkxpy15PXD--cyvNJUQJbk6hwzvvum0',
//   },
// })

axiosRetry(api, { retries: 3 })

export default boot(({ app }) => {
  // for use inside Vue files (Options API) through this.$axios and this.$api

  app.config.globalProperties.$axios = axios
  // ^ ^ ^ this will allow you to use this.$axios (for Vue Options API form)
  //       so you won't necessarily have to import axios in each vue file

  app.config.globalProperties.$api = api
  // ^ ^ ^ this will allow you to use this.$api (for Vue Options API form)
  //       so you can easily perform requests against your app's API
})

export { api }
