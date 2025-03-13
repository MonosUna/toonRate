import { createSlice } from '@reduxjs/toolkit';

const userData = createSlice({
  name: 'user',
  initialState: {
    userInfo: null,
  },
  reducers: {
    login(state, action) {
      state.userInfo = action.payload;
    },
    logout(state) {
      state.userInfo = null;
    },
  },
});

export const { login, logout } = userData.actions;
export default userData.reducer;