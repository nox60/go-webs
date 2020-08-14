import request from '@/utils/request'

export function getSampleData(query) {
  return request({
    url: '/listSampleData',
    method: 'post',
    data: query
  })
}

export function getItem(itemId) {
  return request({
    url: '/getItem/'+itemId,
    method: 'get',
  })
}

export function addOrUpdateItem(data) {
  return request({
    url: '/addOrUpdateItem',
    method: 'post',
    data: data
  })
}

export function deleteItem(data) {
  return request({
    url: '/deleteItem/'+data,
    method: 'delete',
  })
}
