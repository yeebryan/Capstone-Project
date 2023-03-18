import React from 'react';
import {BrowserRouter, Routes, Route } from 'react-router-dom';
import Product from './Product';
import About from './About';
import FeaturedProduct from './FeaturedProduct';

import './App.css';



function App() {
  return (
    <BrowserRouter>
    <div className='App'>
      <Routes>
        <Route path="/" element={<FeaturedProduct userId="123"/>} />
        <Route path="/products/:id" element={<Product/>} />
        <Route path="/about" element={<About/>} />
      </Routes>
    </div>
    </BrowserRouter>
  );
}
export default App;


