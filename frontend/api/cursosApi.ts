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
    const response = await api.get("/teachers");
    console.log("RESPONSE", response.data);
    return response.data;
  } catch (error) {
    console.log("GET Teacher error", error);
    throw error;
  }
};
export const createCourse = async (teacher: any) =>
  api.post("/teachers", teacher);
export const updateCourse = async (teacher: any) => {
  console.log("ID", teacher);
  return api.put(`/teachers/${teacher.ID}`, teacher);
};
export const deleteCourse = async (teacher: any) => {
  console.log("DELETE ID", teacher);
  return api.delete(`/teachers/${teacher.ID}`);
};

export default api;
