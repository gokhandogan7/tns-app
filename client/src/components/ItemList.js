import React, { useState } from "react";
import {useNavigate} from 'react-router-dom';
import { useApi } from "../hooks/useApi";
import { services } from "../services";
import Item from "./ui/Item";
import "./itemlist.css";
import useDebounce from "../hooks/useDebounce";
import SearchBox from "./ui/SearchBox";

function ItemList() {

  const navigate = useNavigate();

  const navigateToAuthors = () => {
    // 👇️ navigate to /contacts
    navigate('/authors');
  };
  const navigateToContents = () => {
    // 👇️ navigate to /
    navigate('/contents');
  };
  const navigateToHighlights = () => {
    // 👇️ navigate to /
    navigate('/highlights');
  };
  const navigateToHome = () => {
    // 👇️ navigate to /
    navigate('/');
  };


  const [value, setValue] = useState("");
  const debouncedValue = useDebounce(value, 700);
  const [state] = useApi(services.getArticles, [],debouncedValue);
  const [articles] = useApi(services.getUsersArticles, [], 1, 2);
  console.log(articles)
  const handleChange = (event) => {
    setValue(event.target.value);
  };

  return (
    <div data-cy="detail-page" style={{ textAlign: "center" }}>
      <h1 style={{ textAlign: "center" }}>List of Articles</h1>
      <SearchBox handleChange={handleChange} value={value} />
      <button className="buttonA buttonAA" onClick={navigateToAuthors}>Author</button>
      <button className="button button3" onClick={navigateToContents}>Content</button>
      <button className="buttonH buttonHH" onClick={navigateToHighlights}>Highlight</button>
      <button className="buttonG buttonGG" onClick={navigateToHome}>Go Back To Home</button>
      <Item items={state} search={debouncedValue} />
    </div>
  );
}

export default ItemList;
