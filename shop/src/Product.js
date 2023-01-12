import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom'
import axios from 'axios';

// individual product
// functional component


const Product = ({ match }) => {
    const [data1, setData] = useState([]); // set initial state (data) to an empty array

    useEffect(() => {
        axios
          .get(`https://fakestoreapi.com/products/?id=${match.params.id}`)
          .then((res) => {
            setData(res.data);
            console.log(res.data);
          })
          .catch((err) => console.log(err));
      }, [match.params.id]);






return (
    <div>
      {data1.map((item) => {
        return (
          <div className='product-container' key={item.id}>
            <div>
              <img className='prod-image' src={item.image} alt='' />
            </div>
            <div>
              <h1 className='brand'>{item.title}</h1>
              <h2>{item.category}</h2>
              <p>{item.description}</p>
              <p>
                <strong>Price:</strong> {item.price}
              </p>
              <p>
                <strong>Rating:</strong> {item.rating.rate}
              </p>
            </div>
          </div>
        );
      })}
      <div className='back'>
        <Link to='/'>Feature Products</Link>
      </div>
    </div>
  );



    
}

export default Product

