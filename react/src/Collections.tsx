import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

const Collections = () => {
    const [collections, setCollections] = useState([]);
    const [currentPage, setCurrentPage] = useState(1);
    const itemsPerPage = 5;

    useEffect(() => {
        const fetchCollections = async () => {
            try {
                const response = await axios.get('http://31.15.18.177:5050/api/get_collections');
                setCollections(response.data.collections || []);
            } catch (err) {}
        };

        fetchCollections();
    }, []);

    const indexOfLastItem = currentPage * itemsPerPage;
    const indexOfFirstItem = indexOfLastItem - itemsPerPage;
    const currentItems = collections.slice(indexOfFirstItem, indexOfLastItem);

    const totalPages = Math.ceil(collections.length / itemsPerPage);

    const goToNextPage = () => setCurrentPage((prev) => Math.min(prev + 1, totalPages));
    const goToPreviousPage = () => setCurrentPage((prev) => Math.max(prev - 1, 1));

    return (
        <div className="discussions">
            <div className="discussion-list">
                {currentItems.length > 0 ? (
                    currentItems.map(collection => (
                        <div key={collection.id} className="discussion-item">
                            <Link to={`/collections/${collection.id}`}>
                                <h3>{collection.topic}</h3>
                            </Link>
                            <p>{collection.entry_message}</p>
                            <p>Автор: {collection.author}</p>
                        </div>
                    ))
                ) : (
                    <p>Нет подборок.</p>
                )}
            </div>

            {collections.length > itemsPerPage && (
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
                <Link to={`/create_collection`}>
                    Создать подборку
                </Link>
            </button>
        </div>
    );
};

export default Collections;
