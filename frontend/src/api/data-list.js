import request from '@/utils/request'

export function getSampleData(query) {
  return request({
    url: '/listSampleData',
    method: 'post',
    params: query
  })
}
