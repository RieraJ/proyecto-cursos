import React, { useState, useEffect, useCallback, useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import Swal from 'sweetalert2';
import './AdminPanel.css';
import { API_URL } from '../config';
import { validateImageFile } from '../utils';

const EMPTY_FORM = {
    name: '', description: '', price: '', instructor: '',
    length: '', requirements: '', categories: [{ name: '' }], imageFile: null,
};

const AdminPanel = () => {
    const navigate = useNavigate();
    const [activeTab, setActiveTab] = useState('courses');
    const [courses, setCourses] = useState([]);
    const [users, setUsers] = useState([]);
    const [userRoles, setUserRoles] = useState({});
    const [showCreateForm, setShowCreateForm] = useState(false);
    const [form, setForm] = useState(EMPTY_FORM);
    const [formError, setFormError] = useState('');
    const [imagePreview, setImagePreview] = useState(null);

    const fetchCourses = useCallback(async () => {
        try {
            const res = await fetch(`${API_URL}/courses`, { credentials: 'include' });
            if (!res.ok) throw new Error();
            const data = await res.json();
            setCourses(data.courses || []);
        } catch {
            setCourses([]);
        }
    }, []);

    const fetchUsers = useCallback(async () => {
        try {
            const res = await fetch(`${API_URL}/users`, { credentials: 'include' });
            if (!res.ok) {
                if (res.status === 403) { navigate('/'); return; }
                throw new Error();
            }
            const data = await res.json();
            const list = Array.isArray(data) ? data : [];
            setUsers(list);
            const roles = {};
            list.forEach(u => { roles[u.id] = u.userType; });
            setUserRoles(roles);
        } catch {
            setUsers([]);
        }
    }, [navigate]);

    useEffect(() => {
        fetchCourses();
        fetchUsers();
    }, [fetchCourses, fetchUsers]);

    // ── Cursos ──

    const handleDeleteCourse = async (id, name) => {
        const result = await Swal.fire({
            title: `¿Eliminar "${name}"?`,
            text: 'Esta acción no se puede deshacer.',
            icon: 'warning',
            showCancelButton: true,
            confirmButtonColor: '#d32f2f',
            cancelButtonText: 'Cancelar',
            confirmButtonText: 'Eliminar',
        });
        if (!result.isConfirmed) return;

        const res = await fetch(`${API_URL}/courses/${id}`, {
            method: 'DELETE',
            credentials: 'include',
        });
        if (res.ok) {
            Swal.fire({ icon: 'success', title: 'Curso eliminado', timer: 1500, showConfirmButton: false });
            fetchCourses();
        } else {
            Swal.fire({ icon: 'error', title: 'Error al eliminar el curso' });
        }
    };

    const handleFormChange = (field, value) => {
        setForm(prev => ({ ...prev, [field]: value }));
    };

    const handleCategoryChange = (index, value) => {
        const cats = [...form.categories];
        cats[index].name = value;
        setForm(prev => ({ ...prev, categories: cats }));
    };

    const handleAddCategory = () => {
        setForm(prev => ({ ...prev, categories: [...prev.categories, { name: '' }] }));
    };

    const handleRemoveCategory = (index) => {
        setForm(prev => ({
            ...prev,
            categories: prev.categories.filter((_, i) => i !== index),
        }));
    };

    const imagePreviewUrlRef = useRef(null);

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
        setForm(prev => ({ ...prev, imageFile: file }));
        setImagePreview(url);
    };

    const handleCreateCourse = async (e) => {
        e.preventDefault();
        setFormError('');
        const { name, description, price, instructor, length, categories, imageFile } = form;
        if (!name || !description || !price || !instructor || !length) {
            setFormError('Completá todos los campos obligatorios.');
            return;
        }
        const validCats = categories.filter(c => c.name.trim() !== '');
        if (validCats.length === 0) {
            setFormError('Agregá al menos una categoría.');
            return;
        }

        const formData = new FormData();
        formData.append('name', name);
        formData.append('description', description);
        formData.append('price', price);
        formData.append('instructor', instructor);
        formData.append('length', length);
        formData.append('requirements', form.requirements);
        formData.append('active', 'true');
        formData.append('categories', JSON.stringify(validCats));
        if (imageFile) formData.append('image', imageFile);

        const res = await fetch(`${API_URL}/courses`, {
            method: 'POST',
            credentials: 'include',
            body: formData,
        });

        if (res.ok) {
            Swal.fire({ icon: 'success', title: 'Curso creado', timer: 1500, showConfirmButton: false });
            setForm(EMPTY_FORM);
            setImagePreview(null);
            setShowCreateForm(false);
            fetchCourses();
        } else {
            const data = await res.json();
            setFormError(data.error || 'Error al crear el curso.');
        }
    };

    // ── Usuarios ──

    const handleRoleChange = (userId, newRole) => {
        setUserRoles(prev => ({ ...prev, [userId]: newRole }));
    };

    const handleSaveRole = async (userId) => {
        const user = users.find(u => u.id === userId);
        const result = await Swal.fire({
            title: `¿Cambiar rol de ${user?.name}?`,
            text: `Nuevo rol: ${userRoles[userId]}`,
            icon: 'warning',
            showCancelButton: true,
            confirmButtonText: 'Confirmar',
            cancelButtonText: 'Cancelar',
        });
        if (!result.isConfirmed) return;

        const res = await fetch(`${API_URL}/update-user-type`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            credentials: 'include',
            body: JSON.stringify({ user_id: userId, user_type: userRoles[userId] }),
        });
        if (res.ok) {
            Swal.fire({ icon: 'success', title: 'Rol actualizado', timer: 1500, showConfirmButton: false });
            fetchUsers();
        } else {
            Swal.fire({ icon: 'error', title: 'Error al actualizar el rol' });
        }
    };

    return (
        <div className="admin-container">
            <h1 className="admin-title">Panel Administrativo</h1>

            <div className="admin-tabs">
                <button
                    className={`admin-tab ${activeTab === 'courses' ? 'admin-tab--active' : ''}`}
                    onClick={() => setActiveTab('courses')}
                >
                    Cursos
                </button>
                <button
                    className={`admin-tab ${activeTab === 'users' ? 'admin-tab--active' : ''}`}
                    onClick={() => setActiveTab('users')}
                >
                    Usuarios
                </button>
            </div>

            {/* ── Sección Cursos ── */}
            {activeTab === 'courses' && (
                <div className="admin-section">
                    <div className="admin-section-header">
                        <h2>Gestión de Cursos</h2>
                        <button
                            className="admin-btn admin-btn--primary"
                            onClick={() => { setShowCreateForm(prev => !prev); setFormError(''); }}
                        >
                            {showCreateForm ? 'Cancelar' : '+ Crear Curso'}
                        </button>
                    </div>

                    {showCreateForm && (
                        <form className="admin-create-form" onSubmit={handleCreateCourse}>
                            <h3>Nuevo Curso</h3>
                            {formError && <p className="admin-form-error">{formError}</p>}

                            <div className="admin-form-row">
                                <div className="admin-form-group">
                                    <label>Nombre *</label>
                                    <input value={form.name} onChange={e => handleFormChange('name', e.target.value)} placeholder="Nombre del curso" />
                                </div>
                                <div className="admin-form-group">
                                    <label>Instructor *</label>
                                    <input value={form.instructor} onChange={e => handleFormChange('instructor', e.target.value)} placeholder="Nombre del instructor" />
                                </div>
                            </div>

                            <div className="admin-form-group">
                                <label>Descripción *</label>
                                <textarea rows={3} value={form.description} onChange={e => handleFormChange('description', e.target.value)} placeholder="Descripción del curso" />
                            </div>

                            <div className="admin-form-row">
                                <div className="admin-form-group">
                                    <label>Precio (USD) *</label>
                                    <input type="number" min="0" step="0.01" value={form.price} onChange={e => handleFormChange('price', e.target.value)} placeholder="0.00" />
                                </div>
                                <div className="admin-form-group">
                                    <label>Duración (HH:MM:SS) *</label>
                                    <input value={form.length} onChange={e => handleFormChange('length', e.target.value)} placeholder="02:30:00" pattern="\d{2}:\d{2}:\d{2}" />
                                </div>
                            </div>

                            <div className="admin-form-group">
                                <label>Requisitos</label>
                                <textarea rows={2} value={form.requirements} onChange={e => handleFormChange('requirements', e.target.value)} placeholder="Opcional" />
                            </div>

                            <div className="admin-form-group">
                                <label>Categorías *</label>
                                {form.categories.map((cat, i) => (
                                    <div key={i} className="admin-category-row">
                                        <input value={cat.name} onChange={e => handleCategoryChange(i, e.target.value)} placeholder={`Categoría ${i + 1}`} />
                                        {form.categories.length > 1 && (
                                            <button type="button" className="admin-btn-icon" onClick={() => handleRemoveCategory(i)}>✕</button>
                                        )}
                                    </div>
                                ))}
                                <button type="button" className="admin-btn admin-btn--ghost" onClick={handleAddCategory}>+ Categoría</button>
                            </div>

                            <div className="admin-form-group">
                                <label>Imagen</label>
                                <label htmlFor="admin-img" className="admin-image-label">
                                    {imagePreview
                                        ? <img src={imagePreview} alt="preview" className="admin-image-preview" />
                                        : <span>📷 Seleccionar imagen</span>}
                                </label>
                                <input id="admin-img" type="file" accept="image/*" onChange={handleImageChange} className="admin-image-hidden" />
                            </div>

                            <button type="submit" className="admin-btn admin-btn--primary">Crear Curso</button>
                        </form>
                    )}

                    <div className="admin-table-wrapper">
                    <table className="admin-table">
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Nombre</th>
                                <th>Instructor</th>
                                <th>Precio</th>
                                <th>Acciones</th>
                            </tr>
                        </thead>
                        <tbody>
                            {courses.length === 0 ? (
                                <tr><td colSpan={5} className="admin-table-empty">Sin cursos</td></tr>
                            ) : courses.map(course => (
                                <tr key={course.id}>
                                    <td>{course.id}</td>
                                    <td>{course.name}</td>
                                    <td>{course.instructor || '—'}</td>
                                    <td>${course.price?.toFixed(2) ?? '—'}</td>
                                    <td>
                                        <button
                                            className="admin-btn admin-btn--danger"
                                            onClick={() => handleDeleteCourse(course.id, course.name)}
                                        >
                                            Eliminar
                                        </button>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                    </div>
                </div>
            )}

            {/* ── Sección Usuarios ── */}
            {activeTab === 'users' && (
                <div className="admin-section">
                    <h2>Gestión de Usuarios</h2>
                    <div className="admin-table-wrapper">
                    <table className="admin-table">
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Nombre</th>
                                <th>Email</th>
                                <th>Rol</th>
                                <th>Acciones</th>
                            </tr>
                        </thead>
                        <tbody>
                            {users.length === 0 ? (
                                <tr><td colSpan={5} className="admin-table-empty">Sin usuarios</td></tr>
                            ) : users.map(user => (
                                <tr key={user.id}>
                                    <td>{user.id}</td>
                                    <td>{user.name} {user.surname}</td>
                                    <td>{user.email}</td>
                                    <td>
                                        <select
                                            className="admin-role-select"
                                            value={userRoles[user.id] || 'student'}
                                            onChange={e => handleRoleChange(user.id, e.target.value)}
                                        >
                                            <option value="student">student</option>
                                            <option value="admin">admin</option>
                                        </select>
                                    </td>
                                    <td>
                                        <button
                                            className="admin-btn admin-btn--primary"
                                            onClick={() => handleSaveRole(user.id)}
                                        >
                                            Guardar
                                        </button>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                    </div>
                </div>
            )}
        </div>
    );
};

export default AdminPanel;
