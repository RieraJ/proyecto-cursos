import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://localhost:4000', // Cambia el puerto si es necesario
  withCredentials: true, // Esto permite que las cookies se envíen con cada solicitud
});

export default instance;
