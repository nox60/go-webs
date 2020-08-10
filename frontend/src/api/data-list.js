import request from '@/utils/request'

export function getSampleData(query) {
  return request({
    url: '/listSampleData',
    method: 'post',
    data: query
  })
}

export function addItem(data) {
  return request({
    url: '/addItem',
    method: 'post',
    data: data
  })
}

export function deleteItem(data) {
  return request({
    url: '/deleteItem',
    method: 'delete',
    data: data
  })
}
