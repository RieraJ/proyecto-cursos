import React, { useState, useEffect } from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';
import Main from './components/Main';
import SignupForm from './components/SignUpForm';
import LoginForm from './components/LoginForm';
import './App.css';
import Navbar from './components/Navbar';
import Courses from './components/Courses';
import Profile from './components/Profile';
import CourseComments from './components/CourseComments';
import CreateCourse from './components/CreateCourse';
import AdminPanel from './components/AdminPanel';
import CourseEdit from './components/CourseEdit';
import UserEdit from './components/UserEdit';
import { ThemeContext } from './ThemeContext';

function App() {
  const [theme, setTheme] = useState(() => localStorage.getItem('theme') || 'dark');

  useEffect(() => {
    document.documentElement.setAttribute('data-theme', theme);
    localStorage.setItem('theme', theme);
  }, [theme]);

  const toggleTheme = () => setTheme(t => t === 'dark' ? 'light' : 'dark');

  return (
    <ThemeContext.Provider value={{ theme, toggleTheme }}>
      <div>
        <Navbar />
        <main className='main-content'>
          <Routes>
            <Route path='/course/:courseId/comments' element={<CourseComments />} />
            <Route path="/signup" element={<SignupForm />} />
            <Route path="/login" element={<LoginForm />} />
            <Route path="/" element={<Main />} />
            <Route path="/courses" element={<Courses />} />
            <Route path="/profile" element={<Profile />} />
            <Route path="/create-course" element={<CreateCourse />} />
            <Route path="/admin" element={<AdminPanel />} />
            <Route path="/courses/:id/edit" element={<CourseEdit />} />
            <Route path="/profile/edit" element={<UserEdit />} />
            <Route path="*" element={<Navigate to="/" replace />} />
          </Routes>
        </main>
      </div>
    </ThemeContext.Provider>
  );
}

export default App;
