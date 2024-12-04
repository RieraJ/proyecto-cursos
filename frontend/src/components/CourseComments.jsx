import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

const CourseComments = () => {
  const { courseId } = useParams(); // Obtiene el ID del curso de la URL
  const [comments, setComments] = useState([]);

  useEffect(() => {
    // Llamada a la API para obtener los comentarios del curso
    fetch(`/api/courses/${courseId}/comments`) // Reemplazar con tu endpoint real
      .then((response) => response.json())
      .then((data) => setComments(data))
      .catch((error) => console.error("Error fetching comments:", error));
  }, [courseId]);

  return (
    <div className="comments-container">
      <h2>Comments for Course {courseId}</h2>
      {comments.length === 0 ? (
        <p>No comments available for this course.</p>
      ) : (
        <ul className="comments-list">
          {comments.map((comment, index) => (
            <li key={index} className="comment-item">
              <p><strong>{comment.author}</strong>: {comment.text}</p>
              <p><em>{new Date(comment.date).toLocaleString()}</em></p>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default CourseComments;
