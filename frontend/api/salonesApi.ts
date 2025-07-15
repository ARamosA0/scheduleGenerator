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
    const response = await api.get("/rooms");
    console.log("RESPONSE", response.data);
    return response.data;
  } catch (error) {
    console.log("GET Teacher error", error);
    throw error;
  }
};
export const createRoom = async (room: any) => api.post("/rooms", room);
export const updateRoom = async (room: any) => {
  console.log("ID", room);
  return api.put(`/rooms/${room.ID}`, room);
};
export const deleteRoom = async (room: any) => {
  console.log("DELETE ID", room);
  return api.delete(`/rooms/${room.ID}`);
};

export default api;
