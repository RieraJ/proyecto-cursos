import React, { useState } from 'react';
import './Login.css'
import { useNavigate } from 'react-router-dom';
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"></link>

function LoginForm() {
  const [user, setUser] = useState({
    email: '',
    password: ''
});

const navigate = useNavigate();

const handleChange = (e) => {
    const { name, value } = e.target;
    setUser(prevState => ({
        ...prevState,
        [name]: value
    }));
};

const handleSubmit = async (e) => {
    e.preventDefault();

    try {
        const response = await fetch('http://localhost:4000/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(user),
            credentials: 'include'
        });

        const data = await response.json();

        if (response.ok) {
            console.log(data);
            alert('You have logged in successfully');
            navigate('/');
        } else {
            alert('Invalid email or password');
        }
    } catch (error) {
        console.error('Error:', error);
        alert('An error occurred');
    }
};


  return (
    <div id="login">
      <form id="formLogin" onSubmit={handleSubmit}>
        <h1>Login</h1>
        <div className="inputContainer">
          <input
            type="email"
            className="inputLogin"
            placeholder=" "
            id="email"
            name='email'
            onChange={handleChange}
            value={user.email}
            required
          />
          <label htmlFor="email" className="labelLogin">Email</label>
        </div>
        <div className="inputContainer">
          <input
            type="password"
            className="inputLogin"
            placeholder=" "
            id="password"
            name='password'
            onChange={handleChange}
            value={user.password}
            required
          />
          <label htmlFor="password" className="labelLogin">Password</label>
        </div>
        <button type="submit" className="submit-btn">Login</button>
      </form>
    </div>
  );
}

export default LoginForm;