import React, { useState, useEffect, useCallback, useRef } from 'react';
import { useParams } from 'react-router-dom';
import { FaUser, FaImage, FaCommentDots } from 'react-icons/fa';
import './CourseComments.css';
import Swal from 'sweetalert2';
import { API_URL } from '../config';
import { validateImageFile } from '../utils';

const CourseComments = () => {
  const { courseId } = useParams();
  const [userId, setUserId] = useState(null);
  const [comments, setComments] = useState([]);
  const [newComment, setNewComment] = useState('');
  const [imageFile, setImageFile] = useState(null);
  const [imageName, setImageName] = useState('');
  const [error, setError] = useState(null);
  const [submitting, setSubmitting] = useState(false);
  const fileInputRef = useRef(null);
  const imagePreviewUrlRef = useRef(null);

  const fetchUserInfo = async () => {
    try {
      const response = await fetch(`${API_URL}/user-info`, { credentials: 'include' });
      if (!response.ok) throw new Error('Error fetching user info');
      const data = await response.json();
      setUserId(data.userInfo.id);
    } catch (err) {
      console.error('Error fetching user info:', err);
    }
  };

  const fetchComments = useCallback(async (_courseId, _pageNum = 1) => {
    try {
      const response = await fetch(`${API_URL}/courses/${courseId}/comments`, {
        credentials: 'include',
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
      const safeData = Array.isArray(data) ? data : [];

      // Se crean URLs de las imágenes al recibir un comentario, si tiene una
      const commentsWithUrls = safeData.map((comment) => {
        if (comment.image) {
          comment.imageUrl = `data:image/png;base64,${comment.image}`;
        }
        if (comment.user_image) {
          comment.userAvatarUrl = `data:image/png;base64,${comment.user_image}`;
        }
        return comment;
      });

      setComments(commentsWithUrls);
      setError(null);
    } catch (err) {
      console.error(err);
      setError(err.message);
      setComments([]);
    }
  }, [courseId]);

  useEffect(() => {
    fetchUserInfo();
    fetchComments(courseId, 1);
  }, [courseId, fetchComments]);

  const handleImageChange = (e) => {
    const file = e.target.files[0];
    if (!file) return;
    const validationError = validateImageFile(file);
    if (validationError) {
      Swal.fire({ icon: 'error', title: 'Imagen inválida', text: validationError });
      e.target.value = '';
      return;
    }
    if (imagePreviewUrlRef.current) URL.revokeObjectURL(imagePreviewUrlRef.current);
    imagePreviewUrlRef.current = URL.createObjectURL(file);
    setImageFile(file);
    setImageName(file.name);
  };

  const handleCreateComment = async (e) => {
    e.preventDefault();

    if (!newComment.trim()) {
      Swal.fire({ icon: 'warning', title: 'Comentario vacío', text: 'El comentario no puede estar vacío.' });
      return;
    }

    setSubmitting(true);
    const formData = new FormData();
    formData.append('user_id', userId);
    formData.append('course_id', courseId);
    formData.append('content', newComment);
    if (imageFile) formData.append('image', imageFile);

    try {
      const response = await fetch(`${API_URL}/comments`, {
        method: 'POST',
        credentials: 'include',
        body: formData,
      });

      if (response.ok) {
        setNewComment('');
        setImageFile(null);
        setImageName('');
        if (fileInputRef.current) fileInputRef.current.value = '';
        if (imagePreviewUrlRef.current) { URL.revokeObjectURL(imagePreviewUrlRef.current); imagePreviewUrlRef.current = null; }
        fetchComments(courseId, 1);
        Swal.fire({ icon: 'success', title: 'Comentario enviado', text: 'Comentario enviado exitosamente!' });
      } else {
        const errorData = await response.json();
        Swal.fire({ icon: 'error', title: 'Error al enviar', text: errorData.error === 'user is not enrolled in this course' ? 'Debés inscribirte al curso para comentar.' : 'No se pudo enviar el comentario.' });
      }
    } catch (err) {
      Swal.fire({ icon: 'error', title: 'Error', text: 'Error al crear el comentario.' });
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className="comments-page">
      <div className="comments-header">
        <FaCommentDots className="comments-header-icon" />
        <div>
          <h2>Comentarios</h2>
          <p>Curso #{courseId}</p>
        </div>
      </div>

      <div className="comments-body">
        <div className="comment-form-card">
          <h3>Dejá tu comentario</h3>
          <form onSubmit={handleCreateComment} className="comment-form">
            <div className="textarea-wrapper">
              <textarea
                value={newComment}
                onChange={(e) => setNewComment(e.target.value)}
                placeholder="Escribí tu comentario aquí..."
                maxLength={500}
                required
              />
              <span className={`char-count${newComment.length >= 450 ? ' char-count--warn' : ''}`}>{newComment.length}/500</span>
            </div>

            <div className="form-actions">
              <label htmlFor="commentImage" className="image-upload-btn">
                <FaImage />
                {imageName ? imageName : 'Adjuntar imagen'}
              </label>
              <input
                ref={fileInputRef}
                type="file"
                id="commentImage"
                accept="image/*"
                onChange={handleImageChange}
                className="hidden-file-input"
              />
              <button type="submit" className="post-comment-btn" disabled={submitting}>
                {submitting ? 'Enviando...' : 'Publicar'}
              </button>
            </div>
          </form>
        </div>

        {error && error !== 'No comments found' && (
          <p className="error-message">{error}</p>
        )}

        <div className="comments-list">
          {comments.length > 0 ? (
            comments.map((comment) => (
              <div key={comment.id} className="comment-card">
                <div className="comment-avatar">
                  {comment.userAvatarUrl
                    ? <img src={comment.userAvatarUrl} alt="avatar" className="comment-avatar-photo" />
                    : comment.user_name
                      ? comment.user_name.charAt(0).toUpperCase()
                      : <FaUser />
                  }
                </div>
                <div className="comment-content">
                  <p className="comment-text">{comment.content}</p>
                  {comment.imageUrl && (
                    <img
                      src={comment.imageUrl}
                      alt="Imagen del comentario"
                      className="comment-image"
                    />
                  )}
                </div>
              </div>
            ))
          ) : (
            <div className="empty-comments">
              <FaCommentDots className="empty-icon" />
              <p>No hay comentarios todavía. ¡Sé el primero en comentar!</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default CourseComments;
