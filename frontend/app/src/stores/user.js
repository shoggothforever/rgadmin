import { defineStore } from 'pinia';
import { getToken, setToken, removeToken } from '../utils/auth';
import { login, getUserInfo, logout } from '../api/user';
import { resetRouter } from '../router';

export const useUserStore = defineStore('user', {
  state: () => ({
    token: getToken(),
    name: '',
    roles: []
  }),
  actions: {
    SET_TOKEN(token) {
      this.token = token;
    },
    SET_NAME(name) {
      this.name = name;
    },
    SET_ROLES(roles) {
      this.roles.push(roles);
    },
    CLEAR_ROLES() {
      this.roles = [];
    },
    // 用户登录
    login(userInfo) {
      const { staffCode, password } = userInfo;
      return new Promise((resolve, reject) => {
        login({ staffCode: staffCode.trim(), password: password })
          .then((res) => {
            this.SET_TOKEN(res.accessToken);
            setToken(res.accessToken);
            resolve();
          })
          .catch((error) => {
            reject(error);
          });
      });
    },

    // 获取用户信息
    getInfo() {
      return new Promise((resolve, reject) => {
        getUserInfo()
          .then((res) => {
            console.log('申请用户信息成功了', res)
            if (!res) {
              reject('认证失败，请重新登录');
            }
            const name = res.name
            const roles = Array.of(res.role)
            if (!roles || roles.length <= 0) {
              reject('getInfo: 角色必须是非空数组!');
            }
            this.SET_NAME(name);
            this.SET_ROLES(roles);
            resolve(res);
          })
          .catch((error) => {
            reject(error);
          });
      });
    },

    // 用户登出
    logout() {
      return new Promise((resolve, reject) => {
        logout()
          .then(() => {
            this.SET_TOKEN('');
            this.CLEAR_ROLES();
            removeToken();
            resetRouter();
            resolve();
          })
          .catch((error) => {
            reject(error);
          });
      });
    },

    // 删除token
    resetToken() {
      return new Promise((resolve) => {
        this.SET_TOKEN('');
        this.CLEAR_ROLES();
        removeToken();
        resolve();
      });
    }
  }
});
