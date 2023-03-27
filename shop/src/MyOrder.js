import React, {useState} from 'react';
import axios from 'axios'
import { useLocation, useNavigate } from 'react-router-dom';
import './MyOrder.css';
import Navbar1 from './Navbar';
import authAxios from './authAxios';
import withAuth from './withAuth';

const MyOrder = () => {
  const location = useLocation();
  const { order, foods } = location.state || { order: {}, foods: [] };
  const [playlistName, setPlaylistName] = useState('');

  const navigate = useNavigate();


  if (!order.food) {
    return (
      <div className="my-order">
        <h2>No order details available</h2>
      </div>
    );
  }
  
  const generateDeliveryDates = (startDate, interval, lifetimeInMonths = 1) => {
    const start = new Date(startDate);
    const endDate = new Date(start);
    endDate.setMonth(endDate.getMonth() + lifetimeInMonths);
    const deliveryDates = [];
  
    while (start < endDate) {
      deliveryDates.push(new Date(start));
      if (interval === 'Weekly') {
        start.setDate(start.getDate() + 7);
      } else if (interval === 'Bi-weekly') {
        start.setDate(start.getDate() + 14);
      } else {
        start.setMonth(start.getMonth() + 1);
      }
    }
  
    return deliveryDates;
  };

  const deliveryDates = generateDeliveryDates(order.startDate, order.interval);

  const renderPlaylistItems = () => {
    return deliveryDates.map((date, index) => (
      <div key={index} className="playlist-item">
        <img src={order.food.image.url} alt={order.food.name} />
        <div className="playlist-item-details">
          <h4>Delivery {index + 1}</h4>
          <p>
            <strong>Date:</strong> {date.toLocaleDateString()}
          </p>
          <p>
            <strong>Time:</strong> {order.time}
          </p>
        </div>
      </div>
    ));
  };
  
  const handleNameChange = (event) => {
    setPlaylistName(event.target.value);
  };

    // saved to mongodb 
    const savePlaylist = async () => {
        try {
            const user = JSON.parse(localStorage.getItem('user'));
            const userId = user.id;
            console.log('user ID:', userId);
            console.log('playlist Name:', playlistName);
            console.log('order:', order);
            console.log('food ID:', order.food._id);
            console.log('food ID 2:', order.foods);
            console.log('food ID 3:', order.foodId)
            console.log('food ID 3:', order.food.id)
            const response = await authAxios.post('http://localhost:3000/food/random/create', {            
            playlistName: playlistName,
            category: order.category,
            foodType: order.foodType,
            interval: order.interval,
            startDate: order.startDate,
            time: order.time,
            foodId: order.food.id,
          });
          if (response.status === 200) {
            console.log('Playlist saved successfully');
          } else {
            console.log('Error saving playlist');
          }
        } catch (error) {
          console.log('Error saving playlist:', error);
        }
      };
      
  const handleProceed = () => {
    savePlaylist();
    // Navigate to another page - checkout
    navigate('/checkout', { state: { order, foods }});
    console.log('you have saved!')
  };



  return (
    <div>
    <Navbar1 showCartIcon={false}/>
    <div className="my-order">
      <h2>Your Order Summary</h2>
      <div className="summary-card">
        <div className="food-details">
          <img src={order.food.image.url} alt={order.food.name} />
          <h4>{order.food.name}</h4>
        </div>
        <div className="order-details">
            {order.category && (
          <p>
            <strong>Category:</strong> {order.category}
          </p>
            )}
            {order.foodType && (
          <p>
            <strong>Food Type:</strong> {order.foodType}
          </p>
            )}
          <p>
            <strong>Interval:</strong> {order.interval}
          </p>
          <p>
            <strong>Start Date:</strong> {order.startDate}
          </p>
          <p>
            <strong>Time:</strong> {order.time}
          </p>
        </div>
      </div>
              {/* Add the playlist items */}
              <div className="playlist-items-container">
          <h3>Delivery Playlist: {playlistName}</h3>
          <input
            type="text"
            value={playlistName}
            onChange={handleNameChange}
            placeholder="Enter playlist name"
            className="playlist-name"
          />
        <div className="playlist-items">{renderPlaylistItems()}</div>
        </div>
    </div>
    <div className='button-container'>
        <button className="proceed-button" onClick={handleProceed}>Proceed</button>
    </div>
</div>
  );
};

export default withAuth(MyOrder);
