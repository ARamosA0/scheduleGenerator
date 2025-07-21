import axios from "axios";
const API_BASE = "http://127.0.0.1:8080/api";

const api = axios.create({
  baseURL: API_BASE,
  headers: {
    "Content-Type": "application/json",
  },
});

export const getAllCourses = async () => {
  try {
    const response = await api.get("/subject");
    console.log("RESPONSE", response.data);
    return response.data;
  } catch (error) {
    console.log("GET Subject error", error);
    throw error;
  }
};
export const createCourse = async (subject: any) =>
  api.post("/subject", subject);
export const updateCourse = async (subject: any) => {
  console.log("ID", subject);
  return api.put(`/subject/${subject.ID}`, subject);
};
export const deleteCourse = async (subject: any) => {
  console.log("DELETE ID", subject);
  return api.delete(`/subject/${subject.ID}`);
};

export default api;
