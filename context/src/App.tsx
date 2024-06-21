import React from 'react';
// import logo from './logo.svg';
import './App.css';
import { NotificationContextProvider } from './context/notificationContext/notificationContext';

import { RandomComponent } from './components/randomComponent/randomComponent';

function App() {
  return (
    <NotificationContextProvider>
      <h2>Hello</h2>
      <RandomComponent></RandomComponent>
    </NotificationContextProvider>
  );
}

export default App;
