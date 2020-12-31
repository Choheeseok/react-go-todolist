import { connect } from "react-redux";
import { Link, Route } from "react-router-dom";
import { reqeustCompleteToDo, requestDeleteToDo } from "../axiosAPI";
import Detail from "../pages/Detail";
import { remove, complete } from "../store";

function ToDo({ text, id, complete, deleteToDoInStore, completeToDoInStore }) {
  const onCompleteCheck = (e) => {
    const data = reqeustCompleteToDo(id);
    if (data === null) {
      console.log("request complete to do is failed");
      return;
    }
    completeToDoInStore(id);
  };
  const onDeleteClick = (e) => {
    const data = requestDeleteToDo(id);
    if (data === null) {
      console.log("request delete to do is failed");
      return;
    }
    deleteToDoInStore(id);
  };

  return (
    <li>
      <input type="checkbox" checked={complete} onChange={onCompleteCheck} />
      <Link to={`detail/${id}`}>{text}</Link>
      <button onClick={onDeleteClick}>delete</button>
    </li>
  );
}

function mapDispatchToProps(dispatch) {
  return {
    deleteToDoInStore: (id) => dispatch(remove(id)),
    completeToDoInStore: (id) => dispatch(complete(id)),
  };
}

export default connect(null, mapDispatchToProps)(ToDo);
