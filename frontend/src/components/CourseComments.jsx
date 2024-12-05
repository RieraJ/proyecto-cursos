import React, { useState, useEffect, useCallback } from 'react';
import Cookies from 'js-cookie';
import { useParams, Link } from 'react-router-dom';
import './CourseComments.css'; // Necesitarás crear este archivo de estilos

const CourseComments = () => {
  const { courseId } = useParams();
  const [comments, setComments] = useState([]);
  const [newComment, setNewComment] = useState('');
  const [imageFile, setImageFile] = useState(null);
  const [error, setError] = useState(null);
  const [page, setPage] = useState(1);
  const [, setTotalComments] = useState(0);

  const userId = Cookies.get('userId');

  const fetchComments = useCallback(async (courseId, pageNum = 1) => {
    try {
      const response = await fetch(`http://localhost:4000/courses/${courseId}/comments`, {
        credentials: 'include'
      });

      if (!response.ok) {
        const errorData = await response.json();
        if (errorData.error === 'No comments found') {
          setComments([]);
          setError(null);
          return;
        }
        throw new Error(errorData.error || 'Error fetching comments');
      }

      const data = await response.json();

      // Se crear URLs de las imagenes al recibir un comentario, si tiene una
      const commentsWithUrls = data.map((comment) => {
        if (comment.image) {
          comment.imageUrl = `data:image/png;base64,${comment.image}`;
        }
        return comment;
      });

      setComments(commentsWithUrls);
      setTotalComments(data ? data.length : 0);
      setError(null);
    } catch (err) {
      console.error(err);
      setError(err.message);
      setComments([]);
    }
  }, []);

  useEffect(() => {
    fetchComments(courseId, page);
  }, [courseId, page, fetchComments]);

  const handleCreateComment = async (e) => {
    e.preventDefault();

    if (!newComment.trim()) {
      alert('Comment cannot be empty');
      return;
    }

    const formData = new FormData();
    formData.append('user_id', userId);
    formData.append('course_id', courseId);
    formData.append('content', newComment);

    if (imageFile) {
      formData.append('image', imageFile);
    }

    try {
      const response = await fetch('http://localhost:4000/comments', {
        method: 'POST',
        credentials: 'include',
        body: formData,
      });

      if (response.ok) {
        setNewComment('');
        setImageFile(null);
        fetchComments(courseId, page);
      } else {
        const errorData = await response.json();
        alert(errorData.error || 'Failed to create comment');
      }
    } catch (err) {
      console.error('Error creating comment:', err);
      alert('Error creating comment');
    }
  };

  /* 
  const handleDeleteComment = async (commentId) => {
    try {
      const response = await fetch(`http://localhost:4000/comments/${commentId}`, {
        method: 'DELETE',
        credentials: 'include'
      });

      if (response.ok) {
        fetchComments(courseId, page);
      } else {
        const errorData = await response.json();
        alert(errorData.error || 'Failed to delete comment');
      }
    } catch (err) {
      console.error('Error deleting comment:', err);
      alert('Error deleting comment');
    }
  };
  */

  return (
    <div className="course-comments-container">
      <h2>Comments for Course</h2>

      {/* Formulario para crear un comentario */}
      <form onSubmit={handleCreateComment} className="comment-form">
        <textarea
          value={newComment}
          onChange={(e) => setNewComment(e.target.value)}
          placeholder="Write your comment..."
          maxLength={500}
          required
        ></textarea>
        <input
          type="file"
          accept="image/*"
          onChange={(e) => setImageFile(e.target.files[0])}
        />
        <button type="submit">Post Comment</button>
      </form>

      {/* Lista de comentarios */}
      {error && error !== 'No comments found' && <p className="error-message">{error}</p>}

      <div className="comments-list">
        {comments.length > 0 ? (
          comments.map((comment) => (
            <div key={comment.id} className="comment-card">
              <p>{comment.content}</p>
              {comment.imageUrl && (
                <img
                  src={comment.imageUrl}
                  alt="Comment"
                  className="comment-image"
                />
              )}
            </div>
          ))
        ) : (
          <p className="no-comments-message">No hay comentarios todavía</p>
        )}
      </div>
    </div>
  );
};

export default CourseComments;
