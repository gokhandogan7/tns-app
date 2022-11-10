import { BASE_URL } from "./constants";
import { errorMapping } from "./errorMessages";

export const getArticles = async (searchKey,limit) => {
 

  const url = `${BASE_URL}/fullarticles?search=${searchKey}&limit=${limit}`;
  console.log(url)
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
  const message = errorMapping[response?.status] ?? errorMapping.default;
  throw new Error(message);
};

export const getAllAuthors = async (searchKey) => {
 
  const url = `${BASE_URL}/authors?search=${searchKey}`;
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
  const message = errorMapping[response?.status] ?? errorMapping.default;
  throw new Error(message);
};

export const getAllContents = async (searchKey) => {
 
  const url = `${BASE_URL}/contents?search=${searchKey}`;
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
  const message = errorMapping[response?.status] ?? errorMapping.default;
  throw new Error(message);
};

export const getAllHighlights = async (searchKey) => {
 
  const url = `${BASE_URL}/highlights?search=${searchKey}`;
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
  const message = errorMapping[response?.status] ?? errorMapping.default;
  throw new Error(message);
};

export const getArticlesByUserId = async (id) => {
 
  const url = `${BASE_URL}/fullarticles/${id}`;
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
  const message = errorMapping[response?.status] ?? errorMapping.default;
  throw new Error(message);
};




export const services = {
  getArticles,
  getAllAuthors,
  getAllContents,
  getAllHighlights,
  getArticlesByUserId
};
