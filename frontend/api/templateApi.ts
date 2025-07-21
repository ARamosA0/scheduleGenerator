import axios from "axios";
const API_BASE = "http://127.0.0.1:8080/api";

const api = axios.create({
  baseURL: API_BASE,
  headers: {
    "Content-Type": "application/json",
  },
});

export const getAllTemplates = async () => {
  try {
    const response = await api.get("/template");
    console.log("RESPONSE", response.data);
    return response.data;
  } catch (error) {
    console.log("GET Subject error", error);
    throw error;
  }
};
export const getTemplate = async (subject: any) => {
  try {
    const response = await api.get(`/template/${subject.ID}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};
export const createTemplate = async (template: any) => {
  console.log("TEMPLTE", template);
  return api.post("/template", template);
};

export const updateTemplate = async (subject: any) => {
  console.log("ID", subject);
  return api.put(`/template/${subject.ID}`, subject);
};
export const deleteTemplate = async (subject: any) => {
  console.log("DELETE ID", subject);
  return api.delete(`/template/${subject.ID}`);
};

export default api;
