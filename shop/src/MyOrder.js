// this will be where user can see their playlist
// so can change to MyPlaylist
// !order ensure that order is defined 
import React from 'react';

const MyOrder = ({ order }) => {
  if (!order || !order.food) {
    return <div>No order yet!</div>;
  }

  return (
    <div>
      <h2>My Order</h2>
      <p>Category: {order.category}</p>
      <p>Food Type: {order.foodType}</p>
      <p>Interval: {order.interval}</p>
      <p>Start Date: {order.startDate}</p>
      <p>Time: {order.time}</p>
      {order.food && (
        <>
          <p><img src={order.food.imageUrl} alt={order.food.name} /></p>
          <p>Selected Food: {order.food.name}</p>
        </>
      )}
    </div>
  );
};

export default MyOrder;


