import React, { useEffect, useState } from 'react';
import axios from 'axios'
import { Link } from 'react-router-dom';

const Profile: React.FC = () => {
    const [pfp, setPfp] = useState<string | null>(null);
    const [username, setUsername] = useState<string | null>(null);
    const [description, setDescription] = useState<string | null>(null);
    const [newDescription, setNewDescription] = useState<string>('');
    const [newPfp, setNewPfp] = useState<string>('');
    const [isEditingDescription, setIsEditingDescription] = useState<boolean>(false);
    const [isEditingPfp, setIsEditingPfp] = useState<boolean>(false);
    const [isHovered, setIsHovered] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);
    const [ratings, setRatings] = useState<number>(0);
    const [discussions, setDiscussions] = useState<number>(0);
    const [reviews, setReviews] = useState<number>(0);
    const [collections, setCollections] = useState<number>(0);

    useEffect(() => {
        const user = localStorage.getItem('user');
        if (user) {
            const userData = JSON.parse(user);
            setPfp(userData.pfp);
            setUsername(userData.username);
            setDescription(userData.description);
            const fetchStatistics = async () => {
                try {
                    const response = await axios.get(
                        `http://31.15.18.177:5050/api/get_statistics/${userData.username}`
                    );
                    const data = response.data;

                    setRatings(data.ratings_count || 0);
                    setDiscussions(data.discussions_count || 0);
                    setReviews(data.reviews_count || 0);
                    setCollections(data.collections_count || 0);
                } catch (err) {
                    console.error('Ошибка при загрузке статистики:', err);
                    setError('Ошибка при загрузке статистики');
                }
            };
            fetchStatistics();
        }
    }, []);

    const handleAddDescription = async () => {
        const user = localStorage.getItem('user');
        if (user) {
            const userData = JSON.parse(user);
            userData.description = newDescription;
    
            try {
                await axios.post(`http://31.15.18.177:5050/api/update_user_description`, {
                    ID: userData.id,
                    Description: newDescription,
                });
                localStorage.setItem('user', JSON.stringify(userData));
                setDescription(newDescription);
                setNewDescription('');
                setIsEditingDescription(false);
            } catch (err) {
                setError('Ошибка при добавлении описания');
            }
        }
    };
    

    const handleAddPfp = async () => {
        if (!newPfp.trim()) {
            setError('URL аватарки не может быть пустым.');
            return;
        }
    
        const user = localStorage.getItem('user');
        if (user) {
            const userData = JSON.parse(user);
            userData.pfp = newPfp;
    
            try {
                await axios.post(`http://31.15.18.177:5050/api/update_user_pfp`, {
                    ID: userData.id,
                    Pfp: newPfp,
                });
                localStorage.setItem('user', JSON.stringify(userData));
                setPfp(newPfp);
                setNewPfp('');
                setIsEditingPfp(false);
                setError(null);
            } catch (err) {
                setError('Ошибка при добавлении аватарки');
            }
        }
    };
    

    return (
        <div>
            <div className="leftBlock">
                <div className="statistics">
                    <h2>Статистика:</h2>
                    <p>Общее число оценок: {ratings}</p>
                    <p>Общее число обсуждений: {discussions}</p>
                    <p>Общее число обзоров: {reviews}</p>
                    <p>Общее число подборок: {collections}</p>
                </div>
                <Link to="/production-status" className="to-status-button">
                    <button className="big-button">Перейти к статусу произведений</button>
                </Link>
            </div>
            <div className="profile-container">
                <div 
                    className="profile-pic-container"
                    onMouseEnter={() => setIsHovered(true)} 
                    onMouseLeave={() => setIsHovered(false)}
                >
                    {pfp ? (
                        <img className="profile-pic" src={pfp} alt="Profile" />
                    ) : (
                        <p>Аватарка не найдена</p>
                    )}
                    {(isHovered || !pfp) && (
                        <button onClick={() => setIsEditingPfp(true)} className="change-pic-button">
                            Сменить аватарку
                        </button>
                    )}
                    {isEditingPfp && (
                        <div className="edit-pfp-container">
                            <input
                                type="text"
                                value={newPfp}
                                onChange={(e) => setNewPfp(e.target.value)}
                                placeholder="Введите URL изображения"
                            />
                            <button onClick={handleAddPfp} className="add-description-button" >Сохранить аватарку</button>
                            <button onClick={() => {
                                setIsEditingPfp(false);
                                setError(null);
                            }} className="cancel-button">✖️</button>
                        </div>
                    )}
                    {error && <p className="error-message">{error}</p>}
                </div>
                <p className="profile-nickname">{username}</p>
                {description ? (
                    <div className="profile-description-container">
                        <p className="profile-description">{description}</p>
                        <button onClick={() => {
                            setIsEditingDescription(true);
                            setNewDescription(description);
                        }} className="add-description-button">
                            Изменить описание
                        </button>
                    </div>
                ) : (
                    <button onClick={() => setIsEditingDescription(true)} className="add-description-button">
                        Добавить описание
                    </button>
                )}

                {isEditingDescription && (
                    <div className="edit-description-container">
                        <input
                            type="text"
                            value={newDescription}
                            onChange={(e) => setNewDescription(e.target.value)}
                            placeholder="Введите описание"
                        />
                        <button onClick={handleAddDescription} className="add-description-button">Сохранить описание</button>
                        <button onClick={() => {
                            setIsEditingDescription(false);
                            setError(null);
                        }} className="cancel-button">✖️</button>
                    </div>
                )}

            </div>
        </div>
    );
};

export default Profile;
