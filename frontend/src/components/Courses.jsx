import React, { useState, useEffect, useCallback } from 'react';
import Cookies from 'js-cookie';
import './Courses.css';
import { useNavigate } from 'react-router-dom';

//const API_BASE_URL = process.env.REACT_APP_API_URL || "http://localhost:4000";

const Courses = () => {
    const navigate = useNavigate();
    const [searchTerm, setSearchTerm] = useState('');
    const [courses, setCourses] = useState([]);
    const [error, setError] = useState(null);
    const [enrollmentMessage, setEnrollmentMessage] = useState('');
    const [userType, setUserType] = useState(null); // Estado para el tipo de usuario

    const userId = Cookies.get('userId');

    const fetchUserInfo = useCallback(async () => {
        try {
            const response = await fetch(`http://localhost:4000/user-info`, {
                credentials: 'include',
            });
            if (!response.ok) throw new Error('Failed to fetch user info');

            const data = await response.json();
            setUserType(data.userInfo.userType);
        } catch (err) {
            console.error('Error fetching user info:', err);
        }
    }, []);

    const formatLength = (length) => {
        if (!length || length === "N/A") return "N/A";

        const parts = length.split(':');
        const hours = parseInt(parts[0], 10) || 0;
        const minutes = parseInt(parts[1], 10) || 0;
        const seconds = parseInt(parts[2], 10) || 0;

        return `${hours} hours ${minutes} minutes ${seconds} seconds`;
    };

    const fetchCourses = useCallback(async (url) => {
        try {
            const response = await fetch(url, { credentials: 'include' });
            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.error || 'Error fetching courses');
            }

            const data = await response.json();
            if (!data || !data.courses) {
                throw new Error('No hay cursos disponibles');
            }

            const formattedCourses = data.courses.map(course => ({
                ...course,
                length: formatLength(course.length),
                categories: course.categories.map(cat => cat.name),
            }));
            setCourses(formattedCourses);
            setError(null);
        } catch (err) {
            console.error(err);
            setError(err.message);
            setCourses([]);
        }
    }, []);

    useEffect(() => {
        fetchUserInfo(); // Llama al endpoint para obtener el tipo de usuario
        fetchCourses(`http://localhost:4000/courses`);
    }, [fetchUserInfo, fetchCourses]);

    const handleSearch = (e) => {
        e.preventDefault();
        const url = searchTerm
            ? `http://localhost:4000/search-courses?name=${encodeURIComponent(searchTerm)}`
            : `http://localhost:4000/courses`;
        fetchCourses(url);
    };

    const handleEnroll = async (courseId) => {
        try {
            const response = await fetch(`http://localhost:4000/enroll`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
                body: JSON.stringify({ user_id: userId, course_id: courseId })
            });
            if (response.ok) {
                setEnrollmentMessage('Successfully enrolled!');
                alert('Successfully enrolled!');
            } else {
                const errorData = await response.json();
                if (response.status === 400 && errorData.error === 'user is already enrolled in this course') {
                    alert('You are already enrolled in this course.');
                } else {
                    alert(errorData.error || 'Failed to enroll');
                }
            }
        } catch (err) {
            console.error('Error enrolling:', err);
            setEnrollmentMessage('Error enrolling');
        }
    };

    const handleViewComments = (courseId) => {
        navigate(`/course/${courseId}/comments`);
    };

    const handleCreateCourse = () => {
        navigate('/create-course');
    };

    return (
        <div className="courses-container">
            {/* Mostrar botón si el usuario es admin */}
            {userType === 'admin' && (
                <button
                    className="create-course-button"
                    onClick={handleCreateCourse}
                >
                    Create Course
                </button>
            )}

            <form className="search-form" onSubmit={handleSearch}>
                <input
                    type="text"
                    value={searchTerm}
                    onChange={(e) => setSearchTerm(e.target.value)}
                    placeholder="Enter course name"
                    className="search-input"
                />
                <button type="submit" className="search-button">Search</button>
            </form>

            {error && <p className="error-message">{error}</p>}
            {enrollmentMessage && <p className="success-message">{enrollmentMessage}</p>}

            <ul className='course-list'>
                {courses.length > 0 ? (
                    courses.map((course) => (
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
                            <p className="price">
                                Price: ${course.price ? course.price.toFixed(2) : "N/A"}
                            </p>
                            <p className="instructor">
                                <strong>Instructor:</strong> {course.instructor ? course.instructor : "N/A"}
                            </p>
                            <p className="length">
                                <strong>Length:</strong> {course.length ? course.length : "N/A"}
                            </p>
                            <p className="requirements">
                                <strong>Requirements:</strong> {course.requirements ? course.requirements : "N/A"}
                            </p>
                            <button
                                className="enroll-button"
                                onClick={() => handleEnroll(course.id)}
                            >
                                Enroll
                            </button>
                            <button
                                className="view-comments-button"
                                onClick={() => handleViewComments(course.id)}
                            >
                                View Comments
                            </button>
                        </li>
                    ))
                ) : (
                    !error && <p>No courses available at the moment.</p>
                )}
            </ul>
        </div>
    );
};

export default Courses;
