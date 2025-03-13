import React from 'react';
import Sidebar from './Sidebar';
import TopMainPage from './TopMainPage';
import LastDiscussion from './LastDiscussion';
import Footer from './Footer';

const MainPage: React.FC = () => {
    return (
        <> 
            <Sidebar /> <TopMainPage />  <LastDiscussion/> <Footer />
        </>
 
    );
};

export default MainPage;
