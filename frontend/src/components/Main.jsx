import React from 'react';
import './Main.css';
import { FaGraduationCap, FaCode, FaShieldAlt, FaLaptopCode, FaCalculator, FaArrowRight } from 'react-icons/fa';

const Main = () => {
    return (
        <div className="main-container">
            {/* Decorative blobs */}
            <div className="hero-blob hero-blob--a" aria-hidden="true" />
            <div className="hero-blob hero-blob--b" aria-hidden="true" />

            <section className="hero-section">
                <div className="hero-badge">
                    <FaGraduationCap className="badge-icon" />
                    Plataforma de Aprendizaje
                </div>

                <h1>
                    Aprende nuevas{' '}
                    <span className="highlight">habilidades</span>
                </h1>

                <p className="hero-subtitle">
                    Explora cursos de vanguardia dictados por profesionales en programación,
                    hacking, computación, matemáticas y más.
                </p>

                <a href="/courses" className="cta-button">
                    Explorar Cursos <FaArrowRight className="btn-icon" />
                </a>
            </section>

            <div className="features-grid">
                <div className="feature-card">
                    <div className="feature-icon-wrapper">
                        <FaCode />
                    </div>
                    <h3>Programación</h3>
                    <p>Desarrollo web, móvil y arquitecturas de software modernas.</p>
                </div>

                <div className="feature-card">
                    <div className="feature-icon-wrapper">
                        <FaShieldAlt />
                    </div>
                    <h3>Hacking & Seguridad</h3>
                    <p>Ciberseguridad práctica, hacking ético y análisis forense.</p>
                </div>

                <div className="feature-card">
                    <div className="feature-icon-wrapper">
                        <FaLaptopCode />
                    </div>
                    <h3>Computación</h3>
                    <p>Ciencias de la computación, hardware, redes y sistemas.</p>
                </div>

                <div className="feature-card">
                    <div className="feature-icon-wrapper">
                        <FaCalculator />
                    </div>
                    <h3>Matemáticas</h3>
                    <p>Cálculo, álgebra, criptografía y ciencia de datos.</p>
                </div>
            </div>
        </div>
    );
};

export default Main;
