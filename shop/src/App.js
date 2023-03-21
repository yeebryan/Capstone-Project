import React from 'react';
import {BrowserRouter, Routes, Route } from 'react-router-dom';
import Product from './Product';
import About from './Pages/About';
import FeaturedProduct from './FeaturedProduct';
import Login from './Login';
import UserPreferenceForm from './UserPreferenceForm';
import './App.css';
import MyOrder from './MyOrder';



function App() {
  return (
    <BrowserRouter>
    <div className='App'>
      <Routes>
        <Route path="/" element={<FeaturedProduct/>} />
        <Route path="/restaurants/:restaurant_id" element={<Product/>} />
        <Route path="/about" element={<About/>} />
        <Route path="/login" element={<Login/>} />
        <Route path="/userform" element = {<UserPreferenceForm/>} />
        <Route path="/myorder" element = {<MyOrder/>} />
      </Routes>
    </div>
    </BrowserRouter>
  );
}
export default App;


