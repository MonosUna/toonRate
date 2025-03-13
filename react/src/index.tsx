import { createRoot } from 'react-dom/client';
import React from "react";
import { Provider } from 'react-redux';
import store from './store/store';
import App from "./App";

// Clear the existing HTML content
document.body.innerHTML = '<div id="app"></div>';

// Render your React component instead
const root = createRoot(document.getElementById('app'));
root.render(
  <Provider store={store} children={undefined}>
    <App />
  </Provider>
);