import React, { useState, useEffect } from 'react';
import "bootstrap/dist/css/bootstrap.min.css"
import {Link} from 'react-router-dom'
import AddToCartButton from './AddtoCart';
import {Modal, Button} from "react-bootstrap"
import Navbar1 from './Navbar';
import axios from 'axios';
import "./App.css";

// this .js is functional hence we can use HOOK

// arrow function 
const FeaturedProduct = (props) => {
    // `products` is the current state // `setProducts` function that updates the state //
    const [products, setProducts] = useState([]);
    const [cartCount, setCartCount] = useState(0);
    const [showCart, setShowCart] = useState(false);
    const [cart, setCart] = useState([])




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
            setProducts(res.data) // for storerestapi, take note it must be res.data.data because res.data is an object not an array // passing data into our setProducts function so that we can set our state to the data
        })
        .catch((err) => {
            console.log(err);
        })
}

// cart count, and keep track of item added

const onAddToCart = (item) => {
    setCartCount(cartCount + 1);
    setCart([...cart, item])
}

// open the cart

const onOpenCart = () => {
    setShowCart(true)
}

// clear cart btn function - setState back to 0

const clearCart = () => {
    setCartCount(0); // the counter to 0 
    setCart([]);
    setShowCart(false);

}


// modal component for Cart

const CartModal = () => {
    let total = 0;
    let quantity = 0;
    cart.forEach((item) => {
        total += item.price;
        quantity++;
    })
    return (
        <Modal show={showCart} onHide={() => setShowCart(false)}>
        <Modal.Header closeButton>
        <Modal.Title>Cart</Modal.Title>
        </Modal.Header>
        <Modal.Body>
        {/* Display items in the cart */}
        {cart.map((item, index) => (
          <div key={index}>
            <p>{item.title}</p>
            <p>{item.price.toFixed(2)}</p>
          </div>
        ))}
        <hr />
        <p>Total: {total}</p>
        <p>Quantity: {quantity}</p>
        </Modal.Body>
        <Modal.Footer>
        <Button onClick={() => clearCart()}>Clear</Button>
        <Button onClick={() => setShowCart(false)}>Close</Button>
        </Modal.Footer>
        </Modal>
        );
}

// display the data (JSX corner)
    return (
        <div>
            <Navbar1 cartCount={cartCount} onOpenCart = {onOpenCart}/>
            <CartModal/>
            <h1>Featured Product</h1>
            <div className='item-container'>
                {products.map((product) =>(
                    <div className='card' key={product.id}> 
                        <img src={product.image} alt='' />
                        <h3>{product.title}</h3>
                        <h4>{`$${product.price.toFixed(2)}`}</h4>
                        {/*<p>{product.description}</p>*/}
                        <Link to={`/products/${product.id}`}>View</Link>
                        <AddToCartButton onAddToCart={() => onAddToCart(product)} />
                    </div>
                ))}
            </div> {/* container div */}
        </div>
    )
}

export default FeaturedProduct

// test 

