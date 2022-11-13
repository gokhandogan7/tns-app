import React, { useState } from "react";
import { useApi } from "../../hooks/useApi";
import { services } from "../../services";
import SearchBox from "../../components/ui/SearchBox";
import { useNavigate } from "react-router-dom";
import "./highlightpage.css"
import useDebounce from "../../hooks/useDebounce";

export const HighlightPage = () => {
  const [value, setValue] = useState("");
  const debouncedValue = useDebounce(value, 700);
  const [highlights] = useApi(services.getAllHighlights, [], debouncedValue);
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
      <div data-cy="highlight-page" style={{ textAlign: "center" }}>
        <h1 data-cy="highlight-header" style={{ textAlign: "center" }}>List of Highlight</h1>
        <SearchBox handleChange={handleChange} value={value} />
        <button data-cy="H-goback-button" className="buttonH buttonHH" onClick={navigateToHome}>Go Back Articles Page</button>
        {highlights.map((highlight, index) => (
          <div data-cy="highlight-container" className="container" key={index}>
            <div className="itemContainer">
              <p data-cy="short-text" className="text">
                {highlight.Short_Text}
              </p>
              <p data-cy="date" className="image">
                {highlight.Date}
              </p>
            </div>
          </div>
        ))}
      </div>
    </>
  );
};
