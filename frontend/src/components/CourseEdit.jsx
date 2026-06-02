import React, { useState, useEffect, useRef } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import Swal from 'sweetalert2';
import './CreateCourse.css';
import './CourseEdit.css';
import { API_URL } from '../config';
import { isValidLength, validateImageFile } from '../utils';

const toBase64 = (file) =>
    new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = () => resolve(reader.result.split(',')[1]);
        reader.onerror = reject;
    });

const CourseEdit = () => {
    const { id } = useParams();
    const navigate = useNavigate();

    const [loading, setLoading] = useState(true);
    const [saving, setSaving] = useState(false);
    const [error, setError] = useState('');
    const [formError, setFormError] = useState('');
    const imagePreviewUrlRef = useRef(null);

    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [price, setPrice] = useState('');
    const [instructor, setInstructor] = useState('');
    const [length, setLength] = useState('');
    const [requirements, setRequirements] = useState('');
    const [categories, setCategories] = useState([{ name: '' }]);
    const [currentImage, setCurrentImage] = useState('');
    const [newImageFile, setNewImageFile] = useState(null);
    const [imagePreview, setImagePreview] = useState(null);

    useEffect(() => {
        const fetchCourse = async () => {
            try {
                const res = await fetch(`${API_URL}/courses`, { credentials: 'include' });
                if (!res.ok) throw new Error('Error al obtener los cursos');
                const data = await res.json();
                const course = (data.courses || []).find(c => c.id === parseInt(id));
                if (!course) throw new Error('Curso no encontrado');

                setName(course.name || '');
                setDescription(course.description || '');
                setPrice(course.price != null ? String(course.price) : '');
                setInstructor(course.instructor || '');
                setLength(course.length || '');
                setRequirements(course.requirements || '');
                setCategories(
                    course.categories && course.categories.length > 0
                        ? course.categories.map(c => ({ name: typeof c === 'string' ? c : c.name }))
                        : [{ name: '' }]
                );
                setCurrentImage(course.image || '');
            } catch (err) {
                setError(err.message);
            } finally {
                setLoading(false);
            }
        };
        fetchCourse();
    }, [id]);

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
        setNewImageFile(file);
        setImagePreview(url);
    };

    const handleCategoryChange = (index, value) => {
        const updated = [...categories];
        updated[index].name = value;
        setCategories(updated);
    };

    const handleAddCategory = () => setCategories(prev => [...prev, { name: '' }]);

    const handleRemoveCategory = (index) => {
        setCategories(prev => prev.filter((_, i) => i !== index));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setFormError('');

        if (!name || !description || !price || !instructor || !length) {
            setFormError('Completá todos los campos obligatorios.');
            return;
        }
        if (parseFloat(price) < 0) {
            setFormError('El precio no puede ser negativo.');
            return;
        }
        if (!isValidLength(length)) {
            setFormError('La duración debe tener el formato HH:MM:SS (ej: 02:30:00).');
            return;
        }
        const validCats = categories.filter(c => c.name.trim() !== '');
        if (validCats.length === 0) {
            setFormError('Agregá al menos una categoría.');
            return;
        }
        setSaving(true);

        let imageBase64 = currentImage;
        if (newImageFile) {
            imageBase64 = await toBase64(newImageFile);
        }

        const body = {
            name,
            description,
            price: parseFloat(price),
            instructor,
            length,
            requirements,
            active: true,
            categories: validCats,
            image: imageBase64,
        };

        const res = await fetch(`${API_URL}/courses/${id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            credentials: 'include',
            body: JSON.stringify(body),
        });

        if (res.ok) {
            Swal.fire({ icon: 'success', title: 'Curso actualizado', timer: 1500, showConfirmButton: false });
            navigate('/courses');
        } else {
            const data = await res.json();
            setFormError(data.error || 'Error al actualizar el curso.');
        }
        setSaving(false);
    };

    if (loading) return <div className="course-edit-loading">Cargando curso...</div>;
    if (error) return <div className="course-edit-error">{error}</div>;

    return (
        <div className="course-edit-container">
            <div className="course-edit-card">
                <div className="course-edit-header">
                    <h1>Editar Curso</h1>
                    <p>Modificá los datos del curso y guardá los cambios</p>
                </div>

                {formError && <div className="course-edit-form-error">{formError}</div>}

                <form className="course-edit-form" onSubmit={handleSubmit}>

                    <div className="form-group full-width">
                        <label htmlFor="edit-name">Nombre del Curso <span className="required">*</span></label>
                        <input
                            id="edit-name"
                            type="text"
                            value={name}
                            onChange={e => setName(e.target.value)}
                            placeholder="Nombre del curso"
                            required
                        />
                    </div>

                    <div className="form-group full-width">
                        <label htmlFor="edit-description">Descripción <span className="required">*</span></label>
                        <textarea
                            id="edit-description"
                            rows={4}
                            value={description}
                            onChange={e => setDescription(e.target.value)}
                            placeholder="Descripción del curso"
                            required
                        />
                    </div>

                    <div className="form-row">
                        <div className="form-group">
                            <label htmlFor="edit-price">Precio (USD) <span className="required">*</span></label>
                            <input
                                id="edit-price"
                                type="number"
                                min="0"
                                step="0.01"
                                value={price}
                                onChange={e => setPrice(e.target.value)}
                                placeholder="0.00"
                                required
                            />
                        </div>
                        <div className="form-group">
                            <label htmlFor="edit-instructor">Instructor <span className="required">*</span></label>
                            <input
                                id="edit-instructor"
                                type="text"
                                value={instructor}
                                onChange={e => setInstructor(e.target.value)}
                                placeholder="Nombre del instructor"
                                required
                            />
                        </div>
                    </div>

                    <div className="form-group full-width">
                        <label htmlFor="edit-length">Duración (HH:MM:SS) <span className="required">*</span></label>
                        <input
                            id="edit-length"
                            type="text"
                            value={length}
                            onChange={e => setLength(e.target.value)}
                            placeholder="02:30:00"
                            pattern="\d{2}:\d{2}:\d{2}"
                            required
                        />
                    </div>

                    <div className="form-group full-width">
                        <label htmlFor="edit-requirements">Requisitos</label>
                        <textarea
                            id="edit-requirements"
                            rows={3}
                            value={requirements}
                            onChange={e => setRequirements(e.target.value)}
                            placeholder="Conocimientos previos necesarios (opcional)"
                        />
                    </div>

                    <div className="form-group full-width">
                        <label>Imagen del Curso</label>
                        <label htmlFor="edit-image" className="image-upload-label">
                            {imagePreview ? (
                                <img src={imagePreview} alt="Nueva imagen" className="image-preview" />
                            ) : currentImage ? (
                                <img src={`data:image/png;base64,${currentImage}`} alt="Imagen actual" className="image-preview" />
                            ) : (
                                <div className="image-upload-placeholder">
                                    <span className="upload-icon">📷</span>
                                    <span>Hacé clic para cambiar la imagen</span>
                                </div>
                            )}
                        </label>
                        <input
                            type="file"
                            id="edit-image"
                            accept="image/*"
                            onChange={handleImageChange}
                            className="image-input-hidden"
                        />
                    </div>

                    <div className="form-group full-width categories-section">
                        <label>Categorías <span className="required">*</span></label>
                        <div className="categories-list">
                            {categories.map((cat, index) => (
                                <div key={index} className="category-input">
                                    <input
                                        type="text"
                                        value={cat.name}
                                        onChange={e => handleCategoryChange(index, e.target.value)}
                                        placeholder={`Categoría ${index + 1}`}
                                    />
                                    {categories.length > 1 && (
                                        <button
                                            type="button"
                                            className="remove-category-btn"
                                            onClick={() => handleRemoveCategory(index)}
                                            aria-label="Eliminar categoría"
                                        >
                                            ✕
                                        </button>
                                    )}
                                </div>
                            ))}
                        </div>
                        <button type="button" className="add-category-btn" onClick={handleAddCategory}>
                            + Agregar Categoría
                        </button>
                    </div>

                    <div className="course-edit-actions">
                        <button type="button" className="cancel-btn" onClick={() => navigate('/courses')}>
                            Cancelar
                        </button>
                        <button type="submit" className="submit-course-btn" disabled={saving}>
                            {saving ? 'Guardando...' : 'Guardar Cambios'}
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
};

export default CourseEdit;
