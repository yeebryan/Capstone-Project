import React from 'react';
import "bootstrap/dist/css/bootstrap.min.css";
import { Navbar, Nav, Container } from 'react-bootstrap';
import Cart from './Cart';
import "./App.css";

const Navbar1 = (props) => {

  const user = JSON.parse(localStorage.getItem('user'));

  const handleLogout = () => {
    localStorage.removeItem('user');
    localStorage.removeItem('token');
    window.location.href = '/login';
  };

  return (
    <Navbar className="navbar-style justify-content-lg-between">
      <Container>
        <Navbar.Brand href="/">
          <img
            src="/imgs/foodpanda_logo.png"
            alt="Logo"
            className="d-none d-md-block logo-img"
          />
          <span className="d-md-none logo-text">foodpanda</span>
        </Navbar.Brand>
        <div className="d-none d-lg-block">
          <Nav className="ml-auto">
            <Nav.Link href="/about">Playlist</Nav.Link>
            <Nav.Link href="/">About Us</Nav.Link>
          </Nav>
        </div>
        <Nav className="ms-auto">
          <Nav>
            <Cart cartCount={props.cartCount} onOpenCart={props.onOpenCart} />
          </Nav>
          {user ? (
            <>
              <Navbar.Text className="me-4">
                Hello, {user.first_name}
              </Navbar.Text>
              <Nav.Link onClick={handleLogout}>Logout</Nav.Link>
            </>
          ) : (
            <Navbar.Text className="me-4">
              <a href="/login">Login</a>
            </Navbar.Text>
          )}
        </Nav>
      </Container>
    </Navbar>
  );
};

export default Navbar1;
