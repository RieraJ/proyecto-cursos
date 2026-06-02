import React, { useState } from 'react';
import './SignUp.css';
import { useNavigate } from 'react-router-dom';
import Swal from 'sweetalert2';
import { API_URL } from '../config';

function SignupForm() {
  const [loading, setLoading] = useState(false);
  const [formData, setFormData] = useState({
    name: '',
    surname: '',
    email: '',
    password: '',
    confirmPassword: ''
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prevState => ({
      ...prevState,
      [name]: value
    }));
  };

  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (formData.password.length < 8) {
      Swal.fire({ icon: 'warning', title: 'Contraseña muy corta', text: 'La contraseña debe tener al menos 8 caracteres.' });
      return;
    }
    if (formData.password !== formData.confirmPassword) {
      Swal.fire({ icon: 'warning', title: 'Las contraseñas no coinciden', text: 'Verificá que ambas contraseñas sean iguales.' });
      return;
    }

    setLoading(true);
    const { confirmPassword, ...payload } = formData;

    try {
      const response = await fetch(`${API_URL}/signup`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
      });

      const data = await response.json();

      if (response.ok) {
        await Swal.fire({
          icon: 'success',
          title: '¡Registro Exitoso!',
          text: 'Te has registrado correctamente.',
          confirmButtonText: 'Ir a Login'
        });
        navigate('/login');
      } else {
        Swal.fire({
          icon: 'error',
          title: 'Error de Registro',
          text: data.error || 'Ya existe un usuario con ese email.'
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
    <div id="signup">
      <form id="formSignUp" onSubmit={handleSubmit}>
        <h1>Crear Cuenta</h1>
        <div className="inputContainer">
          <label htmlFor="name">Nombre</label>
          <input
            type="text"
            className="inputSignUp"
            placeholder="Tu nombre"
            id="name"
            name="name"
            value={formData.name}
            onChange={handleChange}
            required
          />
        </div>
        <div className="inputContainer">
          <label htmlFor="surname">Apellido</label>
          <input
            type="text"
            className="inputSignUp"
            placeholder="Tu apellido"
            id="surname"
            name="surname"
            value={formData.surname}
            onChange={handleChange}
            required
          />
        </div>
        <div className="inputContainer">
          <label htmlFor="email">Email</label>
          <input
            type="email"
            className="inputSignUp"
            placeholder="tu@email.com"
            id="email"
            name="email"
            value={formData.email}
            onChange={handleChange}
            required
          />
        </div>
        <div className="inputContainer">
          <label htmlFor="password">Contraseña (mín. 8 caracteres)</label>
          <input
            type="password"
            className="inputSignUp"
            placeholder="Mínimo 8 caracteres"
            id="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            minLength={8}
            required
          />
        </div>
        <div className="inputContainer">
          <label htmlFor="confirmPassword">Confirmar contraseña</label>
          <input
            type="password"
            className="inputSignUp"
            placeholder="Repetí tu contraseña"
            id="confirmPassword"
            name="confirmPassword"
            value={formData.confirmPassword}
            onChange={handleChange}
            required
          />
        </div>
        <button type="submit" className="submit-btn" disabled={loading}>
          {loading ? 'Registrando...' : 'Crear Cuenta'}
        </button>
      </form>
    </div>
  );
}

export default SignupForm;
