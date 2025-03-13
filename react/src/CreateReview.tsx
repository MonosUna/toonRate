import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link, useNavigate } from 'react-router-dom';

const CreateReview = () => {
    const [user, setUser] = useState<{ username: string; reviews: { [key: string]: number } } | null>(null);
    const [prod, setProd] = useState('');
    const [selectedProduction, setSelectedProduction] = useState(0);
    const [filteredProductions, setFilteredProductions] = useState([]);
    const [productions, setProductions] = useState([]);
    const [reviews, setReviews] = useState([]);
    const [searchTerm, setSearchTerm] = useState('');
    const [showSearch, setShowSearch] = useState(false);
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate()
    const [formData, setFormData] = useState({
        topic: "",
        message: "",
    })
    const handleSearchToggle = () => {
        setShowSearch(!showSearch);
        setSearchTerm('');
    };

    useEffect(() => {
        const storedUser = localStorage.getItem('user');
        if (storedUser) {
            setUser(JSON.parse(storedUser));
        }

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

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        
        if (user && selectedProduction) {
            console.log(reviews)
            const newFormData = {
                author: user.username,
                production: selectedProduction,
                id: reviews.length + 1,
                topic: formData.topic,
                message: formData.message,
            };

            axios.post(`http://31.15.18.177:5050/api/add_review`, newFormData)
                .then(result => {
                    navigate(`/reviews/${String(newFormData.id)}`);
                })
                .catch(err => console.error(err));
        } else {
            setErrorMessage('Пожалуйста, выберите произведение и войдите в систему.');
        }
    };

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

    const handleSelectProduction = (id, name) => {
        setSelectedProduction(id);
        setProd(name)
        setSearchTerm('');
        setShowSearch(false);
    };

    return (
        <div className="create-discussion">
        <h2>Создать новое обзор</h2>
        {errorMessage && <div className="error-message">{errorMessage}</div>}
        <div className="search-create-container">
            <button className="search-button" onClick={handleSearchToggle}>
                {showSearch ? 'Скрыть поиск' : 'Обзор на произведение'}
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
                                    onClick={() => handleSelectProduction(production.id, production.title)}
                                >
                                    {production.title} ({production.year})
                                </div>
                            ))}
                        </div>
                    )}
                </div>
            )}
        </div>
        Обзор на произведение: {prod}
        <form onSubmit={handleSubmit} className="discussion-form">
            <div className="form-group">
                <label htmlFor="topic">Заголовок:</label>
                <input
                    id="topic"
                    type="text"
                    className="form-input"
                    onChange={(event) => setFormData({...formData, topic: event.target.value}) }
                    required
                />
            </div>
            <div className="form-group">
                <label htmlFor="text">Обзор:</label>
                <textarea
                    id="text"
                    className="form-textarea"
                    onChange={(event) => setFormData({...formData, message: event.target.value}) }
                    required
                />
            </div>
            <button type="submit" className="submit-button">Создать</button>
        </form>
    </div>
    );
};

export default CreateReview;
