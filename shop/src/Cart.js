import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faShoppingCart } from '@fortawesome/free-solid-svg-icons'

// cart counter
// use in NavBar

const Cart = (props) => {
  return (
    <div>
      <button onClick={props.onOpenCart} className="cart-icon">
        <FontAwesomeIcon icon={faShoppingCart} />
      <span className='cart-count'>{props.cartCount}</span>
      </button>
    </div>
  );
};

export default Cart;
