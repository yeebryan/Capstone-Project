import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Navbar from './Navbar';
import authAxios from './authAxios';

const UserPlaylist = () => {
  const [playlists, setPlaylists] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    fetchPlaylists();
  }, []);

  const fetchPlaylists = async () => {
    try {
      const response = await authAxios.get('http://localhost:3000/playlists/user');
      setPlaylists(response.data);
      console.log(response.data)
    } catch (error) {
      console.error('Error fetching playlists:', error);
    }
  };

  const updatePlaylistStatus = async (playlist_id, newStatus) => {
    try {
      await authAxios.put(`http://localhost:3000/playlists/update/${playlist_id}`, { status: newStatus });
      fetchPlaylists();
    } catch (error) {
      console.error('Error updating playlist status:', error);
    }
  };

  const renderPlaylists = () => {
    return playlists.map((playlist) => (
      <div key={playlist._id} className="playlist-item">
        <h3>{playlist.name}</h3>
        <p>Status: {playlist.status}</p>
        {/* <button onClick={() => updatePlaylistStatus(playlist._id, 'paused')}>Pause</button>
        <button onClick={() => updatePlaylistStatus(playlist._id, 'stopped')}>Stop</button>
        <button onClick={() => updatePlaylistStatus(playlist._id, 'pending')}>Pending</button>
        <button onClick={() => updatePlaylistStatus(playlist._id, 'delivered')}>Delivered</button>
        <button onClick={() => updatePlaylistStatus(playlist._id, 'completed')}>Completed</button> */}
      </div>
    ));
  };

  return (
    <div>
      <Navbar />
      <div className="user-playlists">
        <h2>Your Playlists & Order History</h2>
        <div className="playlists-container">{renderPlaylists()}</div>
      </div>
    </div>
  );
};

export default UserPlaylist;
