import React, { useState, useEffect } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import {useParams } from "react-router-dom";
import axios from "axios";
import Navbar1 from "./Navbar";
import AddToCartButton from './AddtoCart';
import { Card } from 'react-bootstrap';
import "./App.css";



// individual product
// functional component

const Product = (props) => {
  const [data, setData] = useState([]); // set initial state (data) to an empty array
  const { _id } = useParams();
  const [cartCount, setCartCount] = useState(0);
  const [cart, setCart] = useState([]);
  const [total, setTotal] = useState(0)





// axios 
  useEffect(() => {
    axios
      .get(`http://localhost:3000/restaurants/${_id}`) // change from ?id=${id} to ${id} because API url is .com/products/1  // But couldn't map due to not being array
      .then((res) => {
        console.log(JSON.stringify(res)) 
        setData(res.data);
      })
      .catch((err) => console.log(err));
  }, [_id]);


// cart count, and keep track of item added

const onAddToCart = (item) => {
  setCartCount(cartCount + 1);
  setCart([...cart, item])
  setTotal(total + item.price)
}

// open the cart

/*const onOpenCart = () => {
  setShowCart(true)
}

// close cart
const onCloseCart = () => {
  setShowCart(false);
}

// clear cart btn function - setState back to 0

const clearCart = () => {
  setCartCount(0); // the counter to 0 
  setCart([]);
  setShowCart(false);
}
*/




const ProductItem = (
  <div className="product-container" key={data.id}>
    <div>
      <img className="prod-image" src={data.image} alt="" />
    </div>
    <div>
      <h1 className="brand">{data.title}</h1>
      <h2>{data.category}</h2>
      <p>{data.description}</p>
      {data.menu &&
        data.menu.map((item) => (
          <Card key={item.id}>
            <Card.Body>
              <Card.Title>{item.name}</Card.Title>
              <Card.Text>{`$${item.price.toFixed(2)}`}</Card.Text>

              {item.options &&
                item.options
                  .filter(
                    (option, index, self) =>
                      option.name !== 'none' &&
                      index === self.findIndex((t) => t.name === option.name)
                  )
                  .map((option) => (
                    <div className="form-check" key={option.name}>
                      <input
                        className="form-check-input"
                        type="checkbox"
                        name={option.name}
                        id={option.name}
                        value={option.price}
                      />
                      <label
                        className="form-check-label"
                        htmlFor={option.name}
                      >
                        {option.name} ({`$${option.price.toFixed(2)}`})
                      </label>
                      <br />
                    </div>
                  ))}

<AddToCartButton onClick={() => onAddToCart(item)} />
            </Card.Body>
          </Card>
        ))}
    </div>
  </div>
);



  return (
    <div>
      <Navbar1 cartCount={cartCount}/> 
      <div className="product_item">{ProductItem}</div>
    </div>
  );
  }

export default Product;