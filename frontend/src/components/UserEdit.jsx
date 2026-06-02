import React, { useState, useEffect, useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import { FaCamera } from 'react-icons/fa';
import Swal from 'sweetalert2';
import './UserEdit.css';
import { API_URL } from '../config';
import { validateImageFile } from '../utils';

const UserEdit = () => {
    const navigate = useNavigate();
    const fileInputRef = useRef(null);
    const photoPreviewUrlRef = useRef(null);
    const [form, setForm] = useState({ name: '', surname: '', email: '', password: '', confirmPassword: '' });
    const [currentImage, setCurrentImage] = useState('');
    const [photoFile, setPhotoFile] = useState(null);
    const [photoPreview, setPhotoPreview] = useState('');
    const [error, setError] = useState('');
    const [loading, setLoading] = useState(true);
    const [saving, setSaving] = useState(false);

    useEffect(() => {
        const fetchUserInfo = async () => {
            try {
                const res = await fetch(`${API_URL}/user-info`, { credentials: 'include' });
                if (!res.ok) { navigate('/login'); return; }
                const data = await res.json();
                const { name, surname, email, image } = data.userInfo;
                setForm(prev => ({ ...prev, name: name || '', surname: surname || '', email: email || '' }));
                setCurrentImage(image || '');
            } catch {
                navigate('/login');
            } finally {
                setLoading(false);
            }
        };
        fetchUserInfo();
    }, [navigate]);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setForm(prev => ({ ...prev, [name]: value }));
    };

    const handlePhotoChange = (e) => {
        const file = e.target.files[0];
        if (!file) return;
        const validationError = validateImageFile(file);
        if (validationError) {
            Swal.fire({ icon: 'error', title: 'Imagen inválida', text: validationError });
            e.target.value = '';
            return;
        }
        if (photoPreviewUrlRef.current) URL.revokeObjectURL(photoPreviewUrlRef.current);
        const url = URL.createObjectURL(file);
        photoPreviewUrlRef.current = url;
        setPhotoFile(file);
        setPhotoPreview(url);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');

        if (form.password && form.password.length < 8) {
            setError('La contraseña debe tener al menos 8 caracteres.');
            return;
        }
        if (form.password && form.password !== form.confirmPassword) {
            setError('Las contraseñas no coinciden.');
            return;
        }
        setSaving(true);

        const body = {};
        if (form.name) body.name = form.name;
        if (form.surname) body.surname = form.surname;
        if (form.email) body.email = form.email;
        if (form.password) body.password = form.password;

        try {
            const res = await fetch(`${API_URL}/users/me`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                credentials: 'include',
                body: JSON.stringify(body),
            });

            if (!res.ok) {
                const data = await res.json();
                setError(data.error || 'Error al actualizar el perfil.');
                return;
            }

            if (photoFile) {
                const formData = new FormData();
                formData.append('photo', photoFile);
                const photoRes = await fetch(`${API_URL}/users/me/photo`, {
                    method: 'PUT',
                    credentials: 'include',
                    body: formData,
                });
                if (!photoRes.ok) {
                    setError('Datos actualizados, pero hubo un error al subir la foto.');
                    return;
                }
            }

            await Swal.fire({
                icon: 'success',
                title: '¡Perfil actualizado!',
                text: 'Tus datos fueron guardados correctamente.',
                timer: 1500,
                showConfirmButton: false,
            });
            navigate('/profile');
        } catch {
            setError('Error de red. Intentá de nuevo.');
        } finally {
            setSaving(false);
        }
    };

    const displayPhoto = photoPreview || (currentImage ? `data:image/png;base64,${currentImage}` : '');

    if (loading) return null;

    return (
        <div className="user-edit-container">
            <div className="user-edit-card">
                <div className="user-edit-header">
                    <h1>Editar Perfil</h1>
                    <p>Actualizá tus datos personales</p>
                </div>

                <form className="user-edit-form" onSubmit={handleSubmit}>
                    {error && <div className="user-edit-error">{error}</div>}

                    {/* Foto de perfil */}
                    <div className="user-edit-photo-section">
                        <div className="user-edit-avatar" onClick={() => fileInputRef.current?.click()}>
                            {displayPhoto
                                ? <img src={displayPhoto} alt="Foto de perfil" className="user-edit-avatar-img" />
                                : <span className="user-edit-avatar-initials">
                                    {form.name?.charAt(0).toUpperCase()}{form.surname?.charAt(0).toUpperCase()}
                                  </span>
                            }
                            <div className="user-edit-avatar-overlay">
                                <FaCamera />
                            </div>
                        </div>
                        <input
                            ref={fileInputRef}
                            type="file"
                            accept="image/*"
                            onChange={handlePhotoChange}
                            className="user-edit-photo-input"
                        />
                        <p className="user-edit-photo-hint">Hacé clic en la foto para cambiarla</p>
                    </div>

                    <div className="user-edit-row">
                        <div className="user-edit-group">
                            <label htmlFor="name">Nombre</label>
                            <input
                                id="name"
                                name="name"
                                type="text"
                                value={form.name}
                                onChange={handleChange}
                                placeholder="Tu nombre"
                            />
                        </div>
                        <div className="user-edit-group">
                            <label htmlFor="surname">Apellido</label>
                            <input
                                id="surname"
                                name="surname"
                                type="text"
                                value={form.surname}
                                onChange={handleChange}
                                placeholder="Tu apellido"
                            />
                        </div>
                    </div>

                    <div className="user-edit-group">
                        <label htmlFor="email">Email</label>
                        <input
                            id="email"
                            name="email"
                            type="email"
                            value={form.email}
                            onChange={handleChange}
                            placeholder="tu@email.com"
                        />
                    </div>

                    <div className="user-edit-divider">Cambiar contraseña</div>

                    <div className="user-edit-group">
                        <label htmlFor="password">Nueva contraseña</label>
                        <input
                            id="password"
                            name="password"
                            type="password"
                            value={form.password}
                            onChange={handleChange}
                            placeholder="Dejar vacío para no cambiarla"
                        />
                    </div>

                    <div className="user-edit-group">
                        <label htmlFor="confirmPassword">Confirmar contraseña</label>
                        <input
                            id="confirmPassword"
                            name="confirmPassword"
                            type="password"
                            value={form.confirmPassword}
                            onChange={handleChange}
                            placeholder="Repetí la nueva contraseña"
                        />
                    </div>

                    <div className="user-edit-actions">
                        <button type="button" className="user-edit-cancel" onClick={() => navigate('/profile')}>
                            Cancelar
                        </button>
                        <button type="submit" className="user-edit-submit" disabled={saving}>
                            {saving ? 'Guardando...' : 'Guardar Cambios'}
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
};

export default UserEdit;
