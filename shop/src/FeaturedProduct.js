import React, { useState, useEffect } from 'react';
import "bootstrap/dist/css/bootstrap.min.css"
import {Link} from 'react-router-dom'
//import AddToCartButton from './AddtoCart';
import {Modal, Card, Button} from "react-bootstrap"
import Navbar1 from './Navbar';
import axios from 'axios';
import Playlist from './Playlist';
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
    const [FPplaylists, setFPplaylist] = useState([]);


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


// for Foodpanda Playlists
useEffect(() =>{
  fetchPlaylist();
}, []);

// fetch data from API (restaurants)
const fetchPlaylist = () => {
  axios
      .get('http://localhost:3000/playlists')
      .then((res) =>{
          console.log(res); // response // response.data
          setFPplaylist(res.data) // for storerestapi, take note it must be res.data.data because res.data is an object not an array // passing data into our setProducts function so that we can set our state to the data
      })
      .catch((err) => {
          console.log(err);
      })
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

const PlaylistTile = () => {

  const handleClick = () => {
    window.location.href = '/userform';
  };

  return (
    <div 
      className="tile-main1"
      onClick={handleClick}
    >
      <img 
        src="/imgs/foodpanda-panda.gif"
        alt="DIY playlist" 
        className="tile-image"
      />
      <span className="tile-text-1">Introducing ...</span>
      <span className="tile-text">DIY PLAYLIST</span>
      <span className='tile-para'>Click on me to begin!</span>
    </div>
  );
};


const CarouselPlaylist = () => {

    return (
      <Carousel 
        className="first-carousel"
        responsive={settings}
        swipeable={true}
        draggable={true}
        showDots={false}
        ssr={true} // means rendering carousel on server-side
        infinite={false}
        autoPlay={true}
        autoPlaySpeed={10000}
        keyBoardControl={true}
        customTransition="all .5"
        transitionDuration={500}
        containerClass="carousel-container"
        //removeArrowOnDeviceType={["tablet", "mobile"]}
        dotListClass="custom-dot-list-style">
        {FPplaylists.map((FPplaylist) => (
          <div className='card-wrapper' key={FPplaylist._id}>
            <Link to={`/playlists/${FPplaylist._id}`}>
              <img src={FPplaylist.image.url} alt={FPplaylist.image.url} />
            </Link>
            <div className='card-body' style={{ textAlign: 'center' }}>
              <h3 className='h3-playlist'>{FPplaylist.name}</h3>
              <p>{FPplaylist.description}</p>
            </div>
          </div>
      ))}
      </Carousel>
)}

// 2nd carousel for featured restaurants
const CarouselRestaurants = () => {


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
        <div className='card-wrapper' key={product._id}>
          <Link to={`/restaurants/${product._id}`}>
            <img src={product.image.url} alt={product.image.url} />
          </Link>
          <div className='card-body' style={{ textAlign: 'center' }}>
            <h3 className='h3-playlist'>{product.name}</h3>
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
      {restaurants.map((restaurant) => (
        <div className='card-wrapper' key={restaurant._id}>
          <Link to={`/restaurants/${restaurant._id}`}>
            <img src={restaurant.image.url} alt={restaurant.image.url} />
          </Link>
          <div className='card-body' style={{ textAlign: 'center' }}>
            <h3 className='h3-playlist'>{restaurant.name}</h3>
            <p>{restaurant.description}</p>
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
              <input className="threeCards "type="button" value="Western"  id="btn" onClick={() => handleClickCat("Western")} />
              <input className="threeCards "type="button" value="Chinese"  id="btn" onClick={() => handleClickCat("Chinese")} />
              <input className="threeCards "type="button" value="Japanese"  id="btn" onClick={() => handleClickCat("Japanese")} />
              <input className="threeCards "type="button" value="Indian"  id="btn" onClick={() => handleClickCat("Indian")} />
              <input className="threeCards "type="button" value="Mexican"  id="btn" onClick={() => handleClickCat("Mexican")} />
              <input className="threeCards "type="button" value="Vegetarian"  id="btn" onClick={() => handleClickCat("Vegetarian")} />
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
          <PlaylistTile/>
          <h3>Foodpanda Playlists</h3>
          <CarouselPlaylist />
          <h3>Featured Restaurants</h3>
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

