import React from 'react';
import './Main.css';

const Main = () => {
    return (
        <div className="main-container">
            <header className="main-header">
                <h1>Aprende nuevas habilidades</h1>
                <p>Desde Programacion, Hacking, Computacion, Matematicas, etc...</p>
                <a href="/courses" className="cta-button">Buscar cursos</a>
            </header>
        </div>
    );
};

export default Main;
