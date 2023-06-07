import request from '@/utils/request';

export function login(loginForm) {
  return request({
    url: '/login',
    method: 'post',
    data: loginForm
  });
}

export function getUserInfo() {
  return request({
    url: '/user/userinfo',
    method: 'get'
  });
}

export function logout() {
  return request({
    url: '/logout',
    method: 'post'
  });
}

export function uploadWorkTime(workTime) {
  return request({
    url: '/user/wage',
    method: 'post',
    data: workTime
  })
}

export function queryWage(data) {
  return request({
    url: '/user/querywage',
    method: 'post',
    data: data
  })
}
