import { BASE_URL } from "./constants";
import { errorMapping } from "./errorMessages";

export const getArticles = async () => {
  const url = `${BASE_URL}/articles`;
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
  const message = errorMapping[response?.status] ?? errorMapping.default;
  throw new Error(message);
};

export const services = {
  getArticles,
};
