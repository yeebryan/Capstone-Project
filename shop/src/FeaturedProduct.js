import React, { useState, useEffect } from 'react';
import "bootstrap/dist/css/bootstrap.min.css"
import {Link} from 'react-router-dom'
//import AddToCartButton from './AddtoCart';
import {Modal, Col, Row, Card, Button} from "react-bootstrap"
import Navbar1 from './Navbar';
import axios from 'axios';
import Carousel from "react-multi-carousel";
import "react-multi-carousel/lib/styles.css";

import "./App.css";

// this .js is functional hence we can use HOOK








// arrow function 
const FeaturedProduct = (props) => {
    
    const [products, setProducts] = useState([]); // `products` is the current state // `setProducts` function that updates the state 
    const [cartCount, setCartCount] = useState(0);
    const [showCart, setShowCart] = useState(false);
    const [cart, setCart] = useState([]);
    const [restaurants, setRestaurants] = useState([]);
    const [selectedCat, setSelectedCat] = useState(null);


// rwd
const settings = 
    {
    desktop: {breakpoint: { max: 3000, min: 600}, items: 5, slidesToSlide: 5},
    tablet: {breakpoint: {max: 600, min: 480}, items: 5, slidesToSlide: 5},
    mobile: {breakpoint: {max: 480, min: 0}, items: 2, slidesToSlide: 2}
    }


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

// fetch data from API (restaurants)
const fetchProducts = () => {
    axios
        .get('http://localhost:3000/restaurants')
        .then((res) =>{
            console.log(res); // response // response.data
            setProducts(res.data) // for storerestapi, take note it must be res.data.data because res.data is an object not an array // passing data into our setProducts function so that we can set our state to the data
        })
        .catch((err) => {
            console.log(err);
        })
}


// click category
const handleClickCat = async (category) => {
    setSelectedCat(category);
    const response = await fetch(`http://localhost:3000/restaurants?category=${category}`);
    const data = await response.json();
    setRestaurants(data);
  }


// cart count, and keep track of item added
const onAddToCart = (item) => {
    setCartCount(cartCount + 1);
    const userCart = JSON.parse(localStorage.getItem(props.userId)) || [];
    userCart.push(item);
    localStorage.setItem(props.userId, JSON.stringify(userCart));
  };
  

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
// group items by ID, storing their quantity separately
// group items by ID, storing their quantity separately
const CartModal = () => {
    let total = 0;
    let quantity = 0;

    const userCart = JSON.parse(localStorage.getItem(props.userId)) || [];

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
          {userCart.map((item, index) => (
            <div key={index}>
              <p>{item.title}</p>
            </div>
          ))}
          <hr />
          <p>Total: {total}</p>
          <p>Quantity: {quantity}</p>
        </Modal.Body>
        <Modal.Footer>
          <Button onClick={() => clearCart()}>Clear</Button>
          <Button onClick={() => setShowCart(false)}>Close</Button>
          <Link to="/checkout">Checkout</Link>
        </Modal.Footer>
      </Modal>
    );
  };
  
// carousel for pre-made playlists
const CarouselPlaylist = () => {

    return (
      <Carousel 
        className="first-carousel"
        responsive={settings}
        swipeable={true}
        draggable={true}
        showDots={false}
        ssr={true} // means rendering carousel on server-side
        infinite={true}
        autoPlay={true}
        autoPlaySpeed={10000}
        keyBoardControl={true}
        customTransition="all .5"
        transitionDuration={500}
        containerClass="carousel-container"
        //removeArrowOnDeviceType={["tablet", "mobile"]}
        dotListClass="custom-dot-list-style">
        {products.map((product) => (
          <div className='card-wrapper' key={product.id}>
            <Link to={`/products/${product._id}`}>
              <img src={product.image.url} alt={product.image.url} />
            </Link>
            <div className='card-body' style={{ textAlign: 'center' }}>
              <h3>{product.name}</h3>
              <p>{product.description}</p>
            </div>
          </div>
      ))}
      </Carousel>
)}






