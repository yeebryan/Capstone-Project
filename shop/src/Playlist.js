import React, { useState, useEffect } from 'react';
import axios from 'axios';
import {Link} from 'react-router-dom'
import Carousel from "react-multi-carousel";
import "./App.css";

// get playlist


const Playlist = () => {

    const [products, setProducts] = useState([]); // `products` is the current state // `setProducts` function that updates the state 

    const settings = 
    {
    desktop: {breakpoint: { max: 3000, min: 600}, items: 5, slidesToSlide: 5},
    tablet: {breakpoint: {max: 600, min: 480}, items: 5, slidesToSlide: 5},
    mobile: {breakpoint: {max: 480, min: 0}, items: 2, slidesToSlide: 2}
    }

    useEffect(() =>{
        fetchProducts();
    }, []);
    
    // fetch data from API (restaurants)
    const fetchProducts = () => {
        axios
            .get('http://localhost:3000/playlists')
            .then((res) =>{
                console.log(res); // response // response.data
                setProducts(res.data) // for storerestapi, take note it must be res.data.data because res.data is an object not an array // passing data into our setProducts function so that we can set our state to the data
            })
            .catch((err) => {
                console.log(err);
            })
    }

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
        {products.map((product) => (
          <div className='card-wrapper' key={product._id}>
            <Link to={`/playlists/${product._id}`}>
              <img src={product.image.url} alt={product.image.url} />
            </Link>
            <div className='card-body' style={{ textAlign: 'center' }}>
              <h3>{product.name}</h3>
            </div>
          </div>
      ))}
      </Carousel>
)}

export default Playlist;


