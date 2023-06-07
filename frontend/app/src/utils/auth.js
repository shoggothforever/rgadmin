import Cookie from 'js-cookie';

const Tokenkey = 'jwt';

export function getToken() {
  return Cookie.get(Tokenkey);
}

export function setToken(token) {
  return Cookie.set(Tokenkey, token);
}

export function removeToken() {
  return Cookie.remove(Tokenkey);
}
