import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

const Reviews = () => {
    const [selectedProduction, setSelectedProduction] = useState("0");
    const [productions, setProductions] = useState([]);
    const [reviews, setReviews] = useState([]);
    const [searchTerm, setSearchTerm] = useState('');
    const [filteredProductions, setFilteredProductions] = useState([]);
    const [filteredReviews, setFilteredReviews] = useState([]);
    const [showSearch, setShowSearch] = useState(false);

    const [currentPage, setCurrentPage] = useState(1);
    const itemsPerPage = 5;

    const handleSearchToggle = () => {
        setShowSearch(!showSearch);
        setSearchTerm('');
    };

    useEffect(() => {
        const fetchProductions = async () => {
            try {
                const response = await axios.get('http://31.15.18.177:5050/api/get_productions');
                setProductions(response.data.productions);
            } catch (err) {}
        };

        const fetchReviews = async () => {
            try {
                const response = await axios.get('http://31.15.18.177:5050/api/get_reviews');
                setReviews(response.data.reviews || []);
            } catch (err) {}
        };

        fetchProductions();
        fetchReviews();
    }, []);

    useEffect(() => {
        if (productions?.length > 0) {
            const filtered = productions
                .filter(production =>
                    production.title.toLowerCase().includes(searchTerm.toLowerCase())
                )
                .slice(0, 5);
            setFilteredProductions(filtered);
        }
    }, [productions, searchTerm]);

    const handleSelectProduction = (id) => {
        setSelectedProduction(id);
        setSearchTerm('');
        setShowSearch(false);
        setCurrentPage(1);
    };

    useEffect(() => {
        if (reviews && reviews.length > 0) {
            const filtered = reviews.filter(review => 
                review.production == selectedProduction
            );
            setFilteredReviews(filtered);
        }
    }, [reviews, selectedProduction]);

    const indexOfLastItem = currentPage * itemsPerPage;
    const indexOfFirstItem = indexOfLastItem - itemsPerPage;
    const currentItems = filteredReviews.slice(indexOfFirstItem, indexOfLastItem);

    const totalPages = Math.ceil(filteredReviews.length / itemsPerPage);

    const goToNextPage = () => setCurrentPage((prev) => Math.min(prev + 1, totalPages));
    const goToPreviousPage = () => setCurrentPage((prev) => Math.max(prev - 1, 1));

    return (
        <div className="discussions">
            <div className="search-container">
                <button className="search-button" onClick={handleSearchToggle}>
                    {showSearch ? 'Скрыть поиск' : 'Обзор по произведению'}
                </button>
                
                {showSearch && (
                    <div className="search-box">
                        <input
                            type="text"
                            className="search-input"
                            placeholder="Введите название произведения..."
                            value={searchTerm}
                            onChange={(e) => setSearchTerm(e.target.value)}
                        />
                        {filteredProductions.length > 0 && (
                            <div className="dropdown-list">
                                {filteredProductions.map(production => (
                                    <div
                                        key={production.id}
                                        className="dropdown-item"
                                        onClick={() => handleSelectProduction(production.id)}
                                    >
                                        {production.title} ({production.year})
                                    </div>
                                ))}
                            </div>
                        )}
                    </div>
                )}
            </div>

            {selectedProduction !== "0" ? (
                <div className="discussion-list">
                    {currentItems.length > 0 ? (
                        currentItems.map(review => (
                            <div key={review.id} className="discussion-item">
                                <Link to={`/reviews/${review.id}`}>
                                    <h3>{review.topic}</h3>
                                </Link>
                                <p>Автор: {review.author}</p>
                            </div>
                        ))
                    ) : (
                        <p>Нет обзоров этого произведения.</p>
                    )}

                    {filteredReviews.length > itemsPerPage && (
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
                    )}
                </div>
            ) : (
                <div className="discussion-list">
                    <p>Выберите произведение, на которое хотите посмотреть обзоры или создайте обзор.</p> 
                </div>
            )}

            <button className="create_discussion_button"> 
                <Link to={`/create_review`}>
                    Создать обзор
                </Link>
            </button>
        </div>
    );
};

export default Reviews;
