import React from 'react';

export default function About() {

    return(
        <div>
            <h1>Capstone Project Mar 2023</h1>
            <h2>Bryan Yee</h2>
            <h2>Chekwee Quek</h2>
            <h2>Daniel Lim</h2>
        </div>
    )


}

// checkout 

// import React from 'react';
// import { useLocation } from 'react-router-dom';
// import Navbar1 from './Navbar';
// import './Checkout.css';
// import withAuth from './withAuth';

// const Checkout = () => {
//   const location = useLocation();
//   const { order, deliveryDates } = location.state || { order: { food: undefined }, deliveryDates: [] };
//   if (!order.food) {
//     return (
//       <div className="checkout">
//         <h2>No order details available</h2>
//       </div>
//     );
//   }

//   const renderOrderItems = () => {
//     return deliveryDates.map((date, index) => (
//       <div key={index} className="order-item">
//         <img src={order.food.image.url} alt={order.food.name} />
//         <div className="order-item-details">
//           <p>
//             <strong>Date:</strong> {date.toLocaleDateString()}
//           </p>
//           <p>
//             <strong>Time:</strong> {order.time}
//           </p>
//           <p>
//             <strong>Status:</strong> {order.status}
//           </p>
//         </div>
//       </div>
//     ));
//   };

//   return (
//     <div>
//       <Navbar1 showCartIcon={false} />
//       <div className="checkout">
//         <h2>My Order</h2>
//         <div className="order-items-container">{renderOrderItems()}</div>
//       </div>
//     </div>
//   );
// };

// export default withAuth(Checkout);
