import axios from "axios";
const HTTP = axios.create({
   baseURL: process.env.API_URL,
});

export const request1 = async () => {
   return await HTTP.get("/");
};

export const request2 = async (data: { token: string; _id: string }) =>
   await HTTP.post("/route2", data);
