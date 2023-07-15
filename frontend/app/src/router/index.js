import { createRouter, createWebHistory } from 'vue-router';

export const constantRoutes = [
  {
    path: '/',
    redirect: '/dashboard',
    component: () => import('@/layout/index.vue'),
    children: [
      {
        path: 'dashboard',
        component: () => import('@/views/Dashboard/index.vue'),
        name: 'Dashboard',
        meta: { title: '首页' }
      }
    ]
  },
  {
    path: '/login',
    component: () => import('@/views/login/index.vue'),
    hidden: true
  },
  {
    path: '/profile',
    redirect: '/profile/index',
    component: () => import('@/layout/index.vue'),
    children: [
      {
        path: 'index',
        component: () => import('@/views/Profile/index.vue'),
        name: 'Profile',
        meta: { title: '动态数据展示页面' }
      }
    ]
  }

];

export const asyncRoutes = [
  {
    path: '/admin',
    redirect: '/admin/profile',
    component: () => import('@/layout/index.vue'),
    meta: { roles: ['finance staff'], title: '管理员' },
    children: [
      {
        path: 'wage',
        component: () => import('@/views/Admin/Wage/index.vue'),
        name: 'Wage',
        meta: { title: '获取工资报表' }
      },
      {
        path: 'upload',
        component: () => import('@/views/Admin/upload/index.vue'),
        name: 'Upload',
        meta: { title: '上传工资表' }
      }
    ]
  },
  {
    path: '/user',
    component: () => import('@/layout/index.vue'),
    meta: { title: '员工', roles: ['admissions staff', 'assistant', 'cleaner', 'computer administrator', 'counsellor', 'driver', 'filing clerk', 'hr', 'librarian', 'principal', 'restaurant staff', 'security bureau', 'teacher', 'vice-principal'] },
    children: [
      {
        path: 'wage',
        component: () => import('@/views/User/wage/index.vue'),
        name: 'Userwage',
        meta: { title: '工时上报' }
      },
      {
        path: 'query',
        component: () => import('@/views/User/query/index.vue'),
        name: 'Query',
        meta: { title: '工资查询' }
      }
    ]
  }
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: constantRoutes
});

export function resetRouter() {
  const newRouter = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: constantRoutes
  });
  router.resolve = newRouter.resolve;
}

export default router;
