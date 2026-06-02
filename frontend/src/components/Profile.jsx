import React, { useState, useEffect, useCallback } from 'react';
import { useNavigate } from 'react-router-dom';
import { FaDollarSign, FaUserTie, FaClock, FaListUl, FaComments } from 'react-icons/fa';
import './Profile.css';
import './Courses.css';
import { API_URL } from '../config';
import { formatLength } from '../utils';

const Profile = () => {
    const navigate = useNavigate();
    const [courses, setCourses] = useState([]);
    const [error, setError] = useState(null);
    const [userId, setUserId] = useState(null);
    const [selectedCourse, setSelectedCourse] = useState(null);

    const fetchUserInfo = async () => {
        try {
            const response = await fetch(`${API_URL}/user-info`, { credentials: 'include' });
            if (!response.ok) throw new Error('Error fetching user info');
            const data = await response.json();
            setUserId(data.userInfo.id);
        } catch (err) {
            console.error('Error fetching user info:', err);
        }
    };

    const fetchUserCourses = useCallback(async () => {
        if (!userId) return;
        try {
            const response = await fetch(`${API_URL}/users/${userId}/courses`, { credentials: 'include' });
            if (!response.ok) throw new Error('Error fetching user courses');
            const data = await response.json();
            const formattedCourses = data.courses.map(course => ({
                ...course,
                length: formatLength(course.length),
                categories: course.categories.map(cat => cat.name),
            }));
            setCourses(formattedCourses);
        } catch (err) {
            setError('No courses found for the user');
            setCourses([]);
        }
    }, [userId]);

    useEffect(() => { fetchUserInfo(); }, []);
    useEffect(() => { fetchUserCourses(); }, [userId, fetchUserCourses]);

    const openModal = (course) => setSelectedCourse(course);
    const closeModal = () => setSelectedCourse(null);

    const handleViewComments = (e, courseId) => {
        e.stopPropagation();
        navigate(`/course/${courseId}/comments`);
    };

    return (
        <div className="courses-container">
            <header className="courses-page-header">
                <h1 className="courses-page-title">Mis <span>Cursos</span></h1>
                <p className="courses-page-subtitle">Los cursos en los que estás inscripto</p>
            </header>

            {error && <p className="error-message">{error}</p>}

            <ul className="course-list">
                {courses.length > 0 ? courses.map((course) => (
                    <li
                        key={course.id}
                        className="course-card"
                        onClick={() => openModal(course)}
                    >
                        <div className="course-card-image">
                            {course.image ? (
                                <img
                                    src={`data:image/png;base64,${course.image}`}
                                    alt={course.name}
                                    className="course-img"
                                />
                            ) : (
                                <div className="course-img-placeholder" />
                            )}
                        </div>
                        <div className="course-card-body">
                            <h3 className="course-card-title">{course.name || "Sin título"}</h3>
                        </div>
                    </li>
                )) : (
                    !error && <p className="courses-empty">Todavía no estás inscripto en ningún curso.</p>
                )}
            </ul>

            {selectedCourse && (
                <div
                    className="course-modal-overlay"
                    onClick={closeModal}
                    role="dialog"
                    aria-modal="true"
                    aria-label={`Detalle: ${selectedCourse.name}`}
                >
                    <div className="course-modal" onClick={(e) => e.stopPropagation()}>
                        <button
                            className="course-modal-close"
                            onClick={closeModal}
                            aria-label="Cerrar modal"
                        >
                            ✕
                        </button>

                        <div className="course-modal-image">
                            {selectedCourse.image ? (
                                <img
                                    src={`data:image/png;base64,${selectedCourse.image}`}
                                    alt={selectedCourse.name}
                                    className="course-modal-img"
                                />
                            ) : (
                                <div className="course-img-placeholder course-modal-img-placeholder" />
                            )}
                        </div>

                        <div className="course-modal-body">
                            <h2 className="course-modal-title">{selectedCourse.name}</h2>
                            <p className="course-modal-description">{selectedCourse.description}</p>

                            <div className="course-modal-meta">
                                <span className="course-modal-price">
                                    <FaDollarSign />
                                    {selectedCourse.price ? selectedCourse.price.toFixed(2) : "N/A"}
                                </span>
                                <span className="course-modal-instructor">
                                    <FaUserTie /> {selectedCourse.instructor || "N/A"}
                                </span>
                            </div>

                            <div className="course-modal-fields">
                                <p className="course-modal-field">
                                    <FaClock className="field-icon" />
                                    <span>{selectedCourse.length || "N/A"}</span>
                                </p>
                                {selectedCourse.requirements && (
                                    <p className="course-modal-field">
                                        <FaListUl className="field-icon" />
                                        <span>{selectedCourse.requirements}</span>
                                    </p>
                                )}
                            </div>

                            {selectedCourse.categories.length > 0 && (
                                <div className="course-modal-categories">
                                    {selectedCourse.categories.map((cat, i) => (
                                        <span key={i} className="course-category-pill">{cat}</span>
                                    ))}
                                </div>
                            )}

                            <div className="course-modal-actions">
                                <button
                                    className="view-comments-button"
                                    onClick={(e) => handleViewComments(e, selectedCourse.id)}
                                >
                                    <FaComments /> Comentarios
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
};

export default Profile;
