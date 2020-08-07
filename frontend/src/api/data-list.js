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
    url: '/AddItem',
    method: 'post',
    data: data
  })
}
