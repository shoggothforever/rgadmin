import router from './router';
import { ElMessage } from 'element-plus';
import { getToken } from './utils/auth';
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import { useUserStore } from '@/stores/user.js';
import { usePermissionStore } from '@/stores/permission.js';

NProgress.configure({ showSpinner: false });

const whiteList = ['/login'];

router.beforeEach(async(to, from, next) => {
  const userStore = useUserStore();
  const permissionStore = usePermissionStore();
  // 启动进度条
  NProgress.start();
  // 设置页面标题
  document.title = to.meta.title;
  // 获取Token
  const hasToken = getToken();
  // 判断是否有Token
  if (hasToken) {
    if (to.path === '/login') {
      // 如果有Token，但是要去的页面是登录页，直接跳转到首页
      next({ path: '/' });
      // 关闭进度条
      NProgress.done();
    } else {
      // 获取用户的角色
      const hasRoles = userStore.roles && userStore.roles.length > 0;
      console.log('hasRoles', hasRoles)
      console.log('userStore.roles', userStore.roles)
      console.log('userStore.roles.length', userStore.roles.length)
      // console.log('storeRoles', userStore.roles, 'storeRoleslength', userStore.roles.length)
      if (hasRoles) {
        console.log('登录进来去了哪儿1')
        next();
      } else {
        try {
          console.log('登录进来去了哪儿2')
          // 如果没有角色，就向服务器申请
          // 请注意：角色必须是对象数组！例如：['admin']或，['developer'，'editor']
          const userInfoRes = await userStore.getInfo();
          const roles = Array.of(userInfoRes.role);
          console.log('roles', roles)
          // 动态添加可访问的路由
          permissionStore.addRoutesToRouter(router, roles).then(() => {
            console.log('动态添加可访问的路由成功了')
          });
          console.log(permissionStore.routes)

          next({ ...to, replace: true });
        } catch (error) {
          // 如果出错，就重置Token并跳转到登录页
          await userStore.resetToken();
          ElMessage.error(error || 'Has Error');
          // 跳转到登录页
          next('/login');
          NProgress.done();
        }
      }
    }
  } else {
    // 如果没有Token
    if (whiteList.indexOf(to.path) !== -1) {
      // 如果要去的页面在白名单中，直接跳转
      next();
    } else {
      // 如果要去的页面不在白名单中，跳转到登录页
      next('/login');
      NProgress.done();
    }
  }
});

router.afterEach(() => {
  // 关闭进度条
  NProgress.done();
});
