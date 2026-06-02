import React, { useState, useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import Swal from 'sweetalert2';
import './CreateCourse.css';
import { API_URL } from '../config';
import { isValidLength, validateImageFile } from '../utils';

const CreateCourse = () => {
    const navigate = useNavigate();
    const [courseName, setCourseName] = useState('');
    const [description, setDescription] = useState('');
    const [price, setPrice] = useState('');
    const [instructor, setInstructor] = useState('');
    const [length, setLength] = useState('');
    const [requirements, setRequirements] = useState('');
    const [categories, setCategories] = useState([{ name: '' }]);
    const [imageFile, setImageFile] = useState(null);
    const [imagePreview, setImagePreview] = useState(null);
    const [error, setError] = useState('');
    const [loading, setLoading] = useState(false);
    const imagePreviewUrlRef = useRef(null);

    const handleAddCategory = () => {
        setCategories([...categories, { name: '' }]);
    };

    const handleCategoryChange = (index, value) => {
        const newCategories = [...categories];
        newCategories[index].name = value;
        setCategories(newCategories);
    };

    const handleRemoveCategory = (index) => {
        const newCategories = categories.filter((_, i) => i !== index);
        setCategories(newCategories);
    };

    const handleImageChange = (e) => {
        const file = e.target.files[0];
        if (!file) return;
        const validationError = validateImageFile(file);
        if (validationError) {
            Swal.fire({ icon: 'error', title: 'Imagen inválida', text: validationError });
            e.target.value = '';
            return;
        }
        if (imagePreviewUrlRef.current) URL.revokeObjectURL(imagePreviewUrlRef.current);
        const url = URL.createObjectURL(file);
        imagePreviewUrlRef.current = url;
        setImageFile(file);
        setImagePreview(url);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');

        if (!courseName || !description || !price || !instructor || !length) {
            setError('Por favor completá todos los campos requeridos');
            return;
        }
        if (parseFloat(price) < 0) {
            setError('El precio no puede ser negativo');
            return;
        }
        if (!isValidLength(length)) {
            setError('La duración debe tener el formato HH:MM:SS (ej: 02:30:00)');
            return;
        }

        const validCategories = categories.filter(cat => cat.name.trim() !== '');
        if (validCategories.length === 0) {
            setError('Agregá al menos una categoría');
            return;
        }

        setLoading(true);
        const formData = new FormData();
        formData.append('name', courseName);
        formData.append('description', description);
        formData.append('price', price);
        formData.append('instructor', instructor);
        formData.append('length', length);
        formData.append('requirements', requirements);
        formData.append('active', 'true');
        formData.append('categories', JSON.stringify(validCategories));
        if (imageFile) {
            formData.append('image', imageFile);
        }

        try {
            const response = await fetch(`${API_URL}/courses`, {
                method: 'POST',
                credentials: 'include',
                body: formData,
            });

            const data = await response.json();

            if (response.ok) {
                navigate('/courses');
            } else {
                setError(data.error || 'Error al crear el curso');
            }
        } catch (err) {
            setError('Error de red. Intentá de nuevo.');
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="create-course-container">
            <div className="create-course-card">
                <div className="create-course-header">
                    <h1>Crear Nuevo Curso</h1>
                    <p>Completá los datos para publicar tu curso</p>
                </div>

                {error && <div className="form-error">{error}</div>}

                <form onSubmit={handleSubmit} className="create-course-form">
                    <div className="form-group full-width">
                        <label htmlFor="courseName">
                            Nombre del Curso <span className="required">*</span>
                        </label>
                        <input
                            type="text"
                            id="courseName"
                            value={courseName}
                            onChange={(e) => setCourseName(e.target.value)}
                            placeholder="Ej: Introducción a Python"
                            required
                        />
                    </div>

                    <div className="form-group full-width">
                        <label htmlFor="description">
                            Descripción <span className="required">*</span>
                        </label>
                        <textarea
                            id="description"
                            value={description}
                            onChange={(e) => setDescription(e.target.value)}
                            placeholder="Describí el contenido y objetivos del curso"
                            rows={4}
                            required
                        />
                    </div>

                    <div className="form-row">
                        <div className="form-group">
                            <label htmlFor="price">
                                Precio (USD) <span className="required">*</span>
                            </label>
                            <input
                                type="number"
                                id="price"
                                value={price}
                                onChange={(e) => setPrice(e.target.value)}
                                placeholder="0.00"
                                min="0"
                                step="0.01"
                                required
                            />
                        </div>
                        <div className="form-group">
                            <label htmlFor="instructor">
                                Instructor <span className="required">*</span>
                            </label>
                            <input
                                type="text"
                                id="instructor"
                                value={instructor}
                                onChange={(e) => setInstructor(e.target.value)}
                                placeholder="Nombre del instructor"
                                required
                            />
                        </div>
                    </div>

                    <div className="form-group full-width">
                        <label htmlFor="length">
                            Duración (HH:MM:SS) <span className="required">*</span>
                        </label>
                        <input
                            type="text"
                            id="length"
                            value={length}
                            onChange={(e) => setLength(e.target.value)}
                            placeholder="02:30:00"
                            pattern="\d{2}:\d{2}:\d{2}"
                            required
                        />
                    </div>

                    <div className="form-group full-width">
                        <label htmlFor="requirements">Requisitos</label>
                        <textarea
                            id="requirements"
                            value={requirements}
                            onChange={(e) => setRequirements(e.target.value)}
                            placeholder="Conocimientos previos necesarios (opcional)"
                            rows={3}
                        />
                    </div>

                    <div className="form-group full-width">
                        <label>Imagen del Curso</label>
                        <label htmlFor="courseImage" className="image-upload-label">
                            {imagePreview ? (
                                <img src={imagePreview} alt="Preview" className="image-preview" />
                            ) : (
                                <div className="image-upload-placeholder">
                                    <span className="upload-icon">📷</span>
                                    <span>Hacé clic para seleccionar una imagen</span>
                                </div>
                            )}
                        </label>
                        <input
                            type="file"
                            id="courseImage"
                            accept="image/*"
                            onChange={handleImageChange}
                            className="image-input-hidden"
                        />
                    </div>

                    <div className="form-group full-width categories-section">
                        <label>Categorías <span className="required">*</span></label>
                        <div className="categories-list">
                            {categories.map((category, index) => (
                                <div key={index} className="category-input">
                                    <input
                                        type="text"
                                        value={category.name}
                                        onChange={(e) => handleCategoryChange(index, e.target.value)}
                                        placeholder={`Categoría ${index + 1}`}
                                    />
                                    {categories.length > 1 && (
                                        <button
                                            type="button"
                                            onClick={() => handleRemoveCategory(index)}
                                            className="remove-category-btn"
                                            aria-label="Eliminar categoría"
                                        >
                                            ✕
                                        </button>
                                    )}
                                </div>
                            ))}
                        </div>
                        <button
                            type="button"
                            onClick={handleAddCategory}
                            className="add-category-btn"
                        >
                            + Agregar Categoría
                        </button>
                    </div>

                    <button type="submit" className="submit-course-btn" disabled={loading}>
                        {loading ? 'Publicando...' : 'Publicar Curso'}
                    </button>
                </form>
            </div>
        </div>
    );
};

export default CreateCourse;
