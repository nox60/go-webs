import request from '@/utils/request'
import qs from 'qs'

// export function login(data) {
//   return request({
//     url: '/vue-element-admin/user/login',
//     method: 'post',
//     data
//   })
// }

export function getInfo() {
  return request({
    url: '/info',
    method: 'get',
  })
}

export function logout() {
  return request({
    url: '/vue-element-admin/user/logout',
    method: 'post'
  })
}

export function checkLogin(data) {
  return request({
    url: '/checkLogin',
    method: 'post',
    data: {
      username: data.username,
      password: data.password
    }
  })
}

export function login(data) {
  return request({
    url: '/checkLogin',
    method: 'post',
    data: {
      username: data.username,
      password: data.password
    }
  })
}
