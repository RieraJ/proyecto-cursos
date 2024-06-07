import axios from '../api/axios';

export const signup = async (userData) => {
  return axios.post('/signup', userData);
};

export const login = async (credentials) => {
  return axios.post('/login', credentials);
};

export const getUserCourses = async (userId) => {
  return axios.get(`/users/${userId}/courses`);
};
