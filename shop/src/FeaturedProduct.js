import React, { useState, useEffect } from 'react';
import {Link} from 'react-router-dom'
import axios from 'axios';
import "./App.css";


// this .js is functional hence we can use HOOK

// arrow function 
const FeaturedProduct = () => {
    // `products` is the current state // `setProducts` function that updates the state //
    const [products, setProducts] = useState([]);




// | useEFFECTS | 
// allow to perform fetching data; direct update DOM; console messages; 
// loading data; working with animations


// one thing to thing note is that useEffect takes in 2 arguments
// in this case, the 1st one is : fetchProducts(); // function that will be called whenever we want our effect to happen

// 2nd one is: [] empty array which means the `EFFECT` will not be called again after the 1st render
// it has to do with 'WHEN the effect is actually being called'
// it is also called the `DEPENDENCY ARRAY`
// Because we can use it to define when this should called

useEffect(() =>{
    fetchProducts();
}, []);

// fetch data from API
const fetchProducts = () => {
    axios
        .get('https://fakestoreapi.com/products')
        .then((res) =>{
            console.log(res); // response // response.data
            setProducts(res.data) // take note it must be res.data.data because res.data is an object not an array // passing data into our setProducts function so that we can set our state to the data
        })
        .catch((err) => {
            console.log(err);
        })
}
// display the data (JSX corner)
    return (
        <div>
            <h1>Featured Product</h1>
            <div className='item-container'>
                {products.map((product) =>(
                    <div className='card' key={product.id}> 
                        <img src={product.image} alt='' />
                        <h3>{product.title}</h3>
                        <h3>{`$${product.price.toFixed(2)}`}</h3>
                        <p>{product.description}</p>
                        <Link to={`/products/${product.id}`}>View</Link>
                    </div>
                ))}
            </div> {/* container div */}
        </div>
    )
}

export default FeaturedProduct

