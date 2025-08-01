import React from 'react';
import { createRoot } from 'react-dom/client';
import Layout from './components/layout/layout';
import './style/reset.css';
import './style/style.css';

const container = document.getElementById('root');

const root = createRoot(container!);

root.render(
  <React.StrictMode>
    <Layout />
  </React.StrictMode>
);
