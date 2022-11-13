import React, { useState } from "react";
import { useApi } from "../../hooks/useApi";
import { services } from "../../services";
import SearchBox from "../../components/ui/SearchBox";
import { useNavigate } from "react-router-dom";
import "./contentpage.css"
import useDebounce from "../../hooks/useDebounce";

export const ContentPage = () => {
  const [value, setValue] = useState("");
  const debouncedValue = useDebounce(value, 700);
  const [contents] = useApi(services.getAllContents, [], debouncedValue);
  const handleChange = (event) => {
    setValue(event.target.value);
  };

  const navigate = useNavigate();

  const navigateToHome = () => {
    // 👇️ navigate to /contacts
    navigate("/fullarticles");
  };
  console.log(contents)
  return (
    <>
      <div data-cy="content-page"  style={{ textAlign: "center" }}>
        <h1 data-cy="content-header" style={{ textAlign: "center" }}>List of Content</h1>
        <SearchBox handleChange={handleChange} value={value} />
        <button data-cy="C-goback-button" className="buttonC buttonCC" onClick={navigateToHome}>Go Back Articles Page</button>
        {contents.map((content, index) => (
          <div data-cy="content-container" className="container" key={index}>
            <div className="itemContainer">
              <p data-cy="text" className="text">
                {content.Text}
              </p>
              <img data-cy="image" src={content.Image} style={{width:400, height:300}}/>
            </div>
          </div>
        ))}
      </div>
    </>
  );
};
