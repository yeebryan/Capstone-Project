import React, {useState} from 'react';
import { useLocation } from 'react-router-dom';
import Navbar1 from './Navbar';
import './Checkout.css';
import withAuth from './withAuth';
import authAxios from './authAxios';


const Checkout = () => {
  const location = useLocation();
  const { order, deliveryDates, playlist_id } = location.state || { order: { food: undefined }, deliveryDates: [], playlist_id: null };
  const [showPaymentSuccess, setShowPaymentSuccess] = useState(false);
  const [showConfirmCancel, setShowConfirmCancel] = useState(false);

  console.log('location.state:', location.state);

  console.log('order:', order);
  console.log('deliveryDates:', deliveryDates);
  console.log('playlist_id:', playlist_id);
  if (!order || !order.food) {


    return (
      <div className="checkout">
        <h2>No order details available</h2>
      </div>
    );
  }
  
  const handlePay = () => {
    setShowPaymentSuccess(true);
    setTimeout(() => {
      setShowPaymentSuccess(false);
    }, 3000);
  };

  const handleCancel = () => {
    setShowConfirmCancel(true);
  };

  const handleConfirmCancel = async () => {
    try {
      console.log(`Deleting playlist with ID: ${playlist_id}`);
      const response = await authAxios.delete(`http://localhost:3000/playlists/delete/${playlist_id}`);
  
      if (response.status !== 200) {
        throw new Error('Failed to delete playlist');
      }
  
      // Close the confirmation modal
      setShowConfirmCancel(false);
    } catch (error) {
      console.error('Error deleting playlist:', error);
    }
  };

  const handleCloseCancel = () => {
    setShowConfirmCancel(false);
  };
  
  console.log('order:', order);
console.log('deliveryDates:', deliveryDates);


  const renderOrderItems = () => {
    if (!order || !order.food || !deliveryDates || deliveryDates.length === 0) {
      return <h2>No order details available</h2>;
      
    }
  
    return deliveryDates.map((date, index) => (
      <div key={index} className="order-item">
        <img src={order.food.image.url} alt={order.food.name} />
        <div className="order-item-details">
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
  

  return (
    <div>
      <Navbar1 showCartIcon={false} />
      <div className="checkout">
        <h2>My Order</h2>
        <div className="order-items-container">{renderOrderItems()}</div>
        {/* Add your payment method component here */}
        <div className="checkout-buttons">
          <button onClick={handlePay}>Pay</button>
          <button onClick={handleCancel}>Cancel</button>
        </div>
        {showPaymentSuccess && (
          <div className="payment-success-modal">
            <h2>Payment Success!</h2>
          </div>
        )}
        {showConfirmCancel && (
          <div className="confirm-cancel-modal">
            <h2>Confirm</h2>
            <p>Are you sure you want to cancel your order?</p>
            <div className="confirm-cancel-buttons">
              <button onClick={handleConfirmCancel}>Yes</button>
              <button onClick={handleCloseCancel}>No</button>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default withAuth(Checkout);






