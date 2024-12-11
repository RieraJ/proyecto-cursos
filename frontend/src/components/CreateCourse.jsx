import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './CreateCourse.css';

// const API_BASE_URL = process.env.REACT_APP_API_URL || "http://localhost:4000";

const CreateCourse = () => {
    const navigate = useNavigate();
    const [courseName, setCourseName] = useState('');
    const [description, setDescription] = useState('');
    const [price, setPrice] = useState('');
    const [instructor, setInstructor] = useState('');
    const [length, setLength] = useState('');
    const [requirements, setRequirements] = useState('');
    const [categories, setCategories] = useState([{ name: '' }]);
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');

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

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');
        setSuccess('');

        // Basic validation
        if (!courseName || !description || !price || !instructor || !length) {
            setError('Please fill in all required fields');
            return;
        }

        // Validate categories
        const validCategories = categories.filter(cat => cat.name.trim() !== '');
        if (validCategories.length === 0) {
            setError('Please add at least one category');
            return;
        }

        const courseData = {
            name: courseName,
            description,
            price: parseFloat(price),
            active: true,
            instructor,
            length,
            requirements,
            categories: validCategories
        };

        try {
            const response = await fetch(`http://localhost:4000/courses`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
                body: JSON.stringify(courseData)
            });

            const data = await response.json();

            if (response.ok) {
                setSuccess('Course created successfully!');
                // Reset form
                setCourseName('');
                setDescription('');
                setPrice('');
                setInstructor('');
                setLength('');
                setRequirements('');
                setCategories([{ name: '' }]);
                
                navigate('/courses');
            } else {
                setError(data.error || 'Failed to create course');
            }
        } catch (err) {
            console.error('Error creating course:', err);
            setError('Network error. Please try again.');
        }
    };

    return (
        <div className="create-course-container">
            {error && <p className="error-message">{error}</p>}
            {success && <p className="success-message">{success}</p>}
            <form onSubmit={handleSubmit} className="create-course-form">
                <div className="form-group">
                    <label htmlFor="courseName">Course Name</label>
                    <input
                        type="text"
                        id="courseName"
                        value={courseName}
                        onChange={(e) => setCourseName(e.target.value)}
                        placeholder="Enter course name"
                        required
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="description">Description</label>
                    <textarea
                        id="description"
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                        placeholder="Enter course description"
                        required
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="price">Price ($)</label>
                    <input
                        type="number"
                        id="price"
                        value={price}
                        onChange={(e) => setPrice(e.target.value)}
                        placeholder="Enter course price"
                        min="0"
                        step="0.01"
                        required
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="instructor">Instructor</label>
                    <input
                        type="text"
                        id="instructor"
                        value={instructor}
                        onChange={(e) => setInstructor(e.target.value)}
                        placeholder="Enter instructor name"
                        required
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="length">Course Length (HH:MM:SS)</label>
                    <input
                        type="text"
                        id="length"
                        value={length}
                        onChange={(e) => setLength(e.target.value)}
                        placeholder="Enter course length (e.g., 02:30:00)"
                        pattern="\d{2}:\d{2}:\d{2}"
                        required
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="requirements">Requirements</label>
                    <textarea
                        id="requirements"
                        value={requirements}
                        onChange={(e) => setRequirements(e.target.value)}
                        placeholder="Enter course requirements"
                    />
                </div>

                <div className="form-group">
                    <label>Categories</label>
                    {categories.map((category, index) => (
                        <div key={index} className="category-input">
                            <input
                                type="text"
                                value={category.name}
                                onChange={(e) => handleCategoryChange(index, e.target.value)}
                                placeholder="Enter category"
                            />
                            {categories.length > 1 && (
                                <button 
                                    type="button" 
                                    onClick={() => handleRemoveCategory(index)}
                                    className="remove-category-btn"
                                >
                                    Remove
                                </button>
                            )}
                        </div>
                    ))}
                    <button 
                        type="button" 
                        onClick={handleAddCategory}
                        className="add-category-btn"
                    >
                        Add Category
                    </button>
                </div>

                <button type="submit" className="submit-course-btn">
                    Create Course
                </button>
            </form>
        </div>
    );
};

export default CreateCourse;