import axios from 'axios';

function getToken() {
    const token = localStorage.getItem('auth-token')

    return token
}

// axios.defaults.baseURL = 'http://localhost:1010/'
// axios.defaults.headers.common = {'Authorization': `Bearer ${getToken()}`}

import { BASE_URL } from '@/main';

const SESSION_EXPIRED_STATUS_CODE = 401;

const baseApiClient = axios.create({
    baseURL: BASE_URL,
    headers: {'Authorization': `Bearer ${getToken()}`},
});

const apiClient = ({ ...options }) => {
    const onSuccess = (response: any) => response;

    // const onError = (error: { response: { status: number; }; }) => {
    //     if (error.response.status === SESSION_EXPIRED_STATUS_CODE) {
    //         // Navigate to Login screen
    //     }

    //     return Promise.reject(error);
    // };

    return baseApiClient(options)
        .then(onSuccess);
        // .catch(onError);
};

export default apiClient;

// export default axios;