import request from '@/utils/request'

//     data
// export function login(data) {
//   return request({
//     url: '/vue-element-admin/user/login',
//     method: 'post',
//   })
// }

export function getInfo() {
  return request({
    url: '/userInfo',
    method: 'get'
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

export function listUserData(query) {
  return request({
    url: '/listUserData',
    method: 'post',
    data: query
  })
}

export function addOrUpdateUser(data) {
  return request({
    url: '/addOrUpdateUser',
    method: 'post',
    data: data
  })
}

export function deleteUser(id) {
  return request({
    url: `/deleteUser/${id}`,
    method: 'delete'
  })
}

export function resetUser(id) {
  return request({
    url: `/resetUser/${id}`,
    method: 'put'
  })
}
