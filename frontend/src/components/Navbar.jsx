import React, { useState } from "react";
import { NavLink } from "react-router-dom";
import "./Navbar.css";
import { FaHome, FaUser } from 'react-icons/fa';
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
                <NavLink to="/profile">
                    <FaUser alt="" className="logo" />
                </NavLink>
            </div>
        </div>
    );
};

export default Navbar;