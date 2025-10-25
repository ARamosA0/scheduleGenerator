import axios from "axios";
const API_BASE = "http://3.139.74.24:8080/api";

const api = axios.create({
  baseURL: API_BASE,
  headers: {
    "Content-Type": "application/json",
  },
});

export const getAllTeachers = async () => {
  try {
    const response = await api.get("/teachers");
    console.log("RESPONSE", response.data);
    return response.data;
  } catch (error) {
    console.log("GET Teacher error", error);
    throw error;
  }
};
export const createTeacher = async (teacher: any) =>
  api.post("/teachers", teacher);
export const updateTeacher = async (teacher: any) => {
  console.log("ID", teacher);
  return api.put(`/teachers/${teacher.ID}`, teacher);
};
export const deleteTeacher = async (teacher: any) => {
  console.log("DELETE ID", teacher);
  return api.delete(`/teachers/${teacher.ID}`);
};

export const uploadTeacherExcelData = async (
  file: string,
  selectedMapping: any,
) => {
  console.log("FILEID", file);
  console.log("Selected MAPPING", selectedMapping);
  const data = {
    fileId: file,
    selectedMapping: selectedMapping,
  };
  return api.post("/teachers/bulk", data);
};

export default api;
