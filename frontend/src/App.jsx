// App.jsx

import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Main from './components/Main';
import SignupForm from './components/SignUpForm';
import LoginForm from './components/LoginForm';
import './App.css';
import Navbar from './components/Navbar';
import Courses from './components/Courses';
import Profile from './components/Profile';
import CourseComments from './components/CourseComments';
import CreateCourse from './components/CreateCourse';

function App() {

  return (
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
        </Routes>
      </main>
    </div>
  );
}

export default App;
