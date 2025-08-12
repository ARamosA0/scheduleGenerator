import axios from "axios";
const API_BASE = "http://127.0.0.1:8080/api";

const api = axios.create({
  baseURL: API_BASE,
  headers: {
    "Content-Type": "application/json",
  },
});

export const getAllGroups = async () => {
  try {
    const response = await api.get("/group");
    console.log("RESPONSE", response.data);
    return response.data;
  } catch (error) {
    console.log("GET Subject error", error);
    throw error;
  }
};
export const createGroup = async (group: any) => api.post("/group", group);
export const updateGroup = async (group: any) => {
  console.log("ID", group);
  return api.put(`/group/${group.ID}`, group);
};
export const deleteGroup = async (group: any) => {
  console.log("DELETE ID", group);
  return api.delete(`/group/${group.ID}`);
};

export const uploadGroupExcelData = async (
  file: string,
  selectedMapping: any,
) => {
  console.log("FILEID", file);
  console.log("Selected MAPPING", selectedMapping);
  const data = {
    fileId: file,
    selectedMapping: selectedMapping,
  };
  return api.post("/group/bulk", data);
};

export default api;
