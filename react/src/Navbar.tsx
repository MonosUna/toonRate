import React, { useEffect, useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import logoImg from './images/logo.png';

interface NavbarProps {
    user: { name: string } | null;
    onLogout: () => void;
}

const Navbar: React.FC<NavbarProps> = () => {
    const [user, setUser] = useState<{ name: string } | null>(null);
    const navigate = useNavigate();

    useEffect(() => {
        const storedUser = localStorage.getItem('user');
        if (storedUser) {
            setUser(JSON.parse(storedUser));
        }
    }, []);

    const handleLogout = () => {
        localStorage.removeItem('user');
        setUser(null);
        navigate('/');
        window.location.reload();
    };

    const handleLoginClick = () => {
        navigate('/login');
    };

    const handleRegisterClick = () => {
        navigate('/registration');
    };

    const handleTopClick = () => {
        navigate('/top-productions');
    };

    const handleProfileClick = () => {
        navigate('/profile');
    };

    const handleDiscussionsClick = () => {
        navigate('/discussions');
    };

    const handleReviewsClick = () => {
        navigate('/reviews');
    };

    const handleCollectionsClick = () => {
        navigate('/collections');
    };

    return (
        <nav className="navbar">
            <ul className="nav-list">
                <li>
                    <Link to="/">
                        <img src={logoImg} alt="logo" style={{ height: '30px' }} />
                    </Link>
                </li>
                <li><Link to="/">ToonRate</Link></li>
                <li><a onClick={handleTopClick}>Топы</a></li>
                <li><a onClick={handleDiscussionsClick}>Обсуждения</a></li>
                <li><a onClick={handleReviewsClick}>Обзоры</a></li>
                <li><a onClick={handleCollectionsClick}>Подборки</a></li>
            </ul>
            <div>
                <ul className="nav-list">
                    {user ? (
                        <>
                            <li>
                            <button onClick={handleProfileClick} className="profile-button">{user.username}</button>
                            </li>
                            <li>
                                <button onClick={handleLogout} className="navbar-button">Выйти</button>
                            </li>
                        </>
                    ) : (
                        <>
                            <li>
                                <button onClick={handleLoginClick} className="navbar-button">Войти</button>
                            </li>
                            <li>
                                <button onClick={handleRegisterClick} className="navbar-button">Зарегистрироваться</button>
                            </li>
                        </>
                    )}
                </ul>
            </div>
        </nav>
    );
};

export default Navbar;
