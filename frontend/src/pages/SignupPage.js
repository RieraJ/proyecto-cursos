import React, { useState } from 'react';
import { signup } from '../services/userService';

const SignupPage = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [name, setName] = useState('');
  const [surname, setSurname] = useState('');
  const [userType, setUserType] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const dataToSend = {
        // Your data here, for example:
        name: 'jjjjjj',
        email: 'jjjjjj',
        userType: 'admin',
        surname: 'jjjjjj',
        password: 'jjjjjj',
        // Add other fields as necessary
      };
      // const response = await signup({ email, password, name, surname, userType });
      fetch('http://localhost:4000/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(dataToSend),
    })
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then(data => {
        console.log('Success:', data);
        // setEmail(data.email);
      })
      .catch(error => {
        console.error('corssssssssssssssssssssssssssssss', error);
      });
      
    } catch (error) {
      console.error('Error signing up:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Email:</label>
        <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
      </div>
      <div>
        <label>Password:</label>
        <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
      </div>
      <div>
        <label>Name:</label>
        <input type="text" value={name} onChange={(e) => setName(e.target.value)} />
      </div>
      <div>
        <label>Surname:</label>
        <input type="text" value={surname} onChange={(e) => setSurname(e.target.value)} />
      </div>
      <div>
        <label>User Type:</label>
        <input type="text" value={userType} onChange={(e) => setUserType(e.target.value)} />
      </div>
      <button type="submit">Sign Up</button>
    </form>
  );
};

export default SignupPage;
