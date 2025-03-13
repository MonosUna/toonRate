import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

interface Production {
    id: string;
    title: string;
    average_rating: number;
}

const TopMainPage: React.FC = () => {
    const [productions, setProductions] = useState<Production[]>([]);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchProductions = async () => {
            try {
                const response = await axios.get('http://31.15.18.177:5050/api/top_5_productions');
                setProductions(response.data.productions);
            } catch (err) {
                setError('Ошибка при загрузке данных');
            }
        };

        fetchProductions();
    }, []);

    if (error) {
        return <div>{error}</div>;
    }

    return (
        <div className="top-productions-mainpage">
            <h1>Топ произведений</h1>
            <ul>
                {productions.map((production) => (
                    <li key={production.id}>
                        <Link to={`/production/${production.id}`}>
                            <h2>
                                {production.title} - {production.average_rating.toFixed(2)}
                            </h2>
                        </Link>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default TopMainPage;