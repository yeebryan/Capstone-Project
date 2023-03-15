import React, { useState } from 'react';

const Checkout = (props) => {
  const [interval, setInterval] = useState('');
  const [deliveryTime, setDeliveryTime] = useState('');
  const [paymentMethod, setPaymentMethod] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    const userCart = JSON.parse(localStorage.getItem(props.userId)) || [];
    const order = {
      cart: userCart,
      interval: interval,
      deliveryTime: deliveryTime,
      paymentMethod: paymentMethod,
    };
    // send order to backend to process payment and complete order
    console.log(order);
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>
        Interval:
        <input
          type="text"
          value={interval}
          onChange={(e) => setInterval(e.target.value)}
        />
      </label>
      <label>
        Delivery Time:
        <input
          type="text"
          value={deliveryTime}
          onChange={(e) => setDeliveryTime(e.target.value)}
        />
      </label>
      <label>
        Payment Method:
        <select
          value={paymentMethod}
          onChange={(e) => setPaymentMethod(e.target.value)}
        >
          <option value="credit-card">Credit Card</option>
          <option value="paypal">Paypal</option>
          <option value="bitcoin">Bitcoin</option>
        </select>
      </label>
      <button type="submit">Complete Order</button>
    </form>
  );
};

export default Checkout;
