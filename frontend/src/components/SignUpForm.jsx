import React, { useState } from 'react';
import './SignUp.css';
import { useNavigate } from 'react-router-dom';

function SignupForm() {
  const [formData, setFormData] = useState({
    name: '',
    surname: '',
    email: '',
    password: ''
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

    try {
      const response = await fetch('http://localhost:4000/signup', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData)
      });

      const data = await response.json();

      if (response.ok) {
        console.log(data);
        alert('Te has registrado correctamente. Por favor, logeate para continuar');
        navigate('/login');
      } else {
        alert('Ya existe un usuario con ese email, intenta con otro email');
      }
    } catch (error) {
      console.error('Error:', error);
        alert('Ha ocurrido un error');
    }
  };

  return (
    <div id="signup">
      <form id="formSignUp" onSubmit={handleSubmit}>
        <h1>Sign Up</h1>
        <div className="inputContainer">
          <input
            type="text"
            className="inputSignUp"
            placeholder=" "
            id="name"
            name="name"
            value={formData.name}
            onChange={handleChange}
            required
          />
          <label htmlFor="name" className="labelSignUp">Name</label>
        </div>
        <div className="inputContainer">
          <input
            type="text"
            className="inputSignUp"
            placeholder=" "
            id="surname"
            name="surname"
            value={formData.surname}
            onChange={handleChange}
            required
          />
          <label htmlFor="surname" className="labelSignUp">Surname</label>
        </div>
        <div className="inputContainer">
          <input
            type="email"
            className="inputSignUp"
            placeholder=" "
            id="email"
            name="email"
            value={formData.email}
            onChange={handleChange}
            required
          />
          <label htmlFor="email" className="labelSignUp">Email</label>
        </div>
        <div className="inputContainer">
          <input
            type="password"
            className="inputSignUp"
            placeholder=" "
            id="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            required
          />
          <label htmlFor="password" className="labelSignUp">Password</label>
        </div>
        <button type="submit" className="submit-btn">Sign Up</button>
      </form>
    </div>
  );
}

export default SignupForm;
