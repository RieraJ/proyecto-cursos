import React, { useState } from "react";
import { NavLink } from "react-router-dom";
import "./Navbar.css";
import { FaHome } from 'react-icons/fa';
const Navbar = () => {

    return (
        <div className="navbar">
            <div className="home">
                <NavLink to="/" exact="true">
                    <FaHome alt="" className="logo" />
                </NavLink>
            </div>
            <ul>
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
            <div className="profile">
                    <FaHome alt="" className="logo" />
            </div>
            <div className="options">
                    <FaHome alt="" className="logo" />
            </div>
        </div>
    );
};

export default Navbar;