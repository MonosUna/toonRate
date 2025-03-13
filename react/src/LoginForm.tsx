import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import axios from "axios";

interface LoginFormProps {
    onUserLogin: (user: { name: string }) => void;
}

const LoginForm: React.FC<LoginFormProps> = () => {
    const [formData, setFormData] = useState({
        name: '',
        password: '',
    });

    const [errors, setErrors] = useState({});
    const [valid, setValid] = useState(true);
    const navigate = useNavigate();

    const handleUserLogin = (loggedInUser: { name: string }) => {
        localStorage.setItem('user', JSON.stringify(loggedInUser)); 
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        let errors: { [key: string]: string } = {};
        let isOK = true;
    
        axios.get('http://31.15.18.177:5050/api/get_users')
            .then(result => {
                const users = result.data.users;
                const user = users.find((user: { username: string; password: string }) => 
                    user.username === formData.name && user.password === formData.password
                );
                
                if (user) {
                    handleUserLogin(user);
                    navigate('/');
                    window.location.reload();
                } else {
                    errors.password = "Неправильный логин или пароль";
                    isOK = false;
                    setErrors(errors);
                    setValid(isOK);
                }
            })
            .catch(err => console.error(err));
    };
    return (
        <div className="registerbody">
            <div className="registerForm">
                <h1>ToonRate</h1>
                <h3>Войдите в аккаунт</h3>

                <form onSubmit={handleSubmit}>
                    {!valid && (
                        <span className="text-danger">
                            <p>{errors.password}</p>
                        </span>
                    )}
                    <label htmlFor="name">Имя аккаунта:</label>
                    <input
                        type="text"
                        name="name"
                        placeholder="Введите логин"
                        required
                        onChange={(event) => setFormData({ ...formData, name: event.target.value })}
                    />

                    <label htmlFor="password">Пароль:</label>
                    <input
                        type="password"
                        id="password"
                        name="password"
                        placeholder="Введите пароль"
                        required
                        onChange={(event) => setFormData({ ...formData, password: event.target.value })}
                    />

                    <div className="wrap">
                        <button type="submit">Отправить</button>
                    </div>
                </form>

                <p>
                    Вы новенький?
                    <Link to="/registration"> Создайте аккаунт </Link>
                </p>
            </div>
        </div>
    );
};

export default LoginForm;
