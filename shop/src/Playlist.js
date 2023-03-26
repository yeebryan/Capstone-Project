import React, { useState, useEffect } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import {useParams, useNavigate } from "react-router-dom";
import axios from "axios";
import Navbar1 from "./Navbar";
import { Card } from 'react-bootstrap';
import "./App.css";



// PLAYLIST 
// functional component

const Playlist = () => {
  const [data, setData] = useState([]); // set initial state (data) to an empty array [] not object {}
  const { playlist_id } = useParams();
  const [cartCount, setCartCount] = useState(0);
  const [cart, setCart] = useState([]);
  const [total, setTotal] = useState(0);
  const [playlistName, setPlaylistName] = useState("");

  const navigate = useNavigate();

  const onProceedToCheckout = () => {
    // Redirect to checkout page
    navigate('/checkout');
  }


// axios 
  useEffect(() => {
    console.log(`Fetching restaurant data for ID ${playlist_id}...`);

    axios
      .get(`http://localhost:3000/playlists/${playlist_id}`) // change from ?id=${id} to ${id} because API url is .com/products/1  // But couldn't map due to not being array
      .then((res) => {
        console.log(JSON.stringify(res)) 
        setData(res.data);
        setPlaylistName(res.data.playlistName)
        console.log(`yo mr white: ${res.data.playlistName}`)

      })
      .catch((err) => console.log(err));
  }, [playlist_id]);

// cart count, and keep track of item added

const onAddToCart = (item) => {
  setCartCount(cartCount + 1);
  setCart([...cart, item])
  setTotal(total + item.price)
}

// open the cart

/*const onOpenCart = () => {
  setShowCart(true)
}

// close cart
const onCloseCart = () => {
  setShowCart(false);
}

// clear cart btn function - setState back to 0

const clearCart = () => {
  setCartCount(0); // the counter to 0 
  setCart([]);
  setShowCart(false);
}
*/

const PlaylistItems = data.length > 0 && (
  <div className="playlist-container" key={data.id}>
      {data &&
        data.map((item) => (
          <Card className="item-card" key={item.id}>
            <Card.Img variant="top" src={item.image.url} alt={item.image.url} />
            <div className="item-title">{item.name}</div>
            <div>{item.price}</div>
            <div>{item.description}</div>
              {console.log(`yo log this: ${item} & this ${item.name}`)}
          </Card>
        ))}
  </div>
);


  return (
    <div>
          {/* <Navbar1 cartCount={props.cartCount} onOpenCart = {props.onOpenCart}/> */}
      <h1>{playlistName}</h1>
      <div className="playlist_item">{PlaylistItems}</div>
      <button onClick={onProceedToCheckout}>Proceed to Checkout</button>
    </div>
  );
  }

export default Playlist;
