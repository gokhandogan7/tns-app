import { BASE_URL } from "./constants";
import { errorMapping } from "./errorMessages";

export const getArticles = async (searchKey) => {
 
  const url = `${BASE_URL}/fullarticles?search=${searchKey}`;
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
