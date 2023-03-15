import React from 'react';
import "bootstrap/dist/css/bootstrap.min.css";
import { Navbar, Nav, Container } from 'react-bootstrap';
import Cart from './Cart';


const Navbar1 = (props) => {
    return (
      <Navbar expand="lg">
        <Container>
          <Navbar.Brand href="/">
            <img src="/imgs/panda.png" alt="Logo" />
          </Navbar.Brand>
          <Navbar.Toggle aria-controls="responsive-navbar-nav" />
          <Navbar.Collapse id="responsive-navbar-nav">
            <Nav className="ml-auto">
              <Nav.Link href="/">Home</Nav.Link>
              <Nav.Link href="/about">Playlist</Nav.Link>
              <Nav.Link href="/">About Us</Nav.Link>
            </Nav>
            <Nav className='ms-auto'> 
            <Nav>
            <Cart cartCount={props.cartCount} onOpenCart={props.onOpenCart} />
            </Nav>
                <Navbar.Text className='me-4'>
                    Hello <a href="/">Bryan</a>
                </Navbar.Text>
            </Nav>
          </Navbar.Collapse>
        </Container>
      </Navbar>
    );
  };
  
  // don't use ml-auto, use ms-auto 

  export default Navbar1;
  

  //
  
  
  
  