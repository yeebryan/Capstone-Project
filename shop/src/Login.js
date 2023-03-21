import React, { useState } from 'react';
import axios from 'axios';
import Navbar1 from './Navbar';
import "./App.css";

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loggedIn, setLoggedIn] = useState(false);
  const [user, setUser] = useState({})

  const handleSubmit = async (e) => {
    e.preventDefault();
    console.log('handleSubmit called');


    try {
      const response = await axios.post('http://localhost:3000/users/login', {
        email,
        password,
      });

      console.log('before response:', response); // Add this line
      if (response.status === 200) {
        const { user: user1, token } = response.data;
        console.log(response.data)

        // Update state to indicate user is logged in
        setLoggedIn(true);
        setUser(user1)

        // Do something with the user data, such as storing it in localStorage
        localStorage.setItem('user', JSON.stringify(user1));
        localStorage.setItem('token', token);
        console.log(`this is: ${user1}`)
           // Redirect to the main page
      window.location.href = '/';
      } else {
        setError('Invalid email or password');
      }
    } catch (err) {
      setError('Invalid email or password');
    }
  };

  return (
    <div>
    <Navbar1 loggedIn={loggedIn} user={user.first_name} />
    <div className="login-page">
    <div className="login-container">
      <h2>Login</h2>
      {error && <p>{error}</p>}
      <form onSubmit={handleSubmit}>
        <label>
          Email:
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </label>
        <label s>
          Password:
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </label>
        <button type="submit">Login</button>
      </form>
    </div>
    </div>
    </div>

  );
};

export default Login;
