import request from '@/utils/request'
import qs from 'qs'

// export function login(data) {
//   return request({
//     url: '/vue-element-admin/user/login',
//     method: 'post',
//     data
//   })
// }

export function getInfo(token) {
  return request({
    url: '/vue-element-admin/user/info',
    method: 'get',
    params: { token }
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
