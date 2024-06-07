import React, { useEffect, useState } from 'react';
import { getUserCourses } from '../services/userService';

const CoursesPage = () => {
  const [courses, setCourses] = useState([]);
  const userId = 1; // Replace with the actual user ID or get it from the auth context or state

  useEffect(() => {
    const fetchCourses = async () => {
      try {
        const response = await getUserCourses(userId);
        setCourses(response.data);
      } catch (error) {
        console.error('Error fetching courses:', error);
      }
    };

    fetchCourses();
  }, [userId]);

  return (
    <div>
      <h1>Courses</h1>
      <ul>
        {courses.map((course) => (
          <li key={course.id}>{course.name}</li>
        ))}
      </ul>
    </div>
  );
};

export default CoursesPage;
