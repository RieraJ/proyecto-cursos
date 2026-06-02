export const formatLength = (length) => {
    if (!length || length === 'N/A') return 'N/A';
    const parts = length.split(':');
    const hours = parseInt(parts[0], 10) || 0;
    const minutes = parseInt(parts[1], 10) || 0;
    const seconds = parseInt(parts[2], 10) || 0;
    return `${hours}h ${minutes}m ${seconds}s`;
};

const LENGTH_PATTERN = /^\d{2}:\d{2}:\d{2}$/;
export const isValidLength = (value) => LENGTH_PATTERN.test(value);

const IMAGE_TYPES = ['image/jpeg', 'image/png', 'image/gif', 'image/webp'];
const MAX_IMAGE_SIZE = 5 * 1024 * 1024; // 5 MB

export const validateImageFile = (file) => {
    if (!IMAGE_TYPES.includes(file.type)) return 'Solo se permiten imágenes JPG, PNG, GIF o WebP.';
    if (file.size > MAX_IMAGE_SIZE) return 'La imagen no puede superar 5 MB.';
    return null;
};
