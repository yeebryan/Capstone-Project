import React from 'react';

const AddToCartButton = (props) => {
  return (
    <button onClick={props.onAddToCart}>Add to Cart</button>
  );
};

export default AddToCartButton;