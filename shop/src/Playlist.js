import React, { useState, useEffect } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import {useParams, useNavigate } from "react-router-dom";
import axios from "axios";
import Navbar1 from "./Navbar";
import { Card } from 'react-bootstrap';
import authAxios from './authAxios';
import "./App.css";



// PLAYLIST 
// functional component

const Playlist = (props) => {
  const [data, setData] = useState([]); // set initial state (data) to an empty array [] not object {}
  const { playlist_id } = useParams();
  const [cartCount, setCartCount] = useState(0);
  const [cart, setCart] = useState([]);
  const [total, setTotal] = useState(0);
  const [category, setCategory] = useState('');
  const [foodType, setFoodType] = useState('');
  const [playlistName, setPlaylistName] = useState("");
  const [interval, setInterval] = useState('Weekly');
  const [startDate, setStartDate] = useState('');
  const [time, setTime] = useState('');
  const [loading, setLoading] = useState(false);
  const [foods, setFoods] = useState([]);
  const [order, setOrder] = useState({});
  const [error, setError] = useState('');
  const [submitted, setSubmitted] = useState(false); // track form submit
  const [showError, setShowError] = useState(false);

  const navigate = useNavigate();

  const resetState = () => {
    setError('');
    setFoods([]);
  }
  
  const handleSubmit = async (e) => {
    e.preventDefault();

    resetState();
    setShowError(false);

    if (!category && !foodType) {
      setError('Please select either category or food type');
      return;
    }
  
    if (!interval || !startDate || !time) {
      setError('Please fill in all required fields');
      return;
    }

    setLoading(true);
    setSubmitted(true); // Add this line


    const params = {
      category,
    };
    if (foodType) {
      params.foodType = foodType;
    }
      const now = new Date();
      const selectedDate = new Date(`${startDate} ${time}`);
      if (selectedDate < now) {
        setError('Start date and time must be in the future');
        setLoading(false);
        return;
      }
  

    const url = "http://localhost:3000/food/random?" + new URLSearchParams(params);

    try {
      const response = await authAxios.get(url);
      setFoods(response.data.foods);
      setOrder({
        category,
        foodType,
        interval,
        startDate,
        time,
        food: response.data.foods[0],
      });
      console.error(error.response.data);
      setShowError(true); // Add this line
    } catch (error) {
      console.error(error.response.data);
      setShowError(true); // Add this line
    } finally {
      setLoading(false);
    }
  };


// axios 
  useEffect(() => {
    console.log(`Fetching restaurant data for ID ${playlist_id}...`);

    axios
      .get(`http://localhost:3000/playlists/${playlist_id}`, {
        headers : {
          token : localStorage.getItem('token')
        }
      }) // change from ?id=${id} to ${id} because API url is .com/products/1  // But couldn't map due to not being array
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

const playlistNames = data.length > 0 && ( 
  <div key={data.id}>
    {data &&
        data.map(() => (
          <h1>{data.name}</h1>
        ))}
  </div>
);

const renderFoodCards = () => {

  if (foods.length === 0 && submitted && showError) {
    return <p>No available foods found.</p>;
  }

  return foods.map((food) => (
    <div key={food.id} className='food-card'>
      <img src={food.image.url} alt={food.name} />
      <h3>{food.name}</h3>
      <p>{food.description}</p>
    </div>
  ));
};

const handleProceed = () => {
  // Navigate to another page
  navigate('/myorder', { state: { order, foods }});
};


  return (
    <div>
          <Navbar1 cartCount={props.cartCount} onOpenCart = {props.onOpenCart}/>
      
      <h1>{playlistNames}</h1>
      <table className="premade_playlist_details">
      <tr className="premade_playlist_details_tr">
      <div className="playlist_item">{PlaylistItems}</div>
      </tr>
      <tr className="premade_playlist_details_tr">
      <div className='user-preference-form'>
      <div className='uform-page'>
                {/* Add a container for the title and progress indicator */}
                <div className='title-progress-container'>
      </div>
        <div className='uform-container uform-playlist'>
          <form onSubmit={handleSubmit}>
            <div>
              <label>
                Interval:
                <select
                  value={interval}
                  onChange={(e) => setInterval(e.target.value)}
                >
                  <option value='Weekly'>Weekly</option>
                  <option value='Bi-weekly'>Bi-weekly</option>
                  <option value='Monthly'>Monthly</option>
                </select>
              </label>
            </div>
            <div>
              <label>
                Start Date:
                <input
                  type='date'
                  value={startDate}
                  onChange={(e) => setStartDate(e.target.value)}
                  required
                />
              </label>
            </div>
            <div>
              <label>
                Time:
                <input type='time'
                  className='form-input-time time-input time-playlist'
                  value={time}
                  onChange={(e) => setTime(e.target.value)}
                  required
                />
              </label>
            </div>
          </form>
          <div className='food-cards-container'>{renderFoodCards()}</div>
          {/* <MyOrder order={order} /> */}
        </div>
        <div className='button-container'>
        <button className="proceed-button" onClick={handleProceed} disabled={error || !submitted}>Proceed</button>
        </div>
        {error && <p>{error}</p>}
      </div>
    </div>
    </tr>
    </table>
    </div>
  );
  }

export default Playlist;
