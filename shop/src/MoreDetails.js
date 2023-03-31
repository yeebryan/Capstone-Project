import React, { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import Navbar1 from './Navbar';
import authAxios from './authAxios';
import withAuth from './withAuth';

const MoreDetails = () => {
  const location = useLocation();
  const { playlist_id } = location.state || { playlist_id: null };
  const [playlistDetails, setPlaylistDetails] = useState(null);
  const [playlistData, setPlaylistData] = useState(null);


  useEffect(() => {
    const fetchPlaylistDetails = async () => {
      if (playlist_id) {
        try {
          const user = JSON.parse(localStorage.getItem('user'));
          const userId = user.id;
  
          const response = await authAxios.get(`http://localhost:3000/playlist/${playlist_id}/user/${userId}`);
          setPlaylistDetails(response.data); // Update the playlistDetails state with the fetched data
  
        } catch (error) {
          console.log('Error fetching playlist details:', error);
        }
      }
    };
    fetchPlaylistDetails();
  }, [playlist_id]);

  const Footer = () => {
    return (
      <footer className="bg-primary-500 text-white p-4">
        <div className="container mx-auto">
          <div className="flex flex-wrap items-center justify-between">
            <table className="footer-table">
              <tr className="footer-stuff">
                <td className="footer-td">
                <p className="text-sm">© 2023 Your Company. All rights reserved.</p>
                </td>
                <td className="footer-tdd">
                  <a href="/" className="text-black hover:text-primary-100">
                    Terms of Service
                  </a> 
                </td>
                <td className="footer-tdd">
                  <a href="/" className="text-black hover:text-primary-100">
                    Privacy Policy
                  </a>
                </td>
                <td className="footer-tdd">
                  <a href="/" className="text-black hover:text-primary-100">
                    Contact Us
                  </a>
                </td>
              </tr>
            </table>
          </div>
        </div>
      </footer>
    );
  };
  

  const renderPlaylistDetails = () => {
    if (!playlistDetails) {
        return <div>Loading...</div>;
      }
      
      console.log('Playlist Details:', playlistDetails); // Add this line to debug

      // Use playlistData to render the component
      

    // Use playlistDetails to render the component
    return (
        <>
          <h2>{playlistDetails.playlistName}</h2>
          <div className="playlist-details">
            {playlistDetails.deliveries.map((delivery, index) => (
              <div key={index} className="delivery-item">
                <img src={delivery.food.image.url} alt={delivery.food.name} />
                <div className="delivery-item-details">
                  <h4>Delivery {index + 1}</h4>
                  <p>
                    <strong>Date:</strong> {new Date(delivery.date).toLocaleDateString()}
                  </p>
                  <p>
                    <strong>Time:</strong> {delivery.time}
                  </p>
                </div>
              </div>
            ))}
          </div>
        </>
      );
    };

  return (
    <div>
      <Navbar1 showCartIcon={false} />
      <div className="more-details">{renderPlaylistDetails()}</div>
      <Footer />
    </div>
  );
};

export default withAuth(MoreDetails);
