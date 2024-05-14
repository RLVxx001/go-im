// axios-instance.js  
import axios from 'axios';  
  
// 创建一个 axios 实例  
const service = axios.create({  
});  
  
// 请求拦截器  
service.interceptors.request.use(  
  config => {  
    // 从 cookie、localStorage 或其他地方获取 token  
    const token = localStorage.getItem('token');  
    if (token) {  
      // 如果 token 存在，则将其添加到请求头中  
      config.headers['Authorization'] = `Bearer ${token}`;  
    }  
    return config;  
  },  
  error => {  
    // 对请求错误做些什么  
    return Promise.reject(error);  
  }  
);  
  
// 响应拦截器（如果需要的话）  
// service.interceptors.response.use(response => {}, error => {});  
  
export default service;