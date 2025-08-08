import axios from "axios";
const API_BASE = "http://127.0.0.1:8080/api";

const api = axios.create({
  baseURL: API_BASE,
  headers: {
    "Content-Type": "application/json",
  },
});

export const getAllRooms = async () => {
  try {
    const response = await api.get("/room");
    console.log("RESPONSE", response.data);
    return response.data;
  } catch (error) {
    console.log("GET Teacher error", error);
    throw error;
  }
};
export const createRoom = async (room: any) => {
  console.log("API ROOM CREATE", room);
  return api.post("/room", room);
};
export const updateRoom = async (room: any) =>
  api.put(`/room/${room.ID}`, room);
export const deleteRoom = async (room: any) => api.delete(`/room/${room.ID}`);

export const uploadRoomExcelData = async (
  file: string,
  selectedMapping: any,
) => {
  console.log("FILEID", file);
  console.log("Selected MAPPING", selectedMapping);
  const data = {
    fileId: file,
    selectedMapping: selectedMapping,
  };
  return api.post("/room/bulk", data);
};

export default api;
