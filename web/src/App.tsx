import { useState } from 'react';
// import reactLogo from './assets/react.svg';
// import viteLogo from '/vite.svg';
import './App.css';
import 'csh-material-bootstrap/dist/css/csh-material-bootstrap.css';
import 'csh-material-bootstrap/dist/js/color-modes.js';
import { Button, Navbar, NavbarBrand } from 'reactstrap';
import HomeNavbar from './components/Navbar';

function App() {
  const [count, setCount] = useState(0);

  // Note, this doesn't work right now, you have to go to /auth/login first
  fetch('/api/v1/users/1').then((x) => x.json().then((x) => console.log(x)));
  return (
    <>
      <HomeNavbar/>
      <Button onClick={() => setCount(count + 1)}>
        HI, {count} clicks so far
      </Button>
    </>
  );
}

export default App;
