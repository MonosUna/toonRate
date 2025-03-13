import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';

const Sidebar: React.FC = () => {
    const [randomProductionId, setRandomProductionId] = useState<string | null>(null);
    const [randomDiscussionId, setRandomDiscussionId] = useState<string | null>(null);
    const [randomReviewId, setRandomReviewId] = useState<string | null>(null);
    const [randomCollectionId, setRandomCollectionId] = useState<string | null>(null);

    useEffect(() => {
        const fetchRandomProduction = async () => {
            try {
                const response = await axios.get('http://31.15.18.177:5050/api/random_production');
                if (response.data.id) {
                    setRandomProductionId(response.data.id);
                }
            } catch (err) {
                console.error('Error fetching random production:', err);
            }
        };

        const fetchRandomDiscussion = async () => {
            try {
                const response = await axios.get('http://31.15.18.177:5050/api/random_discussion');
                if (response.data.id) {
                    setRandomDiscussionId(response.data.id);
                }
            } catch (err) {
                console.error('Error fetching random discussion:', err);
            }
        };

        const fetchRandomReview = async () => {
            try {
                const response = await axios.get('http://31.15.18.177:5050/api/random_review');
                if (response.data.id) {
                    setRandomReviewId(response.data.id);
                }
            } catch (err) {
                console.error('Error fetching random review:', err);
            }
        };

        const fetchRandomCollection = async () => {
            try {
                const response = await axios.get('http://31.15.18.177:5050/api/random_collection');
                if (response.data.id) {
                    setRandomCollectionId(response.data.id);
                }
            } catch (err) {
                console.error('Error fetching random collection:', err);
            }
        };

        fetchRandomProduction();
        fetchRandomDiscussion();
        fetchRandomReview();
        fetchRandomCollection();
    }, []);

    return (
        <nav>
            <div className="right-sidebar">
                {randomProductionId && (
                    <Link to={`/production/${randomProductionId}`}>
                        Случайное произведение
                    </Link>
                )}
                {randomDiscussionId && (
                    <Link to={`/discussions/${randomDiscussionId}`}>
                        Случайное обсуждение
                    </Link>
                )}
                {randomCollectionId && (
                    <Link to={`/collections/${randomCollectionId}`}>
                        Случайная подборка
                    </Link>
                )}
                {randomReviewId && (
                    <Link to={`/reviews/${randomReviewId}`}>
                        Случайный обзор
                    </Link>
                )}
            </div>
        </nav>
    );
};

export default Sidebar;