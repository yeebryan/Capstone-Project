import React, { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import './MyOrder.css';
import Navbar1 from './Navbar';
import withAuth from './withAuth';

const MoreDetails = () => {
  const location = useLocation();
  const { order, foods  } = location.state || { order: {},  foods: [] };
  const [playlistName, setPlaylistName] = useState('');
  const [deliveryDates, setDeliveryDates] = useState([]);

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

  useEffect(() => {
    if (order.startDate && order.interval) {
      const dates = generateDeliveryDates(order.startDate, order.interval.interval);
      setDeliveryDates(dates);
      setPlaylistName(order.name);
    }
  }, [order]);
  
  const renderPlaylistItems = () => {
    console.log('order:', order);
    console.log('deliveryDates:', deliveryDates);
    return deliveryDates.map((date, index) => (
      <div key={index} className="playlist-item">
        <img src={foods[index].image.url} alt={foods[index].name} />
        <div className="playlist-item-details">
          <h4>Delivery {index + 1}</h4>
          <p>
            <strong>Date:</strong> {date.toLocaleDateString()}
          </p>
          <p>
            <strong>Time:</strong> {order.time}
          </p>
          <p>
            <strong>Food:</strong> {foods[index].name} - ${foods[index].price}
          </p>
        </div>
      </div>
    ));
  };
  
  

  return (
    <div>
      <Navbar1 showCartIcon={false} />
      <div className="my-order">
        {/* Keep the existing JSX code for displaying order details */}
        <div className="playlist-items-container">
          <h3>Delivery Playlist: {playlistName}</h3>
          <div className="playlist-items">{renderPlaylistItems()}</div>
        </div>
      </div>
    </div>
  );
};

export default withAuth(MoreDetails);
