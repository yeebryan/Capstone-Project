import React from 'react';
import { Routes, Route } from 'react-router-dom';
import FeaturedProduct from './FeaturedProduct';
import Product from './Product';
import './App.css';



function App() {
  return (
    <div className='App'>
      <Routes>
        <Route path="/" element={<FeaturedProduct/>} />
        <Route path="/product/:id" element={<Product/>} />
      </Routes>
    </div>
  );
}
export default App;