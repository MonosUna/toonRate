import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';

interface Production {
    id: string;
    title: string;
    description: string;
    genre: string;
    year: number;
}

const Product: React.FC = () => {
    const { id } = useParams<{ id: string }>();
    const [loading, setLoading] = useState<boolean>(true);
    const [currentStatus, setCurrentStatus] = useState<number>(0);
    const [error, setError] = useState<string | null>(null);
    const [newRating, setNewRating] = useState<number | ''>('');
    const [averageRating, setAverageRating] = useState<number | null>(null);
    const [production, setProduction] = useState<Production | null>(null);
    const [user, setUser] = useState<{ username: string } | null>(null);

    useEffect(() => {
        const storedUser = localStorage.getItem('user');
        if (storedUser) {
            setUser(JSON.parse(storedUser));
        }
    }, []);

    useEffect(() => {
        const fetchRatings = async () => {
            try {
                const response = await axios.get(`http://31.15.18.177:5050/api/get_production_ratings/${id}`);
                const ratings: number[] = response.data.ratings;
    
                if (ratings == null) {
                    setAverageRating(0);
                } else {
                    const total = ratings.reduce((sum, rating) => sum + rating, 0);
                    const average = total / ratings.length;
                    setAverageRating(average);
                }
            } catch (err) {
                setError('Ошибка при загрузке рейтингов');
            }
        };
    
        fetchRatings();
    }, [id]);

    useEffect(() => {
        const fetchStatus = async () => {
            if (user) {
                try {
                    const response = await axios.get(`http://31.15.18.177:5050/api/get_production_status/${id}/${user.username}`);
                    setCurrentStatus(response.data.status);
                } catch (err) {
                    setError('Ошибка при загрузке статуса');
                }
            }
        };

        fetchStatus();
    }, [id, user]);

    useEffect(() => {
        const fetchProduction = async () => {
            try {
                const response = await axios.get(`http://31.15.18.177:5050/api/production/${id}`);
                setProduction(response.data);
            } catch (err) {
                setError('');
            } finally {
                setLoading(false);
            }
        };

        fetchProduction();
    }, [id]);

    const handleRatingSubmit = async () => {
        if (user && production && newRating) {
            if (newRating < 1 || newRating > 10) {
                alert('Пожалуйста, введите рейтинг от 1 до 10.');
                return;
            }

            const ratingData = {
                ID: production.id,
                Author: user.username,
                Rating: newRating,
            };
    
            try {
                const response = await axios.post(`http://31.15.18.177:5050/api/set_new_production_rating`, ratingData);
                if (response.status === 200) {
                    setNewRating('');

                    const fetchRatings = async () => {
                        try {
                            const response = await axios.get(`http://31.15.18.177:5050/api/get_production_ratings/${id}`);
                            const ratings: number[] = response.data.ratings;
    
                            if (ratings == null) {
                                setAverageRating(0);
                            } else {
                                const total = ratings.reduce((sum, rating) => sum + rating, 0);
                                const average = total / ratings.length;
                                setAverageRating(average);
                            }
                        } catch (err) {
                            setError('Ошибка при загрузке рейтингов');
                        }
                    };
    
                    fetchRatings();
                } else {
                    setError('Ошибка при добавлении оценки для производства');
                }
            } catch (err) {
                setError('Ошибка при добавлении оценки для производства');
                console.error(err);
            }
        }
    };
    

    const handleProductionStatusChange = async (status: number) => {
        if (user && production) {
            const ratingData = {
                ID: production.id,
                Author: user.username,
                Rating: status,
            };

            try {
                const response = await axios.post(
                    `http://31.15.18.177:5050/api/set_new_production_status`, 
                    ratingData
                );
                if (response.status === 200) {
                    setCurrentStatus(status);
                }

            } catch (err) {
                setError('Ошибка при обновлении статуса произведения');
            }
        }
    };

    const statusLabels = {
        1: 'Просмотрено',
        2: 'Смотрю',
        3: 'Хочу посмотреть',
    };

    if (loading) {
        return <div>Загрузка...</div>;
    }

    if (error) {
        return <div>{error}</div>;
    }

    if (!production) {
        return <div>Произведение не найдено.</div>;
    }

    return (
        <div className="production-detail">
            <h1>{production.title} ({production.year})</h1>
            <p>{production.description}</p>
            <p>Жанр: {production.genre}</p>
            <p>Средняя оценка: {averageRating.toFixed(2)}</p>

            {user && (
                <div>
                    <h2>Оцените произведение</h2>
                    <input
                        type="number"
                        value={newRating}
                        onChange={(e) => setNewRating(Number(e.target.value))}
                        placeholder="Введите оценку (от 1 до 10)"
                    />
                    <button onClick={handleRatingSubmit}>Оценить</button>

                    <h2>Текущий статус: {statusLabels[currentStatus] || 'Не установлен'}</h2>
                    <button onClick={() => handleProductionStatusChange(0)}>Ничего</button>
                    <button onClick={() => handleProductionStatusChange(1)}>Просмотрено</button>
                    <button onClick={() => handleProductionStatusChange(2)}>Смотрю</button>
                    <button onClick={() => handleProductionStatusChange(3)}>Хочу посмотреть</button>
                </div>
            )}
        </div>
    );
};

export default Product;
