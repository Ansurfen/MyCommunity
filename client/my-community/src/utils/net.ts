import axios from "axios";

let env = 'dev';
let baseURL;
if (env === 'dev') {
    baseURL = 'http://localhost:9090'
}

const service = axios.create({
    baseURL
})

axios.interceptors.request.use(config => {
    console.log(config);
    return config;
}, err => {
    console.log(err);
    return Promise.reject(err);
})

export default service