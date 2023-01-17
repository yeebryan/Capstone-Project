import React, { useState, useEffect } from "react";
import { Link, useParams } from "react-router-dom";
import axios from "axios";

// individual product
// functional component

const Product = () => {
  const [data, setData] = useState([]); // set initial state (data) to an empty array
  const { id } = useParams();

  useEffect(() => {
    axios
      .get(`https://fakestoreapi.com/products/${id}`) // change from ?id=${id} to ${id} because API url is .com/products/1  // But couldn't map due to not being array
      .then((res) => {
        console.log(JSON.stringify(res)) 
        setData(res.data);
      })
      .catch((err) => console.log(err));
  }, [id]);


  return (
    <div>
      <div className="product-container" key={data.id}>
        <div>
          <img className="prod-image" src={data.image} alt="" />
        </div>
        <div>
          <h1 className="brand">{data.title}</h1>
          <h2>{data.category}</h2>
          <p>{data.description}</p>
          <p>
            <strong>Price:</strong> {data.price}
          </p>
        </div>
      </div>
      <div className="back">
        <Link to="/">Feature Products</Link>
      </div>
    </div>
  );
};

export default Product;
