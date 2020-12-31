import { configureStore, createSlice } from "@reduxjs/toolkit";

const toDos = createSlice({
  name: "toDoReducer",
  initialState: [],
  reducers: {
    set: (state, action) => action.payload,
    add: (state, action) => {
      state.push(action.payload);
    },
    remove: (state, action) =>
      state.filter((todo) => todo.id !== action.payload),
    complete: (state, action) =>
      state.map((todo) =>
        todo.id === action.payload
          ? { ...todo, complete: !todo.complete }
          : todo
      ),
  },
});

export const { set, add, remove, complete } = toDos.actions;

export default configureStore({ reducer: toDos.reducer });
