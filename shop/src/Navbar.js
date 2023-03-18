import React from 'react';
import "bootstrap/dist/css/bootstrap.min.css";
import { Navbar, Nav, Container } from 'react-bootstrap';
import Cart from './Cart';
import "./App.css";

const Navbar1 = (props) => {
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
          <Navbar.Text className="me-4">
            Hello <a href="/">Bryan</a>
          </Navbar.Text>
        </Nav>
      </Container>
    </Navbar>
  );
};

export default Navbar1;
