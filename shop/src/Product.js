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
  const [data, setData] = useState([]); // set initial state (data) to an empty array [] not object {}
  const { restaurant_id } = useParams();
  const [cartCount, setCartCount] = useState(0);
  const [cart, setCart] = useState([]);
  const [total, setTotal] = useState(0)





// axios 
  useEffect(() => {
    console.log(`Fetching restaurant data for ID ${restaurant_id}...`);

    axios
      .get(`http://localhost:3000/restaurants/${restaurant_id}`) // change from ?id=${id} to ${id} because API url is .com/products/1  // But couldn't map due to not being array
      .then((res) => {
        console.log(JSON.stringify(res)) 
        setData(res.data);

      })
      .catch((err) => console.log(err));
  }, [restaurant_id]);

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
      {data &&
        data.map((item) => (
          <Card className="item-card" key={item.id}>
            <Card.Img variant="top" src={item.image.url} alt={item.image.url} />
            <div className="item-title">{item.name}</div>
            <div>{item.price}</div>
            <div>{item.description}</div>
              {console.log(`yo log this: ${item} & this ${item.name}`)}
            <div className="center">
            <AddToCartButton onClick={() => onAddToCart(item)} />
            </div>
          </Card>
        ))}
  </div>
);

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


  return (
    <div>
          <Navbar1 cartCount={props.cartCount} onOpenCart = {props.onOpenCart}/>
      <div className="product_item">{ProductItem}</div>
      <Footer />
    </div>
  );
  }

export default Product;
