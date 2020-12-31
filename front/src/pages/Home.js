import React, { useState, useEffect } from "react";
import { connect } from "react-redux";
import { requestGetToDos, requestAddToDo } from "../axiosAPI";
import ToDo from "../components/ToDo";
import { set, add } from "../store";

function Home({ toDos, setToDoInStore, addTodoInStore }) {
  const [input, setInput] = useState("");
  useEffect(() => {
    const fetchData = async () => {
      const data = await requestGetToDos();
      if (data === null) {
        console.log("request get to dos is failed");
        return;
      }
      setToDoInStore(data);
    };
    fetchData();
  }, []);

  const onChange = (e) => {
    setInput(e.target.value);
  };

  const addToDo = async () => {
    const id = Date.now();
    const newToDo = {
      text: input,
      complete: false,
      id: id,
    };

    const data = await requestAddToDo(newToDo);
    if (data === null) {
      console.log("request add to do is failed");
      return;
    }
    addTodoInStore(data);
  };

  const onSubmit = (e) => {
    e.preventDefault();
    addToDo();
    setInput("");
  };

  return (
    <>
      <h1>Your To dos</h1>
      <form onSubmit={onSubmit}>
        <input type="text" value={input} onChange={onChange} />
        <button>Add</button>
      </form>
      <ul>
        {toDos.map((toDo) => (
          <ToDo {...toDo} key={toDo.id} />
        ))}
      </ul>
    </>
  );
}

function mapStateToProps(state) {
  return { toDos: state };
}

function mapDispatchToProps(dispatch) {
  return {
    addTodoInStore: ({ id, text }) => dispatch(add({ id, text })),
    setToDoInStore: (toDos) => dispatch(set(toDos)),
  };
}

export default connect(mapStateToProps, mapDispatchToProps)(Home);
