import React, { useState } from "react";
import { useApi } from "../../hooks/useApi";
import { services } from "../../services";
import SearchBox from "../../components/ui/SearchBox";
import { useNavigate } from "react-router-dom";
import "./contentpage.css"

export const ContentPage = () => {
  const [value, setValue] = useState("");
  const [contents] = useApi(services.getAllContents, [], value);
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
        <h1 style={{ textAlign: "center" }}>List of Content</h1>
        <SearchBox handleChange={handleChange} value={value} />
        <button className="button button3" onClick={navigateToHome}>Go Back Articles Page</button>
        {contents.map((content, index) => (
          <div data-cy="content" className="container" key={index}>
            <div className="itemContainer">
              <p data-cy="text" className="text">
                {content.Text}
              </p>
              <p data-cy="image" className="image">
                {content.Image}
              </p>
            </div>
          </div>
        ))}
      </div>
    </>
  );
};
