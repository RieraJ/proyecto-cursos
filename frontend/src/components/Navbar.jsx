import React, { useState, useEffect, useRef, useCallback } from "react";
import { NavLink, useNavigate, useLocation } from "react-router-dom";
import "./Navbar.css";
import { FaGraduationCap, FaUser, FaChevronDown, FaSignOutAlt, FaBookOpen, FaShieldAlt, FaUserEdit } from 'react-icons/fa';
import { API_URL } from '../config';
import Swal from 'sweetalert2';
import { useTheme } from '../ThemeContext';
import { Sun, Moon } from 'lucide-react';

const Navbar = () => {
    const navigate = useNavigate();
    const location = useLocation();
    const { theme, toggleTheme } = useTheme();
    const [dropdownOpen, setDropdownOpen] = useState(false);
    const [userInfo, setUserInfo] = useState(null);
    const dropdownRef = useRef(null);

    const fetchUserInfo = useCallback(async () => {
        try {
            const response = await fetch(`${API_URL}/user-info`, {
                credentials: 'include',
            });
            if (!response.ok) {
                setUserInfo(null);
                return;
            }
            const data = await response.json();
            setUserInfo({
                name: data.userInfo.name,
                surname: data.userInfo.surname,
                userType: data.userInfo.userType,
                image: data.userInfo.image || '',
            });
        } catch {
            setUserInfo(null);
        }
    }, []);

    useEffect(() => {
        fetchUserInfo();
    }, [fetchUserInfo, location]);

    useEffect(() => {
        const handleMouseDown = (e) => {
            if (dropdownRef.current && !dropdownRef.current.contains(e.target)) {
                setDropdownOpen(false);
            }
        };
        document.addEventListener('mousedown', handleMouseDown);
        return () => document.removeEventListener('mousedown', handleMouseDown);
    }, []);

    const handleLogout = async () => {
        const result = await Swal.fire({
            title: '¿Cerrar sesión?',
            icon: 'question',
            showCancelButton: true,
            confirmButtonText: 'Sí, salir',
            cancelButtonText: 'Cancelar',
        });
        if (!result.isConfirmed) return;
        try {
            await fetch(`${API_URL}/logout`, {
                method: 'POST',
                credentials: 'include',
            });
        } catch {
            // ignorar errores de red; redirigir igual
        }
        setUserInfo(null);
        setDropdownOpen(false);
        navigate('/login');
    };

    return (
        <nav className="navbar">
            {/* Brand */}
            <NavLink to="/" className="navbar-brand">
                <FaGraduationCap className="nav-logo-icon" />
                <span className="nav-brand-text">EduCursos</span>
            </NavLink>

            {/* Links */}
            <ul className="navbar-links">
                <li>
                    <NavLink to="/login" className={({ isActive }) => `nav-link${isActive ? ' active' : ''}`}>
                        Iniciar Sesión
                    </NavLink>
                </li>
                <li>
                    <NavLink to="/courses" className={({ isActive }) => `nav-link${isActive ? ' active' : ''}`}>
                        Cursos
                    </NavLink>
                </li>
            </ul>

            {/* Theme toggle */}
            <button
                className="theme-toggle"
                onClick={toggleTheme}
                aria-label={theme === 'dark' ? 'Cambiar a modo claro' : 'Cambiar a modo oscuro'}
                title={theme === 'dark' ? 'Modo claro' : 'Modo oscuro'}
            >
                {theme === 'dark' ? <Sun /> : <Moon />}
            </button>

            {/* User menu */}
            <div className="profile" ref={dropdownRef}>
                <button
                    id="user-menu-button"
                    className="user-menu-button"
                    onClick={() => setDropdownOpen((prev) => !prev)}
                    aria-haspopup="true"
                    aria-expanded={dropdownOpen}
                >
                    <span className="user-avatar">
                        {userInfo?.image
                            ? <img src={`data:image/png;base64,${userInfo.image}`} alt="avatar" className="user-avatar-photo" />
                            : <FaUser />
                        }
                    </span>
                    <FaChevronDown className={`chevron-icon ${dropdownOpen ? 'chevron-open' : ''}`} />
                </button>

                {dropdownOpen && (
                    <div className="dropdown-menu" role="menu">
                        {userInfo ? (
                            <>
                                <div className="dropdown-user-name">
                                    <span className="dropdown-user-initial">
                                        {userInfo.image
                                            ? <img src={`data:image/png;base64,${userInfo.image}`} alt="avatar" className="dropdown-user-photo" />
                                            : userInfo.name?.charAt(0).toUpperCase()
                                        }
                                    </span>
                                    <span>{userInfo.name} {userInfo.surname}</span>
                                </div>
                                <div className="dropdown-divider" />
                                <button
                                    id="dropdown-mis-cursos"
                                    className="dropdown-item"
                                    role="menuitem"
                                    onClick={() => { setDropdownOpen(false); navigate('/profile'); }}
                                >
                                    <FaBookOpen className="dropdown-icon" />
                                    Mis Cursos
                                </button>
                                <button
                                    id="dropdown-editar-perfil"
                                    className="dropdown-item"
                                    role="menuitem"
                                    onClick={() => { setDropdownOpen(false); navigate('/profile/edit'); }}
                                >
                                    <FaUserEdit className="dropdown-icon" />
                                    Editar Perfil
                                </button>
                                {userInfo.userType === 'admin' && (
                                    <button
                                        id="dropdown-admin"
                                        className="dropdown-item"
                                        role="menuitem"
                                        onClick={() => { setDropdownOpen(false); navigate('/admin'); }}
                                    >
                                        <FaShieldAlt className="dropdown-icon" />
                                        Panel Admin
                                    </button>
                                )}
                                <div className="dropdown-divider" />
                                <button
                                    id="dropdown-logout"
                                    className="dropdown-item dropdown-item--danger"
                                    role="menuitem"
                                    onClick={handleLogout}
                                >
                                    <FaSignOutAlt className="dropdown-icon" />
                                    Cerrar Sesión
                                </button>
                            </>
                        ) : (
                            <div className="dropdown-user-name dropdown-not-logged">
                                No has iniciado sesión
                            </div>
                        )}
                    </div>
                )}
            </div>
        </nav>
    );
};

export default Navbar;
