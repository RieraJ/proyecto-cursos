import React, { useState, useEffect } from 'react';
import Cookies from 'js-cookie';
import './Courses.css';

const Courses = () => {
    const [searchTerm, setSearchTerm] = useState('');
    const [courses, setCourses] = useState([]);
    const [error, setError] = useState(null);
    const [enrollmentMessage, setEnrollmentMessage] = useState('');

    const userId = Cookies.get('userId');
    
    const fetchCourses = async (url) => {
        try {
            const response = await fetch(url, { credentials: 'include' });
            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.error || 'Error fetching courses');
            }
            const data = await response.json();
            const formattedCourses = (data.courses || []).map(course => ({
                ...course,
                length: formatLength(course.length),
            }));
            setCourses(formattedCourses);
            setError(null);
        } catch (err) {
            if (err.message === 'no courses found') {
                setError('No courses found');
            } else {
                setError(err.message);
            }
            setCourses([]);
        }
    };

    const formatLength = (length) => {
        if (!length || length === "N/A") return "N/A";
        
        const parts = length.split(':');
        const hours = parseInt(parts[0], 10) || 0;
        const minutes = parseInt(parts[1], 10) || 0;
        const seconds = parseInt(parts[2], 10) || 0;
    
        return `${hours} hours ${minutes} minutes ${seconds} seconds`;
    };

    const handleSearch = (e) => {
        e.preventDefault();
        const url = searchTerm
            ? `http://localhost:4000/search-courses?name=${encodeURIComponent(searchTerm)}`
            : `http://localhost:4000/courses`;
        fetchCourses(url);
    };

    useEffect(() => {
        fetchCourses('http://localhost:4000/courses');
    }, []);

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

    return (
        <div className="courses-container">
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
                            <p className="price">Price: ${course.price ? course.price.toFixed(2) : "N/A"}</p>
                            <p className="instructor"><strong>Instructor:</strong> {course.instructor ? course.instructor : "N/A"}</p>
                            <p className="length"><strong>Length:</strong> {course.length ? course.length : "N/A"}</p>
                            <p className="requirements"><strong>Requirements:</strong> {course.requirements ? course.requirements : "N/A"}</p>
                            <button className="enroll-button" onClick={() => handleEnroll(course.id)}>Enroll</button>
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
