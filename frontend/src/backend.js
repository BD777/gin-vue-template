import axios from 'axios'
import router from './router'

let $axios = axios.create({
  baseURL: '/api/',
  timeout: 20000,
  headers: { 'Content-Type': 'application/json' }
})

// Request Interceptor
$axios.interceptors.request.use(function (config) {
  // config.headers['Authorization'] = 'Fake Token'
  return config
})

$axios.interceptors.response.use(function (response) {
  return response
}, function (error) {
  if (error.response) {
    if (error.response.status === 401) {
      router.replace({ path: '/login' })
    } else if (error.response.status === 403) {
      router.replace({ path: '/' })
    }
  }
  return Promise.reject(error)
})

export default {
  async getLoginInfo() {
    return $axios.get(`admin/logininfo`).then(response => response.data)
  },

  async login(username, password) {
    return $axios.post(`admin/login`, {
      username,
      password
    }).then(response => response.data)
  },

  async getPageAuth() {
    return $axios.get(`admin/pageauth`).then(response => response.data)
  }
}
