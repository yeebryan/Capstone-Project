import React, { useState } from 'react';
import Navbar1 from './Navbar';
// import MyOrder from './MyOrder';
import authAxios from './authAxios';
import withAuth from './withAuth';
import { useNavigate } from 'react-router-dom';
import "./UPF.css";


const UserPreferenceForm = ({ onSubmit, cartCount, onOpenCart }) => {
  const [category, setCategory] = useState('');
  const [foodType, setFoodType] = useState('');
  const [interval, setInterval] = useState('Weekly');
  const [startDate, setStartDate] = useState('');
  const [time, setTime] = useState('');
  const [loading, setLoading] = useState(false);
  const [foods, setFoods] = useState([]);
  const [order, setOrder] = useState({});
  const [error, setError] = useState('');
  const [step, setStep] = useState(1); // for multi-step pro bar
  const [submitted, setSubmitted] = useState(false); // track form submit
  const [showError, setShowError] = useState(false);

  const resetState = () => {
    setError('');
    setFoods([]);
  }
  
  const navigate = useNavigate();


const handleSubmit = async (e) => {
    e.preventDefault();

    resetState();
    setShowError(false);

    if (!category && !foodType) {
      setError('Please select either category or food type');
      return;
    }
  
    if (!interval || !startDate || !time) {
      setError('Please fill in all required fields');
      return;
    }

    setLoading(true);
    setSubmitted(true); // Add this line


    const params = {
      category,
    };
    if (foodType) {
      params.foodType = foodType;
    }
      const now = new Date();
      const selectedDate = new Date(`${startDate} ${time}`);
      if (selectedDate < now) {
        setError('Start date and time must be in the future');
        setLoading(false);
        return;
      }
  

    const url = "http://localhost:3000/food/random?" + new URLSearchParams(params);

    try {
      const response = await authAxios.get(url);
      setFoods(response.data.foods);
      setOrder({
        category,
        foodType,
        interval,
        startDate,
        time,
        food: response.data.foods[0],
      });
      console.error(error.response.data);
      setShowError(true); // Add this line
    } catch (error) {
      setShowError(true); // Add this line
    } finally {
      setLoading(false);
    }
  };

  const renderFoodCards = () => {
    if (loading) {
      return <img src="/imgs/loading_gif.gif" alt='loading...' />;
    }


    if (foods.length === 0 && submitted && showError) {
      return <p>No available foods found.</p>;
    }

    return foods.map((food) => (
      <div key={food.id} className='food-card'>
        <img src={food.image.url} alt={food.name} />
        <h3>{food.name}</h3>
        <p>{food.description}</p>
      </div>
    ));
  };

  const handleProceed = () => {
    // Navigate to another page
    navigate('/myorder', { state: { order, foods }});
  };

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
    <div className='user-preference-form'>
    <Navbar1 showCartIcon={false}/>
      <div className='uform-page'>
                {/* Add a container for the title and progress indicator */}
                <div className='title-progress-container'>
          <h2 className='form-title'>User Preferences</h2>
      </div>
        <div className='uform-container'>
          <form onSubmit={handleSubmit}>
            <div>
              <label>
                Category:
                <select
                  value={category}
                  onChange={(e) => setCategory(e.target.value)}
                >
                <option value="">Select a category</option>
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
                <option value="">Select a food type</option>
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
                  required
                />
              </label>
            </div>
            <div>
              <label>
                Time:
                <input type='time'
                  className='form-input-time time-input'
                  value={time}
                  onChange={(e) => setTime(e.target.value)}
                  required
                />
              </label>
            </div>
            <button type='submit'>Submit</button>
          </form>
          {!loading && foods.length === 0 && !showError ? (
            <div className="initial-gif-container">
              <img src="/imgs/happy.gif" alt="happy_panda" />
            </div>
          ): null}
          <div className='food-cards-container'>{renderFoodCards()}</div>
          {/* <MyOrder order={order} /> */}
        </div>
        <div className='button-container'>
        <button className="proceed-button" onClick={handleProceed} disabled={error || !submitted}>Proceed</button>
        </div>
        {error && <p>{error}</p>}
      </div>
      <Footer />
    </div>
  );
};

export default withAuth(UserPreferenceForm);