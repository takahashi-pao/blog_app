import React from 'react';
import { BrowserRouter, Route, Routes, Link } from 'react-router-dom';
import Home from "./components/home/Home";
import Add from './components/register/Register';
import SignIn from './components/signIn/SignIn';
import SignUp from './components/signUp/SignUp';

function App(){
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/add" element={<Add />} />
                <Route path="/signIn" element={<SignIn />} />
                <Route path="/signUp" element={<SignUp />} />
            </Routes>
            
        </BrowserRouter>
    )
}

export default App