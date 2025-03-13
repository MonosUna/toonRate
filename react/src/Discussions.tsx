import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

const Discussions = () => {
    const [selectedProduction, setSelectedProduction] = useState(0);
    const [productions, setProductions] = useState([]);
    const [discussions, setDiscussions] = useState([]);
    const [searchTerm, setSearchTerm] = useState('');
    const [filteredProductions, setFilteredProductions] = useState([]);
    const [filteredDiscussions, setFilteredDiscussions] = useState([]);
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

        const fetchDiscussions = async () => {
            try {
                const response = await axios.get('http://31.15.18.177:5050/api/get_discussions');
                setDiscussions(response.data.discussions || []);
            } catch (err) {}
        };

        fetchProductions();
        fetchDiscussions();
    }, []);

    useEffect(() => {
        if (productions?.length > 0) {
            const filtered = productions
                .filter(production =>
                    production.title.toLowerCase().includes(searchTerm.toLowerCase())
                )
                .slice(0, 4);
            setFilteredProductions(filtered);
        }
    }, [productions, searchTerm]);

    const handleSelectProduction = (id) => {
        setSelectedProduction(id);
        setSearchTerm('');
        setShowSearch(false);
    };

    useEffect(() => {
        if (discussions && discussions.length > 0) {
            const filtered = discussions.filter(discussion => 
                discussion.production == selectedProduction
            );
            setFilteredDiscussions(filtered);
        }
    }, [discussions, selectedProduction]);

    const indexOfLastItem = currentPage * itemsPerPage;
    const indexOfFirstItem = indexOfLastItem - itemsPerPage;
    const currentItems = filteredDiscussions.slice(indexOfFirstItem, indexOfLastItem);
    const totalPages = Math.ceil(filteredDiscussions.length / itemsPerPage);

    const goToNextPage = () => setCurrentPage((prev) => Math.min(prev + 1, totalPages));
    const goToPreviousPage = () => setCurrentPage((prev) => Math.max(prev - 1, 1));

    return (
        <div className="discussions">
            <div className="search-container">
                <button className="search-button" onClick={handleSearchToggle}>
                    {showSearch ? 'Скрыть поиск' : 'Обсуждение по произведению'}
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
                                <div
                                    className="dropdown-item"
                                    onClick={() => handleSelectProduction(0)}
                                >
                                    Общие
                                </div>
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

            <div className="discussion-list">
                {currentItems.length > 0 ? (
                    currentItems.map(discussion => (
                        <div key={discussion.id} className="discussion-item">
                            <Link to={`/discussions/${discussion.id}`}>
                                <h3>{discussion.topic}</h3>
                            </Link>
                            <p>Автор: {discussion.author}</p>
                            <p>{discussion.entry_message}</p>
                        </div>
                    ))
                ) : (
                    <p>Нет обсуждений для этого произведения.</p>
                )}
            </div>

            {filteredDiscussions.length > itemsPerPage && (
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

            <button className="create_discussion_button"> 
                <Link to={`/create_discussion`}>
                    Создать обсуждение
                </Link>
            </button>
        </div>
    );
};

export default Discussions;
