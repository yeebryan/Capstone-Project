// authAxios.js

import axios from 'axios';

const authAxios = axios.create();

authAxios.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers['token'] = token;
  }
  return config;
}, (error) => {
  return Promise.reject(error);
});

export default authAxios;