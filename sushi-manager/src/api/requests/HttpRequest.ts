/* eslint-disable */
import type { AxiosRequestConfig, AxiosResponse } from "axios";
import axios from "axios";

const instance = axios.create();

// Add a request interceptor
instance.interceptors.request.use(
  function (config) {
    // TODO: externalize this??
    config.baseURL = `http://localhost:8081`;
    return config;
  },
  function (error) {
    // Do something with request error
    return Promise.reject(error);
  }
);

export default {
  get: <T = any, R = AxiosResponse<T>>(
    url: string,
    config?: AxiosRequestConfig
  ): Promise<R> => instance.get<T, R>(url, config),
  post: <T = any, R = AxiosResponse<T>>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig
  ): Promise<R> => instance.post<T, R>(url, data, config),
  put: <T = any, R = AxiosResponse<T>>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig
  ): Promise<R> => instance.put<T, R>(url, data, config),
  delete: <T = any, R = AxiosResponse<T>>(
    url: string,
    config?: AxiosRequestConfig
  ): Promise<R> => instance.delete<T, R>(url, config),
  patch: <T = any, R = AxiosResponse<T>>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig
  ): Promise<R> => instance.patch<T, R>(url, data, config),
};
