import React, { useState } from "react";
import { useApi } from "../../hooks/useApi";
import { services } from "../../services";
import SearchBox from "../../components/ui/SearchBox";
import { useNavigate } from "react-router-dom";
import "./highlightpage.css"

export const HighlightPage = () => {
  const [value, setValue] = useState("");
  const [highlights] = useApi(services.getAllHighlights, [], value);
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
        <h1 style={{ textAlign: "center" }}>List of Highlight</h1>
        <SearchBox handleChange={handleChange} value={value} />
        <button className="buttonH buttonHH" onClick={navigateToHome}>Go Back Articles Page</button>
        {highlights.map((highlight, index) => (
          <div data-cy="highlight" className="container" key={index}>
            <div className="itemContainer">
              <p data-cy="text" className="text">
                {highlight.Short_Text}
              </p>
              <p data-cy="image" className="image">
                {highlight.Date}
              </p>
            </div>
          </div>
        ))}
      </div>
    </>
  );
};
