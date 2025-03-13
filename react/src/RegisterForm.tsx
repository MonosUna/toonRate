import React, {useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import axios from "axios"

const RegisterForm: React.FC = () => {
    const [formData, setFormData] = useState({
        username: '',
        email: '',
        password: '',
        rpassword: '',
        pfp: '',
        description: ''
    })

    const [errors, setErrors] = useState({})
    const [valid, setValid] = useState(true)
    const navigate = useNavigate()

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        let errors: { [key: string]: string } = {};
        let isOK = true;

        if (formData.password != formData.rpassword) {
            alert('Пароли не совпадают попробуйте снова.');
            return;
        }
    
        axios.get('http://31.15.18.177:5050/api/get_users')
            .then(result => {
                const users = result.data.users;
                const maxId = users.reduce((max: number, user: { id: number }) => Math.max(max, user.id), 0);

                const newUser = {
                    id: maxId + 1,
                    username: formData.username,
                    password: formData.password,
                    email: formData.email,
                    pfp: 'https://avatars.mds.yandex.net/get-entity_search/2040540/791428223/S600xU_2x',
                    description: 'Нет описания',
                };
    
                const usernameExists = users.some((user: { username: string }) => user.username === formData.username);
                const emailExists = users.some((user: { email: string }) => user.email === formData.email);
    
                if (usernameExists) {
                    errors.name = 'Имя пользователя уже занято';
                    alert('Имя пользователя уже занято');
                    isOK = false;
                } else {
                    if (emailExists) {
                        errors.email = 'Email уже используется';
                        alert('Email уже используется');
                        isOK = false;
                    } else {
                        axios.post('http://31.15.18.177:5050/api/add_user', newUser)
                        .then(() => {
                            alert('Регистрация пройдена');
                            navigate('/login');
                        })
                        .catch(err => console.error('Ошибка при добавлении пользователя:', err));
                    }
                }
    
                setErrors(errors);
                setValid(isOK);
            })
            .catch(err => console.error('Ошибка при получении пользователей:', err));
    };
    


    return (
        <div className="registerbody">
            <div className="registerForm">
                <h1>ToonRate</h1>
                <h3>Создайте аккаунт</h3>

                <form action="" onSubmit={handleSubmit}>
                    {
                        valid ? <></>:
                        <span className="text-danger">
                            <p>{errors.email}</p>
                            <p>{errors.rpassword}</p>
                        </span>
                    }
                    <label htmlFor="username">Имя аккаунта:</label>
                    <input
                        type="text"
                        name="username"
                        placeholder="Введите логин"
                        required
                        onChange={(event) => setFormData({...formData, username:event.target.value}) }
                    />

                    <label htmlFor="email">Email:</label>
                    <input
                        type="email"
                        name="email"
                        placeholder="Введите адрес электронной почты"
                        required
                        onChange={(event) => setFormData({...formData, email:event.target.value}) }
                    />

                    <label htmlFor="password">Пароль:</label>
                    <input
                        type="password"
                        name="password"
                        placeholder="Введите пароль"
                        required
                        onChange={(event) => setFormData({...formData, password:event.target.value}) }
                    />

                    <label htmlFor="rpassword">Повторите пароль:</label>
                    <input
                        type="password"
                        name="rpassword"
                        placeholder="Повторите пароль"
                        required
                        onChange={(event) => setFormData({...formData, rpassword:event.target.value}) }
                    />

                    <div className="wrap">
                        <button type="submit">Отправить</button>
                    </div>
                </form>

                <p>
                    Уже в потоке?
                    <Link to="/login">
                        Заходи в аккаунт
                    </Link>
                </p>
            </div>
        </div>
    );
};

export default RegisterForm;
