import axios from "axios";
const API_BASE = "http://127.0.0.1:8080/api";

const api = axios.create({
  baseURL: API_BASE,
  headers: {
    "Content-Type": "application/json",
  },
});

export const getScheduleById = async (scheduleId: string) => {
  try {
    console.log("SCHEDULEID", scheduleId);
    const response = await api.get(`/schedule/${scheduleId}`);
    console.log("RESPONSE", response);
    return response.data;
  } catch (error) {
    console.log("ERRRO API SCHEDULE", error);
    throw error;
  }
};
