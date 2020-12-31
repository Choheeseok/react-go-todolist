import axios from "axios";

const axiosAPI = axios.create({
  baseURL: `http://localhost:8000`,
});

// request get/post/delete/put . . . ("/link") 형식으로 해도 좋을듯?
const requestGetToDos = async () => {
  try {
    const res = await axiosAPI.get("/");
    return res.data;
  } catch (error) {
    console.log(error);
    return null;
  }
};

const requestAddToDo = async (toDo) => {
  try {
    const res = await axiosAPI.post("/", toDo);
    return res.data;
  } catch (error) {
    console.log(error);
    return null;
  }
};

const requestDeleteToDo = async (id) => {
  try {
    const res = await axiosAPI.delete(`/${id}`);
    return res.data;
  } catch (error) {
    console.log(error);
    return null;
  }
};

const reqeustCompleteToDo = async (id) => {
  try {
    const res = await axiosAPI.put(`/${id}`);
    return res.data;
  } catch (error) {
    console.log(error);
    return null;
  }
};

const requestDetailToDo = async (id) => {
  try {
    const res = await axiosAPI.get(`/detail/${id}`);
    return res.data;
  } catch (error) {
    console.log(error);
    return null;
  }
};

export {
  requestGetToDos,
  requestAddToDo,
  requestDeleteToDo,
  reqeustCompleteToDo,
  requestDetailToDo,
};
