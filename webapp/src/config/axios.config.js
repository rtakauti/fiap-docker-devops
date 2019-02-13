import axios from 'axios';

axios.interceptors.request.use(config => {
  config.url = `/api${config.url}`;
  return config;
});
