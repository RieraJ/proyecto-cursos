import React, { useState, useEffect } from 'react';
import Cookies from 'js-cookie';
import './Profile.css';

const Profile = () => {
    const [courses, setCourses] = useState([]);
    const [error, setError] = useState(null);
    const [userId, setUserId] = useState(null);

    const fetchUserInfo = async () => {
        try {
            const response = await fetch('http://localhost:4000/user-info', { credentials: 'include' });
            if (!response.ok) {
                throw new Error('Error fetching user info');
            }
            const data = await response.json();
            setUserId(data.userInfo.id); // Guardar el userId en el estado
        } catch (err) {
            console.error('Error fetching user info:', err);
        }
    };

    const fetchUserCourses = async () => {
        if (!userId) return; // Asegurarse de que userId esté disponible
        try {
            const response = await fetch(`http://localhost:4000/users/${userId}/courses`, { credentials: 'include' });
            if (!response.ok) {
                throw new Error('Error fetching user courses');
            }
            const data = await response.json();
            setCourses(data.courses);
        } catch (err) {
            setError('No courses found for the user');
            setCourses([]);
        }
    };

    useEffect(() => {
        fetchUserInfo();
    }, []);

    useEffect(() => {
        fetchUserCourses();
    }, [userId]);

    return (
        <div className="profile-container">
            <h1>Your Courses</h1>
            {error && <p className="error-message">{error}</p>}
            <ul className="course-list">
                {courses.map((course) => (
                    <li key={course.id} className="course-card">
                        <h3>{course.name ? course.name : "No name available"}</h3>
                        <p>{course.description ? course.description : "No description available"}</p>
                        <p className="price">Price: ${course.price ? course.price.toFixed(2) : "N/A"}</p>
                        <p className="instructor"><strong>Instructor:</strong> {course.instructor ? course.instructor : "N/A"}</p>
                        <p className="length"><strong>Length:</strong> {course.length ? course.length : "N/A"}</p>
                        <p className="requirements"><strong>Requirements:</strong> {course.requirements ? course.requirements : "N/A"}</p>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default Profile;
