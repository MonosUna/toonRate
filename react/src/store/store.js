import { configureStore } from '@reduxjs/toolkit';
import userReducer from './userData';

const store = configureStore({
  reducer: {
    user: userReducer,
  },
});

export default store;
