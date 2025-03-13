import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link, useNavigate } from 'react-router-dom';

const CreateCollection = () => {
    const [user, setUser] = useState<{ username: string; collections: string[] } | null>(null);
    const [prod, setProd] = useState('');
    const [selectedProductions, setSelectedProductions] = useState([]);
    const [productions, setProductions] = useState([]);
    const [collections, setCollections] = useState([]);
    const [collectionProductions, setcollectionProductions] = useState([]);
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

    const handleSelectProduction = (id, title) => {
        const alreadySelected = selectedProductions.some(item => item.id === id);
        if (!alreadySelected) {
            setSelectedProductions([
                ...selectedProductions,
                { id, title, comment: '' }
            ]);
        }
        setShowSearch(false);
    };

    const handleDeselectProduction = (id) => {
        setSelectedProductions(selectedProductions.filter(item => item.id !== id));
    };

    const handleCommentChange = (id, comment) => {
        setSelectedProductions(selectedProductions.map(item => 
            item.id === id ? { ...item, comment } : item
        ));
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

        const fetchCollections = async () => {
            try {
                const response = await axios.get('http://31.15.18.177:5050/api/get_collections');
                setCollections(response.data.collections || []);
            } catch (err) {}
        };

        const fetchCollectionProductions = async () => {
            try {
                const response = await axios.get('http://31.15.18.177:5050/api/get_collections_productions');
                setcollectionProductions(response.data.collectionsProductions || []);
            } catch (err) {}
        };

        fetchProductions();
        fetchCollections();
        fetchCollectionProductions();
    }, []);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        
        if (user && selectedProductions.length > 0) {
            const newFormData = {
                author: user.username,
                id: collections.length + 1,
                topic: formData.topic,
                message: formData.message,
            };

            await axios.post(`http://31.15.18.177:5050/api/add_collection`, newFormData);

            for (const prod of selectedProductions) {
                const itemData = {
                    id: collectionProductions.length + 1,
                    collection_id: newFormData.id,
                    production_id: prod.id,
                    comment: prod.comment,
                };
    
                await axios.post(`http://31.15.18.177:5050/api/add_collection_production`, itemData);

                collectionProductions.push(itemData);
            }
            navigate(`/collections/${String(newFormData.id)}`);
        } else {
            alert("выберите произведения или войдите в аккаунт")
        }
    };
    

    const filteredProductions = productions
    .filter(production => 
        production.title.toLowerCase().includes(searchTerm.toLowerCase())
    )
    .slice(0, 4);


    return (
            <div className="create-discussion">
                <h1>Создать коллекцию</h1>
                
                {errorMessage && <div className="error-message">{errorMessage}</div>}
        
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
                        <label htmlFor="text">Описание коллекции:</label>
                        <textarea
                            id="text"
                            className="form-textarea"
                            onChange={(event) => setFormData({...formData, message: event.target.value}) }
                            required
                        />
                    </div>
        
                    <div className="search-create-container">
                        <button  type="button" className="search-button" onClick={handleSearchToggle}>
                            {showSearch ? 'Скрыть поиск' : 'Выбрать произведение'}
                        </button>
                        {showSearch && (
                            <div className="search-box">
                                <input
                                    type="text"
                                    className="search-input"
                                    placeholder="Поиск по названию"
                                    value={searchTerm}
                                    onChange={e => setSearchTerm(e.target.value)}
                                />
                                <div>
                                    {filteredProductions.map(production => (
                                        <li key={production.id}>
                                            <button
                                                type="button"
                                                className="dropdown-item"
                                                onClick={() => handleSelectProduction(production.id, production.title)}
                                            >
                                                {production.title}
                                            </button>
                                        </li>
                                    ))}
                                </div>
                            </div>
                        )}
                    </div>
        
                    <div class="selected-works-container">
                        <h3 class="selected-works-title">Выбранные произведения</h3>
                        <ul class="selected-works-list">
                            {selectedProductions.map(production => (
                                <li key={production.id} class="selected-works-item">
                                    <span class="selected-works-title-item">{production.title}</span>
                                    <button
                                        type="button"
                                        onClick={() => handleDeselectProduction(production.id)}
                                        class="selected-works-remove-btn"
                                    >
                                        Отменить выбор
                                    </button>
                                    <div>
                                        <textarea
                                            class="selected-works-textarea"
                                            placeholder="Добавьте комментарий"
                                            value={production.comment}
                                            onChange={(e) => handleCommentChange(production.id, e.target.value)}
                                        />
                                    </div>
                                </li>
                            ))}
                        </ul>
                    </div>

        
                    <button type="submit" className="submit-button">Создать</button>
                </form>
            </div>
    );
};

export default CreateCollection;
