import React, { useState } from "react";
import { NavLink } from "react-router-dom";
import "./Navbar.css";
import { FaHome } from 'react-icons/fa';
const Navbar = () => {

    return (
        <div className="navbar">
            <NavLink to="/" exact="true">
                <FaHome alt="" className="logo" />
            </NavLink>
            <ul>
                <li>
                    <NavLink to="/signup" className="custom-link">
                        Registrarse
                    </NavLink>
                </li>
                <li>
                    <NavLink to="/login" className="custom-link">
                        Iniciar Sesion
                    </NavLink>
                </li>
                <li>
                    <NavLink to="/courses" className="custom-link">
                        Cursos
                    </NavLink>
                </li>
            </ul>
        </div>
    );
};

export default Navbar;