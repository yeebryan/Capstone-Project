import React from 'react';

// cart counter

const Cart = (props) => {
  return (
    <div>
      <span>{props.cartCount}</span>
      <button onClick={props.onOpenCart}>View Cart</button>
    </div>
  );
};

export default Cart;
