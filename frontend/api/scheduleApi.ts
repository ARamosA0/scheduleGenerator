import axios from "axios";
const API_BASE = "http://3.139.74.24:8080/api";

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
    console.log("RESPONSE", response.data);
    return response.data;
  } catch (error) {
    console.log("ERRRO API SCHEDULE", error);
    throw error;
  }
};
