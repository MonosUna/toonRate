import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

interface Production {
    id: string;
    title: string;
    description: string;
    genre: string;
    year: number;
    average_rating: number;
}

const TopProductions: React.FC = () => {
    const [productions, setProductions] = useState<Production[]>([]);
    const [error, setError] = useState<string | null>(null);
    const [currentPage, setCurrentPage] = useState<number>(1);
    const itemsPerPage = 4;

    useEffect(() => {
        const fetchProductions = async () => {
            try {
                const response = await axios.get('http://31.15.18.177:5050/api/top_productions');
                setProductions(response.data.top_productions);
            } catch (err) {
                setError('Ошибка при загрузке данных');
            }
        };

        fetchProductions();
    }, []);

    const indexOfLastItem = currentPage * itemsPerPage;
    const indexOfFirstItem = indexOfLastItem - itemsPerPage;
    const currentItems = productions.slice(indexOfFirstItem, indexOfLastItem);

    const totalPages = Math.ceil(productions.length / itemsPerPage);

    const goToNextPage = () => setCurrentPage((prev) => Math.min(prev + 1, totalPages));
    const goToPreviousPage = () => setCurrentPage((prev) => Math.max(prev - 1, 1));

    if (error) {
        return <div>{error}</div>;
    }

    return (
        <div className="top-productions">
            <h1>Топ произведений</h1>
            <ul>
                {currentItems.map((production) => (
                    <li key={production.id}>
                        <Link to={`/production/${production.id}`}>
                            <h2>{production.title} ({production.year})</h2>
                        </Link>
                        <p>{production.description}</p>
                        <p>Жанр: {production.genre}</p>
                        <p>Средняя оценка: {production.average_rating.toFixed(2)}</p>
                    </li>
                ))}
            </ul>

            <div className="pagination">
                <button onClick={goToPreviousPage} disabled={currentPage === 1}>
                    Предыдущая
                </button>
                <span>
                    Страница {currentPage} из {totalPages}
                </span>
                <button onClick={goToNextPage} disabled={currentPage === totalPages}>
                    Следующая
                </button>
            </div>
        </div>
    );
};

export default TopProductions;
