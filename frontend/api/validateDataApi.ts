import axios from "axios";
const API_BASE = "http://127.0.0.1:8080/api";

const api = axios.create({
  baseURL: API_BASE,
  headers: {
    "Content-Type": "multipart/form-data",
  },
});

export const validateFile = async (file: any) => {
  console.log("FILE", file);
  const formData = new FormData();
  formData.append("file", file);
  const result = await api.post("/validate", formData);
  console.log("VALIDATE RESULT", result);
  return result.data;
};

export default api;
