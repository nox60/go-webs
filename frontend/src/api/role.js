import request from '@/utils/request'

export function getRoutes() {
  return request({
    url: '/vue-element-admin/routes',
    method: 'get'
  })
}

export function getRoles() {
  return request({
    url: '/vue-element-admin/roles',
    method: 'get'
  })
}

export function addRole(data) {
  return request({
    url: '/vue-element-admin/role',
    method: 'post',
    data
  })
}

export function updateRole(id, data) {
  return request({
    url: `/vue-element-admin/role/${id}`,
    method: 'put',
    data
  })
}

export function getFunctions(id) {
  return request({
    url: `/getFunctions/${id}`,
    method: 'get',
  })
}

export function getNodes(id) {
  return request({
    url: `/getNodes/${id}`,
    method: 'get',
  })
}

export function deleteFunction(id) {
  return request({
    url: `/deleteFunction/${id}`,
    method: 'delete'
  })
}

export function addOrUpdateFunction(data) {
  return request({
    url: '/addOrUpdateFunction',
    method: 'post',
    data: data
  })
}

export function addOrUpdateFunctionItem(data) {
  return request({
    url: '/addOrUpdateFunctionItem',
    method: 'post',
    data: data
  })
}

export function getFunctionById(id) {
  return request({
    url: `/getFunctionById/${id}`,
    method: 'get',
  })
}

export function addOrUpdateRole(data) {
  return request({
    url: '/addOrUpdateRole',
    method: 'post',
    data: data
  })
}

export function listRoleData(query) {
  return request({
    url: '/listRoleData',
    method: 'post',
    data: query
  })
}

export function deleteRole(id) {
  return request({
    url: `/deleteRole/${id}`,
    method: 'delete'
  })
}

export function getRoleById(id) {
  return request({
    url: `/getRoleById/${id}`,
    method: 'get',
  })
}

export function getAllFuncs() {
  return request({
    url: `/pid/0`,
    method: 'get',
  })
}

export function deleteFunctionItem(id) {
  return request({
    url: `/deleteFunctionItem/${id}`,
    method: 'delete'
  })
}

export function getFunctionItemById(id) {
  return request({
    url: `/getFunctionItemById/${id}`,
    method: 'get',
  })
}
