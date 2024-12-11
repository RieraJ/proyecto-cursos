import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import './Profile.css';

// const API_BASE_URL = process.env.REACT_APP_API_URL || "http://localhost:4000";

const Profile = () => {
    const navigate = useNavigate();
    const [courses, setCourses] = useState([]);
    const [error, setError] = useState(null);
    const [userId, setUserId] = useState(null);

    const formatLength = (length) => {
        if (!length || length === "N/A") return "N/A";

        const parts = length.split(':');
        const hours = parseInt(parts[0], 10) || 0;
        const minutes = parseInt(parts[1], 10) || 0;
        const seconds = parseInt(parts[2], 10) || 0;

        return `${hours} hours ${minutes} minutes ${seconds} seconds`;
    };

    const fetchUserInfo = async () => {
        try {
            const response = await fetch(`http://localhost:4000/user-info`, { credentials: 'include' });
            if (!response.ok) {
                throw new Error('Error fetching user info');
            }
            const data = await response.json();
            setUserId(data.userInfo.id);
        } catch (err) {
            console.error('Error fetching user info:', err);
        }
    };

    const fetchUserCourses = async () => {
        if (!userId) return;
        try {
            const response = await fetch(`http://localhost:4000/users/${userId}/courses`, { credentials: 'include' });
            if (!response.ok) {
                throw new Error('Error fetching user courses');
            }
            const data = await response.json();
            
            const formattedCourses = data.courses.map(course => ({
                ...course,
                length: formatLength(course.length),
                categories: course.categories.map(cat => cat.name), // Extraer solo nombres de categorías
            }));

            setCourses(formattedCourses);
        } catch (err) {
            setError('No courses found for the user');
            setCourses([]);
        }
    };

    const handleViewComments = (courseId) => {
        navigate(`/course/${courseId}/comments`);
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
            
            <ul className='course-list'>
                {courses.map((course) => (
                    <li key={course.id} className="course-card">
                        <h3>{course.name ? course.name : "No name available"}</h3>
                        <p>{course.description ? course.description : "No description available"}</p>
                        <p className="category">
                            <strong>Categories:</strong>
                            {course.categories.length > 0 ? (
                                <ul>
                                    {course.categories.map((cat, index) => (
                                        <li key={index}>{cat}</li>
                                    ))}
                                </ul>
                            ) : (
                                "No categories available"
                            )}
                        </p>
                        <p className="price">Price: ${course.price ? course.price.toFixed(2) : "N/A"}</p>
                        <p className="instructor"><strong>Instructor:</strong> {course.instructor ? course.instructor : "N/A"}</p>
                        <p className="length"><strong>Length:</strong> {course.length ? course.length : "N/A"}</p>
                        <p className="requirements"><strong>Requirements:</strong> {course.requirements ? course.requirements : "N/A"}</p>
                        <div className="course-actions">
                            <button 
                                className="view-comments-button" 
                                onClick={() => handleViewComments(course.id)}
                            >
                                View Comments
                            </button>
                        </div>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default Profile;