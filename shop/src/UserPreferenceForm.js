import React, { useState } from 'react';
import axios from 'axios';
import Navbar1 from './Navbar';
import MyOrder from './MyOrder';
import "./App.css";


const UserPreferenceForm = ({ onSubmit, cartCount, onOpenCart }) => {
  const [category, setCategory] = useState('Western');
  const [foodType, setFoodType] = useState('');
  const [interval, setInterval] = useState('Weekly');
  const [startDate, setStartDate] = useState('');
  const [time, setTime] = useState('');
  const [loading, setLoading] = useState(false);
  const [foods, setFoods] = useState([]);
  const [order, setOrder] = useState({});


const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);

    const params = {
      category,
    };
    if (foodType) {
      params.foodType = foodType;
    }

    const url = "http://localhost:3000/food/random?" + new URLSearchParams(params);

    try {
      const response = await axios.get(url);
      setFoods(response.data.foods);
      setOrder({
        category,
        foodType,
        interval,
        startDate,
        time,
        food: response.data.foods[0],
      });
    } catch (error) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const renderFoodCards = () => {
    if (loading) {
      return <img src="/imgs/loading_gif.gif" alt='loading...' />;
    }


    if (foods.length === 0) {
      return <p>No available foods found.</p>;
    }

    return foods.map((food) => (
      <div key={food.id} className='food-card'>
        <img src={food.imageUrl} alt={food.name} />
        <h3>{food.name}</h3>
        <p>{food.description}</p>
      </div>
    ));
  };

  return (
    <div className='user-preference-form'>
      <Navbar1 cartCount={cartCount} onOpenCart={onOpenCart} />
      <div className='uform-page'>
        <div className='uform-container'>
          <form onSubmit={handleSubmit}>
            <div>
              <label>
                Category:
                <select
                  value={category}
                  onChange={(e) => setCategory(e.target.value)}
                >
                <option value='Western'>Western</option>
                <option value='Chinese'>Chinese</option>
                <option value='Japanese'>Japanese</option>
                <option value='Indian'>Indian</option>
                <option value='Mexican'>Mexican</option>
                <option value='Vegetarian'>Vegetarian</option>
              </select>
            </label>
          </div>
          <div>
            <label>
              Food Type:
              <select
                  value={foodType}
                  onChange={(e) => setFoodType(e.target.value)}
                >
                <option value='Burger'>Burger</option>
                <option value='Pizza'>Pizza</option>
                <option value='Sushi'>Sushi</option>
                <option value='Taco'>Taco</option>
                <option value='Chicken'>Chicken</option>
                <option value='Noodle'>Noodle</option>
                <option value='Breakfast'>Breakfast</option>
                <option value='Rice'>Rice</option>
                <option value='Salad'>Salad</option>
              </select>
              {/* <input
                type='text'
                value={foodType}
                onChange={(e) => setFoodType(e.target.value)}
              /> */}
            </label>
            </div>
            <div>
              <label>
                Interval:
                <select
                  value={interval}
                  onChange={(e) => setInterval(e.target.value)}
                >
                  <option value='Weekly'>Weekly</option>
                  <option value='Bi-weekly'>Bi-weekly</option>
                  <option value='Monthly'>Monthly</option>
                </select>
              </label>
            </div>
            <div>
              <label>
                Start Date:
                <input
                  type='date'
                  value={startDate}
                  onChange={(e) => setStartDate(e.target.value)}
                />
              </label>
            </div>
            <div>
              <label>
                Time:
                <input type='time'
                  value={time}
                  onChange={(e) => setTime(e.target.value)}
                />
              </label>
            </div>
            <button type='submit'>Submit</button>
          </form>
          <div className='food-cards-container'>{renderFoodCards()}</div>
          <MyOrder order={order} />
        </div>
      </div>
    </div>
  );
};

export default UserPreferenceForm;