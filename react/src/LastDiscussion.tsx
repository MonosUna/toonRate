import React, { useState, useEffect } from 'react';
import axios from 'axios';

interface User {
    username: string;
    pfp: string;
}

interface LastComment {
    text: string;
    username: string;
    pfp: string;
}

const LastDiscussion: React.FC = () => {
    const [lastComment, setLastComment] = useState<LastComment | null>(null);

    useEffect(() => {
        const fetchLastComment = async () => {
            try {
                const response = await axios.get('http://31.15.18.177:5050/api/last_discussion');
                if (response.data && response.data.text) {
                    setLastComment(response.data);
                } else {
                    setLastComment(null);
                }
            } catch (err) {
                console.error('Не могу получить данные о последнем комментарии:', err);
            }
        };

        fetchLastComment();
    }, []);

    return (
        <div className="last-discussion">
            <h1>Главное обсуждение:</h1>
            <h3>Последнее сообщение:</h3>
            {lastComment ? (
                <div className="comment">
                    <div className="user-info">
                        <img
                            src={lastComment.pfp}
                            alt={`${lastComment.username}'s profile`}
                            className="user-pfp"
                        />
                        <div>
                            <strong>{lastComment.username}</strong>
                        </div>
                    </div>
                    <p>{lastComment.text}</p>
                </div>
            ) : (
                <p>Нет комментариев.</p>
            )}
        </div>
    );
};

export default LastDiscussion;
