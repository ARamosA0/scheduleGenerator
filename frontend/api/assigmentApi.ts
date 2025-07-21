import axios from "axios";
const API_BASE = "http://127.0.0.1:8080/api";

const api = axios.create({
  baseURL: API_BASE,
  headers: {
    "Content-Type": "application/json",
  },
});

export const getAllAssigment = async () => {
  try {
    const response = await api.get("/assigment");
    console.log("RESPONSE", response.data);
    return response.data;
  } catch (error) {
    console.log("GET Subject error", error);
    throw error;
  }
};
export const createAssigment = async (assigment: any) =>
  api.post("/assigment", assigment);
export const updateAssigment = async (assigment: any) => {
  console.log("ID", assigment);
  return api.put(`/assigment/${assigment.ID}`, assigment);
};
export const deleteAssigment = async (assigment: any) => {
  console.log("DELETE ID", assigment);
  return api.delete(`/assigment/${assigment.ID}`);
};

export default api;
