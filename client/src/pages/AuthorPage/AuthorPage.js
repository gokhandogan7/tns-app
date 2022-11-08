import React, { useState } from "react";
import { useApi } from "../../hooks/useApi";
import { services } from "../../services";
import SearchBox from "../../components/ui/SearchBox";
import { useNavigate } from "react-router-dom";
import "./authorpage.css"
import useDebounce from "../../hooks/useDebounce";

export const AuthorPage = () => {
  const [value, setValue] = useState("");
  const debouncedValue = useDebounce(value, 700);
  const [authors] = useApi(services.getAllAuthors, [], debouncedValue);
  const handleChange = (event) => {
    setValue(event.target.value);
  };
  const navigate = useNavigate();

  const navigateToHome = () => {
    // 👇️ navigate to /contacts
    navigate("/fullarticles");
  };
  return (
    <>
      <div style={{ textAlign: "center" }}>
      <h1 style={{ textAlign: "center" }}>List of Author</h1>
        <SearchBox handleChange={handleChange} value={value} />
        <button className="buttonA buttonAA" onClick={navigateToHome}>Go Back Articles Page</button>
        {authors.map((author, index) => (
          <div data-cy="article" className="container" key={index}>
            <div className="itemContainer">
              <p data-cy="name" className="name">
                {author.Name}
              </p>
              <p data-cy="email" className="email">
                {author.Email}
              </p>
            </div>
          </div>
        ))}
      </div>
    </>
  );
};
