import axios from 'axios';
import { getToken } from '@/utils/auth';
import { useUserStore } from '../stores/user';
import { ElMessage, ElMessageBox } from 'element-plus';

// 创建axios实例
const service = axios.create({
  baseURL: 'http://rg.sfct.top', // url = base url + request url
  // http://rg.sfct.top
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000
});

// 请求体拦截器
service.interceptors.request.use(
  (config) => {
    const store = useUserStore();
    if (store.token) {
      //   console.log('把token放到请求头里面了!');
      config.headers['Authorization'] = `Bearer ${getToken()}`;
    }
    return config;
  },
  (error) => {
    console.log(error);
    return Promise.reject(error);
  }
);

// 返回体拦截器
service.interceptors.response.use(
  (response) => {
    const store = useUserStore();
    const res = response.data;
    console.log(res)
    // 如果返回的状态码不是200，就判断为错误
    if (res.code !== 200) {
      ElMessage({
        message: res.message || 'Error',
        type: 'error',
        duration: 5 * 1000
      });
      // 返回401说明token已经失效了
      if (res.code === 401) {
        ElMessageBox.confirm(
          '出现了一些问题，你需要刷新一下吗?',
          {
            confirmButtonText: '刷新',
            cancelButtonText: '取消',
            type: 'warning'
          }
        ).then(() => {
          // store.dispatch('user/resetToken').then(() => {
          //   location.reload();
          // });
          store.resetToken().then(() => {
            location.reload();
          });
        });
      }
      return Promise.reject(new Error(res.msg || 'Error'));
    } else {
      return res;
    }
  },
  (error) => {
    console.log('err' + error); // for debug
    ElMessage({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    });
    return Promise.reject(error);
  }
);

export default service;
