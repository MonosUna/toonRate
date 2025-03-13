import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';

interface Comment {
    text: string;
    author: string;
    entity_id: number;
    id: number
}

interface Big_Comment {
    text: string;
    author: string;
    entity_id: number;
    id: number
    ratings: {}
}

interface Review {
    id: string;
    production: string;
    owner: string;
    topic: string;
    message: string;
}

const Review = () => {
    const { id } = useParams<{ id: string }>();
    const [user, setUser] = useState<{ username: string; discussions: string[] } | null>(null);
    const [review, setReview] = useState<Review | null>(null);
    const [comments, setComments] = useState<[]>([]);
    const [newCommentText, setNewCommentText] = useState('');

    useEffect(() => {
        const storedUser = localStorage.getItem('user');
        if (storedUser) {
            setUser(JSON.parse(storedUser));
        }

        const fetchReview = async () => {
            try {
                const response = await axios.get(`http://31.15.18.177:5050/api/get_review/${id}`);
                setReview(response.data || null);
            } catch (err) {
                console.error('not found review:', err);
            }
        };

        fetchReview();
    }, [id]);

    useEffect(() => {
        const fetchComments = async () => {
            try {
                const response = await axios.get(`http://31.15.18.177:5050/api/get_review_comments/${id}`);
                const commentsData = response.data || [];
    
                for (let comment of commentsData) {
                    const ratingResponse = await axios.get(`http://31.15.18.177:5050/api/get_comment_rating/${comment.id}`);
                    const ratingsData = ratingResponse.data?.ratings || [];

                    const ratings = ratingsData.reduce((acc: { [key: string]: number }, rating: { user: string, rating: number }) => {
                        acc[rating.user] = rating.rating;
                        return acc;
                    }, {});
    
                    comment.ratings = ratings;
                    console.log(ratings)
                }
    
                setComments(commentsData);
            } catch (err) {
                console.error('', err);
            }
        };
    
        fetchComments();
    }, [id]);

    const handleCommentSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        const newComment: Comment = {
            text: newCommentText,
            author: user.username,
            entity_id: Number(id),
            id: 0
        };

        try {
            await axios.post(`http://31.15.18.177:5050/api/add_review_comment`, newComment);

            try {
                const response = await axios.get(`http://31.15.18.177:5050/api/count_of_comments`);
                newComment.id = response.data.comment_count
            } catch (err) {}

            const newBigComment: Big_Comment = {
                text: newCommentText,
                author: user.username,
                entity_id: Number(id),
                id: newComment.id,
                ratings: {}
            };

            const updatedComments = [
                ...comments,
                newBigComment,
            ];

            setComments(updatedComments);
        } catch (err) {}

        setNewCommentText('');
    };

    const handleRating = async (index: number, rating: number) => {
        const commentId = comments[index].id;
        const updatedRatings = {
            ...comments[index].ratings,
            [user.username]: rating,
        };

        const newRating = {
            author: user.username,
            rating: rating,
            id: Number(commentId)
        };
    
        try {
            await axios.post(`http://31.15.18.177:5050/api/update_comment_rating`, newRating);
            
            const updatedComments = [...comments];
            updatedComments[index].ratings = updatedRatings;
            console.log(updatedComments[index].ratings)
            setComments(updatedComments);
        } catch (err) {
            console.error('Ошибка при обновлении рейтинга', err);
        }
    };


    return (
        <div className="discussion">
            {review ? (
                <>
                    <h1>{review.topic}</h1>
                    <span><strong>Создатель:</strong> {review.author}</span>
                    <div className="msg">{review.message}</div>

                    <h3>Комментарии</h3>
                    {comments.length === 0 ? (
                        <span>Без комментариев.</span>
                    ) : (
                        comments.map((comment, index) => {
                            const totalRating = Object.values(comment.ratings).reduce((sum: number, rating) => sum +  (rating as number), 0)

                            return (
                                <div key={comment.id} className="comment">
                                    <strong>{comment.author}</strong>: {comment.text} 
                                    <span className="rating">Рейтинг: {totalRating}</span>
                                    {user && (
                                        <div className="rating-buttons">
                                            <button onClick={() => handleRating(index, 1)}>↑</button>
                                            <button onClick={() => handleRating(index, -1)}>↓</button>
                                        </div>
                                    )}
                                </div>
                            );
                        })
                    )}

                    {user && (
                        <form onSubmit={handleCommentSubmit}>
                            <textarea
                                value={newCommentText}
                                onChange={(e) => setNewCommentText(e.target.value)}
                                placeholder="Напишите комментарий..."
                                required
                            />
                            <button type="submit">Отправить комментарий</button>
                        </form>
                    )}
                </>
            ) : (
                <span>Loading...</span>
            )}
        </div>
    );
};

export default Review;
