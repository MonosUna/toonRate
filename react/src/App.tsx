import React from 'react';
import Navbar from './Navbar';
import LoginForm from './LoginForm';
import RegisterForm from './RegisterForm';
import TopProductions from './TopProductions';
import Product from './Product';
import Profile from './Profile';
import Discussions from './Discussions'
import Discussion from './Discussion';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import ProductionStatus from './ProductionStatus';
import CreateDiscussion from './CreateDiscussion';
import Reviews from './Reviews';
import Review from './Review';
import CreateReview from './CreateReview';
import Collections from './Collections';
import Collection from './Collection';
import CreateCollection from './CreateCollection';
import MainPage from './MainPage';
import './styles.css';

const App: React.FC = () => {

    return (
        <BrowserRouter>
            <div>
            <Navbar /> 
                <Routes>
                    <Route path="/" element={< MainPage />} />
                    <Route path="/login" element={< LoginForm />} />
                    <Route path="/registration" element={< RegisterForm />} />
                    <Route path="/top-productions" element={< TopProductions />}/>
                    <Route path="/production/:id" element={< Product />}/>
                    <Route path="/profile" element = {< Profile />} />
                    <Route path="/production-status" element = {< ProductionStatus />}/>
                    <Route path="/discussions" element = {< Discussions />}/>
                    <Route path="/discussions/:id" element = {< Discussion />}/>
                    <Route path="/create_discussion" element = {< CreateDiscussion />}/>
                    <Route path="/reviews" element = {< Reviews />}/>
                    <Route path="/reviews/:id" element = {< Review />}/>
                    <Route path="/create_review" element = {< CreateReview />}/>
                    <Route path="/collections" element = {< Collections />}/>
                    <Route path="/collections/:id" element = {< Collection />}/>
                    <Route path="/create_collection" element = {< CreateCollection />}/>
                </Routes>
            </div>
        </BrowserRouter>
    );
};

export default App;