// 2nd carousel for featured restaurants
const CarouselRestaurants = () => {

  return (
    <Carousel 
      className="my-carousel"
      responsive={settings}
      swipeable={true}
      draggable={true}
      showDots={false}
      ssr={true} // means rendering carousel on server-side
      infinite={true}
      autoPlay={true}
      autoPlaySpeed={10000}
      keyBoardControl={true}
      customTransition="all .5"
      transitionDuration={500}
      containerClass="carousel-container"
      //removeArrowOnDeviceType={["tablet", "mobile"]}
      dotListClass="custom-dot-list-style">
      {products.map((product) => (
        <div className='card-wrapper' key={product.id}>
          <Link to={`/products/${product._id}`}>
            <img src={product.image.url} alt={product.image.url} />
          </Link>
          <div className='card-body' style={{ textAlign: 'center' }}>
            <h3>{product.name}</h3>
            <p>{product.description}</p>
          </div>
        </div>
    ))}
    </Carousel>
)}


// carousel for category
const CarouselCategory = () => {

  return (
    <Carousel 
      className="my-carousel"
      responsive={settings}
      swipeable={true}
      draggable={true}
      showDots={false}
      ssr={true} // means rendering carousel on server-side
      infinite={true}
      autoPlay={true}
      autoPlaySpeed={10000}
      keyBoardControl={true}
      customTransition="all .5"
      transitionDuration={500}
      containerClass="carousel-container"
      //removeArrowOnDeviceType={["tablet", "mobile"]}
      dotListClass="custom-dot-list-style">
      {products.map((product) => (
        <div className='card-wrapper' key={product.id}>
          <Link to={`/products/${product._id}`}>
            <img src={product.image.url} alt={product.image.url} />
          </Link>
          <div className='card-body' style={{ textAlign: 'center' }}>
            <h3>{product.name}</h3>
            <p>{product.description}</p>
          </div>
        </div>
    ))}
    </Carousel>
)}


// footer
const Footer = () => {
  return (
    <footer className="bg-primary-500 text-white p-4">
      <div className="container mx-auto">
        <div className="flex flex-wrap items-center justify-between">
          <p className="text-sm">© 2023 Your Company. All rights reserved.</p>
          <ul className="flex space-x-4">
            <li>
              <a href="/" className="text-black hover:text-primary-100">
                Terms of Service
              </a>
            </li>
            <li>
              <a href="/" className="text-black hover:text-primary-100">
                Privacy Policy
              </a>
            </li>
            <li>
              <a href="/" className="text-black hover:text-primary-100">
                Contact Us
              </a>
            </li>
          </ul>
        </div>
      </div>
    </footer>
  );
};


// CARD MENU home page - PENDING
const ThreeColumnCard = ({ handleClickCat }) => {
    return (
      <Card border="light" className='border-0'>
        <Card.Body>
          <Row className="row-three">
            <Col className='column-three'>
              <div className="threeCards" onClick={() => handleClickCat("Western")}>Western</div>
            </Col>
            <Col className='row-three'>
              <div className="threeCards" onClick={() => handleClickCat("Chinese")}>Chinese</div>
            </Col>
            <Col className='row-three'>
              <div className="threeCards" onClick={() => handleClickCat("Japanese")}>Japanese</div>
            </Col>
            <Col className='row-three'>
              <div className="threeCards" onClick={() => handleClickCat("Indian")}>Indian</div>
            </Col>
            <Col className='row-three'>
              <div className="threeCards" onClick={() => handleClickCat("Mexican")}>Mexican</div>
            </Col>
            <Col className='row-three'>
              <div className="threeCards" onClick={() => handleClickCat("Vegetarian")}>Vegetarian</div>
            </Col>
          </Row>
        </Card.Body>
      </Card>
    );
  };
  





// display the data (JSX corner)
    return (
        <div className='page'>
          <Navbar1 cartCount={cartCount} onOpenCart = {onOpenCart}/>
          <CartModal/>
          <div className='main-body'>
            <h1 className="title-mainpage" style={{ textAlign: 'center' }}>Featured Playlists</h1>
            <h2 className="title2-mainpage" style={{ textAlign: 'center' }}>DIY your own playlists or choose popular pre-made playlists!</h2>
          <CarouselPlaylist />
          <CarouselRestaurants />
          <ThreeColumnCard handleClickCat={handleClickCat}/>
          {selectedCat && (CarouselCategory)}
        </div>
          <Footer/>
        </div>
          )
}

export default FeaturedProduct

// test 

