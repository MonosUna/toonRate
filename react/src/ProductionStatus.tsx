import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

interface Production {
    id: string;
    title: string;
    year: number;
}

const ProductionStatus: React.FC = () => {
    const [viewed, setViewed] = useState<Production[]>([]);
    const [watching, setWatching] = useState<Production[]>([]);
    const [toWatch, setToWatch] = useState<Production[]>([]);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const user = localStorage.getItem('user');
        const userData = JSON.parse(user);
        const fetchProductionsStatus = async () => {
            try {
                const response = await axios.get(
                    `http://31.15.18.177:5050/api/all_product_status/${userData.username}`
                );
                if (response.status == 200) {
                    setViewed(response.data.viewed || []);
                    setWatching(response.data.watching || []);
                    setToWatch(response.data.to_watch || []);
                }
            } catch (err) {
                setError('Ошибка при загрузке данных о статусах произведений.');
            }
        };
        fetchProductionsStatus();
    }, []);

    if (error) {
        return <div>{error}</div>;
    }

    return (
        <div className="production-status">
            <h1>Статусы произведений</h1>
            <div className="production-columns">
                <div className="column">
                    <h2>Просмотрено</h2>
                    <ul>
                        {viewed.map((prod) => (
                            <li key={prod.id}>
                                <Link to={`/production/${prod.id}`}>
                                    {prod.title} ({prod.year})
                                </Link>
                            </li>
                        ))}
                    </ul>
                </div>

                <div className="column">
                    <h2>Смотрю</h2>
                    <ul>
                        {watching.map((prod) => (
                            <li key={prod.id}>
                                <Link to={`/production/${prod.id}`}>
                                    {prod.title} ({prod.year})
                                </Link>
                            </li>
                        ))}
                    </ul>
                </div>

                <div className="column">
                    <h2>Хочу посмотреть</h2>
                    <ul>
                        {toWatch.map((prod) => (
                            <li key={prod.id}>
                                <Link to={`/production/${prod.id}`}>
                                    {prod.title} ({prod.year})
                                </Link>
                            </li>
                        ))}
                    </ul>
                </div>
            </div>
        </div>
    );
};

export default ProductionStatus;
