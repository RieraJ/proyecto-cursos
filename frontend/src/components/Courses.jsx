import React, { useState, useEffect, useCallback } from 'react';
import Cookies from 'js-cookie';
import './Courses.css';
import { useNavigate } from 'react-router-dom';
import Swal from 'sweetalert2';
import { FaSearch, FaPlus, FaDollarSign, FaUserTie, FaClock, FaListUl, FaComments, FaEdit } from 'react-icons/fa';
import { API_URL } from '../config';
import { formatLength } from '../utils';

const Courses = () => {
    const navigate = useNavigate();
    const [searchTerm, setSearchTerm] = useState('');
    const [courses, setCourses] = useState([]);
    const [error, setError] = useState(null);
    const [, setEnrollmentMessage] = useState('');
    const [userType, setUserType] = useState(null);
    const [selectedCourse, setSelectedCourse] = useState(null);

    const userId = Cookies.get('userId');

    const fetchUserInfo = useCallback(async () => {
        try {
            const response = await fetch(`${API_URL}/user-info`, {
                credentials: 'include',
            });
            if (!response.ok) throw new Error('Failed to fetch user info');
            const data = await response.json();
            setUserType(data.userInfo.userType);
        } catch (err) {
            console.error('Error fetching user info:', err);
        }
    }, []);

    const fetchCourses = useCallback(async (url) => {
        try {
            const response = await fetch(url, { credentials: 'include' });
            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.error || 'Error fetching courses');
            }
            const data = await response.json();
            if (!data || !data.courses) throw new Error('No hay cursos disponibles');
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
        fetchUserInfo();
        fetchCourses(`${API_URL}/courses`);
    }, [fetchUserInfo, fetchCourses]);

    const handleSearch = (e) => {
        e.preventDefault();
        const url = searchTerm
            ? `${API_URL}/search-courses?name=${encodeURIComponent(searchTerm)}`
            : `${API_URL}/courses`;
        fetchCourses(url);
    };

    const handleEnroll = async (e, courseId) => {
        e.stopPropagation();
        try {
            const response = await fetch(`${API_URL}/enroll`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                credentials: 'include',
                body: JSON.stringify({ user_id: userId, course_id: courseId })
            });
            if (response.ok) {
                setEnrollmentMessage('Successfully enrolled!');
                Swal.fire({
                    icon: 'success',
                    title: '¡Inscripción Exitosa!',
                    text: 'Te has inscrito en este curso correctamente.',
                    timer: 2000,
                    showConfirmButton: false
                });
            } else {
                const errorData = await response.json();
                if (response.status === 400 && errorData.error === 'user is already enrolled in this course') {
                    Swal.fire({ icon: 'warning', title: 'Ya Inscrito', text: 'Ya te encuentras registrado en este curso.' });
                } else {
                    Swal.fire({ icon: 'error', title: 'Error de Inscripción', text: errorData.error || 'Failed to enroll' });
                }
            }
        } catch (err) {
            console.error('Error enrolling:', err);
            setEnrollmentMessage('Error enrolling');
        }
    };

    const handleViewComments = (e, courseId) => {
        e.stopPropagation();
        navigate(`/course/${courseId}/comments`);
    };

    const handleCreateCourse = () => navigate('/create-course');

    const openModal = (course) => setSelectedCourse(course);
    const closeModal = () => setSelectedCourse(null);

    const handleEditCourse = (courseId) => {
        closeModal();
        navigate(`/courses/${courseId}/edit`);
    };

    return (
        <div className="courses-container">

            <header className="courses-page-header">
                <h1 className="courses-page-title">Explorar <span>Cursos</span></h1>
                <p className="courses-page-subtitle">Descubrí cursos para potenciar tus habilidades</p>
            </header>

            <div className="courses-toolbar">
                <form className="search-form" onSubmit={handleSearch}>
                    <div className="search-input-wrapper">
                        <FaSearch className="search-icon" />
                        <input
                            type="text"
                            value={searchTerm}
                            onChange={(e) => setSearchTerm(e.target.value)}
                            placeholder="Buscar por nombre de curso..."
                            className="search-input"
                        />
                    </div>
                    <button type="submit" className="search-button">Buscar</button>
                </form>

                {userType === 'admin' && (
                    <button className="create-course-button" onClick={handleCreateCourse}>
                        <FaPlus /> Nuevo Curso
                    </button>
                )}
            </div>

            {error && <p className="error-message">{error}</p>}

            <ul className="course-list">
                {courses.length > 0 ? (
                    courses.map((course) => (
                        <li key={course.id} className="course-card" onClick={() => openModal(course)}>
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
                    ))
                ) : (
                    !error && <p className="courses-empty">No hay cursos disponibles por el momento.</p>
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
                            id="close-course-modal"
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
                                    id="modal-enroll-button"
                                    className="enroll-button"
                                    onClick={(e) => handleEnroll(e, selectedCourse.id)}
                                >
                                    Inscribirse
                                </button>
                                <button
                                    id="modal-comments-button"
                                    className="view-comments-button"
                                    onClick={(e) => handleViewComments(e, selectedCourse.id)}
                                >
                                    <FaComments /> Comentarios
                                </button>
                                {userType === 'admin' && (
                                    <button
                                        id="modal-edit-button"
                                        className="edit-course-button"
                                        onClick={() => handleEditCourse(selectedCourse.id)}
                                    >
                                        <FaEdit /> Editar
                                    </button>
                                )}
                            </div>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
};

export default Courses;
