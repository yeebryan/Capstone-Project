import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Navbar from './Navbar';
import authAxios from './authAxios';
import './UserPlaylist.css';


const UserOrders = () => {
  const [orders, setOrders] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    fetchOrders();
  }, []);

  const fetchOrders = async () => {
    try {
      const storedUser = localStorage.getItem('user');
      if (!storedUser) {
        console.log('User not logged in');
        return;
      }

      const user = JSON.parse(storedUser);
      const userId = user.id;
      console.log('userId:', userId);
      const response = await authAxios.get(`http://localhost:3000/order/me`);
      console.log('Response:', response);
      setOrders(response.data);
      console.log(response.data)
    } catch (error) {
      console.error('Error fetching orders:', error);
    }
  };

  const renderOrders = () => {
    return (
      <table>
        <thead>
          <tr>
            <th>Order ID</th>
            <th>Start Date</th>
            <th>Delivery Time</th>
            <th>Status</th>
            <th>Items</th>
          </tr>
        </thead>
        <tbody>
          {orders.map((order) => (
            <tr key={order._id}>
              <td data-label="Order ID">{order._id}</td>
              <td data-label="Start Date">{order.start_date}</td>
              <td data-label="Delivery Time">{order.delivery_time}</td>
              <td data-label="Status">{order.status}</td>
              <td data-label="Items">
                <ul>
                  {order.Items.map((item) => (
                    <li key={item.food_id}>
                      {item.name} - {item.quantity} x ${item.price}
                    </li>
                  ))}
                </ul>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    );
  };
  

  return (
    <div>
      <Navbar />
      <div className="user-orders">
        <h2>Your Orders</h2>
        <div className="orders-container">{renderOrders()}</div>
      </div>
    </div>
  );
  
};

export default UserOrders;
