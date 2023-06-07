import request from '@/utils/request';

export function checkDetil(obj) {
  return request({
    url: '/admin/wage',
    method: 'post',
    data: obj
  })
}

export function lookUp() {
  return request({
    url: '/admin/lookup',
    method: 'get'
  })
}
