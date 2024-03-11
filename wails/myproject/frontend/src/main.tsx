import React from 'react'
import {createRoot} from 'react-dom/client'
import Layout from "./components/layout/Layout";
import './style.css'

const container = document.getElementById('root')

const root = createRoot(container!)

root.render(
  <Layout/>
)
