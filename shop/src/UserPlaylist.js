import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Navbar from './Navbar';
import authAxios from './authAxios';
import './UserPlaylist.css';


const UserPlaylist = () => {
  const [orders, setOrders] = useState([]);
  const [playlists, setPlaylists] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    fetchOrders();
    fetchPlaylists();
  }, []);

  const toggleOrdersVisibility = () => {
    const ordersContainer = document.querySelector('.orders-container');
    ordersContainer.classList.toggle('hidden');
  };

  const viewPlaylist = (playlist) => {
    navigate('/moredetails', { state: { order: playlist } });
};

const pauseUnpausePlaylist = async (playlistId, currentStatus) => {
    try {
      const updatedStatus = currentStatus === 'ongoing' ? 'paused' : 'ongoing';
      const response = await authAxios.put(`http://localhost:3000/playlist/${playlistId}`, { status: updatedStatus });
      console.log('Response:', response);
      // Update the playlists state with the updated playlist
      const updatedPlaylists = playlists.map(playlist => {
        if (playlist._id === playlistId) {
          return { ...playlist, status: updatedStatus };
        }
        return playlist;
      });
      setPlaylists(updatedPlaylists);
    } catch (error) {
      console.error('Error updating playlist status:', error);
    }
  };
  


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

  const fetchPlaylists = async () => {
    try {
      const storedUser = localStorage.getItem('user');
            if (!storedUser) {
        console.log('User not logged in');
        return;
      }

      const user = JSON.parse(storedUser);
      const userId = user.id;
      console.log('userId:', userId);
      const response = await authAxios.get(`http://localhost:3000/playlist/me`);
      console.log('Response:', response);
      setPlaylists(response.data);
      console.log(response.data);
    } catch (error) {
      console.error('Error fetching playlists:', error);
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

  const renderPlaylists = () => {
    return (
      <table className="user-playlists">
        <thead>
          <tr>
            <th>Playlist name</th>
            <th>Status</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {playlists.map((playlist) => (
            <tr key={playlist._id}>
              <td data-label="Name">{playlist.name}</td>
              <td data-label="Status">{playlist.status}</td>
              <td data-label="Actions">
                <button className="btn-secondary" onClick={() => pauseUnpausePlaylist(playlist._id, playlist.status)}>
                  {playlist.status === 'paused' ? 'Unpause' : 'Pause'}
                </button>
                <button className="btn-primary" onClick={() => viewPlaylist(playlist)}>View</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    );
  };
  

  return (
    <div>
      <Navbar showCartIcon={false}/>
      <div className="user-playlists">
        <h2>Your Playlists History</h2>
        <div className="playlists-container">{renderPlaylists()}</div>
      </div>
      <div className="user-orders">
        <h2>Your Orders History</h2>
        <button onClick={toggleOrdersVisibility}>V</button>
        <div className="orders-container hidden">{renderOrders()}</div>
      </div>
    </div>
  );
};




export default UserPlaylist;