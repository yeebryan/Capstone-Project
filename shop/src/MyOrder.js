import React from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import './MyOrder.css';
import Navbar1 from './Navbar';

const MyOrder = () => {
  const location = useLocation();
  const { order, foods } = location.state || { order: {}, foods: [] };
  const navigate = useNavigate();


  if (!order.food) {
    return (
      <div className="my-order">
        <h2>No order details available</h2>
      </div>
    );
  }
  
  const handleProceed = () => {
    // Navigate to another page
    navigate('/checkout', { state: { order, foods }});
  };

  return (
    <div>
    <Navbar1 />
    <div className="my-order">
      <h2>Your Order Summary</h2>
      <div className="summary-card">
        <h3>Food</h3>
        <div className="food-details">
          <img src={order.food.image.url} alt={order.food.name} />
          <h4>{order.food.name}</h4>
        </div>
        <div className="order-details">
          <p>
            <strong>Category:</strong> {order.category}
          </p>
          <p>
            <strong>Food Type:</strong> {order.foodType}
          </p>
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
    </div>
    <div className='button-container'>
        <button className="proceed-button" onClick={handleProceed}>Proceed</button>
    </div>
</div>
  );
};

export default MyOrder;
