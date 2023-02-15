import React from 'react';
import {Routes, Route } from 'react-router-dom';
import Product from './Product';
import FeaturedProduct from './FeaturedProduct';
import './App.css';



function App() {
  return (
    <div className='App'>
      <Routes>
        <Route path="/" element={<FeaturedProduct/>} />
        <Route path="/products/:id" element={<Product/>} />
      </Routes>
    </div>
  );
}
export default App;


