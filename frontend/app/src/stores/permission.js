import { defineStore } from 'pinia';
import { constantRoutes, asyncRoutes } from '../router';

function hasPermission(roles, route) {
  if (route.meta && route.meta.roles) {
    return roles.some((role) => route.meta.roles.includes(role));
  } else {
    return true;
  }
}

export function filterAsyncRoutes(routes, roles) {
  const res = [];
  routes.forEach((route) => {
    const tmp = { ...route };
    if (hasPermission(roles, tmp)) {
      if (tmp.children) {
        tmp.children = filterAsyncRoutes(tmp.children, roles);
      }
      res.push(tmp);
    }
  });
  return res;
}

export const usePermissionStore = defineStore('permission', {
  state: () => ({
    routes: [],
    addRoutes: [],
    routesAdded: false // Add this
  }),
  actions: {
    SET_ROUTES(routes) {
      this.addRoutes = routes;
      this.routes = constantRoutes.concat(routes);
    },
    RESET_ROUTES() {
      this.addRoutes = [];
      this.routes = [];
      this.routesAdded = false;
    },
    generateRoutes(roles) {
      return new Promise((resolve) => {
        const accessedRoutes = filterAsyncRoutes(asyncRoutes, roles);
        console.log('accessedRoutes', accessedRoutes)
        console.log('可以添加的路由是', accessedRoutes)
        this.SET_ROUTES(accessedRoutes);
        resolve(accessedRoutes);
      });
    },

    async addRoutesToRouter(router, roles) {
    // Check if routes have been added
      console.log('进入了addRoutesToRouter')
      if (this.routesAdded) return;

      // get dynamic routes
      console.log('准备执行generateRoutes')
      const routes = await this.generateRoutes(roles);

      // Define addRouteRecursive function here
      const addRouteRecursive = (route) => {
        console.log('Adding route: ', route);

        try {
          router.addRoute(route);
        } catch (error) {
          console.error('Error adding route: ', route, error);
        }
      };

      routes.forEach((route) => addRouteRecursive(route));

      this.routesAdded = true; // Set this to true after adding routes

      console.log('路由有:', router.getRoutes());
    }
  }
});
