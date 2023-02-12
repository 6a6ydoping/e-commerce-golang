import React, {useState,setState} from 'react';
import { Link } from 'react-router-dom';
import Header from './header';
import './style.css';

function Home(){
    return(
        <>
        <Header headerName="Home page"/>
        <Link to="/menu">All items page</Link>
        </>
    );
}

export default Home;