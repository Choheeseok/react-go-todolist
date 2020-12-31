import React, { useState, useEffect } from "react";
import { requestDetailToDo, requestDeleteToDo } from "../axiosAPI";

function Detail({ match, history }) {
  const [toDo, setToDo] = useState({ id: 0, text: "default", complete: false });
  useEffect(() => {
    const fetchData = async () => {
      const { id, text, complete } = await requestDetailToDo(match.params.id);
      setToDo({ id, text, complete });
    };
    fetchData();
  }, []);

  const onDeleteClick = (e) => {
    requestDeleteToDo(toDo.id);
    history.push("/");
  };

  return (
    <>
      <div>Detail page : {toDo.id}</div>
      <div>{toDo.text}</div>
      <div>{toDo.complete ? "COMPLETED" : "UNCOMPLETED"}</div>
      <button onClick={onDeleteClick}>delete</button>
    </>
  );
}

export default Detail;
