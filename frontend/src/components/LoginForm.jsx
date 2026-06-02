import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './Login.css';
import Swal from 'sweetalert2';
import { API_URL } from '../config';

function LoginForm() {
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);
  const [user, setUser] = useState({
    email: '',
    password: ''
});

const handleChange = (e) => {
    const { name, value } = e.target;
    setUser(prevState => ({
        ...prevState,
        [name]: value
    }));
};

const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);

    try {
        const response = await fetch(`${API_URL}/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(user),
            credentials: 'include'
        });

        const data = await response.json();

        if (response.ok) {
            await Swal.fire({
                icon: 'success',
                title: '¡Ingreso Exitoso!',
                text: 'Has iniciado sesión correctamente.',
                timer: 1500,
                showConfirmButton: false
            });
            navigate('/');
        } else {
            Swal.fire({
                icon: 'error',
                title: 'Error de Login',
                text: data.error || 'Email o contraseña inválidos.'
            });
        }
    } catch (error) {
        Swal.fire({
            icon: 'error',
            title: 'Error',
            text: 'Ha ocurrido un error inesperado.'
        });
    } finally {
        setLoading(false);
    }
};


  return (
    <div id="login">
      <form id="formLogin" onSubmit={handleSubmit}>
        <h1>Iniciar Sesión</h1>
        <p className="form-subtitle">Accedé a tu cuenta de EduCursos</p>
        <div className="inputContainer">
          <label htmlFor="email">Email</label>
          <input
            type="email"
            className="inputLogin"
            placeholder="tu@email.com"
            id="email"
            name='email'
            onChange={handleChange}
            value={user.email}
            required
          />
        </div>
        <div className="inputContainer">
          <label htmlFor="password">Contraseña</label>
          <input
            type="password"
            className="inputLogin"
            placeholder="Tu contraseña"
            id="password"
            name='password'
            onChange={handleChange}
            value={user.password}
            required
          />
        </div>
        <button type="submit" className="submit-btn" disabled={loading}>
          {loading ? 'Ingresando...' : 'Iniciar Sesión'}
        </button>
        <p className='signup'>¿No tenés cuenta? <a href="/signup">Registrate</a></p>
      </form>
    </div>
  );
}

export default LoginForm;