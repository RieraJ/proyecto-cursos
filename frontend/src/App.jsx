// App.jsx

import React, { useState } from 'react';
import { Routes, Route } from 'react-router-dom';
import Main from './components/Main';
import SignupForm from './components/SignUpForm';
import LoginForm from './components/LoginForm';
import './App.css';
import Navbar from './components/Navbar';
import Courses from './components/Courses';

function App() {

  return (
    <div>
      <Navbar />
      <main className='main-content'>
        <Routes>
          <Route path="/signup" element={<SignupForm />} />
          <Route path="/login" element={<LoginForm />} />
          <Route path="/" element={<Main />} />
          <Route path="/courses" element={<Courses />} />
        </Routes>
      </main>
    </div>
  );
}

export default App;
